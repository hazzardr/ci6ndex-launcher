/// <reference types="vitest" />
import path from "node:path";
import { defineConfig } from "vite";
import Vue from "@vitejs/plugin-vue";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import VueMacros from "unplugin-vue-macros/vite";

export default defineConfig({
  resolve: {
    alias: {
      "~/": `${path.resolve(__dirname, "src")}/`,
    },
  },
  plugins: [
    VueMacros({
      defineOptions: false,
      defineModels: false,
      plugins: {
        vue: Vue({
          script: {
            propsDestructure: true,
            defineModel: true,
          },
        }),
      },
    }),

    AutoImport({
      imports: ["vue", "@vueuse/core"],
      dts: true,
      dirs: ["./src/composables"],
      vueTemplate: true,
    }),

    Components({
      dts: true,
    }),
  ],

  test: {
    environment: "jsdom",
  },
});
