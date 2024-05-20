import { resolve } from "path";
import { defineConfig } from "vite";

const config = defineConfig({
  build: {
    lib: {
      entry: [resolve(__dirname, "src/htmx.js")],
      formats: ["es"],
      name: "[name]",
      fileName: "[name]",
    },
    outDir: "static",
    emptyOutDir: false,
  },
});

module.exports = config;
