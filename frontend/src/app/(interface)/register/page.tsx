"use client";
import RegisterForm from "@/components/forms/RegisterForm";
import { motion } from "framer-motion";
import FeaturesSection from "@/components/features/FeaturesCollapsible";

const page = () => {
    return (
        <section className="min-h-screen w-full flex ">
            {/* Right: Form */}
            <div className="w-full md:w-[60%] flex flex-col justify-center items-center px-6 py-12 bg-white">
                <motion.div
                    initial={{ opacity: 0, y: 10 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ duration: 0.5 }}
                    className="max-w-md w-full"
                >
                    <h1 className="text-3xl font-bold mb-2 text-gray-800">
                        Create your account
                    </h1>
                    <p className="text-gray-500 mb-6">
                        Join{" "}
                        <span className="text-blue-600 font-semibold">
                            Jobbotic
                        </span>{" "}
                        today and simplify your job application journey.
                    </p>

                    <RegisterForm />
                </motion.div>
            </div>
            {/* Left: Illustration / Image */}
            <div className="hidden md:flex w-1/2  items-center justify-center bg-gray-100/80">
                <motion.div
                    initial={{ opacity: 0, x: -30 }}
                    animate={{ opacity: 1, x: 0 }}
                    transition={{ duration: 0.7 }}
                    className="p-4"
                >
                    <p className="text-2xl font-bold">
                        Applications tracker for productive people
                    </p>
                    <p className="w-96 text-gray-600">
                        Jobbotic helps you track and manage your job
                        applications automatically — all in one clean dashboard.
                    </p>
                    <FeaturesSection />
                </motion.div>
            </div>
        </section>
    );
};

export default page;
