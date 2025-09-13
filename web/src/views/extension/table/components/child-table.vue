<template>
  <div style="padding: 0 16px">
    <ele-pro-table row-key="dictDataId" :columns="dictDataColumns" :datasource="dictDataDatasource"
      :show-overflow-tooltip="true" :toolbar="false" :pagination="false" :empty-props="false"
      :selections="selections || []" @update:selections="updateSelections" @done="onTableDone">
      <template #action="{ row: d }">
        <el-link type="primary" :underline="false" @click="editDictData(d)">
          Edit
        </el-link>
        <el-divider direction="vertical" />
        <el-link type="danger" :underline="false" @click="removeDictData(d)">
          Delete
        </el-link>
      </template>
    </ele-pro-table>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { listDictionaryData } from '@/api/system/dictionary-data';

const props = defineProps({
  dictId: {
    type: Number,
    required: true
  },
  selections: Array
});

const emit = defineEmits(['update:selections', 'setDictionaryData']);

/** Data Column Configuration */
const dictDataColumns = ref([
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
    align: 'center'
  },
  {
    prop: 'dictDataName',
    label: 'Data Name',
    sortable: 'custom',
    minWidth: 140
  },
  {
    prop: 'dictDataCode',
    label: 'Data Code',
    sortable: 'custom',
    minWidth: 140
  },
  {
    prop: 'sortNumber',
    label: 'Sort Number',
    sortable: 'custom',
    width: 110,
    align: 'center'
  },
  {
    prop: 'createTime',
    label: 'Create Time',
    sortable: 'custom',
    minWidth: 140,
    align: 'center'
  },
  {
    columnKey: 'action',
    label: 'Action',
    width: 120,
    align: 'center',
    slot: 'action'
  }
]);

/** Data Source */
const dictDataDatasource = ({ where, orders }) => {
  return listDictionaryData({ ...where, ...orders, dictId: props.dictId });
};

/** Delete Data */
const removeDictData = (row) => {
  EleMessage.error('Delete: ' + row.dictDataName);
};

/** Edit Data */
const editDictData = (row) => {
  EleMessage.info({ message: 'Edit: ' + row.dictDataName, plain: true });
};

/** Update Selections */
const updateSelections = (value) => {
  emit('update:selections', value);
};

/** Table Data Request Complete Event */
const onTableDone = (result) => {
  emit('setDictionaryData', props.dictId, result.data);
};
</script>
