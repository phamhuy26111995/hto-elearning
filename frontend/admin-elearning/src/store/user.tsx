// src/stores/userStore.ts
import studentServices from "@/services/student";
import { User } from "@/types/user";
import { create } from "zustand";
import { devtools } from "zustand/middleware";

interface UserStore {
  currentUserLogin?: User;
  setCurrentUserLogin: (userInfo: User) => void;
}

const useUserStore = create<UserStore>()(
  devtools(
    (set) => ({
      currentUserLogin: undefined,

      setCurrentUserLogin(userInfo) {
        set({ currentUserLogin: userInfo });
      },
    }),
    { name: "UserStore" } // shows up as “UserStore” in DevTools
  )
);

export default useUserStore;
