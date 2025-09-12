import { ref, reactive, unref, provide, watch, onActivated, onDeactivated, nextTick } from 'vue';
import { storeToRefs } from 'pinia';
import { THEME_KEY } from 'vue-echarts';
import { ChartTheme, ChartThemeDark } from 'ele-admin-plus/es';
import { useThemeStore } from '@/store/modules/theme';
import { useBodyResize } from './use-body-resize';

/**
 * Echarts auto switch theme/reset size hook
 * @param refs Chart ref
 */
export function useEcharts(refs) {
  /** Whether it is a dark theme */
  const themeStore = useThemeStore();
  const { darkMode } = storeToRefs(themeStore);
  /** Whether it is in deactivated state */
  const deactivated = ref(false);
  /** Whether the current chart is a dark theme */
  const isDark = ref(unref(darkMode));
  /** Current chart theme */
  const chartsTheme = reactive({
    ...(unref(isDark) ? ChartThemeDark : ChartTheme)
  });

  // Set chart theme
  provide(THEME_KEY, chartsTheme);

  /** Reset chart size */
  const resizeCharts = () => {
    refs.forEach((chartRef) => {
      unref(chartRef)?.resize?.();
    });
  };

  /** Screen size change listener */
  useBodyResize(() => {
    if (!unref(deactivated)) {
      resizeCharts();
    }
  });

  /** Change chart theme */
  const changeTheme = (dark) => {
    isDark.value = dark;
    Object.assign(chartsTheme, dark ? ChartThemeDark : ChartTheme);
  };

  onActivated(() => {
    deactivated.value = false;
    nextTick(() => {
      if (unref(isDark) !== unref(darkMode)) {
        changeTheme(unref(darkMode));
      } else {
        resizeCharts();
      }
    });
  });

  onDeactivated(() => {
    deactivated.value = true;
  });

  watch(darkMode, (dark) => {
    if (!unref(deactivated)) {
      changeTheme(dark);
    }
  });
}
