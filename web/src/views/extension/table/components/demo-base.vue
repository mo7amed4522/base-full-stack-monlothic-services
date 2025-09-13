<template>
  <div>
    <div class="option-wrapper">
      <option-item label="Toolbar background" :responsive="false">
        <el-switch v-model="toolbarBg" size="small" />
      </option-item>
      <option-item label="Border" :responsive="false">
        <el-switch v-model="border" size="small" @change="onConfigChange" />
      </option-item>
      <option-item label="Stripe" :responsive="false">
        <el-switch v-model="stripe" size="small" @change="onConfigChange" />
      </option-item>
      <option-item label="Header" :responsive="false">
        <el-switch v-model="showHeader" size="small" />
      </option-item>
      <option-item label="Table Footer" :responsive="false">
        <el-switch v-model="showSummary" size="small" />
      </option-item>
      <option-item label="Row click highlight" :responsive="false">
        <el-switch v-model="highlightCurrentRow" size="small" />
      </option-item>
      <option-item label="Row click check" :responsive="false">
        <el-switch v-model="rowClickChecked" size="small" />
      </option-item>
      <option-item label="No count" :responsive="false">
        <el-switch v-model="noCount" size="small" @change="reload()" />
      </option-item>
    </div>
    <ele-pro-table ref="tableRef" row-key="pieceId" :columns="columns" :datasource="datasource"
      :show-overflow-tooltip="true" :highlight-current-row="highlightCurrentRow" :row-click-checked="rowClickChecked"
      v-model:selections="selections" v-model:current="current" :pagination="{
        pageSize: 10,
        pageSizes: [10, 20, 40]
      }" :toolbar="{ theme: toolbarBg ? 'default' : 'plain' }" :border="border" :stripe="stripe"
      :show-header="showHeader" :show-summary="showSummary" :summary-method="getSummaries"
      :export-config="{ fileName: 'Piece data', datasource: exportSource }"
      :print-config="{ datasource: exportSource }">
      <template #toolbar>
        <el-button type="primary" class="ele-btn-icon" @click="getSelections">
          Get multiple selected rows
        </el-button>
        <el-button v-if="highlightCurrentRow" type="primary" class="ele-btn-icon" @click="getCurrent">
          Get single selected row
        </el-button>
      </template>
      <template #action="{ row }">
        <el-link type="primary" :underline="false" @click="openEdit(row)">
          Edit
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
import { ref, nextTick } from 'vue';
import { EleMessage, random } from 'ele-admin-plus/es';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';

const data = Array.from({ length: 39 }).map((_, i) => {
  const pieceId = i + 1;
  const no = String(pieceId).padStart(3, '0');
  const secrets = ['Confidential', 'Secret', 'Public', 'Top Secret', 'Internal'];
  const retentions = ['Permanent', '10 years', '20 years'];
  const retentionNos = ['Y', 'D10', 'D20'];
  const retentionRandom = random(0, retentions.length);
  const carriers = ['Paper', 'CD'];
  const year = '2020';
  const item = {
    pieceId,
    title: `Teaching archive${no}`,
    pieceNo: `ELE101-${retentionNos[retentionRandom]}-ZM•JX-${year}-${no}`,
    secret: secrets[random(0, secrets.length)],
    location: `001-001-${Math.random() < 0.5 ? 'A' : 'B'}0${random(1, 7)}`,
    type: 'Teaching organization',
    retention: retentions[retentionRandom],
    carrier: carriers[random(0, carriers.length)],
    year,
    amount: random(2, 10)
  };
  return item;
});

/** Table Example */
const tableRef = ref(null);

/** Table Columns */
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
    prop: 'pieceNo',
    label: 'Piece No',
    minWidth: 260
  },
  {
    prop: 'title',
    label: 'Title',
    minWidth: 160
  },
  {
    prop: 'retention',
    label: 'Retention',
    minWidth: 110,
    align: 'center',
    filters: [
      { text: 'Permanent', value: 'Permanent' },
      { text: '10 years', value: '10 years' },
      { text: 'Regular 20 years', value: 'Regular 20 years' }
    ],
    filterMultiple: false
  },
  {
    prop: 'amount',
    label: 'Amount',
    minWidth: 110,
    align: 'center'
  },
  {
    prop: 'secret',
    label: 'Secret',
    minWidth: 110,
    align: 'center',
    filters: [
      { text: 'Confidential', value: 'Confidential' },
      { text: 'Secret', value: 'Secret' },
      { text: 'Public', value: 'Public' },
      { text: 'Top Secret', value: 'Top Secret' },
      { text: 'Internal', value: 'Internal' }
    ],
    filterMultiple: true
  },
  {
    prop: 'year',
    label: 'Year',
    minWidth: 110,
    align: 'center'
  },
  {
    prop: 'type',
    label: 'Type',
    minWidth: 140,
    align: 'center'
  },
  {
    prop: 'location',
    label: 'Location',
    minWidth: 140,
    align: 'center'
  },
  {
    prop: 'carrier',
    label: 'Carrier',
    minWidth: 110,
    align: 'center'
  },
  {
    columnKey: 'action',
    label: 'Action',
    width: 160,
    align: 'center',
    slot: 'action',
    fixed: 'right',
    hideInPrint: true,
    hideInExport: true
  }
]);

/** Table Data Source */
const datasource = async ({ pages, filter }) => {
  await new Promise((resolve) => setTimeout(resolve, 100));
  const secretFilter = filter.secret?.length ? filter.secret : null;
  const retentFilter = filter.retention?.length ? filter.retention : null;
  const list = data.filter((d) => {
    if (secretFilter != null) {
      if (!d.secret || !secretFilter.includes(d.secret)) {
        return false;
      }
    }
    if (retentFilter != null) {
      if (!d.retention || !retentFilter.includes(d.retention)) {
        return false;
      }
    }
    return true;
  });
  const start = ((pages.page || 1) - 1) * (pages.limit || 10);
  const end = start + (pages.limit || 10);
  return {
    count: noCount.value ? '*' : list.length,
    list: list.slice(start, end > list.length ? list.length : end)
  };
};

/** Table Multi Select */
const selections = ref([]);

/** Table Single Select */
const current = ref(null);

/** More Config */
const toolbarBg = ref(false);
const border = ref(false);
const stripe = ref(false);
const showHeader = ref(true);
const showSummary = ref(true);
const highlightCurrentRow = ref(false);
const rowClickChecked = ref(false);
const noCount = ref(false);

/** Open Edit Dialog */
const openEdit = (row) => {
  EleMessage.success('Edit:' + row.pieceNo);
};

/** 删除单个 */
const remove = (row) => {
  EleMessage.error('Delete:' + row.pieceNo);
};

/** Table Summary */
const getSummaries = ({ columns, data }) => {
  const sums = [];
  const labelIndex = columns[0]?.type === 'selection' ? 2 : 1;
  columns.forEach((column, index) => {
    if (column.property === 'amount') {
      sums[index] = data
        .map((item) => Number(item[column.property]))
        .reduce((prev, curr) => {
          const value = Number(curr);
          if (!isNaN(value)) {
            return prev + curr;
          } else {
            return prev;
          }
        }, 0);
    } else if (index === labelIndex) {
      sums[index] = 'Total';
    }
  });
  return sums;
};

/** Get Multi Select Data */
const getSelections = () => {
  console.log(JSON.parse(JSON.stringify(selections.value)));
  EleMessage.success(
    '共 ' + selections.value.length + ' The data has been printed on the console'
  );
};

/** 获取选中行数据 */
const getCurrent = () => {
  if (!current.value) {
    EleMessage.success('No data selected');
  } else {
    console.log(JSON.parse(JSON.stringify(current.value)));
    EleMessage.success(current.value.pieceNo + ' The data has been printed on the console');
  }
};

/** Search */
const reload = () => {
  selections.value = [];
  tableRef.value?.reload?.({ page: 1 });
};

/** Table Config Change */
const onConfigChange = () => {
  nextTick(() => {
    tableRef.value && tableRef.value.doLayout();
  });
};

/** Export and Print All Data Source */
const exportSource = async ({ filter }) => {
  const res = await datasource({
    filter,
    pages: { page: 1, limit: data.length }
  });
  return res.list;
};
</script>

<style lang="scss" scoped>
.option-wrapper {
  display: flex;
  align-items: flex-start;
  flex-wrap: wrap;
  padding-top: 12px;

  :deep(.option-item) {
    margin: 0 20px 8px 0;

    .option-item-label+.option-item-body {
      padding-left: 4px;
    }
  }
}
</style>
