import axiosClient from "@/api/axiosClient";
import { ENDPOINT } from "@/consts/const";
import { Paging, User } from "@/types/user";

const userServices = {
  getAllTeacher: async (paging: Paging) => {
    return await axiosClient.get(`${ENDPOINT}/user/teacher`, {
      params: paging,
    });
  },

  getAllStudent: async (paging: Paging) => {
    return await axiosClient.get(`${ENDPOINT}/user/student`, {
      params: paging,
    });
  },

  getById: async (id: number) => {
    return await axiosClient.get(`${ENDPOINT}/user/${id}`);
  },

  createUser: async (data: User) => {
    return await axiosClient.post(`${ENDPOINT}/user/create`, data);
  },

  updateUser: async (data: User) => {
    return await axiosClient.put(`${ENDPOINT}/user/update`, data);
  },
};

export default userServices;
