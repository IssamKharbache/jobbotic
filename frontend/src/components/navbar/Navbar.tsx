"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";
import JobboticLogo from "../logos/JobboticLogo";
import MobileNav from "./MobileNav";
import { useUserStore } from "../../../store/UserStore";
import UserAvatar from "./UserAvatar";

export const navbarLinks = [
    { name: "Home", href: "/" },
    { name: "Features", href: "/features" },
    { name: "Pricing", href: "/pricing" },
    { name: "How it Works", href: "/how-it-works" },
    { name: "About", href: "/about" },
    { name: "Contact", href: "/contact" },
];

const Navbar = () => {
    const pathname = usePathname();
    const { user } = useUserStore();
    return (
        <nav className="flex h-24 items-center justify-between border py-5 px-7">
            <JobboticLogo />
            <div className="hidden lg:flex items-center gap-5">
                {navbarLinks.map((item) => {
                    const isActive = pathname === item.href;
                    return (
                        <div key={item.href} className="relative group pb-1">
                            <Link
                                href={item.href}
                                className={`relative inline-block ${isActive ? "font-medium" : "text-gray-600 hover:text-gray-900"}`}
                            >
                                {item.name}
                                {/* Single SVG that handles both states */}
                                <svg
                                    className="absolute left-0 -bottom-2.5 w-full h-5"
                                    viewBox="0 0 200 30"
                                    preserveAspectRatio="none"
                                >
                                    <defs>
                                        <linearGradient
                                            id="linkGradient"
                                            x1="0%"
                                            y1="0%"
                                            x2="100%"
                                            y2="0%"
                                        >
                                            <stop
                                                offset="0%"
                                                stopColor="#3B82F6"
                                            />
                                            <stop
                                                offset="50%"
                                                stopColor="#8B5CF6"
                                            />
                                            <stop
                                                offset="100%"
                                                stopColor="#EC4899"
                                            />
                                        </linearGradient>
                                    </defs>
                                    <path
                                        d="M0,20 C25,8 50,25 75,12 C100,25 125,8 150,20 C160,25 175,15 200,20"
                                        stroke="url(#linkGradient)"
                                        fill="none"
                                        strokeWidth="4"
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        strokeDasharray={isActive ? "0" : "250"}
                                        strokeDashoffset={
                                            isActive ? "0" : "250"
                                        }
                                        className={
                                            isActive
                                                ? ""
                                                : "opacity-0 group-hover:opacity-100 transition-opacity duration-200 group-hover:animate-[fast-draw_0.6s_forwards]"
                                        }
                                    />
                                </svg>
                            </Link>
                        </div>
                    );
                })}
            </div>

            <div className="hidden lg:flex items-center gap-6">
                {user.id ? (
                    <>
                        <UserAvatar />
                    </>
                ) : (
                    <>
                        <Link
                            href="/register"
                            className="border border-black hover:bg-black hover:text-white duration-300 rounded-full py-3 px-6 cursor-pointer"
                        >
                            Get started
                        </Link>

                        <Link
                            href="/sign-in"
                            className="text-gray-600 hover:text-gray-900"
                        >
                            Sign in
                        </Link>
                    </>
                )}
            </div>
            <MobileNav />
        </nav>
    );
};

export default Navbar;
