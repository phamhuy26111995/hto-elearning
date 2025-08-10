import { Suspense, useEffect } from "react";

import { BrowserRouter, Navigate, Route, Routes } from "react-router";
import Layout from "./components/project/layout/Layout";
import Home from "./pages/Home";

import { ROUTES } from "./consts/const";
import Course from "./pages/course/Course";
import CourseDetail from "./pages/course/CourseDetail";
import Login from "./pages/login/Login";
import NotFound from "./pages/NotFound";
import TestingPage from "./pages/testing/TestingPage";
import UserDetail from "./pages/user/UserDetail";
import Users from "./pages/user/Users";
import { ProtectedRoute } from "./routes/ProtectedRoute";
import loginUserServices from "./services/loginUser";
import useUserStore from "./store/user";

import { ReactQueryDevtools } from '@tanstack/react-query-devtools'

import {
  QueryClient,
  QueryClientProvider
} from "@tanstack/react-query";

const queryClient = new QueryClient();

function App() {
  const { setCurrentUserLogin } = useUserStore();

  useEffect(() => {
    getUserInfo();
  }, []);

  async function getUserInfo() {
    const token = localStorage.getItem("token");
    if (!token) {
      return;
    }

    try {
      const response = await loginUserServices.getCurrentUserLogin();
      const { data }: any = response;
      setCurrentUserLogin(data.userInfo);
    } catch (error) {
      console.log(error);
    }
  }

  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Suspense fallback={<div>Loading...</div>}>
          <Routes>
            <Route path={ROUTES.LOGIN} element={<Login />} />
            <Route path="/" element={<Navigate to={ROUTES.COURSE} replace />} />
            <Route element={<ProtectedRoute />}>
              <Route element={<Layout />}>
                {/* ROUTE_HOME */}
                <Route path="/home" element={<Home />} />

                {/* ROUTE_COURSE */}
                <Route path={ROUTES.COURSE} element={<Course />} />
                <Route
                  path={ROUTES.COURSE + "/add"}
                  element={<CourseDetail />}
                />
                <Route
                  path={ROUTES.COURSE + "/edit/:courseId"}
                  element={<CourseDetail />}
                />

                {/* ROUTE_STUDENT */}
                <Route path={ROUTES.USER} element={<Users />} />
                <Route path={ROUTES.USER + "/add"} element={<UserDetail />} />
                <Route
                  path={ROUTES.USER + "/edit/:studentId"}
                  element={<UserDetail />}
                />

                {/* ROUTE_TESTING */}
                <Route path="/testing" element={<TestingPage />} />
              </Route>
            </Route>
            <Route path="*" element={<NotFound />} />
          </Routes>
        </Suspense>
      </BrowserRouter>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}

export default App;
