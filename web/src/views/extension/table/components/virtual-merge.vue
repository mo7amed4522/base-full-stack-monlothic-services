<template>
  <div>
    <ele-pro-table :height="370" :border="true" :virtual="true" row-key="userId" :columns="columns"
      :datasource="datasource" :span-method="spanMethod" :toolbar="{ theme: 'default' }" :show-overflow-tooltip="true"
      :export-config="{ fileName: 'User score data' }" />
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { pageUserScores } from '@/api/example';

/** 表格列配置 */
const columns = ref([
  {
    type: 'index',
    columnKey: 'index',
    width: 50,
    align: 'center'
  },
  {
    columnKey: 'userName',
    prop: 'userName',
    label: 'User name'
  },
  {
    prop: 'courseName',
    label: 'Course name'
  },
  {
    prop: 'score',
    label: 'Score'
  }
]);

/** 表格数据源 */
const datasource = () => {
  return pageUserScores();
};

/** 合并表格单元格 */
const spanMethod = ({ row, column }) => {
  if (column.columnKey === 'userName') {
    return [row.userNameRowSpan, 1];
  }
  return [1, 1];
};
</script>
