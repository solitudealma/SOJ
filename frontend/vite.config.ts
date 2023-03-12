import { UserConfigExport, ConfigEnv, loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve } from "path";
import Icons from "unplugin-icons/vite";
import IconsResolver from "unplugin-icons/resolver";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";

function pathResolve(dir: string) {
  return resolve(process.cwd(), ".", dir);
}

// https://vitejs.dev/config/
export default ({ command, mode }: ConfigEnv): UserConfigExport => {
  const env = loadEnv(mode, process.cwd());

  return {
    plugins: [
      vue(),
      AutoImport({
        imports: ["vue"],
        resolvers: [
          ElementPlusResolver(),
          IconsResolver({
            prefix: "Icon",
          }),
        ],
        dts: resolve(pathResolve("src") + "/", "auto-imports.d.ts"),
      }),
      Components({
        resolvers: [
          IconsResolver({
            enabledCollections: ["ep"],
          }),
          ElementPlusResolver(),
        ],
        dts: resolve(pathResolve("src") + "/", "components.d.ts"),
      }),
      Icons({
        autoInstall: true,
      }),
    ],
    resolve: {
      alias: [
        {
          find: "@",
          replacement: pathResolve("src") + "/",
        },
        {
          find: "#",
          replacement: pathResolve("types") + "/",
        },
      ],
    },
    server: {
      host: "0.0.0.0",
      hmr: true,
      port: 8000,
      proxy: {
        "/api/v1/user": {
          target: "http://127.0.0.1:8888",
          changeOrigin: true,
          rewrite: (path) => path.replace("/^\/api/v1", ""),
        },
        "/api/v1/problem": {
          target: "http://127.0.0.1:8889",
          changeOrigin: true,
          rewrite: (path) => path.replace("/^\/api/v1", ""),
        },
      },
    },
  };
};
