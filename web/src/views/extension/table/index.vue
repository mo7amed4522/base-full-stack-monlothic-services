<template>
  <ele-page :multi-card="false">
    <ele-card :body-style="{ padding: 0, overflow: 'hidden' }">
      <ele-split-panel ref="splitRef" space="0px" size="180px" :allow-collapse="mobile"
        :custom-style="{ borderWidth: '0 1px 0 0', padding: '10px 0' }"
        :body-style="{ padding: '10px 20px 14px 16px', overflow: 'hidden' }" :collapse-style="{ marginLeft: '4px' }"
        style="min-height: 708px; border-radius: var(--ele-card-radius)">
        <ele-menus :items="items" :default-active="active" class="demo-menu" @itemClick="onMenuItemClick" />
        <template #body>
          <transition name="slide-right" mode="out-in">
            <demo-base v-if="active == 'base'" />
            <tree-table v-else-if="active == 'tree'" />
            <default-sorter v-else-if="active == 'sorter'" />
            <nested-table v-else-if="active == 'nested'" />
            <expandable-table v-else-if="active == 'expandable'" />
            <merge-cell v-else-if="active == 'merge'" />
            <editable-table v-else-if="active == 'editable'" />
            <sortable-table v-else-if="active == 'sortable'" />
            <demo-selections v-else-if="active == 'selections'" />
            <virtual-base v-else-if="active == 'virtual-base'" />
            <virtual-header v-else-if="active == 'virtual-header'" />
            <virtual-merge v-else-if="active == 'virtual-merge'" />
            <virtual-tree v-else-if="active == 'virtual-tree'" />
            <virtual-expandable v-else-if="active == 'virtual-expandable'" />
          </transition>
        </template>
      </ele-split-panel>
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref } from 'vue';
import { useMobile } from '@/utils/use-mobile';
import DemoBase from './components/demo-base.vue';
import TreeTable from './components/tree-table.vue';
import DefaultSorter from './components/default-sorter.vue';
import NestedTable from './components/nested-table.vue';
import ExpandableTable from './components/expandable-table.vue';
import MergeCell from './components/merge-cell.vue';
import EditableTable from './components/editable-table.vue';
import SortableTable from './components/sortable-table.vue';
import DemoSelections from './components/demo-selections.vue';
import VirtualBase from './components/virtual-base.vue';
import VirtualHeader from './components/virtual-header.vue';
import VirtualMerge from './components/virtual-merge.vue';
import VirtualTree from './components/virtual-tree.vue';
import VirtualExpandable from './components/virtual-expandable.vue';

/** 是否是移动端 */
const { mobile } = useMobile();

/** 分割面板组件 */
const splitRef = ref(null);

/** 选中 */
const active = ref('base');

/** 菜单 */
const items = ref([
  {
    index: 'base',
    title: 'Basic example'
  },
  {
    index: 'tree',
    title: 'Tree table lazy loading'
  },
  {
    index: 'sorter',
    title: 'Default sort filter'
  },
  {
    index: 'nested',
    title: 'Nested table'
  },
  {
    index: 'expandable',
    title: 'Expandable row'
  },
  {
    index: 'merge',
    title: 'Merge cell'
  },
  {
    index: 'editable',
    title: 'Editable table'
  },
  {
    index: 'sortable',
    title: 'Drag sort'
  },
  {
    index: 'selections',
    title: 'Single selection'
  },
  {
    index: 'virtual',
    title: 'Virtual scroll table',
    group: true,
    children: [
      {
        index: 'virtual-base',
        title: 'Basic example'
      },
      {
        index: 'virtual-header',
        title: 'Multi-level header'
      },
      {
        index: 'virtual-merge',
        title: 'Merge cell'
      },
      {
        index: 'virtual-tree',
        title: 'Tree table'
      },
      {
        index: 'virtual-expandable',
        title: 'Expandable row'
      }
    ]
  }
]);

/** 菜单点击事件 */
const onMenuItemClick = (item) => {
  active.value = item.index;
  // 移动端自动收起左侧
  if (mobile.value) {
    splitRef.value?.toggleCollapse?.(true);
  }
};
</script>

<script>
export default {
  name: 'ExtensionTable'
};
</script>

<style lang="scss" scoped>
.demo-menu {
  width: 100%;
  background: none;
  border: none;

  :deep(> .el-menu-item-group > .el-menu-item-group__title) {
    border-top: 1px solid var(--el-border-color-light);
    padding-bottom: 8px;
    padding-top: 12px;
  }
}
</style>
