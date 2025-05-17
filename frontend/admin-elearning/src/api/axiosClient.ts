// src/api/axiosClient.ts
import axios from "axios";

// 1. Khởi tạo instance
const axiosClient = axios.create({
  baseURL: `${import.meta.env.VITE_API_URL}`,  // ví dụ "https://api.example.com"
  headers: {
    "Content-Type": "application/json",
  },
  timeout: 10000,  // timeout 10 giây
});

// 2. Request interceptor (thêm token nếu có)
axiosClient.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("token");
    if (token && config.headers) {
      config.headers.Authorization = `${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// 3. Response interceptor (xử lý lỗi chung)
axiosClient.interceptors.response.use(
  (response) => response,
  (error) => {
    // ví dụ: nếu 401 Unauthorized thì chuyển về trang login
    localStorage.removeItem("token");
    if (error.response?.status === 401) {
      window.location.href = "/login";
    }
    return Promise.reject(error);
  }
);

export default axiosClient;
