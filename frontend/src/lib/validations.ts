import { z } from "zod";

export const registerSchema = z.object({
    first_name: z.string().min(1, "your first name is required"),
    last_name: z.string().min(1, "your last name is required"),
    email: z
        .string()
        .min(1, { message: "Your Email address is required" })
        .email("This is not a valid email."),
    password: z
        .string()
        .min(6, { message: "Password must be at least 6 characters" }),
});

export type registerType = z.infer<typeof registerSchema>;

export const signInSchema = z.object({
    email: z.string().min(1, {
        message: "Your Email address is required",
    }),
    password: z.string().min(1, "Your password is required"),
});

export type signInType = z.infer<typeof signInSchema>;
