"use client";
import { motion } from "framer-motion";

const JobboticLogo = () => {
    // Wave path variations
    const wavePaths = [
        "M0,8 C30,0 60,16 90,8 C120,0 150,16 180,8 C190,6 200,12 200,8",
        "M0,8 C30,16 60,0 90,8 C120,16 150,0 180,8 C190,12 200,4 200,8",
        "M0,8 C15,2 45,14 75,8 C105,2 135,14 165,8 C175,6 200,12 200,8",
    ];

    return (
        <div className="inline-block">
            {/* Text Logo */}
            <h1 className="text-5xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
                Jobbotic
            </h1>

            {/* Animated Wavy Underline */}
            <div className="relative h-5 w-[100%]">
                <motion.svg
                    width="100%"
                    height="30"
                    viewBox="0 0 200 16"
                    preserveAspectRatio="none"
                    className="absolute -bottom-1"
                >
                    <motion.path
                        initial={{ d: wavePaths[0] }}
                        animate={{
                            d: wavePaths,
                        }}
                        transition={{
                            duration: 15,
                            repeat: Infinity,
                            repeatType: "reverse",
                            ease: "easeInOut",
                        }}
                        stroke="url(#waveGradient)"
                        strokeWidth="1.5"
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
