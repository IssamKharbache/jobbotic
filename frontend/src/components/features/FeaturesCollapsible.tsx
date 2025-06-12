"use client";
import { useState } from "react";
import {
    ChevronDown,
    ChevronRight,
    CheckCircle,
    BarChart3,
    Bell,
} from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";

const features = [
    {
        id: 1,
        title: "Application Tracking",
        description:
            "Keep track of all your job applications in one organized dashboard.",
        icon: CheckCircle,
        details: [
            "Track application status",
            "Set follow-up reminders",
            "Store company information",
            "Upload documents",
        ],
    },
    {
        id: 2,
        title: "Analytics & Insights",
        description: "Get detailed insights about your job search progress.",
        icon: BarChart3,
        details: [
            "Application success rates",
            "Response time analytics",
            "Industry trends",
            "Performance metrics",
        ],
    },
    {
        id: 3,
        title: "Smart Notifications",
        description: "Never miss important deadlines or follow-ups.",
        icon: Bell,
        details: [
            "Application deadlines",
            "Interview reminders",
            "Follow-up alerts",
            "Custom notifications",
        ],
    },
];

const FeaturesSection = () => {
    const [expandedFeature, setExpandedFeature] = useState<number | null>(1);

    const toggleFeature = (featureId: number) => {
        setExpandedFeature(expandedFeature === featureId ? null : featureId);
    };

    return (
        <div className="mt-8 space-y-4">
            <h3 className="text-lg font-semibold text-gray-800 mb-4">
                Key Features
            </h3>
            {features.map((feature) => {
                const Icon = feature.icon;
                const isExpanded = expandedFeature === feature.id;

                return (
                    <div
                        key={feature.id}
                        className="border border-gray-200 rounded-lg overflow-hidden"
                    >
                        <button
                            onClick={() => toggleFeature(feature.id)}
                            className="w-full p-4 text-left hover:bg-gray-50 transition-colors duration-200 flex items-center justify-between"
                        >
                            <div className="flex items-center space-x-3">
                                <Icon className="h-5 w-5 text-blue-600" />
                                <div>
                                    <h4 className="font-medium text-gray-800">
                                        {feature.title}
                                    </h4>
                                    <p className="text-sm text-gray-600">
                                        {feature.description}
                                    </p>
                                </div>
                            </div>
                            {isExpanded ? (
                                <ChevronDown className="h-4 w-4 text-gray-400" />
                            ) : (
                                <ChevronRight className="h-4 w-4 text-gray-400" />
                            )}
                        </button>

                        <AnimatePresence>
                            {isExpanded && (
                                <motion.div
                                    initial={{ height: 0, opacity: 0 }}
                                    animate={{ height: "auto", opacity: 1 }}
                                    exit={{ height: 0, opacity: 0 }}
                                    transition={{ duration: 0.3 }}
                                    className="border-t border-gray-200 bg-gray-50"
                                >
                                    <div className="p-4">
                                        <ul className="space-y-2">
                                            {feature.details.map(
                                                (detail, index) => (
                                                    <li
                                                        key={index}
                                                        className="flex items-center space-x-2 text-sm text-gray-600"
                                                    >
                                                        <div className="w-1.5 h-1.5 bg-blue-600 rounded-full"></div>
                                                        <span>{detail}</span>
                                                    </li>
                                                ),
                                            )}
                                        </ul>
                                    </div>
                                </motion.div>
                            )}
                        </AnimatePresence>
                    </div>
                );
            })}
        </div>
    );
};

export default FeaturesSection;
