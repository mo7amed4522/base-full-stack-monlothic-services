<template>
  <div>
    <div class="option-wrapper">
      <option-item label="Reserve Multi Select" :responsive="false">
        <el-switch v-model="reserveSelection" size="small" @change="onReserveSelectionChange" />
      </option-item>
      <option-item label="Reserve Single Select" :responsive="false">
        <el-switch v-model="reserveCurrent" size="small" />
      </option-item>
    </div>
    <ele-pro-table ref="tableRef" row-key="pieceId" :columns="columns" :datasource="datasource"
      :show-overflow-tooltip="true" v-model:selections="selections" v-model:current="current" highlight-current-row
      :reserveCurrent="reserveCurrent" :pagination="{
        pageSize: 10,
        pageSizes: [10, 20, 40]
      }" :export-config="{ fileName: 'Archive Data', datasource: exportSource }"
      :print-config="{ datasource: exportSource }">
      <template #toolbar>
        <el-space :size="12" wrap>
          <el-button type="primary" class="ele-btn-icon" @click="getSelections">
            Get Multi Select
          </el-button>
          <el-button type="warning" class="ele-btn-icon" @click="setSelections">
            Modify Multi Select
          </el-button>
          <el-button type="danger" class="ele-btn-icon" @click="clearSelections">
            Clear Multi Select
          </el-button>
          <el-button plain type="primary" class="ele-btn-icon" @click="getCurrent">
            Get Single Select
          </el-button>
          <el-button plain type="warning" class="ele-btn-icon" @click="setCurrent">
            Modify Single Select
          </el-button>
          <el-button plain type="danger" class="ele-btn-icon" @click="removeCurrent">
            Clear Single Select
          </el-button>
        </el-space>
      </template>
    </ele-pro-table>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { EleMessage, random } from 'ele-admin-plus/es';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';
const data = Array.from({ length: 39 }).map((_, i) => {
  const pieceId = i + 1;
  const no = String(pieceId).padStart(3, '0');
  const secrets = ['Confidential', 'Secret', 'Public', 'Top Secret', 'Internal'];
  const retentions = ['Permanent', '10-Year Retention', '20-Year Retention'];
  const retentionNos = ['Y', 'D10', 'D20'];
  const retentionRandom = random(0, retentions.length);
  const year = '2020';
  const item = {
    pieceId,
    title: `Teaching Archive${no}`,
    pieceNo: `ELE101-${retentionNos[retentionRandom]}-ZM•JX-${year}-${no}`,
    secret: secrets[random(0, secrets.length)],
    retention: retentions[retentionRandom],
    year,
    amount: random(2, 10)
  };
  return item;
});

/** Keep multiple selections when turning pages */
const reserveSelection = ref(false);

/** Keep single selection when turning pages */
const reserveCurrent = ref(false);

/** Table instance */
const tableRef = ref(null);

/** Table column configuration */
const columns = ref([
  {
    type: 'selection',
    columnKey: 'selection',
    width: 50,
    align: 'center',
    fixed: 'left',
    reserveSelection: reserveSelection.value
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
    label: 'Archive Number',
    minWidth: 260
  },
  {
    prop: 'title',
    label: 'Archive Title',
    minWidth: 160
  },
  {
    prop: 'retention',
    label: 'Retention Period',
    minWidth: 110,
    align: 'center',
    filters: [
      { text: 'Permanent', value: 'Permanent' },
      { text: '10-Year Retention', value: '10-Year Retention' },
      { text: '20-Year Retention', value: '20-Year Retention' }
    ],
    filterMultiple: false
  },
  {
    prop: 'amount',
    label: 'Number of Pieces',
    minWidth: 110,
    align: 'center'
  },
  {
    prop: 'secret',
    label: 'Secret Level',
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
  }
]);

/** Table datasource */
const datasource = async ({ pages, filter }) => {
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
    count: list.length,
    list: list.slice(start, end > list.length ? list.length : end)
  };
};

/** Table multiple selection data */
const selections = ref([]);

/** Table single selection data */
const current = ref(null);

/** Get multiple selection */
const getSelections = () => {
  console.log(JSON.parse(JSON.stringify(selections.value)));
  EleMessage.success(
    'Common ' + selections.value.length + ' data items have been printed in the console'
  );
};

/** 修改多选 */
const setSelections = () => {
  tableRef.value?.setSelectedRows?.([data[1], data[2], data[5]]);
};

/** 清空多选 */
const clearSelections = () => {
  selections.value = [];
};

/** 获取单选 */
const getCurrent = () => {
  if (!current.value) {
    EleMessage.success('No data item has been selected');
  } else {
    console.log(JSON.parse(JSON.stringify(current.value)));
    EleMessage.success(current.value.pieceNo + ' The data has been printed on the console');
  }
};

/** 修改单选 */
const setCurrent = () => {
  tableRef.value?.setCurrentRow?.(data[3]);
};

/** 清空单选 */
const removeCurrent = () => {
  current.value = null;
};

/** 翻页保留多选改变 */
const onReserveSelectionChange = () => {
  columns.value[0].reserveSelection = reserveSelection.value;
};

/** 导出和打印全部数据的数据源 */
const exportSource = async ({ filter }) => {
  const res = await datasource({
    filter,
    pages: { page: 1, limit: 1000 }
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
