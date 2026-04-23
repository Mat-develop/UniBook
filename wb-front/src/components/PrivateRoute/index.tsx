import React from "react";
import { Navigate } from "react-router-dom";
import { getStoredToken, isTokenExpired, setAuthToken } from "../../utils/auth";

interface PrivateRouteProps {
  children: React.ReactNode;
}

const PrivateRoute: React.FC<PrivateRouteProps> = ({ children }) => {
  const token = getStoredToken();
  const expired = token ? isTokenExpired(token) : true;

  if (!token || expired) {
    setAuthToken(null);
    return <Navigate to="/login" replace />;
  }

  return <>{children}</>;
};

export default PrivateRoute;