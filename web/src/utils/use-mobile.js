import { watch } from 'vue';
import { useMediaQuery } from '@vueuse/core';

/**
 * Get whether it is a mobile small screen hook
 * @param onChange Value change callback
 */
export function useMobile(onChange) {
  const mobile = useMediaQuery('(max-width: 768px)');

  watch(mobile, (isMobile) => {
    onChange && onChange(isMobile);
  });

  return { mobile };
}

/**
 * Get whether it is a mobile touch device hook
 * @param onChange Value change callback
 */
export function useMobileDevice(onChange) {
  const mobileDevice = useMediaQuery('(pointer: coarse)');

  watch(mobileDevice, (isMobile) => {
    onChange && onChange(isMobile);
  });

  return { mobileDevice };
}
