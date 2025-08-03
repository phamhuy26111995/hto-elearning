import axiosClient from "@/api/axiosClient";
import { ENDPOINT } from "@/consts/const";
import { CreateCourseBody } from "@/types/course";



const courseServices = {
  getAll: async () => {
    return await axiosClient.get(`${ENDPOINT}/get-all`);
  },

  getById: async (id: number) => {
    return await axiosClient.get(`${ENDPOINT}/${id}`);
  },

  create: async (data: any) => {
    return await axiosClient.post(`${ENDPOINT}/create`, data);
  },

  update: async (data: any) => {
    return await axiosClient.put(`${ENDPOINT}/update`, data);
  },

  createBulk: async (data: CreateCourseBody) => {
    return await axiosClient.post(`/${ENDPOINT}/modules/create-bulk`, data);
  },

};

export default courseServices;
