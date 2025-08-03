import axiosClient from "@/api/axiosClient";
import { COMMON_ENDPOINT } from "@/consts/const";


const loginUserServices = {
  getCurrentUserLogin: async () => {
    return await axiosClient.get(`/${COMMON_ENDPOINT}/current-user`);
  },

};

export default loginUserServices;
