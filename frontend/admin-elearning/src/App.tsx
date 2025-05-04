import { Suspense } from "react";

import { BrowserRouter, Navigate, Route, Routes } from "react-router";
import Layout from "./components/project/layout/Layout";
import Home from "./pages/Home";

import CourseDetail from "./pages/course/CourseDetail";
import Student from "./pages/student/Student";
import { ProtectedRoute } from "./routes/ProtectedRoute";
import Login from "./pages/login/Login";
import NotFound from "./pages/NotFound";
import Course from "./pages/course/Course";

function App() {
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
              <Route path="/student/:studentId" element={<Student />} />
              <Route path="/courses" element={<Course />} />
            </Route>
          </Route>
          <Route path="*" element={<NotFound />} />
        </Routes>
      </Suspense>
    </BrowserRouter>
  );
}

export default App;
