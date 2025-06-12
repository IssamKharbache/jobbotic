import { Loader2Icon } from "lucide-react";
import { Button } from "../ui/button";
interface ButtonProps {
    loading: boolean;
    text: string;
}
const GradientButton = ({ loading = false, text = "" }: ButtonProps) => {
    return (
        <Button
            type="submit"
            className="w-full bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500 text-white font-medium hover:scale-105 duration-300 cursor-pointer"
        >
            {loading ? <Loader2Icon className="animate-spin size-6" /> : text}
        </Button>
    );
};

export default GradientButton;
