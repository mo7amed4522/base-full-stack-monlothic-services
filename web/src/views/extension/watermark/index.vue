<template>
  <ele-page>
    <ele-card header="Watermark" style="z-index: 1; position: relative">
      <ele-watermark content="Ele Admin Plus" :gap="[60, 40]">
        <div style="height: 200px">
          <div>
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
          </div>
          <div style="margin-top: 22px">
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
          </div>
          <div style="margin-top: 22px">
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
          </div>
        </div>
      </ele-watermark>
    </ele-card>
    <ele-card header="Multiple lines of watermark" style="z-index: 1; position: relative">
      <ele-watermark :content="['Ele Admin Plus', 'Happy Working']" :gap="[60, 40]">
        <div style="height: 200px">
          <div>
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
          </div>
          <div style="margin-top: 22px">
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
          </div>
          <div style="margin-top: 22px">
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
          </div>
        </div>
      </ele-watermark>
    </ele-card>
    <ele-card header="Image watermark" style="z-index: 1; position: relative">
      <ele-watermark :height="30" :width="130"
        image="https://mdn.alipayobjects.com/huamei_7uahnr/afts/img/A*lkAoRbywo0oAAAAAAAAAAAAADrJ8AQ/original"
        :gap="[60, 40]" :custom-style="darkMode ? { filter: 'invert(1) !important' } : void 0">
        <div style="height: 200px">
          <div>
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
          </div>
          <div style="margin-top: 22px">
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
          </div>
          <div style="margin-top: 22px">
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
            <el-button>I am the content</el-button>
          </div>
        </div>
      </ele-watermark>
    </ele-card>
    <ele-card header="Combined with other components" style="z-index: 1; position: relative">
      <ele-alert show-icon :closable="false"
        title="Watermark component also implements the function of preventing deletion and tampering of styles, when the node is deleted or the default style is modified, the watermark will be regenerated again"
        style="margin-bottom: 12px" />
      <ele-watermark content="Ele Admin Plus" :gap="[60, 40]">
        <ele-pro-table row-key="id" :columns="columns" :datasource="datasource" :pagination="false" :toolbar="false"
          :cell-style="tableCellStyle">
          <template #money="{ row }">
            <el-input v-model="row.money" style="margin: 4px 0" />
          </template>
          <template #status="{ row }">
            <ele-text v-if="row.status === 0" type="success">
              In progress
            </ele-text>
            <ele-text v-else-if="row.status === 1" type="danger">
              Delayed
            </ele-text>
            <ele-text v-else-if="row.status === 2" type="warning">
              Not started
            </ele-text>
            <ele-text v-else-if="row.status === 3" type="info">
              Ended
            </ele-text>
          </template>
          <template #progress="{ row }">
            <el-progress :percentage="row.progress" />
          </template>
        </ele-pro-table>
      </ele-watermark>
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref } from 'vue';
import { getProjectList } from '@/api/example';
import { useThemeStore } from '@/store/modules/theme';
import { storeToRefs } from 'pinia';

/** 是否是暗黑主题 */
const themeStore = useThemeStore();
const { darkMode } = storeToRefs(themeStore);

/** 表格列配置 */
const columns = ref([
  {
    type: 'index',
    columnKey: 'index',
    width: 50,
    align: 'center'
  },
  {
    label: 'Project name',
    prop: 'projectName',
    align: 'center',
    minWidth: 140
  },
  {
    label: 'Amount (万元)',
    prop: 'money',
    slot: 'money',
    align: 'center',
    minWidth: 120
  },
  {
    label: 'Start time',
    prop: 'startDate',
    align: 'center',
    minWidth: 120
  },
  {
    label: 'End time',
    prop: 'endDate',
    align: 'center',
    minWidth: 120
  },
  {
    label: 'Status',
    prop: 'status',
    slot: 'status',
    align: 'center',
    width: 100
  },
  {
    label: 'Progress',
    prop: 'progress',
    slot: 'progress',
    align: 'center',
    minWidth: 140
  }
]);

/** 表格数据源 */
const datasource = () => {
  return getProjectList();
};

/** 表格单元格样式 */
const tableCellStyle = ({ column }) => {
  if (column.columnKey === 'money') {
    return { padding: '0' };
  }
};
</script>

<script>
export default {
  name: 'ExtensionWatermark'
};
</script>
