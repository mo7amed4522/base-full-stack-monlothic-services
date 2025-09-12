<template>
  <ele-page>
    <ele-card header="Basic Usage">
      <ele-avatar-group :data="data" :max-count="4" :overflow-popover="false" :item-style="{ borderWidth: '2px' }" />
    </ele-card>
    <ele-card header="Exceed Display">
      <ele-avatar-group :data="data" :max-count="4" :item-style="{ borderWidth: '2px' }" />
    </ele-card>
    <ele-card header="Mouse Hover to Expand">
      <ele-avatar-group hover-open :data="data" :item-style="{ borderWidth: '2px' }" />
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref } from 'vue';

/** 数据 */
const data = ref([
  {
    label: 'SunSmile',
    value:
      'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg'
  },
  {
    label: 'Your Name is Nice',
    value:
      'https://cdn.eleadmin.com/20200609/b6a811873e704db49db994053a5019b2.jpg'
  },
  {
    label: 'The Hope of the Entire Village',
    value:
      'https://cdn.eleadmin.com/20200609/948344a2a77c47a7a7b332fe12ff749a.jpg'
  },
  {
    label: 'Jasmine',
    value:
      'https://cdn.eleadmin.com/20200609/f6bc05af944a4f738b54128717952107.jpg'
  },
  {
    label: 'Cool uncle',
    value:
      'https://cdn.eleadmin.com/20200609/2d98970a51b34b6b859339c96b240dcd.jpg'
  },
  {
    label: 'administrator',
    value: 'https://cdn.eleadmin.com/20200610/avatar.jpg'
  }
]);
</script>

<script>
export default {
  name: 'ExtensionAvatar'
};
</script>
