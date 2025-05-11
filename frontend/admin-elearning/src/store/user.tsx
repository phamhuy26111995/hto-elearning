// src/stores/userStore.ts
import studentServices from "@/services/student";
import { User } from "@/types/user";
import { create } from "zustand";
import { devtools } from "zustand/middleware";

interface UserStore {
  users: User[];
  currentUserLogin?: User;
  setUsers: () => void;
  setCurrentUserLogin: (userInfo: User) => void;
}

const useUserStore = create<UserStore>()(
  devtools(
    (set) => ({
      users: [],
      currentUserLogin: undefined,

      async setUsers() {
        const { data } = await studentServices.getAll();
        set({ users: data });
      },

      setCurrentUserLogin(userInfo) {
        set({ currentUserLogin: userInfo });
      },
    }),
    { name: "UserStore" } // shows up as “UserStore” in DevTools
  )
);

export default useUserStore;
