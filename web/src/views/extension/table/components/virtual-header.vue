<template>
  <div>
    <ele-pro-table :height="418" :virtual="true" row-key="userId" :columns="columns" :datasource="datasource"
      :show-overflow-tooltip="true" :pagination="{
        pageSize: 1000,
        pageSizes: [1000, 2000]
      }" :border="true" :toolbar="{ theme: 'default' }" :table-style="{ overflow: 'hidden' }"
      :export-config="{ fileName: 'User data', datasource: exportSource }"
      :print-config="{ datasource: exportSource }" />
  </div>
</template>

<script setup>
import { ref } from 'vue';
import dayjs from 'dayjs';
const addresses = [
  ['Zhejiang', 'Hangzhou', 'Xihu district'],
  ['Jiangsu', 'Suzhou', 'Guhu district'],
  ['Jiangxi', 'Nanchang', 'Qingshan district'],
  ['Fujian', 'Quanzhou', 'Fengze district'],
  ['Hubei', 'Wuhan', 'Wuchang district'],
  ['Anhui', 'Huangshan', 'Huangshan district']
];

const data = [];

/** 表格列配置 */
const columns = ref([
  {
    type: 'index',
    columnKey: 'index',
    width: 68,
    align: 'center',
    fixed: 'left'
  },
  {
    label: 'User account',
    prop: 'username',
    minWidth: 160,
    sortable: 'custom'
  },
  {
    label: 'Username',
    prop: 'nickname',
    minWidth: 160,
    sortable: 'custom'
  },
  {
    label: 'Home address',
    prop: 'address',
    align: 'center',
    children: [
      {
        label: 'Province',
        prop: 'province',
        width: 120,
        sortable: 'custom'
      },
      {
        label: 'City',
        prop: 'city',
        width: 120,
        sortable: 'custom'
      },
      {
        label: 'District',
        prop: 'district',
        width: 180,
        sortable: 'custom'
      }
    ]
  },
  {
    label: 'Create time',
    prop: 'createTime',
    width: 180,
    sortable: 'custom',
    align: 'center'
  }
]);

/** 表格数据源 */
const datasource = async ({ pages, orders }) => {
  await new Promise((resolve) => setTimeout(resolve, 100));
  if (!data.length) {
    Array.from({ length: 2000 }).forEach((_, i) => {
      const userId = i + 1;
      const username = 'user-' + String(userId).padStart(6, '0');
      const index = Math.floor(
        Math.random() * (0 - addresses.length) + addresses.length
      );
      const [province, city, zone] = addresses[index];
      const item = {
        userId,
        username,
        nickname: 'User ' + userId,
        createTime: dayjs()
          .subtract(45 * userId, 's')
          .format('YYYY-MM-DD HH:mm:ss'),
        province,
        city,
        zone
      };
      data.push(item);
    });
  }
  const list = [...data];
  if (orders != null && orders.sort) {
    list.sort((a, b) => {
      const aValue = a[orders.sort];
      const bValue = b[orders.sort];
      if (aValue == bValue) {
        return 0;
      }
      const r = aValue < bValue ? -1 : 1;
      return orders.order === 'desc' ? -r : r;
    });
  }
  const start = ((pages.page || 1) - 1) * (pages.limit || 10);
  const end = start + (pages.limit || 10);
  return {
    count: list.length,
    list: list.slice(start, end > list.length ? list.length : end)
  };
};

/** 导出和打印全部数据的数据源 */
const exportSource = async ({ orders }) => {
  const res = await datasource({
    orders,
    pages: { page: 1, limit: data.length }
  });
  return res.list;
};
</script>
