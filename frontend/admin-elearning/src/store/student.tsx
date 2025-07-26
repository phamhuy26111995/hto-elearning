// src/stores/userStore.ts
import studentServices from "@/services/student";
import { User } from "@/types/user";
import { create } from "zustand";
import { devtools } from "zustand/middleware";

interface StudentStore {
  students: User[];
  fetchStudents: () => void;
  deleteStudent: (id: number) => void;
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

      // API_DELETE_STUDENT
      async deleteStudent(id: number) {
        try {
          await studentServices.deleteStudent(id);
          set((state) => ({
            students: state.students.filter((student) => student.userId !== id),
          }));
        } catch (error) {
          console.error("Failed to delete student:", error);
        }
      }
     
    }),
    { name: "StudentStore" } // shows up as “UserStore” in DevTools
  )
);

export default useStudentStore;
