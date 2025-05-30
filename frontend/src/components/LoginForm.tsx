import axios from "axios";
import { useState } from "react";
const LoginForm = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const login = async () => {
        try {
            const res = await axios.post(
                "http://localhost:4000/api/auth/login",
                {
                    email,
                    password,
                },
            );
            localStorage.setItem("token", res.data.token);
            console.log(res);
        } catch (error) {
            console.log(error);
        }
    };

    const handleLinkGoogle = async () => {
        const token = localStorage.getItem("token");
        try {
            const res = await axios.get(
                "http://localhost:4000/api/auth/google/link",
                {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                },
            );
            window.location.href = res.data.url;
        } catch (error) {
            console.log(error);
        }
    };

    return (
        <div>
            <form>
                <p>Login</p>
                <input
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    type="text"
                    placeholder="Email"
                />
                <input
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <button type="button" onClick={login}>
                    Login
                </button>

                <button type="button" onClick={handleLinkGoogle}>
                    Link Google Account
                </button>
            </form>
        </div>
    );
};

export default LoginForm;
