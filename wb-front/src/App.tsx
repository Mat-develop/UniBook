import { useEffect } from 'react';
import axios from 'axios';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import MainLayout from './components/Layout/Layout';
import Feed from './components/Feed/Feed';
import CommunityFeed from './components/Communities';
import Login from './pages/Login/Login';
import Register from './pages/Login/Register';
import IndexPage from './pages/Index';
import PrivateRoute from './components/PrivateRoute';
import { getStoredToken, isTokenExpired, setAuthToken } from './utils/auth';

function App() {
  useEffect(() => {
    const token = getStoredToken();

    if (token && !isTokenExpired(token)) {
      setAuthToken(token);
    } else {
      setAuthToken(null);
    }

    const interceptor = axios.interceptors.response.use(
      (response) => response,
      (error) => {
        if (error.response?.status === 401) {
          setAuthToken(null);
          window.location.href = '/login';
        }
        return Promise.reject(error);
      }
    );

    return () => {
      axios.interceptors.response.eject(interceptor);
    };
  }, []);

  return (
    <Router>
      <Routes>
        <Route path="/" element={<IndexPage />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route
          element={
            <PrivateRoute>
              <MainLayout />
            </PrivateRoute>
          }
        >
          <Route path="/home" element={<Feed />} />
          <Route path="/communities" element={<CommunityFeed />} />
          <Route path="/popular" element={<Feed />} />
          <Route path="/new" element={<Feed />} />
        </Route>
        <Route path="*" element={<Navigate to="/login" replace />} />
      </Routes>
      <ToastContainer
        position="top-center"
        autoClose={3000}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick
        pauseOnFocusLoss
        draggable
        pauseOnHover
        theme="colored"
      />
    </Router>
  );
}

export default App;
