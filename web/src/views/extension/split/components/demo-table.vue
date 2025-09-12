<template>
  <ele-split-panel :flex-table="true" space="0px" size="360px" :min-size="180" :max-size="-180" :vertical="true"
    :resizable="true" :allow-collapse="true" :custom-style="{ overflow: 'hidden', border: 'none' }" :responsive="false"
    style="background: var(--el-bg-color)" class="demo-split-panel">
    <ele-pro-table ref="tableRef" row-key="organizationId" :columns="columns" :datasource="datasource"
      v-model:selections="selections" v-model:current="current" :highlight-current-row="true" :border="true"
      :toolbar="{ theme: 'default', style: { borderRadius: 0 } }" :footer-style="{
        borderWidth: '0 1px 0 1px',
        borderStyle: 'solid',
        borderColor: 'var(--el-border-color-light)',
        marginTop: '-0.45px'
      }" :pagination="{ style: { padding: '0 0 8px 0' } }" :show-overflow-tooltip="true" :table-style="{
        '--ele-table-radius': 0,
        '--ele-table-sm-radius': 0,
        '--ele-table-lg-radius': 0
      }" :export-config="{ fileName: 'Organization data' }" @done="onDone">
      <template #toolbar>
        <el-button size="small" type="primary" @click="back">Back</el-button>
      </template>
      <template #action>
        <el-link type="primary" :underline="false">Modify</el-link>
        <el-divider direction="vertical" />
        <el-link type="danger" :underline="false">Delete</el-link>
      </template>
    </ele-pro-table>
    <template #body>
      <ele-pro-table ref="table2Ref" row-key="userId" :columns="columns2" :datasource="datasource2"
        v-model:selections="selections2" :border="true" :toolbar="{ theme: 'default', style: { borderRadius: 0 } }"
        :footer-style="{
          borderWidth: '0 1px 1px 1px',
          borderStyle: 'solid',
          borderColor: 'var(--el-border-color-light)',
          marginTop: '-0.45px'
        }" :pagination="{ style: { padding: '0 0 6px 0' } }" :show-overflow-tooltip="true" :load-on-created="false"
        :table-style="{
          '--ele-table-radius': 0,
          '--ele-table-sm-radius': 0,
          '--ele-table-lg-radius': 0
        }" :export-config="{ fileName: 'User data' }">
        <template #toolbar>
          <span>Upper and lower tables will adapt to height</span>
        </template>
        <template #action>
          <el-link type="primary" :underline="false">Modify</el-link>
          <el-divider direction="vertical" />
          <el-link type="danger" :underline="false">Delete</el-link>
        </template>
      </ele-pro-table>
    </template>
  </ele-split-panel>
</template>

<script setup>
import { ref, nextTick, watch } from 'vue';
import { pageOrganizations } from '@/api/system/organization';
import { pageUsers } from '@/api/system/user';

const emit = defineEmits(['change']);

/** Back */
const back = () => {
  emit('change', 0);
};

/** Table Example */
const tableRef = ref(null);

/** Table Column Configuration */
const columns = ref([
  {
    type: 'selection',
    columnKey: 'selection',
    width: 50,
    align: 'center',
    fixed: 'left'
  },
  {
    type: 'index',
    columnKey: 'index',
    width: 50,
    align: 'center',
    fixed: 'left'
  },
  {
    prop: 'organizationName',
    label: 'Organization Name',
    sortable: 'custom',
    align: 'center',
    minWidth: 140
  },
  {
    prop: 'organizationTypeName',
    label: 'Organization Type',
    sortable: 'custom',
    align: 'center',
    minWidth: 140
  },
  {
    prop: 'sortNumber',
    label: 'Sort Number',
    sortable: 'custom',
    align: 'center',
    width: 100
  },
  {
    prop: 'createTime',
    label: 'Create Time',
    sortable: 'custom',
    align: 'center',
    minWidth: 200
  },
  {
    columnKey: 'action',
    label: 'Action',
    width: 128,
    align: 'center',
    resizable: false,
    slot: 'action',
    hideInPrint: true,
    hideInExport: true
  }
]);

/** Table Example Selected Data */
const selections = ref([]);

/** Table Example Single Selection Selected Data */
const current = ref(null);

/** Table Example Data Source */
const datasource = ({ pages, where, orders }) => {
  return pageOrganizations({ ...where, ...orders, ...pages });
};

/** Table Example Data Request Complete Event */
const onDone = ({ data }) => {
  if (data?.length) {
    nextTick(() => {
      tableRef.value?.setCurrentRow?.(data[0]);
    });
  }
};

/** Refresh Table 2 Data */
watch(current, (row) => {
  table2Ref.value?.reload?.({
    where: { organizationId: row.organizationId }
  });
});

/** Table Example 2 */
const table2Ref = ref(null);

/** Table Example 2 Column Configuration */
const columns2 = ref([
  {
    type: 'selection',
    columnKey: 'selection',
    width: 50,
    align: 'center',
    fixed: 'left'
  },
  {
    type: 'index',
    columnKey: 'index',
    width: 50,
    align: 'center',
    fixed: 'left'
  },
  {
    prop: 'username',
    label: 'Username',
    sortable: 'custom',
    align: 'center',
    minWidth: 140
  },
  {
    prop: 'nickname',
    label: 'Nickname',
    sortable: 'custom',
    align: 'center',
    minWidth: 140
  },
  {
    prop: 'sexName',
    label: 'Sex',
    sortable: 'custom',
    align: 'center',
    width: 100
  },
  {
    prop: 'createTime',
    label: 'Create Time',
    sortable: 'custom',
    align: 'center',
    minWidth: 200
  },
  {
    columnKey: 'action',
    label: 'Action',
    width: 128,
    align: 'center',
    resizable: false,
    slot: 'action',
    hideInPrint: true,
    hideInExport: true
  }
]);

/** Table Example 2 Selected Data */
const selections2 = ref([]);

/** Table Example 2 Data Source */
const datasource2 = ({ pages, where, orders }) => {
  return pageUsers({ ...where, ...orders, ...pages });
};
</script>

<style lang="scss" scoped>
.demo-split-panel.is-collapse :deep(.ele-split-collapse-button) {
  margin-top: 0;
}
</style>
