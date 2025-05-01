import { Suspense } from "react";

import { BrowserRouter, Navigate, Route, Routes } from "react-router";
import AdminLayout from "./components/project/layout/AdminLayout";
import Home from "./pages/Home";



function App() {

  return (
    <BrowserRouter>
      <Suspense fallback={<div>Loading...</div>}>
        <Routes>
          <Route path="/" element={<Navigate to="/home" replace />} />
          <Route element={<AdminLayout />}>
            <Route path="/home" element={<Home />} />
          </Route>
        </Routes>
      </Suspense>
    </BrowserRouter>
  );
}

export default App;
