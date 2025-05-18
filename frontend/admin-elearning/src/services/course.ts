import axiosClient from "@/api/axiosClient";
import { TEACHER_ENDPOINT } from "@/consts/const";

const COURSE_ENDPOINT = `/${TEACHER_ENDPOINT}/course`;

const courseServices = {
  getAll: async () => {
    return await axiosClient.get(`${COURSE_ENDPOINT}/get-all`);
  },

  getById: async (id: number) => {
    return await axiosClient.get(`${COURSE_ENDPOINT}/${id}`);
  },

  create: async (data: any) => {
    return await axiosClient.post(`${COURSE_ENDPOINT}/create`, data);
  },

  update: async (data: any) => {
    return await axiosClient.put(`${COURSE_ENDPOINT}/update`, data);
  },

  delete: async (id: number) => {
    return await axiosClient.delete(`${COURSE_ENDPOINT}/${id}`);
  },

};

export default courseServices;
