import Link from "next/link";
import JobboticLogo from "../logos/JobboticLogo";

const navbarLinks = [
    { name: "Home", href: "/" },
    { name: "Features", href: "/features" },
    { name: "Pricing", href: "/pricing" },
    { name: "How it Works", href: "/how-it-works" },
    { name: "About", href: "/about" },
    { name: "Contact", href: "/contact" },
];
const Navbar = () => {
    return (
        <nav className="flex items-center justify-between border py-5 px-7">
            {/* logo here  */}
            <JobboticLogo />
            {/* Links */}
            <div className="flex items-center gap-5">
                {navbarLinks.map((item, idx) => (
                    <div key={idx} className="relative group pb-1">
                        <Link
                            href={item.href}
                            className="relative inline-block"
                        >
                            {item.name}
                            <svg
                                className="absolute left-0 -bottom-2.5 w-full h-5 opacity-0 group-hover:opacity-100 transition-opacity duration-300"
                                viewBox="0 0 200 30"
                                preserveAspectRatio="none"
                            >
                                <path
                                    d="M0,20 C25,8 50,25 75,12 C100,25 125,8 150,20 C160,25 175,15 200,20"
                                    stroke="currentColor"
                                    fill="none"
                                    strokeWidth="4"
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                    strokeDasharray="250"
                                    strokeDashoffset="250"
                                    className="group-hover:animate-[smooth-draw_1.5s_cubic-bezier(0.16,0.8,0.3,1)_forwards]"
                                />
                            </svg>
                        </Link>
                    </div>
                ))}
            </div>{" "}
            {/* Buttons (Get started and sign in)  */}
            <div className="flex items-center gap-6">
                <button className="border border-black hover:bg-black hover:text-white duration-300 rounded-full py-3 px-6 cursor-pointer">
                    Get started
                </button>
                <Link href="/">Sign in</Link>
            </div>
        </nav>
    );
};

export default Navbar;
