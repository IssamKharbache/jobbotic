import { z } from "zod";

export const registerSchema = z.object({
    first_name: z.string().min(1, "your first name is required"),
    last_name: z.string().min(1, "your first name is required"),
    email: z.string().min(1, "your first name is required"),
    password: z.string().min(1, "your first name is required"),
});

export type registerType = z.infer<typeof registerSchema>;
