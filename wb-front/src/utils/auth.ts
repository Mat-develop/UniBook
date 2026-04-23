import axios from 'axios';

export const getStoredToken = (): string | null => {
  return localStorage.getItem('authToken') ?? sessionStorage.getItem('authToken');
};

export const removeStoredToken = (): void => {
  localStorage.removeItem('authToken');
  sessionStorage.removeItem('authToken');
};

export const isTokenExpired = (token: string): boolean => {
  try {
    const payloadBase64 = token.split('.')[1];
    if (!payloadBase64) return true;

    const payloadJson = atob(payloadBase64.replace(/-/g, '+').replace(/_/g, '/'));
    const payload = JSON.parse(payloadJson);

    if (!payload.exp) return false;
    return Date.now() / 1000 >= payload.exp;
  } catch {
    return true;
  }
};

export const getUserIdFromToken = (): number | null => {
  const token = getStoredToken();
  if (!token) return null;
  try {
    const payloadBase64 = token.split('.')[1];
    if (!payloadBase64) return null;
    const payload = JSON.parse(atob(payloadBase64.replace(/-/g, '+').replace(/_/g, '/')));
    return payload.userId != null ? Number(payload.userId) : null;
  } catch {
    return null;
  }
};

export const setAuthToken = (token: string | null, remember = false): void => {
  if (token) {
    if (remember) {
      localStorage.setItem('authToken', token);
    } else {
      sessionStorage.setItem('authToken', token);
    }
    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
  } else {
    localStorage.removeItem('authToken');
    sessionStorage.removeItem('authToken');
    delete axios.defaults.headers.common['Authorization'];
  }
};
