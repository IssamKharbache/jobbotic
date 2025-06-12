import { useLogout } from "@/app/hooks/useLogout";
import { User, useUserStore } from "../../../store/UserStore";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuTrigger,
} from "../ui/dropdown-menu";

export interface UserAvatarProps {
    user: User;
}
const avatarColors = [
    "bg-red-500",
    "bg-blue-500",
    "bg-green-500",
    "bg-yellow-500",
    "bg-purple-500",
    "bg-pink-500",
    "bg-indigo-500",
    "bg-teal-500",
];
const UserAvatar = () => {
    const { user } = useUserStore();
    const logout = useLogout();
    if (!user) return null;
    if (user.email === "") return null;
    const bgColor = getColorClassForUser(user.id);
    //logout hoook
    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <button
                    className={`flex items-center justify-center  rounded-full w-8 h-8 font-semibold ${bgColor}`}
                >
                    {getFirstInitial(user.first_name)}
                </button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="w-56 mr-4" align="start">
                <DropdownMenuLabel>Profile</DropdownMenuLabel>
                <DropdownMenuItem>My account</DropdownMenuItem>
                <DropdownMenuItem onClick={logout}>Log out</DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    );
};

export default UserAvatar;

const getFirstInitial = (first_name: string) => {
    const firstInitial = first_name.charAt(0);
    return firstInitial;
};
const getColorClassForUser = (userId: string): string => {
    let hash = 0;
    for (let i = 0; i < userId.length; i++) {
        hash = userId.charCodeAt(i) + ((hash << 5) - hash);
    }
    const index = Math.abs(hash) % avatarColors.length;
    return avatarColors[index];
};
