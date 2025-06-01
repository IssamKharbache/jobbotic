"use client";
import { motion } from "framer-motion";

const JobboticLogo = () => {
    // More fluid wave paths
    const wavePaths = [
        "M0,8 C25,0 75,16 100,8 C125,0 175,16 200,8",
        "M0,8 C25,16 75,0 100,8 C125,16 175,0 200,8",
        "M0,8 C20,4 80,12 100,8 C120,4 180,12 200,8",
    ];

    return (
        <div className="">
            {/* Text Logo */}
            <h1 className="text-3xl font-extrabold bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500 bg-clip-text text-transparent drop-shadow-sm">
                Jobbotic
            </h1>

            {/* Animated Wavy Underline */}
            <div className="relative h-6 w-full">
                <motion.svg
                    width="100%"
                    height="36"
                    viewBox="0 0 200 16"
                    preserveAspectRatio="none"
                    className="absolute bottom-0 left-0"
                >
                    <motion.path
                        initial={{ d: wavePaths[0] }}
                        animate={{
                            d: wavePaths,
                        }}
                        transition={{
                            duration: 6,
                            repeat: Infinity,
                            repeatType: "reverse",
                            ease: "easeInOut",
                        }}
                        stroke="url(#waveGradient)"
                        strokeWidth="2.5"
                        fill="none"
                        strokeLinecap="round"
                    />
                    <defs>
                        <linearGradient
                            id="waveGradient"
                            x1="0%"
                            y1="0%"
                            x2="100%"
                            y2="0%"
                        >
                            <stop offset="0%" stopColor="#3B82F6" />
                            <stop offset="50%" stopColor="#8B5CF6" />
                            <stop offset="100%" stopColor="#EC4899" />
                        </linearGradient>
                    </defs>
                </motion.svg>
            </div>
        </div>
    );
};

export default JobboticLogo;
