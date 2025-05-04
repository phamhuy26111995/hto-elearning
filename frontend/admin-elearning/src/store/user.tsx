// src/stores/userStore.ts
import { create } from "zustand";
import { devtools } from "zustand/middleware";

interface User {
  userId: number;
  username: string;
  email: string;
  role: string;
}

interface UserStore {
  users: User[];
  currentUserLogin?: User;
  setUsers: (users: User[]) => void;
  setCurrentUserLogin: (userInfo: User) => void;
}

const useUserStore = create<UserStore>()(
  devtools(
    (set) => ({
      users: [],
      currentUserLogin: undefined,

      setUsers(users) {
        set({ users });
      },

      setCurrentUserLogin(userInfo) {
        set({ currentUserLogin: userInfo });
      },
    }),
    { name: "UserStore" } // shows up as “UserStore” in DevTools
  )
);

export default useUserStore;
