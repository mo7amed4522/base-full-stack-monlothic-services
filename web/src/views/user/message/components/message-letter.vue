<template>
  <div>
    <ele-pro-table ref="tableRef" row-key="id" :columns="columns" :datasource="datasource" :show-overflow-tooltip="true"
      v-model:selections="selections" :highlight-current-row="true"
      :export-config="{ fileName: 'Message letter data' }">
      <template #toolbar>
        <el-button type="primary" class="ele-btn-icon" @click="readBatch">
          Mark as read
        </el-button>
        <el-button type="danger" class="ele-btn-icon" @click="remove()">
          Delete message
        </el-button>
      </template>
      <template #status="{ row }">
        <ele-text v-if="row.status === 0" type="warning">Unread</ele-text>
        <ele-text v-else-if="row.status === 1" type="info">Read</ele-text>
      </template>
      <template #action="{ row }">
        <el-link type="primary" :underline="false" @click="reply(row)">
          Reply
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
import { pageLetters } from '@/api/example';

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
    label: 'Message content',
    minWidth: 220
  },
  {
    prop: 'time',
    label: 'Send time',
    width: 180,
    align: 'center'
  },
  {
    prop: 'status',
    label: 'Status',
    width: 100,
    align: 'center',
    slot: 'status',
    formatter: (row) => (row.status == 0 ? 'Unread' : 'Read')
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
  return pageLetters({ ...where, ...orders, ...pages });
};

/** 回复 */
const reply = (row) => {
  console.log(JSON.parse(JSON.stringify(row)));
  EleMessage.success('Click reply');
};

/** 删除 */
const remove = (row) => {
  const rows = row == null ? selections.value : [row];
  if (!rows.length) {
    EleMessage.error('Please select at least one data');
    return;
  }
  EleMessage.error('Click delete');
};

/** 标记已读 */
const readBatch = () => {
  if (!selections.value.length) {
    EleMessage.error('Please select at least one data');
    return;
  }
  EleMessage.success('Mark as read');
  updateUnReadNum();
};

/** 触发更新未读数量事件 */
const updateUnReadNum = () => {
  emit('reload');
};
</script>
