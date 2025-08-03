// src/stores/userStore.ts
import { create } from "zustand";
import { devtools } from "zustand/middleware";

type ModalStore = {
    isOpenModal : boolean
    modalContent : React.ReactNode | null
    openModal: (content : React.ReactNode) => void
    closeModal: () => void

}

const useModalStore = create<ModalStore>()(
  devtools(
    (set) => ({
      isOpenModal : false,
      openModal: (content : React.ReactNode) => set({ isOpenModal : true, modalContent : content }),
      closeModal: () => set({ isOpenModal : false, modalContent : null }),
    }),
    { name: "ModalStore", enabled: false } // shows up as “UserStore” in DevTools
  )
);

export default useModalStore;
