import axiosClient from "@/api/axiosClient";
import { TEACHER_ENDPOINT } from "@/consts/const";

const STUDENT_ENDPOINT = `/${TEACHER_ENDPOINT}/users`;

const studentServices = {
  getAll: async () => {
    return await axiosClient.get(`${STUDENT_ENDPOINT}/get-students`);
  },

  getById: async (id: number) => {
    return await axiosClient.get(`${STUDENT_ENDPOINT}/${id}`);
  },

  createStudent: async (data: any) => {
    return await axiosClient.post(`${STUDENT_ENDPOINT}/create-student`, data);
  },

  deleteStudent: async (id: any) => {
    return await axiosClient.delete(`${STUDENT_ENDPOINT}/student/${id}`);
  },

  enrollment: async (data: any) => {
    return await axiosClient.post(`${STUDENT_ENDPOINT}/enrollment`, data);
  },
};

export default studentServices;
