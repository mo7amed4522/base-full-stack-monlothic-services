import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { resolve } from 'path';
import Compression from 'vite-plugin-compression';
import Components from 'unplugin-vue-components/vite';
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers';
import { EleAdminResolver } from 'ele-admin-plus/es/utils/resolvers';

export default defineConfig(({ command, mode }) => {
  const isBuild = command === 'build';
  const alias = {
    '@/': resolve('src') + '/',
    'vue-i18n': 'vue-i18n/dist/vue-i18n.cjs.js'
  };
  const plugins = [vue()];

  if (isBuild) {
    // Component introduction on demand
    plugins.push(
      Components({
        dts: false,
        resolvers: [
          ElementPlusResolver({
            importStyle: 'sass'
          }),
          EleAdminResolver({
            importStyle: 'sass'
          })
        ]
      })
    );

    // gzip compression
    plugins.push(
      Compression({
        disable: !isBuild,
        threshold: 10240,
        algorithm: 'gzip',
        ext: '.gz'
      })
    );
  } else {
    // Component introduction on demand in development environment
    alias['./as-needed'] = './global-import';
  }

  // Development environment proxy configuration
  const serverConfig =
    mode === 'development'
      ? {
          port: 3040,
          proxy: {
            '/api': {
              target: 'http://localhost:8888',
              changeOrigin: true,
              rewrite: (path) => path.replace(/^\/api/, '')
            }
          }
        }
      : {};

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
      include: ['echarts/core', 'echarts/charts', 'echarts/renderers', 'echarts/components', 'vue-echarts']
    },
    build: {
      target: 'chrome63',
      chunkSizeWarningLimit: 2000
    },
    server: serverConfig
  };
});
