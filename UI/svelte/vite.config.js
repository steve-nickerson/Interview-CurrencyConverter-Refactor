import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

export default defineConfig({
  plugins: [
    svelte({
      compilerOptions: {
        runes: true,
        compatibility: {
          componentApi: 4,
        },
      },
    }),
  ],
  server: {
    port: 5173,
    strictPort: true,
  },
  build: {
    target: "esnext",
    minify: "esbuild",
  },
});
