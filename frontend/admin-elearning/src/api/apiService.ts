// src/api/apiService.ts
import axiosClient from "./axiosClient";
import { AxiosResponse } from "axios";

export const apiService = {
  get: <T>(url: string, params?: object): Promise<AxiosResponse<T>> => {
    return axiosClient.get<T>(url, { params });
  },

  post: <T>(url: string, data?: object): Promise<AxiosResponse<T>> => {
    return axiosClient.post<T>(url, data);
  },

  put: <T>(url: string, data?: object): Promise<AxiosResponse<T>> => {
    return axiosClient.put<T>(url, data);
  },

  delete: <T>(url: string): Promise<AxiosResponse<T>> => {
    return axiosClient.delete<T>(url);
  },
};