<template>
  <div>
    <ele-pro-table ref="tableRef" row-key="userId" :columns="columns" :datasource="datasource"
      :default-sort="{ prop: 'createTime', order: 'descending' }" :show-overflow-tooltip="true"
      :export-config="{ fileName: 'User data', datasource: exportSource }" :print-config="{ datasource: exportSource }">
      <template #toolbar>
        <el-button type="primary" class="ele-btn-icon" @click="setSorter">
          Set sorter
        </el-button>
        <el-button type="primary" class="ele-btn-icon" @click="resetAll">
          Reset sorter and filter
        </el-button>
      </template>
    </ele-pro-table>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { pageUsers, listUsers } from '@/api/system/user';

/** Table Instance */
const tableRef = ref(null);

/** Table Columns */
const columns = ref([
  {
    type: 'index',
    columnKey: 'index',
    width: 50,
    align: 'center'
  },
  {
    prop: 'username',
    label: 'User account',
    sortable: 'custom',
    minWidth: 140
  },
  {
    prop: 'nickname',
    label: 'Username',
    sortable: 'custom',
    minWidth: 140
  },
  {
    prop: 'sexName',
    label: 'Sex',
    sortable: 'custom',
    width: 110,
    align: 'center',
    filterMultiple: false,
    filters: [
      {
        text: 'Male',
        value: 'male'
      },
      {
        text: 'Female',
        value: 'female'
      }
    ],
    filteredValue: ['male']
  },
  {
    prop: 'createTime',
    label: 'Create time',
    sortable: 'custom',
    minWidth: 180,
    align: 'center'
  }
]);

/** Table Data Source */
const datasource = ({ pages, orders, filters }) => {
  return pageUsers({ ...orders, ...filters, ...pages });
};

/** Set Sorter */
const setSorter = () => {
  tableRef.value?.sort?.('username', 'ascending');
};

/** Reset Sorter and Filter */
const resetAll = () => {
  tableRef.value?.clearSort?.();
  tableRef.value?.clearFilter?.();
  tableRef.value?.reload?.({ sorter: {}, filter: {} });
};

/** Export and Print All Data Source */
const exportSource = ({ orders, filters }) => {
  return listUsers({ ...orders, ...filters });
};
</script>
