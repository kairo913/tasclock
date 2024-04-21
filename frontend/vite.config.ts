import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";
import VitePluginBrowserSync from "vite-plugin-browser-sync";

export default defineConfig({
    plugins: [sveltekit(), VitePluginBrowserSync()],
    server: {
        watch: {
            usePolling: true,
            interval: 1000,
        },
    },
});
