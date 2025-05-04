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
      const response = await apiService.get("/api/v1/teacher/users/current-user");
      const { data } : any = response;
      setCurrentUserLogin(data.userInfo);
    } catch (error) {
      console.log(error);
    }
    
  }

  return (
    <BrowserRouter>
      <Suspense fallback={<div>Loading...</div>}>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/" element={<Navigate to="/home" replace />} />
          <Route element={<ProtectedRoute />}>
            <Route element={<Layout />}>
              <Route path="/home" element={<Home />} />
              <Route path="/course/:courseId" element={<CourseDetail />} />
              <Route path="/courses" element={<Course />} />
              <Route path="/student/:studentId" element={<StudentDetail />} />
              <Route path="/students" element={<Student />} />
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
