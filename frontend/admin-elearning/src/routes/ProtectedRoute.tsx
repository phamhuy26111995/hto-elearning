// src/components/auth/ProtectedRoute.tsx
import { ReactNode } from "react";
import { Navigate, Outlet } from "react-router";



export function ProtectedRoute() {
  const token = localStorage.getItem("token");

  if (!token) {
    // no token → redirect to login
    return <Navigate to="/login" replace />;
  }

  // otherwise render matched child routes
  return <Outlet />;
}
