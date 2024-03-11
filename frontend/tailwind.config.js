/** @type {import('tailwindcss').Config} */
export default {
    content: ["./src/**/*.{html,js,svelte,ts}"],
    theme: {
        extend: {
            keyframes: {
                "zoom": {
                    "0%": {
                        transform: "scale(0.99)",
                        opacity: "0",
                    },
                    "100%": {
                        transform: "scale(1)",
                        opacity: "1",
                    },
                },
                "fade": {
                    "0%": {
                        opacity: "0",
                    },
                    "100%": {
                        opacity: "1",
                    },
                },
            },
            animation: {
                "zoom": "zoom 0.2s ease-out",
                "fade": "fade 0.2s ease-out",
            },
        },
    },
    plugins: [],
};
