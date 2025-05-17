// src/components/auth/ProtectedRoute.tsx
import useUserStore from "@/store/user";
import { ReactNode } from "react";
import { Navigate, Outlet } from "react-router";

export function ProtectedRoute() {
  const token = localStorage.getItem("token");
  const { currentUserLogin } = useUserStore();

  if (!token) {
    // no token → redirect to login
    return <Navigate to="/login" replace />;
  }

  // otherwise render matched child routes
  return currentUserLogin && <Outlet />;
}
