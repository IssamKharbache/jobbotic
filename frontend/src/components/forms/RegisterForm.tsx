"use client";
import { registerSchema, type registerType } from "@/lib/validations";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useRouter } from "next/navigation";
import { motion } from "framer-motion";
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "../ui/form";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { Separator } from "../ui/separator";
import axios, { type AxiosError } from "axios";
import { toast } from "sonner";
import { User, Mail, Lock } from "lucide-react";
import Link from "next/link";
import GradientButton from "../buttons/GradientButton";
import DangerAlert from "../alerts/DangerAlert";

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

    const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;
    const handleSubmit = async (data: registerType) => {
        setLoading(true);
        try {
            const response = await axios.post(
                `${backendUrl}/auth/register`,
                data,
            );
            if (response.status === 200) {
                setLoading(false);
                toast("Your account was created successfully", {
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
                            "The email you trying to use is already signed into jobbotic, try a different email address",
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

    const googleSignup = async () => {
        try {
            const res = await axios.get(`${backendUrl}/auth/google/login`);
            router.push(res.data.url);
        } catch (error) {
            console.log(error);
        }
    };

    return (
        <div className="space-y-6">
            {/* Google Sign Up Button */}
            <Button
                type="button"
                variant="outline"
                className="w-full h-12 text-gray-700 border-gray-300 hover:bg-gray-50"
                onClick={googleSignup}
            >
                <svg className="w-5 h-5 mr-3" viewBox="0 0 24 24">
                    <path
                        fill="currentColor"
                        d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
                    />
                    <path
                        fill="currentColor"
                        d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
                    />
                    <path
                        fill="currentColor"
                        d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
                    />
                    <path
                        fill="currentColor"
                        d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
                    />
                </svg>
                Continue with Google
            </Button>

            <div className="relative">
                <div className="absolute inset-0 flex items-center">
                    <Separator className="w-full" />
                </div>
                <div className="relative flex justify-center text-xs uppercase">
                    <span className="bg-white px-2 text-gray-500">
                        Or continue with email
                    </span>
                </div>
            </div>

            {/* Error Display */}
            {error ? <DangerAlert message={error} /> : null}
            {/* Email/Password Form */}
            <Form {...form}>
                <motion.form
                    onSubmit={form.handleSubmit(handleSubmit)}
                    initial={{ opacity: 0, y: 20 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ duration: 0.5 }}
                    className="space-y-4"
                >
                    {/* First Name */}
                    <FormField
                        control={form.control}
                        name="first_name"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel className="text-sm font-medium text-gray-700">
                                    First Name
                                </FormLabel>
                                <div className="relative">
                                    <User className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 h-4 w-4" />
                                    <FormControl>
                                        <Input
                                            placeholder="Enter your first name"
                                            className="pl-10 h-12 border-gray-300 focus:border-blue-500 focus:ring-blue-500"
                                            {...field}
                                        />
                                    </FormControl>
                                </div>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    {/* Last Name */}
                    <FormField
                        control={form.control}
                        name="last_name"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel className="text-sm font-medium text-gray-700">
                                    Last Name
                                </FormLabel>
                                <div className="relative">
                                    <User className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 h-4 w-4" />
                                    <FormControl>
                                        <Input
                                            placeholder="Enter your last name"
                                            className="pl-10 h-12 border-gray-300 focus:border-blue-500 focus:ring-blue-500"
                                            {...field}
                                        />
                                    </FormControl>
                                </div>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    {/* Email */}
                    <FormField
                        control={form.control}
                        name="email"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel className="text-sm font-medium text-gray-700">
                                    Email address
                                </FormLabel>
                                <div className="relative">
                                    <Mail className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 h-4 w-4" />
                                    <FormControl>
                                        <Input
                                            type="email"
                                            placeholder="Enter your email"
                                            className="pl-10 h-12 border-gray-300 focus:border-blue-500 focus:ring-blue-500"
                                            {...field}
                                        />
                                    </FormControl>
                                </div>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    {/* Password */}
                    <FormField
                        control={form.control}
                        name="password"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel className="text-sm font-medium text-gray-700">
                                    Password
                                </FormLabel>
                                <div className="relative">
                                    <Lock className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 h-4 w-4" />
                                    <FormControl>
                                        <Input
                                            type="password"
                                            placeholder="Enter your password"
                                            className="pl-10 h-12 border-gray-300 focus:border-blue-500 focus:ring-blue-500"
                                            {...field}
                                        />
                                    </FormControl>
                                </div>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    {/* Terms and Conditions */}
                    <div className="flex items-center">
                        <input
                            id="terms"
                            name="terms"
                            type="checkbox"
                            required
                            className="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                        />
                        <label
                            htmlFor="terms"
                            className="ml-2 text-sm text-gray-700"
                        >
                            I agree to the{" "}
                            <Link
                                href="/terms"
                                className="text-blue-600 hover:text-blue-500 font-medium"
                            >
                                Terms of Service
                            </Link>{" "}
                            and{" "}
                            <Link
                                href="/privacy"
                                className="text-blue-600 hover:text-blue-500 font-medium"
                            >
                                Privacy Policy
                            </Link>
                        </label>
                    </div>
                    <GradientButton loading={loading} text="Join now" />
                </motion.form>
            </Form>

            <div className="text-center">
                <p className="text-sm text-gray-600">
                    Already have an account?{" "}
                    <Link
                        href="/sign-in"
                        className="text-blue-600 hover:text-blue-500 font-medium"
                    >
                        Sign in
                    </Link>
                </p>
            </div>
        </div>
    );
};

export default RegisterForm;
