<template>
  <ele-card header="Excel export" :body-style="{ padding: '18px 20px' }">
    <ele-pro-table row-key="key" :columns="columns" :datasource="data" :show-overflow-tooltip="true" :pagination="false"
      v-model:selections="selections" :tools="['export', 'print', 'size', 'columns', 'maximized']"
      :toolbar="{ theme: 'default' }" :border="true">
      <template #toolbar>
        <el-button type="primary" class="ele-btn-icon" @click="exportBas">
          Export
        </el-button>
        <el-button type="primary" class="ele-btn-icon" @click="exportAdv">
          Export with merged headers
        </el-button>
        <el-button type="primary" class="ele-btn-icon" @click="exportSel">
          Export selected
        </el-button>
      </template>
    </ele-pro-table>
  </ele-card>
</template>

<script setup>
import { ref } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import ExcelJS from 'exceljs';
import { download } from '@/utils/common';

/** table data */
const data = ref([
  {
    key: 1,
    username: 'Zhang San',
    amount: 18,
    province: 'Zhejiang',
    city: 'Hangzhou',
    zone: 'Xihu Qu',
    street: 'Xihu Street',
    address: 'Xihu Garden 30 Building 1 Unit'
  },
  {
    key: 2,
    username: 'Li Si',
    amount: 39,
    province: 'Jiangsu',
    city: 'Suzhou',
    zone: 'Gugong Qu',
    street: 'Silk Road',
    address: 'Tianzhuang City 9 Building 2 Unit'
  },
  {
    key: 3,
    username: 'Wang Wu',
    amount: 8,
    province: 'Jiangxi',
    city: 'Nanchang',
    zone: 'Qingshan Qu',
    street: 'Aixi Lake Office',
    address: 'Zhongtexuan 1 Building 3 Unit'
  },
  {
    key: 4,
    username: 'Zhao Liu',
    amount: 16,
    province: 'Fujian',
    city: 'Quanzhou',
    zone: 'Fengze Qu',
    street: 'Nanyang Street',
    address: 'Nanyang Village 6 Building 1 Unit'
  },
  {
    key: 5,
    username: 'Sun Qiao',
    amount: 12,
    province: 'Hubei',
    city: 'Wuhan',
    zone: 'Wuchang Qu',
    street: 'Wuchang Road',
    address: 'Luwu Garden 16 Building 2 Unit'
  },
  {
    key: 6,
    username: 'Zhou Ba',
    amount: 11,
    province: 'Anhui',
    city: 'Huangshan',
    zone: 'Huangshan Qu',
    street: 'Tongtou Township',
    address: 'Wenshu Village 21'
  }
]);

/** 表格列配置 */
const columns = ref([
  {
    type: 'selection',
    columnKey: 'selection',
    width: 56,
    align: 'center'
  },
  {
    type: 'index',
    columnKey: 'index',
    width: 56,
    align: 'center'
  },
  {
    prop: 'username',
    label: 'Username',
    align: 'center',
    minWidth: 110
  },
  {
    columnKey: 'cityAddress',
    label: 'Address',
    align: 'center',
    children: [
      {
        prop: 'province',
        label: 'Province',
        align: 'center',
        minWidth: 100
      },
      {
        prop: 'city',
        label: 'City',
        align: 'center',
        minWidth: 100
      },
      {
        prop: 'zone',
        label: 'Zone',
        align: 'center',
        minWidth: 120
      },
      {
        prop: 'street',
        label: 'Street',
        align: 'center',
        minWidth: 120
      },
      {
        prop: 'address',
        label: 'Address',
        align: 'center',
        minWidth: 160
      }
    ]
  },
  {
    prop: 'amount',
    label: 'Amount',
    align: 'center',
    minWidth: 80
  }
]);

/** Select data */
const selections = ref([]);

/** Export to Excel */
const exportBas = () => {
  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet('Sheet1');
  sheet.addRow(['Username', 'Province', 'City', 'District', 'Street', 'Detailed Address', 'Amount']);
  data.value.forEach((d) => {
    sheet.addRow([
      d.username,
      d.province,
      d.city,
      d.zone,
      d.street,
      d.address,
      d.amount
    ]);
  });
  // Set column width
  [16, 16, 16, 16, 20, 30, 12].forEach((width, index) => {
    sheet.getColumn(index + 1).width = width;
  });
  // Set style
  sheet.eachRow({ includeEmpty: true }, (row, rowIndex) => {
    row.height = 20;
    row.eachCell({ includeEmpty: true }, (cell) => {
      cell.border = {
        top: { style: 'thin' },
        left: { style: 'thin' },
        bottom: { style: 'thin' },
        right: { style: 'thin' }
      };
      cell.alignment = {
        vertical: 'middle',
        horizontal: 'center'
      };
      cell.font = { size: 12, bold: rowIndex === 1 };
    });
  });
  // Download file
  workbook.xlsx.writeBuffer().then((data) => {
    download(data, 'UserData.xlsx');
  });
};

/** Export with cell merging */
const exportAdv = () => {
  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet('Sheet1');
  sheet.addRow(['Username', 'Address', null, null, null, null, 'Amount']);
  sheet.addRow([null, 'Province', 'City', 'District', 'Street', 'Detailed Address', null]);
  data.value.forEach((d) => {
    sheet.addRow([
      d.username,
      d.province,
      d.city,
      d.zone,
      d.street,
      d.address,
      d.amount
    ]);
  });
  // Set column width
  [16, 16, 16, 16, 20, 30, 12].forEach((width, index) => {
    sheet.getColumn(index + 1).width = width;
  });
  // Set style
  sheet.eachRow({ includeEmpty: true }, (row, rowIndex) => {
    row.height = 20;
    row.eachCell({ includeEmpty: true }, (cell) => {
      cell.border = {
        top: { style: 'thin' },
        left: { style: 'thin' },
        bottom: { style: 'thin' },
        right: { style: 'thin' }
      };
      cell.alignment = {
        vertical: 'middle',
        horizontal: 'center'
      };
      cell.font = { size: 12, bold: rowIndex < 3 };
    });
  });
  // Merge cells
  sheet.mergeCells(1, 2, 1, 6);
  sheet.mergeCells(1, 1, 2, 1);
  sheet.mergeCells(1, 7, 2, 7);
  // Download file
  workbook.xlsx.writeBuffer().then((data) => {
    download(data, 'UserData.xlsx');
  });
};

/** Export selected data */
const exportSel = () => {
  if (selections.value.length === 0) {
    EleMessage.error('Please select at least one data');
    return;
  }
  //
  const workbook = new ExcelJS.Workbook();
  const sheet = workbook.addWorksheet('Sheet1');
  sheet.addRow(['Username', 'Province', 'City', 'District', 'Street', 'Detailed Address', 'Amount']);
  selections.value.forEach((d) => {
    sheet.addRow([
      d.username,
      d.province,
      d.city,
      d.zone,
      d.street,
      d.address,
      d.amount
    ]);
  });
  // Set column width
  [16, 16, 16, 16, 20, 30, 12].forEach((width, index) => {
    sheet.getColumn(index + 1).width = width;
  });
  // Set style
  sheet.eachRow({ includeEmpty: true }, (row, rowIndex) => {
    row.height = 20;
    row.eachCell({ includeEmpty: true }, (cell) => {
      cell.border = {
        top: { style: 'thin' },
        left: { style: 'thin' },
        bottom: { style: 'thin' },
        right: { style: 'thin' }
      };
      cell.alignment = {
        vertical: 'middle',
        horizontal: 'center'
      };
      cell.font = { size: 12, bold: rowIndex === 1 };
    });
  });
  // Download File
  workbook.xlsx.writeBuffer().then((data) => {
    download(data, 'SelectedUserData.xlsx');
  });
};
</script>
