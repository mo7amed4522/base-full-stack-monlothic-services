// vite.config.js
import { defineConfig } from "file:///app/node_modules/.pnpm/vite@5.4.19_sass@1.91.0/node_modules/vite/dist/node/index.js";
import vue from "file:///app/node_modules/.pnpm/@vitejs+plugin-vue@5.2.4_vite@5.4.19_sass@1.91.0__vue@3.5.20/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import { resolve } from "path";
import Compression from "file:///app/node_modules/.pnpm/vite-plugin-compression@0.5.1_vite@5.4.19_sass@1.91.0_/node_modules/vite-plugin-compression/dist/index.mjs";
import Components from "file:///app/node_modules/.pnpm/unplugin-vue-components@0.27.5_@babel+parser@7.28.3_rollup@4.48.1_vue@3.5.20/node_modules/unplugin-vue-components/dist/vite.js";
import { ElementPlusResolver } from "file:///app/node_modules/.pnpm/unplugin-vue-components@0.27.5_@babel+parser@7.28.3_rollup@4.48.1_vue@3.5.20/node_modules/unplugin-vue-components/dist/resolvers.js";
import { EleAdminResolver } from "file:///app/node_modules/.pnpm/ele-admin-plus@1.4.1_element-plus@2.11.1_vue@3.5.20__vue@3.5.20/node_modules/ele-admin-plus/es/utils/resolvers.js";
var vite_config_default = defineConfig(({ command, mode }) => {
  const isBuild = command === "build";
  const alias = {
    "@/": resolve("src") + "/",
    "vue-i18n": "vue-i18n/dist/vue-i18n.cjs.js"
  };
  const plugins = [vue()];
  if (isBuild) {
    plugins.push(
      Components({
        dts: false,
        resolvers: [
          ElementPlusResolver({
            importStyle: "sass"
          }),
          EleAdminResolver({
            importStyle: "sass"
          })
        ]
      })
    );
    plugins.push(
      Compression({
        disable: !isBuild,
        threshold: 10240,
        algorithm: "gzip",
        ext: ".gz"
      })
    );
  } else {
    alias["./as-needed"] = "./global-import";
  }
  const serverConfig = mode === "development" ? {
    port: 5180,
    proxy: {
      "/api": {
        target: "http://172.232.96.159:13340/api/",
        // target: 'https://newpayn.indpy.com/api/',
        // target: 'https://knympy.com/api/',
        changeOrigin: true,
        // rewrite: (path) => path.replace(/^\/api/, '')
        rewrite: (path) => path.replace(new RegExp("^/api"), "")
      }
    }
  } : {};
  return {
    resolve: { alias },
    plugins,
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: `@use "@/styles/variables.scss" as *;`
        }
      }
    },
    optimizeDeps: {
      include: ["echarts/core", "echarts/charts", "echarts/renderers", "echarts/components", "vue-echarts"]
    },
    build: {
      target: "chrome63",
      chunkSizeWarningLimit: 2e3
    },
    server: serverConfig
  };
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcuanMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCIvYXBwXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ZpbGVuYW1lID0gXCIvYXBwL3ZpdGUuY29uZmlnLmpzXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ltcG9ydF9tZXRhX3VybCA9IFwiZmlsZTovLy9hcHAvdml0ZS5jb25maWcuanNcIjtpbXBvcnQgeyBkZWZpbmVDb25maWcgfSBmcm9tICd2aXRlJztcbmltcG9ydCB2dWUgZnJvbSAnQHZpdGVqcy9wbHVnaW4tdnVlJztcbmltcG9ydCB7IHJlc29sdmUgfSBmcm9tICdwYXRoJztcbmltcG9ydCBDb21wcmVzc2lvbiBmcm9tICd2aXRlLXBsdWdpbi1jb21wcmVzc2lvbic7XG5pbXBvcnQgQ29tcG9uZW50cyBmcm9tICd1bnBsdWdpbi12dWUtY29tcG9uZW50cy92aXRlJztcbmltcG9ydCB7IEVsZW1lbnRQbHVzUmVzb2x2ZXIgfSBmcm9tICd1bnBsdWdpbi12dWUtY29tcG9uZW50cy9yZXNvbHZlcnMnO1xuaW1wb3J0IHsgRWxlQWRtaW5SZXNvbHZlciB9IGZyb20gJ2VsZS1hZG1pbi1wbHVzL2VzL3V0aWxzL3Jlc29sdmVycyc7XG5cbmV4cG9ydCBkZWZhdWx0IGRlZmluZUNvbmZpZygoeyBjb21tYW5kLCBtb2RlIH0pID0+IHtcbiAgY29uc3QgaXNCdWlsZCA9IGNvbW1hbmQgPT09ICdidWlsZCc7XG4gIGNvbnN0IGFsaWFzID0ge1xuICAgICdALyc6IHJlc29sdmUoJ3NyYycpICsgJy8nLFxuICAgICd2dWUtaTE4bic6ICd2dWUtaTE4bi9kaXN0L3Z1ZS1pMThuLmNqcy5qcydcbiAgfTtcbiAgY29uc3QgcGx1Z2lucyA9IFt2dWUoKV07XG5cbiAgaWYgKGlzQnVpbGQpIHtcbiAgICAvLyBcdTdFQzRcdTRFRjZcdTYzMDlcdTk3MDBcdTVGMTVcdTUxNjVcbiAgICBwbHVnaW5zLnB1c2goXG4gICAgICBDb21wb25lbnRzKHtcbiAgICAgICAgZHRzOiBmYWxzZSxcbiAgICAgICAgcmVzb2x2ZXJzOiBbXG4gICAgICAgICAgRWxlbWVudFBsdXNSZXNvbHZlcih7XG4gICAgICAgICAgICBpbXBvcnRTdHlsZTogJ3Nhc3MnXG4gICAgICAgICAgfSksXG4gICAgICAgICAgRWxlQWRtaW5SZXNvbHZlcih7XG4gICAgICAgICAgICBpbXBvcnRTdHlsZTogJ3Nhc3MnXG4gICAgICAgICAgfSlcbiAgICAgICAgXVxuICAgICAgfSlcbiAgICApO1xuXG4gICAgLy8gZ3ppcFx1NTM4Qlx1N0YyOVxuICAgIHBsdWdpbnMucHVzaChcbiAgICAgIENvbXByZXNzaW9uKHtcbiAgICAgICAgZGlzYWJsZTogIWlzQnVpbGQsXG4gICAgICAgIHRocmVzaG9sZDogMTAyNDAsXG4gICAgICAgIGFsZ29yaXRobTogJ2d6aXAnLFxuICAgICAgICBleHQ6ICcuZ3onXG4gICAgICB9KVxuICAgICk7XG4gIH0gZWxzZSB7XG4gICAgLy8gXHU1RjAwXHU1M0QxXHU3M0FGXHU1ODgzXHU1MTY4XHU1QzQwXHU1Qjg5XHU4OEM1XG4gICAgYWxpYXNbJy4vYXMtbmVlZGVkJ10gPSAnLi9nbG9iYWwtaW1wb3J0JztcbiAgfVxuXG4gIC8vIFx1NzUyOFx1NEU4RVx1NUYwMFx1NTNEMVx1NzNBRlx1NTg4M1x1NEVFM1x1NzQwNlx1OTE0RFx1N0Y2RVxuICBjb25zdCBzZXJ2ZXJDb25maWcgPVxuICAgIG1vZGUgPT09ICdkZXZlbG9wbWVudCdcbiAgICAgID8ge1xuICAgICAgICAgIHBvcnQ6IDUxODAsXG4gICAgICAgICAgcHJveHk6IHtcbiAgICAgICAgICAgICcvYXBpJzoge1xuICAgICAgICAgICAgICB0YXJnZXQ6ICdodHRwOi8vMTcyLjIzMi45Ni4xNTk6MTMzNDAvYXBpLycsXG4gICAgICAgICAgICAgIC8vIHRhcmdldDogJ2h0dHBzOi8vbmV3cGF5bi5pbmRweS5jb20vYXBpLycsXG4gICAgICAgICAgICAgIC8vIHRhcmdldDogJ2h0dHBzOi8va255bXB5LmNvbS9hcGkvJyxcbiAgICAgICAgICAgICAgY2hhbmdlT3JpZ2luOiB0cnVlLFxuICAgICAgICAgICAgICAvLyByZXdyaXRlOiAocGF0aCkgPT4gcGF0aC5yZXBsYWNlKC9eXFwvYXBpLywgJycpXG4gICAgICAgICAgICAgIHJld3JpdGU6IChwYXRoKSA9PiBwYXRoLnJlcGxhY2UobmV3IFJlZ0V4cCgnXicgKyAnL2FwaScpLCAnJylcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIDoge307XG5cbiAgcmV0dXJuIHtcbiAgICByZXNvbHZlOiB7IGFsaWFzIH0sXG4gICAgcGx1Z2lucyxcbiAgICBjc3M6IHtcbiAgICAgIHByZXByb2Nlc3Nvck9wdGlvbnM6IHtcbiAgICAgICAgc2Nzczoge1xuICAgICAgICAgIGFkZGl0aW9uYWxEYXRhOiBgQHVzZSBcIkAvc3R5bGVzL3ZhcmlhYmxlcy5zY3NzXCIgYXMgKjtgXG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9LFxuICAgIG9wdGltaXplRGVwczoge1xuICAgICAgaW5jbHVkZTogWydlY2hhcnRzL2NvcmUnLCAnZWNoYXJ0cy9jaGFydHMnLCAnZWNoYXJ0cy9yZW5kZXJlcnMnLCAnZWNoYXJ0cy9jb21wb25lbnRzJywgJ3Z1ZS1lY2hhcnRzJ11cbiAgICB9LFxuICAgIGJ1aWxkOiB7XG4gICAgICB0YXJnZXQ6ICdjaHJvbWU2MycsXG4gICAgICBjaHVua1NpemVXYXJuaW5nTGltaXQ6IDIwMDBcbiAgICB9LFxuICAgIHNlcnZlcjogc2VydmVyQ29uZmlnXG4gIH07XG59KTtcbiJdLAogICJtYXBwaW5ncyI6ICI7QUFBOEwsU0FBUyxvQkFBb0I7QUFDM04sT0FBTyxTQUFTO0FBQ2hCLFNBQVMsZUFBZTtBQUN4QixPQUFPLGlCQUFpQjtBQUN4QixPQUFPLGdCQUFnQjtBQUN2QixTQUFTLDJCQUEyQjtBQUNwQyxTQUFTLHdCQUF3QjtBQUVqQyxJQUFPLHNCQUFRLGFBQWEsQ0FBQyxFQUFFLFNBQVMsS0FBSyxNQUFNO0FBQ2pELFFBQU0sVUFBVSxZQUFZO0FBQzVCLFFBQU0sUUFBUTtBQUFBLElBQ1osTUFBTSxRQUFRLEtBQUssSUFBSTtBQUFBLElBQ3ZCLFlBQVk7QUFBQSxFQUNkO0FBQ0EsUUFBTSxVQUFVLENBQUMsSUFBSSxDQUFDO0FBRXRCLE1BQUksU0FBUztBQUVYLFlBQVE7QUFBQSxNQUNOLFdBQVc7QUFBQSxRQUNULEtBQUs7QUFBQSxRQUNMLFdBQVc7QUFBQSxVQUNULG9CQUFvQjtBQUFBLFlBQ2xCLGFBQWE7QUFBQSxVQUNmLENBQUM7QUFBQSxVQUNELGlCQUFpQjtBQUFBLFlBQ2YsYUFBYTtBQUFBLFVBQ2YsQ0FBQztBQUFBLFFBQ0g7QUFBQSxNQUNGLENBQUM7QUFBQSxJQUNIO0FBR0EsWUFBUTtBQUFBLE1BQ04sWUFBWTtBQUFBLFFBQ1YsU0FBUyxDQUFDO0FBQUEsUUFDVixXQUFXO0FBQUEsUUFDWCxXQUFXO0FBQUEsUUFDWCxLQUFLO0FBQUEsTUFDUCxDQUFDO0FBQUEsSUFDSDtBQUFBLEVBQ0YsT0FBTztBQUVMLFVBQU0sYUFBYSxJQUFJO0FBQUEsRUFDekI7QUFHQSxRQUFNLGVBQ0osU0FBUyxnQkFDTDtBQUFBLElBQ0UsTUFBTTtBQUFBLElBQ04sT0FBTztBQUFBLE1BQ0wsUUFBUTtBQUFBLFFBQ04sUUFBUTtBQUFBO0FBQUE7QUFBQSxRQUdSLGNBQWM7QUFBQTtBQUFBLFFBRWQsU0FBUyxDQUFDLFNBQVMsS0FBSyxRQUFRLElBQUksT0FBTyxPQUFZLEdBQUcsRUFBRTtBQUFBLE1BQzlEO0FBQUEsSUFDRjtBQUFBLEVBQ0YsSUFDQSxDQUFDO0FBRVAsU0FBTztBQUFBLElBQ0wsU0FBUyxFQUFFLE1BQU07QUFBQSxJQUNqQjtBQUFBLElBQ0EsS0FBSztBQUFBLE1BQ0gscUJBQXFCO0FBQUEsUUFDbkIsTUFBTTtBQUFBLFVBQ0osZ0JBQWdCO0FBQUEsUUFDbEI7QUFBQSxNQUNGO0FBQUEsSUFDRjtBQUFBLElBQ0EsY0FBYztBQUFBLE1BQ1osU0FBUyxDQUFDLGdCQUFnQixrQkFBa0IscUJBQXFCLHNCQUFzQixhQUFhO0FBQUEsSUFDdEc7QUFBQSxJQUNBLE9BQU87QUFBQSxNQUNMLFFBQVE7QUFBQSxNQUNSLHVCQUF1QjtBQUFBLElBQ3pCO0FBQUEsSUFDQSxRQUFRO0FBQUEsRUFDVjtBQUNGLENBQUM7IiwKICAibmFtZXMiOiBbXQp9Cg==
