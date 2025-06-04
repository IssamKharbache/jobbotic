import {
    Accordion,
    AccordionContent,
    AccordionItem,
    AccordionTrigger,
} from "@/components/ui/accordion";
import { ChevronDown } from "lucide-react";

const jobboticFeatures = [
    {
        title: "Auto-Track Applications from Gmail",
        description:
            "Connect your Gmail account and Jobbotic will automatically detect job applications from your sent emails. No manual entry needed.",
    },
    {
        title: "Application Insights",
        description:
            "View useful stats like your application success rate, top companies you apply to, and your most active times.",
    },
    {
        title: "Smart Duplicate Detection",
        description:
            "Jobbotic warns you when you're about to apply to the same job again, helping you avoid redundancy and stay efficient.",
    },
    {
        title: "Stay Organized with Status Tags",
        description:
            "Tag applications as Applied, Interviewed, Offer, or Rejected. Get automatic reminders for follow-ups.",
    },
    {
        title: "Centralized Dashboard",
        description:
            "See all your job applications in one clean interface. Filter by company, date, or status to keep track effortlessly.",
    },
];

const FeaturesSection = () => {
    return (
        <section className="mt-10 space-y-4 max-w-2xl mx-auto">
            <Accordion type="single" collapsible>
                {jobboticFeatures.map((feature, idx) => (
                    <AccordionItem key={idx} value={`item-${idx}`}>
                        <div className="max-w-xl mx-auto rounded-lg overflow-hidden">
                            <AccordionTrigger className="flex justify-between items-center w-full p-4 font-semibold text-lg cursor-pointer bg-gray-100 hover:bg-gray-200 duration-200">
                                {feature.title}
                            </AccordionTrigger>
                            <AccordionContent className="p-4 text-gray-700 bg-gray-50">
                                <p>{feature.description}</p>
                            </AccordionContent>
                        </div>
                    </AccordionItem>
                ))}
            </Accordion>
        </section>
    );
};
export default FeaturesSection;
