"use client";
import { motion } from "framer-motion";
import { Button } from "@/components/ui/button";
import { ArrowRight } from "lucide-react";

export default function Hero() {
    return (
        <section className="w-full bg-white dark:bg-gray-950 py-24 px-6 md:px-12 lg:px-24">
            <div className="max-w-7xl mx-auto grid grid-cols-1 md:grid-cols-2 gap-10 items-center">
                {/* Left - Text */}
                <motion.div
                    initial={{ opacity: 0, x: -50 }}
                    animate={{ opacity: 1, x: 0 }}
                    transition={{ duration: 0.7, delay: 0.1 }}
                    className="space-y-6"
                >
                    <h1 className="text-4xl sm:text-5xl lg:text-6xl font-bold tracking-tight text-gray-900 dark:text-white">
                        Job Tracking.
                        <br />
                        Automated. Smart. Effortless.
                    </h1>

                    <motion.p
                        initial={{ opacity: 0, y: 30 }}
                        animate={{ opacity: 1, y: 0 }}
                        transition={{ duration: 0.6, delay: 0.3 }}
                        className="text-lg text-gray-600 dark:text-gray-300 max-w-xl"
                    >
                        Jobbotic connects to your Gmail, detects applications
                        you send, tracks them automatically, and makes sure you
                        never double apply or miss an opportunity again.
                    </motion.p>

                    <motion.div
                        initial={{ opacity: 0, y: 20 }}
                        animate={{ opacity: 1, y: 0 }}
                        transition={{ duration: 0.5, delay: 0.6 }}
                    >
                        <Button size="lg" className="gap-2">
                            Get Started Free
                            <ArrowRight className="w-5 h-5" />
                        </Button>
                    </motion.div>
                </motion.div>

                {/* Right - Image or Animation */}
                <motion.div
                    initial={{ opacity: 0, scale: 0.95 }}
                    animate={{ opacity: 1, scale: 1 }}
                    transition={{ duration: 0.8, delay: 0.4 }}
                    className="flex justify-center"
                >
                    <img
                        src="/dashboard-preview.png"
                        alt="Jobbotic dashboard preview"
                        className="rounded-2xl shadow-xl w-full max-w-md"
                    />
                </motion.div>
            </div>
        </section>
    );
}
