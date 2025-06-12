import React from "react";

interface DangerAlertProps {
    message: string;
}

const DangerAlert: React.FC<DangerAlertProps> = ({ message }) => {
    return (
        <div
            className="flex items-center p-4 mb-4 text-sm text-red-800 border border-red-300 rounded-lg bg-red-50"
            role="alert"
        >
            <svg
                className="flex-shrink-0 inline w-5 h-5 mr-3"
                fill="currentColor"
                viewBox="0 0 20 20"
            >
                <path
                    fillRule="evenodd"
                    d="M8.257 3.099c.765-1.36 2.721-1.36 3.486 0l6.518 11.59c.75 1.335-.213 2.986-1.743 2.986H3.482c-1.53 0-2.493-1.651-1.743-2.986L8.257 3.1zM9 7a1 1 0 012 0v4a1 1 0 01-2 0V7zm1 8a1.5 1.5 0 100-3 1.5 1.5 0 000 3z"
                    clipRule="evenodd"
                />
            </svg>
            <span>{message}</span>
        </div>
    );
};

export default DangerAlert;
