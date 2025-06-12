import axios from "axios";
import { useUserStore } from "../../../store/UserStore";

export const useLogout = () => {
    const { setUser, setHasHydrated } = useUserStore();

    const logout = async () => {
        try {
            await axios.post(
                `${process.env.NEXT_PUBLIC_BACKEND_URL}/auth/logout`,
                {},
                { withCredentials: true },
            );
            setUser({
                id: "",
                first_name: "",
                last_name: "",
                email: "",
            });
            setHasHydrated(false);
        } catch (error) {
            console.error("Logout failed:", error);
        }
    };

    return logout;
};
