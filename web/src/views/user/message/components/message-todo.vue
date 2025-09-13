<template>
  <div>
    <ele-pro-table ref="tableRef" row-key="id" :columns="columns" :datasource="datasource" :show-overflow-tooltip="true"
      v-model:selections="selections" :highlight-current-row="true" :export-config="{ fileName: 'Todo data' }">
      <template #toolbar>
        <el-button type="primary" class="ele-btn-icon" @click="ok()">
          Batch complete
        </el-button>
        <el-button type="danger" class="ele-btn-icon" @click="remove()">
          Delete todo
        </el-button>
      </template>
      <template #status="{ row }">
        <ele-text v-if="row.status === 0" type="warning">Uncompleted</ele-text>
        <ele-text v-else-if="row.status === 1" type="info">Completed</ele-text>
      </template>
      <template #action="{ row }">
        <el-link type="primary" :underline="false" @click="ok(row)">
          Complete
        </el-link>
        <el-divider direction="vertical" />
        <el-link type="danger" :underline="false" @click="remove(row)">
          Delete
        </el-link>
      </template>
    </ele-pro-table>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useMessage } from 'ele-admin-plus/es';
import { pageTodos } from '@/api/example';

const EleMessage = useMessage({ inner: true });

const emit = defineEmits(['reload']);

/** 表格实例 */
const tableRef = ref(null);

/** 表格列配置 */
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
    prop: 'title',
    label: 'Todo content',
    minWidth: 220
  },
  {
    prop: 'time',
    label: 'End time',
    width: 180,
    align: 'center'
  },
  {
    prop: 'status',
    label: 'Status',
    width: 90,
    align: 'center',
    slot: 'status',
    formatter: (row) => (row.status == 0 ? 'Uncompleted' : 'Completed')
  },
  {
    columnKey: 'action',
    label: 'Action',
    width: 120,
    align: 'center',
    slot: 'action',
    hideInPrint: true,
    hideInExport: true
  }
]);

/** 列表选中数据 */
const selections = ref([]);

/** 表格数据源 */
const datasource = ({ pages, where, orders }) => {
  return pageTodos({ ...where, ...orders, ...pages });
};

/** 完成 */
const ok = (row) => {
  const rows = row == null ? selections.value : [row];
  if (!rows.length) {
    EleMessage.error('Please select at least one data');
    return;
  }
  EleMessage.success('Click complete');
  updateUnReadNum();
};

/** 删除单个 */
const remove = (row) => {
  const rows = row == null ? selections.value : [row];
  if (!rows.length) {
    EleMessage.error('Please select at least one data');
    return;
  }
  EleMessage.error('Click delete');
};

/** 触发更新未读数量事件 */
const updateUnReadNum = () => {
  emit('reload');
};
</script>
