<template>
  <ele-card header="Import Excel">
    <el-space>
      <el-upload action="" accept=".xls,.xlsx" :show-upload-list="false" :before-upload="importFile">
        <el-button type="primary" class="ele-btn-icon">Import</el-button>
      </el-upload>
      <el-upload action="" accept=".xls,.xlsx" :show-upload-list="false" :before-upload="importFile2">
        <el-button type="primary" class="ele-btn-icon">
          Import Split and Merge
        </el-button>
      </el-upload>
      <el-upload action="" accept=".xls,.xlsx" :show-upload-list="false" :before-upload="importFile3">
        <el-button type="primary" class="ele-btn-icon">
          Import Keep Merge
        </el-button>
      </el-upload>
    </el-space>
    <div style="overflow: auto">
      <ele-table :border="true" style="min-width: max-content; table-layout: fixed">
        <colgroup>
          <col width="52" />
          <col v-for="item in importTitle" :key="item" />
        </colgroup>
        <thead>
          <tr>
            <th></th>
            <th v-for="item in importTitle" :key="item" style="text-align: center">
              {{ item }}
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in importData" :key="index">
            <td style="text-align: center">{{ index + 1 }}</td>
            <template v-for="key in importTitle">
              <td v-if="item[key].colspan !== 0 && item[key].rowspan !== 0" :key="key" :colspan="item[key].colspan"
                :rowspan="item[key].rowspan" style="text-align: center">
                {{ item[key].value }}
              </td>
            </template>
          </tr>
          <tr v-if="!importData.length">
            <td :colspan="importTitle.length + 1" style="text-align: center; background: none">
              No data available
            </td>
          </tr>
        </tbody>
      </ele-table>
    </div>
  </ele-card>
</template>

<script setup>
import { ref } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import ExcelJS from 'exceljs';

/** Import data column */
const importTitle = ref(['A', 'B', 'C', 'D', 'E', 'F', 'G']);

/** Import data */
const importData = ref([]);

/** Import local excel file */
const importFile = (file) => {
  if (
    ![
      'application/vnd.ms-excel',
      'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
    ].includes(file.type)
  ) {
    EleMessage.error('Only excel files are supported');
    return false;
  }
  if (file.size / 1024 / 1024 > 20) {
    EleMessage.error('Size cannot exceed 20MB');
    return false;
  }
  const loading = EleMessage.loading({ message: 'Loading...', plain: true });
  const reader = new FileReader();
  reader.onload = (e) => {
    const workbook = new ExcelJS.Workbook();
    workbook.xlsx.load(e.target?.result).then(() => {
      const title = [];
      const list = [];
      workbook.eachSheet((sheet) => {
        list.splice(0, list.length);
        sheet.eachRow({ includeEmpty: true }, (row, rowIndex) => {
          title.splice(0, title.length);
          const item = {};
          row.eachCell({ includeEmpty: true }, (cell) => {
            const key = cell.address.slice(0, -String(rowIndex).length);
            title.push(key);
            const isMerged =
              cell.isMerged && cell.address !== cell.master.address;
            item[key] = {
              value: isMerged ? '' : cell.value,
              colspan: 1,
              rowspan: 1
            };
          });
          list.push(item);
        });
      });
      importTitle.value = title;
      importData.value = list;
      loading.close();
    });
  };
  reader.readAsArrayBuffer(file);
  return false;
};

/** Import excel split and merge cells */
const importFile2 = (file) => {
  if (
    ![
      'application/vnd.ms-excel',
      'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
    ].includes(file.type)
  ) {
    EleMessage.error('Only excel files are supported');
    return false;
  }
  if (file.size / 1024 / 1024 > 20) {
    EleMessage.error('Size cannot exceed 20MB');
    return false;
  }
  const loading = EleMessage.loading({ message: 'Loading...', plain: true });
  const reader = new FileReader();
  reader.onload = (e) => {
    const workbook = new ExcelJS.Workbook();
    workbook.xlsx.load(e.target?.result).then(() => {
      const title = [];
      const list = [];
      workbook.eachSheet((sheet) => {
        list.splice(0, list.length);
        sheet.eachRow({ includeEmpty: true }, (row, rowIndex) => {
          title.splice(0, title.length);
          const item = {};
          row.eachCell({ includeEmpty: true }, (cell) => {
            const key = cell.address.slice(0, -String(rowIndex).length);
            title.push(key);
            item[key] = {
              value: cell.value,
              colspan: 1,
              rowspan: 1
            };
          });
          list.push(item);
        });
      });
      importTitle.value = title;
      importData.value = list;
      loading.close();
    });
  };
  reader.readAsArrayBuffer(file);
  return false;
};

/** Import excel read merged information */
const importFile3 = (file) => {
  if (
    ![
      'application/vnd.ms-excel',
      'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
    ].includes(file.type)
  ) {
    EleMessage.error('Only excel files are supported');
    return false;
  }
  if (file.size / 1024 / 1024 > 20) {
    EleMessage.error('Size cannot exceed 20MB');
    return false;
  }
  const loading = EleMessage.loading({ message: 'Loading...', plain: true });
  const reader = new FileReader();
  reader.onload = (e) => {
    const workbook = new ExcelJS.Workbook();
    workbook.xlsx.load(e.target?.result).then(() => {
      const title = [];
      const list = [];
      workbook.eachSheet((sheet) => {
        list.splice(0, list.length);
        sheet.eachRow({ includeEmpty: true }, (row) => {
          title.splice(0, title.length);
          const item = {};
          row.eachCell({ includeEmpty: true }, (cell) => {
            const key = cell.address.slice(0, -String(cell.row).length);
            title.push(key);
            const isMerged =
              cell.isMerged && cell.address !== cell.master.address;
            item[key] = {
              value: cell.value,
              colspan: isMerged ? 0 : 1,
              rowspan: isMerged ? 0 : 1
            };
          });
          list.push(item);
        });
        // Generate cross-row and cross-column tables
        sheet.eachRow({ includeEmpty: true }, (row) => {
          row.eachCell({ includeEmpty: true }, (cell) => {
            const master = cell.master;
            if (cell.isMerged && cell.address !== master.address) {
              const mk = master.address.slice(0, -String(master.row).length);
              const mItem = list[Number(master.row) - 1][mk];
              const rs = Math.abs(Number(cell.row) - Number(master.row)) + 1;
              rs > mItem.rowspan && (mItem.rowspan = rs);
              const cs = Math.abs(Number(cell.col) - Number(master.col)) + 1;
              cs > mItem.colspan && (mItem.colspan = cs);
            }
          });
        });
      });
      importTitle.value = title;
      importData.value = list;
      loading.close();
    });
  };
  reader.readAsArrayBuffer(file);
  return false;
};
</script>
