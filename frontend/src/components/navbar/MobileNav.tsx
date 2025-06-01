"use client";
import { useState, useEffect, useRef } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { navbarLinks } from "./Navbar";
import { FaInstagram } from "react-icons/fa";
import { FaXTwitter } from "react-icons/fa6";

const MobileNav = () => {
    const [isOpen, setIsOpen] = useState(false);
    const menuRef = useRef<HTMLDivElement>(null);
    const buttonRef = useRef<HTMLButtonElement>(null);
    // Close menu when clicking outside
    useEffect(() => {
        const handleClickOutside = (event: MouseEvent) => {
            if (
                menuRef.current &&
                !menuRef.current.contains(event.target as Node) &&
                buttonRef.current &&
                !buttonRef.current.contains(event.target as Node)
            ) {
                setIsOpen(false);
            }
        };

        if (isOpen) {
            document.addEventListener("mousedown", handleClickOutside);
            document.body.style.overflow = "hidden";
        } else {
            document.body.style.overflow = "auto";
        }

        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
            document.body.style.overflow = "auto";
        };
    }, [isOpen]);
    return (
        <>
            {/* Toggle Button */}
            <button
                onClick={() => setIsOpen(!isOpen)}
                className="p-2 focus:outline-none block lg:hidden z-50 relative"
                aria-label="Toggle menu"
            >
                <svg
                    width="32"
                    height="32"
                    viewBox="0 0 24 24"
                    className="overflow-visible"
                >
                    {/* Top line */}
                    <path
                        d="M4 6H20"
                        stroke="currentColor"
                        strokeWidth="2"
                        strokeLinecap="round"
                        className={`transition-transform duration-300 ease-in-out 
              ${
                  isOpen
                      ? "rotate-45 translate-y-[6px]"
                      : "rotate-0 translate-y-0"
              }`}
                        style={{
                            transformOrigin: "center",
                            transformBox: "fill-box",
                        }}
                    />

                    {/* Middle line */}
                    <path
                        d="M4 12H20"
                        stroke="currentColor"
                        strokeWidth="2"
                        strokeLinecap="round"
                        className={`transition-opacity duration-200 ease-in-out ${
                            isOpen ? "opacity-0" : "opacity-100"
                        }`}
                    />

                    {/* Bottom line */}
                    <path
                        d="M4 18H20"
                        stroke="currentColor"
                        strokeWidth="2"
                        strokeLinecap="round"
                        className={`transition-transform duration-300 ease-in-out 
              ${
                  isOpen
                      ? "-rotate-45 -translate-y-[6px]"
                      : "rotate-0 translate-y-0"
              }`}
                        style={{
                            transformOrigin: "center",
                            transformBox: "fill-box",
                        }}
                    />
                </svg>
            </button>

            {/* Overlay that starts below navbar */}
            <AnimatePresence>
                {isOpen && (
                    <motion.div
                        onClick={() => setIsOpen(false)}
                        className={`fixed inset-0 z-40  bg-black block lg:hidden`}
                        style={{ top: "96px" }}
                        initial={{ opacity: 0 }}
                        animate={{ opacity: 0.5 }}
                        exit={{ opacity: 0 }}
                        transition={{ duration: 0.3 }}
                    />
                )}
            </AnimatePresence>
            {/* Slide-in Menu */}
            <AnimatePresence>
                {isOpen && (
                    <motion.div
                        ref={menuRef}
                        className="fixed block lg:hidden right-0 z-50 h-[calc(100vh-64px)] top-24 rounded-l-lg w-[80%] bg-gradient-to-br from-blue-500 via-purple-500 to-pink-500 text-white p-6 shadow-lg"
                        initial={{ x: "100%" }}
                        animate={{ x: 0 }}
                        exit={{ x: "100%" }}
                        transition={{ type: "tween", duration: 0.3 }}
                    >
                        <nav className="flex flex-col h-full">
                            <div className="space-y-6 flex-1">
                                {navbarLinks.map((link) => (
                                    <a
                                        key={link.name}
                                        href={link.href}
                                        className="block text-xl font-semibold py-3 hover:bg-white hover:text-black hover:bg-opacity-20 px-4 rounded-lg transition-all duration-200"
                                        onClick={() => setIsOpen(false)}
                                    >
                                        {link.name}
                                    </a>
                                ))}
                            </div>

                            {/* Optional footer/bottom content */}
                            <div className="flex items-center justify-between pb-6">
                                <h1 className="text-3xl font-extrabold bg-gradient-to-r from-white via-white to-white bg-clip-text text-transparent drop-shadow-sm">
                                    Jobbotic
                                </h1>
                                <div className="flex items-center gap-2">
                                    <FaInstagram />
                                    <FaXTwitter />
                                </div>
                            </div>
                        </nav>
                    </motion.div>
                )}
            </AnimatePresence>
        </>
    );
};

export default MobileNav;
