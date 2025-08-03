// src/stores/userStore.ts
import { UserInfo } from "@/types/user";
import { create } from "zustand";
import { devtools } from "zustand/middleware";

interface UserStore {
  currentUserLogin?: UserInfo;
  setCurrentUserLogin: (userInfo: UserInfo) => void;
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
