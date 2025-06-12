"use client";
import { useEffect } from "react";
import axios from "axios";
import { useUserStore } from "../../../store/UserStore";

const UserProvider: React.FC<{ children: React.ReactNode; token: string }> = ({
    children,
    token,
}) => {
    const { user, setUser, hasHydrated } = useUserStore();

    useEffect(() => {
        if (!hasHydrated) return;
        if (!token) {
            setUser({
                id: "",
                email: "",
                first_name: "",
                last_name: "",
            });
            return;
        }
        if (user.email !== "") return;

        const fetchUser = async () => {
            try {
                console.log("Fetching user...");
                const res = await axios.get(
                    `${process.env.NEXT_PUBLIC_BACKEND_URL}/users/user/get`,
                    {
                        headers: {
                            Authorization: `Bearer ${token}`,
                        },
                        withCredentials: true,
                    },
                );
                setUser(res.data.user);
            } catch (err) {
                console.error("Failed to fetch user:", err);
                setUser({
                    id: "",
                    first_name: "",
                    last_name: "",
                    email: "",
                });
            }
        };

        fetchUser();
    }, [token, user.email, hasHydrated, setUser]);

    return <>{children}</>;
};

export default UserProvider;
