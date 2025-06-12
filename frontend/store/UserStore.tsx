import { create } from "zustand";
import { persist } from "zustand/middleware";

export interface User {
    id: string;
    email: string;
    first_name: string;
    last_name: string;
}

interface UserStore {
    user: User;
    setUser: (user: User) => void;
    hasHydrated: boolean;
    setHasHydrated: (state: boolean) => void;
}

export const useUserStore = create<UserStore>()(
    persist(
        (set) => ({
            user: {
                id: "",
                email: "",
                first_name: "",
                last_name: "",
            },
            setUser: (user) => set({ user }),
            hasHydrated: false,
            setHasHydrated: (state) => set({ hasHydrated: state }),
        }),
        {
            name: "user",
            onRehydrateStorage: () => (state) => {
                state?.setHasHydrated(true);
            },
        },
    ),
);
