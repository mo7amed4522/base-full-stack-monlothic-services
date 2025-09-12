<template>
  <div>
    <ele-pro-table :height="378" :virtual="true" row-key="userId" :columns="columns" :datasource="datasource"
      :show-overflow-tooltip="true" :pagination="{
        pageSize: 1000,
        pageSizes: [1000, 2000]
      }" :export-config="{ fileName: 'User data', datasource: exportSource }"
      :print-config="{ datasource: exportSource }">
      <template #status="{ row }">
        <ele-dot v-if="row.status == 0" text="Active" size="8px" />
        <ele-dot v-else-if="row.status == 1" text="Frozen" type="danger" :ripple="false" size="8px" />
      </template>
      <template #expand="{ row }">
        <div style="padding: 12px 16px 12px 32px">
          <div>
            <span>horn&emsp;&emsp;color：</span>
            <span>
              <el-tag v-for="item in row.roles" :key="item" size="small" :disable-transitions="true"
                style="margin-right: 6px">
                {{ item }}
              </el-tag>
            </span>
          </div>
          <div style="margin-top: 12px">
            <span>Birthday：</span>
            <span>{{ row.birthday }}</span>
          </div>
          <div style="margin-top: 12px">
            <span>Email：</span>
            <span>{{ row.email }}</span>
          </div>
        </div>
      </template>
    </ele-pro-table>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import dayjs from 'dayjs';

const data = [];

/** 表格列配置 */
const columns = ref([
  {
    type: 'expand',
    columnKey: 'expand',
    width: 52,
    align: 'center',
    slot: 'expand',
    fixed: 'left',
    formatter: (row) => {
      return [
        `horn&emsp;&emsp;color：${row.roles.join()}`,
        `Birthday：${row.birthday}`,
        `Email：${row.email}`
      ].join('\n');
    }
  },
  {
    type: 'index',
    columnKey: 'index',
    width: 68,
    fixed: 'left'
  },
  {
    label: 'Username',
    prop: 'username',
    minWidth: 160,
    sortable: 'custom'
  },
  {
    label: 'Nickname',
    prop: 'nickname',
    minWidth: 160,
    sortable: 'custom'
  },
  {
    label: 'Sex',
    prop: 'sex',
    width: 120,
    sortable: 'custom',
    align: 'center'
  },
  {
    label: 'Phone',
    prop: 'phone',
    minWidth: 160,
    sortable: 'custom',
    align: 'center'
  },
  {
    prop: 'status',
    label: 'Status',
    width: 120,
    sortable: 'custom',
    slot: 'status',
    align: 'center',
    filters: [
      { text: 'Active', value: '0' },
      { text: 'Frozen', value: '1' }
    ],
    filterMultiple: false,
    formatter: (row) => (row.status == 0 ? 'Active' : 'Frozen')
  },
  {
    label: 'Create Time',
    prop: 'createTime',
    width: 180,
    sortable: 'custom',
    align: 'center'
  }
]);

/** 表格数据源 */
const datasource = async ({ pages, filter, orders }) => {
  await new Promise((resolve) => setTimeout(resolve, 100));
  if (!data.length) {
    Array.from({ length: 2000 }).forEach((_, i) => {
      const userId = i + 1;
      const username = 'user-' + String(userId).padStart(6, '0');
      const random = Math.random();
      const item = {
        userId,
        username,
        nickname: 'Nickname' + userId,
        sex: random < 0.8 ? 'Male' : 'Female',
        roles: [random < 0.3 ? 'Admin' : 'User'],
        phone: String(12345678900 + userId),
        email: username + '@eleadmin.com',
        status: random < 0.2 ? '1' : '0',
        birthday: dayjs()
          .subtract(3 * userId, 'h')
          .format('YYYY-MM-DD'),
        createTime: dayjs()
          .subtract(45 * userId, 's')
          .format('YYYY-MM-DD HH:mm:ss')
      };
      data.push(item);
    });
  }
  const statusFilter = filter.status?.length ? filter.status : null;
  const list = data.filter((d) => {
    if (statusFilter != null) {
      if (d.status == null || !statusFilter.includes(d.status)) {
        return false;
      }
    }
    return true;
  });
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
const exportSource = async ({ filter, orders }) => {
  const res = await datasource({
    filter,
    orders,
    pages: { page: 1, limit: data.length }
  });
  return res.list;
};
</script>
