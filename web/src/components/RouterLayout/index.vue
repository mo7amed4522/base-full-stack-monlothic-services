<!-- Router outlet -->
<template>
  <router-view v-slot="{ route, Component }">
    <transition :name="transitionName" mode="out-in" appear>
      <div>
        <!-- Wrap a div to keep a single root node -->
        <keep-alive v-if="TAB_KEEP_ALIVE && tabBar" :include="keepAliveInclude" :max="10">
          <component :key="route.path" :is="Component" />
        </keep-alive>
        <component v-else :key="route.path" :is="Component" />
      </div>
    </transition>
  </router-view>
</template>

<script setup>
import { storeToRefs } from 'pinia';
import { useThemeStore } from '@/store/modules/theme';
import { TAB_KEEP_ALIVE } from '@/config/setting';

const themeStore = useThemeStore();
const { keepAliveInclude, transitionName, tabBar } = storeToRefs(themeStore);
</script>

<script>
export default {
  name: 'RouterLayout'
};
</script>
