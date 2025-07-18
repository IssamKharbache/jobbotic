import type { Metadata } from "next";
import { Geist, Geist_Mono, Space_Grotesk } from "next/font/google";
import "../globals.css";
import Navbar from "@/components/navbar/Navbar";
import { Toaster } from "@/components/ui/sonner";
import UserProvider from "@/components/providers/AuthProvider";
import { cookies } from "next/headers";
const geistSans = Geist({
    variable: "--font-geist-sans",
    subsets: ["latin"],
});

const geistMono = Geist_Mono({
    variable: "--font-geist-mono",
    subsets: ["latin"],
});

const roboto = Space_Grotesk({
    subsets: ["latin"],
});

export const metadata: Metadata = {
    title: "Jobbotic",
    description:
        "Jobbotic is a smart job application tracker that connects to your Gmail to automatically detect and organize your job applications. Stay on top of your job hunt, get helpful insights, and avoid duplicate submissions.",
};

export default async function InterfaceLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    const cookieStore = await cookies();
    const token = cookieStore.get("token")?.value || null;
    return (
        <html lang="en">
            <body
                className={`${geistSans.variable} ${geistMono.variable} ${roboto.className} antialiased`}
            >
                <Toaster position="top-right" />
                <Navbar />
                <UserProvider token={token || ""}>
                    <div>{children}</div>
                </UserProvider>
            </body>
        </html>
    );
}
