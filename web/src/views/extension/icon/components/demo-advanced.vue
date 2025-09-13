<template>
  <ele-card header="Icon advanced usage">
    <div style="max-width: 260px">
      <ele-icon-select clearable filterable :data="icons" v-model="selectedIcon" :hide-on-single-tab="true"
        placeholder="Please select" :popper-options="{ strategy: 'fixed' }">
        <template #icon="{ icon, prefix }">
          <img :src="`/ele-file-list/ic_file_${icon}_sm.png`" :style="prefix
              ? { maxWidth: '22px', maxHeight: '22px' }
              : { maxWidth: '100%', maxHeight: '100%' }
            " />
        </template>
      </ele-icon-select>
    </div>
  </ele-card>
</template>

<script setup>
import { ref } from 'vue';

/** Selected value */
const selectedIcon = ref('');

/** Icon data */
const icons = ref([
  {
    title: '',
    children: [
      {
        title: 'Document',
        icons: [
          'folder',
          'word',
          'excel',
          'ppt',
          'pdf',
          'visio',
          'text',
          'code',
          'htm',
          'fonts',
          'zip',
          'ps',
          'cad',
          'bt',
          'misc'
        ]
      },
      {
        title: 'Media',
        icons: ['flash', 'music', 'video', 'picture']
      },
      {
        title: 'Application',
        icons: ['android', 'exe', 'ipa']
      }
    ]
  }
]);
</script>
