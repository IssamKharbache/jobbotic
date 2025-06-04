"use client";
import { registerSchema, registerType } from "@/lib/validations";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useRouter } from "next/navigation";
import { motion } from "framer-motion";
import { FcGoogle } from "react-icons/fc";
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "../ui/form";
import { Input } from "../ui/input";
import { z } from "zod";
import axios, { AxiosError } from "axios";
import GradientButton from "../buttons/GradientButton";
import { toast } from "sonner";
import { Button } from "../ui/button";

const RegisterForm = () => {
    const router = useRouter();
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);

    const form = useForm<registerType>({
        resolver: zodResolver(registerSchema),
        defaultValues: {
            first_name: "",
            last_name: "",
            email: "",
            password: "",
        },
    });

    const handleSubmit = async (data: registerType) => {
        setLoading(true);
        const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;
        try {
            const response = await axios.post(
                `${backendUrl}/auth/register`,
                data,
            );
            if (response.status === 200) {
                setLoading(false);
                toast("Your account was created succesfully", {
                    style: {
                        background: "green",
                        color: "white",
                    },
                });
                setError("");
                router.push("/dashboard");
            }
        } catch (error) {
            setLoading(false);
            const axiosError = error as AxiosError;
            if (axiosError.response) {
                const status = axiosError.response.status;
                const errorMessage =
                    (axiosError.response.data as { error?: string })?.error ||
                    "An error occurred";

                if (status === 400) {
                    setError(errorMessage);
                    toast(errorMessage, {
                        style: {
                            background: "red",
                            color: "white",
                        },
                        description:
                            "The email you trying to use is already signed into jobbotic , try a different email address",
                    });
                } else {
                    toast("Internal server error");
                }
            } else if (axiosError.request) {
                console.log("No response from server");
            } else {
                console.log("Error", axiosError.message);
            }
        }
    };
    return (
        <Form {...form}>
            <motion.form
                onSubmit={form.handleSubmit(handleSubmit)}
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.5 }}
                className="space-y-5 min-w-md p-6 bg-white border rounded-2xl"
            >
                {error ? (
                    <div className="text-white bg-red-500 text-center rounded-lg py-2">
                        {error}
                    </div>
                ) : null}
                {["first_name", "last_name", "email", "password"].map(
                    (fieldName) => (
                        <FormField
                            key={fieldName}
                            control={form.control}
                            name={
                                fieldName as keyof z.infer<
                                    typeof registerSchema
                                >
                            }
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel className="mb-2">
                                        {fieldName
                                            .split("_")
                                            .map(
                                                (word) =>
                                                    word
                                                        .charAt(0)
                                                        .toUpperCase() +
                                                    word.slice(1),
                                            )
                                            .join(" ")}
                                    </FormLabel>
                                    <FormControl>
                                        <Input
                                            type={
                                                fieldName === "password"
                                                    ? "password"
                                                    : "text"
                                            }
                                            placeholder={`Enter your ${fieldName.replace("_", " ")}`}
                                            {...field}
                                        />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                    ),
                )}
                <GradientButton loading={loading} />
                <div className="flex items-cente justify-center text-gray-500 w-full">
                    <p>------</p>
                    <p className="text-center ">OR</p>
                    <p>------</p>
                </div>
                <Button className="flex items-center gap-5 w-full">
                    Continue with Google
                    <FcGoogle />
                </Button>
            </motion.form>
        </Form>
    );
};

export default RegisterForm;
