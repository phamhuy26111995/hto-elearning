import axiosClient from "@/api/axiosClient";
import { TEACHER_ENDPOINT } from "@/consts/const";
import { User } from "@/types/user";

const STUDENT_ENDPOINT = `/${TEACHER_ENDPOINT}/users`;

const studentServices = {
  getAll: async () => {
    return await axiosClient.get(`${STUDENT_ENDPOINT}/get-students`);
  },

  getById: async (id: number) => {
    return await axiosClient.get(`${STUDENT_ENDPOINT}/${id}`);
  },

  createStudent: async (data: User) => {
    return await axiosClient.post(`${STUDENT_ENDPOINT}/create-student`, data);
  },

  updateStudent: async (data: User) => {
    return await axiosClient.put(`${STUDENT_ENDPOINT}/update-student`, data);
  },

  deleteStudent: async (id: number) => {
    return await axiosClient.delete(`${STUDENT_ENDPOINT}/${id}`);
  },

  enrollment: async (data: User) => {
    return await axiosClient.post(`${STUDENT_ENDPOINT}/enrollment`, data);
  },
};

export default studentServices;
