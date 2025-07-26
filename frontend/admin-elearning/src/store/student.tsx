// src/stores/userStore.ts
import studentServices from "@/services/student";
import { User } from "@/types/user";
import { create } from "zustand";
import { devtools } from "zustand/middleware";

interface StudentStore {
  students: User[];
  fetchStudents: () => void;
}

const useStudentStore = create<StudentStore>()(
  devtools(
    (set) => ({
      students: [],

      // API_LIST_ALL_STUDENTS
      async fetchStudents() {
        const { data } = await studentServices.getAll();
        const userList = data.users as User[];
        const usersWithNo = userList.map((user, index) => ({ ...user, numberNo: index + 1,key : user.userId }));

        set({ students : usersWithNo }, undefined, "fetchStudents");
      },
     
    }),
    { name: "StudentStore" } // shows up as “UserStore” in DevTools
  )
);

export default useStudentStore;
