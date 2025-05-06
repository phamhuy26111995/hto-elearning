import axiosClient from "@/api/axiosClient";
import { TEACHER_ENDPOINT } from "@/consts/const";

const STUDENT_ENDPOINT = `/${TEACHER_ENDPOINT}/users`;

const studentServices = {
  getAll: async () => {
    const { data } = await axiosClient.get(`${STUDENT_ENDPOINT}/get-students`);
    return data;
  },
};

export default studentServices;
