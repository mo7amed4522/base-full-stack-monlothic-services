<template>
  <div>
    <ele-pro-table :height="378" :virtual="true" row-key="menuId" :columns="columns" :datasource="datasource"
      :lazy="true" :pagination="false" :show-overflow-tooltip="true"
      :export-config="{ fileName: 'Menu data', datasource: exportSource }"
      :print-config="{ datasource: exportSource }" />
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { toTree } from 'ele-admin-plus/es';
import { listMenus } from '@/api/system/menu';

/** 表格列配置 */
const columns = ref([
  {
    type: 'index',
    columnKey: 'index',
    width: 50,
    align: 'center'
  },
  {
    prop: 'title',
    label: 'Menu name',
    minWidth: 160
  },
  {
    prop: 'path',
    label: 'Route address',
    minWidth: 180
  },
  {
    prop: 'sortNumber',
    label: 'Sort',
    width: 100,
    align: 'center'
  },
  {
    prop: 'createTime',
    label: 'Create time',
    width: 220,
    align: 'center'
  }
]);

/** 表格数据源 */
const datasource = ({ where, parent }) => {
  return listMenus({ ...where, parentId: parent?.menuId || 0 });
};

/** 导出和打印全部数据的数据源 */
const exportSource = async ({ where }) => {
  const data = await listMenus({ ...where });
  return toTree({
    data,
    idField: 'menuId',
    parentIdField: 'parentId'
  });
};
</script>
