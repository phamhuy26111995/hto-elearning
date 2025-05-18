import axiosClient from "@/api/axiosClient";
import { TEACHER_ENDPOINT } from "@/consts/const";

const ENDPOINT = `/${TEACHER_ENDPOINT}/users`;

const userServices = {
  getCurrentUserLogin: async () => {
    return await axiosClient.get(`${ENDPOINT}/current-user`);
  },

};

export default userServices;
