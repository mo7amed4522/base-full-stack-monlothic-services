<template>
  <div>
    <ele-pro-table row-key="dictId" :columns="columns" :datasource="datasource" :show-overflow-tooltip="true"
      :export-config="{ fileName: 'Dictionary data' }">
      <template #toolbar>
        <el-button type="primary" @click="getChildSelections" class="ele-btn-icon">
          Get all selected
        </el-button>
      </template>
      <template #action="{ row }">
        <el-link type="primary" :underline="false" @click="editDict(row)">
          Edit
        </el-link>
        <el-divider direction="vertical" />
        <el-link type="danger" :underline="false" @click="removeDict(row)">
          Delete
        </el-link>
      </template>
      <template #expand="{ row }">
        <child-table :dictId="row.dictId" v-model:selections="childSelections[row.dictId]"
          @setDictionaryData="onSetDictionaryData" />
      </template>
      <template #expandPrint="{ row }">
        <child-print-table :data="dictionaryDatas[row.dictId]" />
      </template>
    </ele-pro-table>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { pageDictionaries } from '@/api/system/dictionary';
import ChildTable from './child-table.vue';
import ChildPrintTable from './child-print-table.vue';

/** 表格列配置 */
const columns = ref([
  {
    type: 'expand',
    columnKey: 'expand',
    width: 50,
    align: 'center',
    slot: 'expand',
    fixed: 'left',
    // 用于打印
    printSlot: 'expandPrint',
    // 用于导出
    formatter: (row) => {
      const rows = [
        ['#', 'Dictionary data name', 'Dictionary data code', 'Sorting number', 'Creation time']
      ];
      dictionaryDatas[row.dictId]?.forEach?.((d, i) => {
        rows.push([
          i + 1,
          d.dictDataName ?? '',
          d.dictDataCode ?? '',
          d.sortNumber ?? '',
          d.createTime ?? ''
        ]);
      });
      return rows.map((cells) => cells.join(',')).join('\n');
    }
  },
  {
    type: 'index',
    columnKey: 'index',
    width: 50,
    align: 'center',
    fixed: 'left'
  },
  {
    prop: 'dictName',
    label: 'Dictionary name',
    sortable: 'custom',
    minWidth: 130
  },
  {
    prop: 'dictCode',
    label: 'Dictionary code',
    sortable: 'custom',
    minWidth: 160
  },
  {
    prop: 'sortNumber',
    label: 'Sorting number',
    sortable: 'custom',
    width: 120,
    align: 'center'
  },
  {
    prop: 'createTime',
    label: 'Creation time',
    sortable: 'custom',
    minWidth: 180,
    align: 'center'
  },
  {
    columnKey: 'action',
    label: 'Operation',
    width: 120,
    align: 'center',
    slot: 'action',
    fixed: 'right',
    hideInPrint: true,
    hideInExport: true
  }
]);

/** 表格数据源 */
const datasource = ({ pages, orders }) => {
  return pageDictionaries({ ...orders, ...pages });
};

/** 删除字典 */
const removeDict = (row) => {
  EleMessage.error('Delete: ' + row.dictName);
};

/** 修改字典 */
const editDict = (row) => {
  EleMessage.info({ message: 'Delete: ' + row.dictName, plain: true });
};

/** 子表格多选选中数据 */
const childSelections = reactive({});

/** 获取子表格多选选中数据 */
const getChildSelections = () => {
  console.log(JSON.parse(JSON.stringify(childSelections)));
  const length = Object.keys(childSelections)
    .map((key) => childSelections[key].length)
    .reduce((a, b) => a + b, 0);
  EleMessage.success(`Total ${length} data selected, printed in console`);
};

/** 存储已加载的子表格数据 */
const dictionaryDatas = reactive({});

const onSetDictionaryData = (dictId, data) => {
  dictionaryDatas[dictId] = data;
};
</script>
