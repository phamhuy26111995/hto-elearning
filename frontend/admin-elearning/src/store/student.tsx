// src/stores/userStore.ts
import studentServices from "@/services/student";
import { User } from "@/types/user";
import { create } from "zustand";
import { devtools } from "zustand/middleware";

interface StudentStore {
  students: User[];
  studentDetail?: User;
  fetchStudents: () => void;
  fetchStudentDetail: (id: number) => void;
}

const useStudentStore = create<StudentStore>()(
  devtools(
    (set) => ({
      students: [],


      async fetchStudents() {
        const { data } = await studentServices.getAll();
        const userList = data.users as User[];
        console.log("🚀 ~ data:", data)
        const usersWithNo = userList.map((user, index) => ({ ...user, numberNo: index + 1,key : user.userId }));

        set({ students : usersWithNo }, undefined, "fetchStudents");
      },

      async fetchStudentDetail(id: number) {
        const { data } = await studentServices.getById(id);
        const studentDetail = data.user as User;
        set({ studentDetail }, undefined, "fetchStudentDetail");
      },



    }),
    { name: "StudentStore" } // shows up as “UserStore” in DevTools
  )
);

export default useStudentStore;
