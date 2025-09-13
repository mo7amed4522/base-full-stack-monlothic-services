<!-- Route refresh page -->
<template>
  <div></div>
</template>

<script setup>
  import { useThemeStore } from '@/store/modules/theme';
  import { storeToRefs } from 'pinia';
  import { useRoute, useRouter } from 'vue-router';

  const { replace } = useRouter();
  const { params, query } = useRoute();
  const themeStore = useThemeStore();
  const { tabs } = storeToRefs(themeStore);

  const from = Array.isArray(params.path) ? params.path.join('/') : params.path;

  setTimeout(() => {
    tabs.value.forEach((t) => {
      if (t.refresh) {
        themeStore.tabSetItem({ key: t.key, refresh: false });
      }
    });
    replace({ path: '/' + from, query });
  }, 100);
</script>

<script>
  export default {
    name: 'RedirectLayout'
  };
</script>
