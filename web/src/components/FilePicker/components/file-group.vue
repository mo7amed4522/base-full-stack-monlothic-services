<template>
  <ElTree ref="treeRef" :data="groupData" :highlightCurrent="true" nodeKey="id" :props="{ label: 'name' }"
    :expandOnClickNode="false" :defaultExpandAll="true" class="file-picker-left" @node-click="handleGroupSelect"
    @node-contextmenu="handleGroupContextmenu">
    <template #default="scope">
      <img src="/ele-file-list/ic_file_folder.png" class="file-picker-tree-icon" />
      <span class="el-tree-node__label" :title="scope.data.name">
        {{ scope.data.name }}
      </span>
      <ElIcon v-if="scope.data.id !== 0 && scope.data.id !== -1" class="file-picker-tree-more"
        @click.stop="(e) => handleMoreIconClick(e, scope.data)">
        <MoreOutlined style="transform: scale(1.1)" />
      </ElIcon>
    </template>
  </ElTree>
  <ElLink type="primary" :underline="false" class="file-picker-left-add" @click="handleAddBtnClick">
    Add Group
  </ElLink>
</template>

<script setup>
import { ref } from 'vue';
import { queryChild } from 'ele-admin-plus/es';
import { MoreOutlined } from '@/components/icons';

defineProps({
  /** File group data */
  groupData: Array
});

const emit = defineEmits([
  'groupSelect',
  'addBtnClick',
  'moreIconClick',
  'hideGroupCtxMenu'
]);

/** Tree component */
const treeRef = ref(null);

/** Group selected event */
const handleGroupSelect = (item) => {
  emit('groupSelect', item, true);
};

/** Add group button click event */
const handleAddBtnClick = () => {
  emit('addBtnClick');
};

/** More icon click event */
const handleMoreIconClick = (e, item) => {
  emit('moreIconClick', e.currentTarget, item);
};

/** Group right-click event */
const handleGroupContextmenu = (e, item) => {
  if (item.id === 0 || item.id === -1) {
    emit('hideGroupCtxMenu');
    return;
  }
  const triggerEl = queryChild(
    e.currentTarget,
    'el-tree-node__content'
  )?.querySelector?.('.file-picker-tree-more');
  // @ts-ignore
  triggerEl && triggerEl.click();
};

/** Set selected group */
const setSelectedGroup = (item) => {
  treeRef.value?.setCurrentKey?.(item.id);
};

defineExpose({ setSelectedGroup });
</script>
