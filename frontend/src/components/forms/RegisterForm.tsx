"use client";
import { registerSchema, registerType } from "@/lib/validations";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useRouter } from "next/navigation";
const RegisterForm = () => {
    const router = useRouter();
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);

    //react hook RegisterForm

    const form = useForm<registerType>({
        resolver: zodResolver(registerSchema),
    });
    async function handleSubmit(data: registerType) {
        console.log(data);
    }

    return <div>Form</div>;
};
export default RegisterForm;
