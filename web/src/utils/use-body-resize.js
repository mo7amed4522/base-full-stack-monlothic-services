import { watch } from 'vue';
import { storeToRefs } from 'pinia';
import { useThemeStore } from '@/store/modules/theme';

/**
 * Body area size change hook
 * @param hook Change callback
 */
export function useBodyResize(hook) {
  if (!hook) {
    return;
  }
  const themeStore = useThemeStore();
  const { contentWidth } = storeToRefs(themeStore);

  watch(contentWidth, (value, oldValue) => {
    if (value != null && oldValue != null) {
      hook();
    }
  });
}
