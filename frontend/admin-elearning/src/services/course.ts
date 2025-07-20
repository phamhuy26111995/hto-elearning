import axiosClient from "@/api/axiosClient";
import { TEACHER_ENDPOINT } from "@/consts/const";
import { CreateCourseBody, FormCourse } from "@/types/course";

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

  createBulk: async (data: CreateCourseBody) => {
    return await axiosClient.post(`/${TEACHER_ENDPOINT}/modules/create-bulk`, data);
  },

};

export default courseServices;
