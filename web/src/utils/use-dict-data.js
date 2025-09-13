import { computed } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { storeToRefs } from 'pinia';
import { useUserStore } from '@/store/modules/user';
import { listDictionaryData } from '@/api/system/dictionary-data';

/**
 * Get dictionary data hook
 * @param codes Dictionary code
 */
export function useDictData(codes) {
  const result = [];

  // Cached dictionaries
  const userStore = useUserStore();
  const { dicts } = storeToRefs(userStore);

  codes.forEach((code) => {
    result.push(computed(() => dicts.value[code] || []));
    // If it has not been cached yet, get the dictionary data
    if (dicts.value[code] != null) {
      return;
    }
    userStore.setDicts([], code);
    listDictionaryData({ dictCode: code })
      .then((list) => {
        userStore.setDicts(list, code);
      })
      .catch((e) => {
        EleMessage.error(e.message);
      });
  });

  return result;
}
