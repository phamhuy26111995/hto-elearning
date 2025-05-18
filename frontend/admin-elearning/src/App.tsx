import { Suspense, useEffect } from "react";

import { BrowserRouter, Navigate, Route, Routes } from "react-router";
import Layout from "./components/project/layout/Layout";
import Home from "./pages/Home";

import CourseDetail from "./pages/course/CourseDetail";
import Student from "./pages/student/Student";
import { ProtectedRoute } from "./routes/ProtectedRoute";
import Login from "./pages/login/Login";
import NotFound from "./pages/NotFound";
import Course from "./pages/course/Course";
import TestingPage from "./pages/testing/TestingPage";
import useUserStore from "./store/user";
import { apiService } from "./api/apiService";
import StudentDetail from "./pages/student/StudentDetail";
import userServices from "./services/user";
import { ROUTES } from "./consts/const";

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
      const response = await userServices.getCurrentUserLogin();
      const { data }: any = response;
      setCurrentUserLogin(data.userInfo);
    } catch (error) {
      console.log(error);
    }
  }

  return (
    <BrowserRouter>
      <Suspense fallback={<div>Loading...</div>}>
        <Routes>
          <Route path={ROUTES.LOGIN} element={<Login />} />
          <Route path="/" element={<Navigate to={ROUTES.COURSE} replace />} />
          <Route element={<ProtectedRoute />}>
            <Route element={<Layout />}>
              <Route path="/home" element={<Home />} />
              <Route
                path={ROUTES.COURSE + "/:courseId"}
                element={<CourseDetail />}
              />
              <Route path={ROUTES.COURSE} element={<Course />} />
              <Route
                path={ROUTES.STUDENT + "/:studentId"}
                element={<StudentDetail />}
              />
              <Route path={ROUTES.STUDENT} element={<Student />} />
              <Route path="/testing" element={<TestingPage />} />
            </Route>
          </Route>
          <Route path="*" element={<NotFound />} />
        </Routes>
      </Suspense>
    </BrowserRouter>
  );
}

export default App;
