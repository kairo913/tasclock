/** @type {import('tailwindcss').Config} */
export default {
    content: ["./src/**/*.{html,js,svelte,ts}"],
    theme: {
        extend: {
            keyframes: {
                zoom: {
                    "0%": "transform: scale(0.95)",
                    "100%": "transform: scale(1)",
                },
                fade: {
                    "0%": "opacity: 0",
                    "100%": "opacity: 1",
                },
            },
            animation: {
                zoom: "zoom 0.3s cubic-bezier(0.34, 1.56, 0.64, 1)",
                fade: "fade 0.2s ease-out",
            },
        },
    },
    plugins: [],
};
