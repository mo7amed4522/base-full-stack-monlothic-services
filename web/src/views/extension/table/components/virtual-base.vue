<template>
  <div>
    <div class="option-wrapper">
      <option-item label="Toolbar background" :responsive="false">
        <el-switch v-model="toolbarBg" size="small" />
      </option-item>
      <option-item label="Display border" :responsive="false">
        <el-switch v-model="border" size="small" />
      </option-item>
      <option-item label="Display zebra striping" :responsive="false">
        <el-switch v-model="stripe" size="small" />
      </option-item>
      <option-item label="Display header" :responsive="false">
        <el-switch v-model="showHeader" size="small" />
      </option-item>
      <option-item label="Display footer" :responsive="false">
        <el-switch v-model="showSummary" size="small" />
      </option-item>
      <option-item label="Row click highlight" :responsive="false">
        <el-switch v-model="highlightCurrentRow" size="small" />
      </option-item>
      <option-item label="Row click tick" :responsive="false">
        <el-switch v-model="rowClickChecked" size="small" />
      </option-item>
    </div>
    <ele-pro-table :height="418" :virtual="true" row-key="userId" :columns="columns" :datasource="datasource"
      :show-overflow-tooltip="true" :highlight-current-row="highlightCurrentRow" :row-click-checked="rowClickChecked"
      v-model:selections="selections" v-model:current="current" :pagination="{
        pageSize: 2000,
        pageSizes: [1000, 2000, 5000, 10000]
      }" :toolbar="{ theme: toolbarBg ? 'default' : 'plain' }" :border="border" :stripe="stripe"
      :show-header="showHeader" :show-summary="showSummary" :summary-method="getSummaries"
      :export-config="{ fileName: 'User data', datasource: exportSource }" :print-config="{ datasource: exportSource }">
      <template #toolbar>
        <el-button type="primary" class="ele-btn-icon" @click="getSelections">
          Get selected rows
        </el-button>
        <el-button v-if="highlightCurrentRow" type="primary" class="ele-btn-icon" @click="getCurrent">
          Get single row
        </el-button>
      </template>
      <template #roles="{ row }">
        <el-tag v-for="item in row.roles" :key="item" size="small" :disable-transitions="true"
          style="margin-right: 6px">
          {{ item }}
        </el-tag>
      </template>
      <template #status="{ row }">
        <ele-dot v-if="row.status == 0" text="正常" size="8px" />
        <ele-dot v-else-if="row.status == 1" text="冻结" type="danger" :ripple="false" size="8px" />
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
import { ref } from 'vue';
import dayjs from 'dayjs';
import { EleMessage } from 'ele-admin-plus/es';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';

const data = [];

/** 表格列配置 */
const columns = ref([
  {
    type: 'selection',
    columnKey: 'selection',
    width: 52,
    align: 'center',
    fixed: 'left'
  },
  {
    type: 'index',
    columnKey: 'index',
    width: 68,
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
    label: 'Sex',
    prop: 'sex',
    width: 120,
    sortable: 'custom',
    align: 'center'
  },
  {
    label: 'Role',
    prop: 'roles',
    slot: 'roles',
    minWidth: 160,
    filters: [
      { text: 'Admin', value: 'administrator' },
      { text: 'Normal user', value: 'Ordinary users' },
      { text: 'Visitor', value: 'Tourism' }
    ],
    filterMultiple: true,
    align: 'center'
  },
  {
    label: 'Phone number',
    prop: 'phone',
    minWidth: 160,
    sortable: 'custom',
    align: 'center'
  },
  {
    label: 'Email',
    prop: 'email',
    minWidth: 180,
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
      { text: 'Normal', value: '0' },
      { text: 'Frozen', value: '1' }
    ],
    filterMultiple: false,
    formatter: (row) => (row.status == 0 ? 'Normal' : 'Frozen')
  },
  {
    label: 'Birthday',
    prop: 'birthday',
    width: 160,
    align: 'center',
    sortable: 'custom'
  },
  {
    label: 'Create time',
    prop: 'createTime',
    width: 180,
    sortable: 'custom',
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

/** 表格数据源 */
const datasource = async ({ pages, filter, orders }) => {
  await new Promise((resolve) => setTimeout(resolve, 100));
  if (!data.length) {
    Array.from({ length: 10000 }).forEach((_, i) => {
      const userId = i + 1;
      const username = 'user-' + String(userId).padStart(6, '0');
      const random = Math.random();
      const item = {
        userId,
        username,
        nickname: 'user' + userId,
        sex: random < 0.8 ? 'male' : 'female',
        roles: [random < 0.3 ? 'administrator' : 'Ordinary users'],
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
  const rolesFilter = filter.roles?.length ? filter.roles : null;
  const statusFilter = filter.status?.length ? filter.status : null;
  const list = data.filter((d) => {
    if (rolesFilter != null) {
      if (!d.roles || !d.roles.some((r) => rolesFilter.includes(r))) {
        return false;
      }
    }
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

/** 表格多选选中数据 */
const selections = ref([]);

/** 表格单选选中数据 */
const current = ref(null);

/** 更多配置 */
const toolbarBg = ref(false);
const border = ref(false);
const stripe = ref(false);
const showHeader = ref(true);
const showSummary = ref(true);
const highlightCurrentRow = ref(false);
const rowClickChecked = ref(false);

/** 打开编辑弹窗 */
const openEdit = (row) => {
  EleMessage.success('Edit:' + row.nickname);
};

/** 删除单个 */
const remove = (row) => {
  EleMessage.error('Delete:' + row.nickname);
};

/** 表尾合计行 */
const getSummaries = ({ columns, data }) => {
  const sums = [];
  columns.forEach((column, i) => {
    if (column.property === 'sex') {
      const num = data.filter((d) => d.sex == 'female').length;
      sums[i] = ((num / data.length) * 100).toFixed(2) + '% (Female)';
    } else if (column.property === 'status') {
      sums[i] = data.filter((d) => d.status == '1').length + ' (Frozen)';
    } else if (column.property === 'username') {
      sums[i] = 'Total';
    }
  });
  return sums;
};

/** 获取多选行数据 */
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
    EleMessage.success(current.value.nickname + ' The data has been printed on the console');
  }
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
