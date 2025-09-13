<template>
  <org-user-search ref="searchRef" style="margin-bottom: -14px" @search="reload" />
  <ele-pro-table ref="tableRef" row-key="userId" :columns="columns" :datasource="datasource"
    :show-overflow-tooltip="true" v-model:selections="selections" :highlight-current-row="true"
    :export-config="{ fileName: 'User data', datasource: exportSource }" :print-config="{ datasource: exportSource }"
    :footer-style="{ paddingBottom: '16px' }" cache-key="systemOrgUserTable">
    <template #toolbar>
      <el-button type="primary" class="ele-btn-icon" :icon="PlusOutlined" @click="openEdit()">
        New user
      </el-button>
      <el-button type="danger" class="ele-btn-icon" :icon="DeleteOutlined" @click="remove()">
        Delete user
      </el-button>
    </template>
    <template #roles="{ row }">
      <el-tag v-for="item in row.roles" :key="item.roleId" size="small" :disable-transitions="true">
        {{ item.roleName }}
      </el-tag>
    </template>
    <template #status="{ row }">
      <el-switch size="small" :model-value="row.status === 0" @change="(checked) => editStatus(checked, row)" />
    </template>
    <template #action="{ row }">
      <el-link type="primary" :underline="false" @click="openEdit(row)">
        Edit
      </el-link>
      <el-divider direction="vertical" />
      <ele-dropdown :items="[
        { title: 'Reset password', command: 'password' },
        { title: 'Delete user', command: 'delete', danger: true }
      ]" style="display: inline" @command="(key) => dropClick(key, row)">
        <el-link type="primary" :underline="false">
          <span>More</span>
          <el-icon :size="12" style="vertical-align: -1px; margin-left: 2px">
            <ArrowDown />
          </el-icon>
        </el-link>
      </ele-dropdown>
    </template>
  </ele-pro-table>
  <org-user-edit :data="current" v-model="showEdit" :organization-id="organizationId" @done="reload" />
</template>

<script setup>
import { ref, watch } from 'vue';
import { ElMessageBox } from 'element-plus/es';
import { EleMessage } from 'ele-admin-plus/es';
import { PlusOutlined, DeleteOutlined, ArrowDown } from '@/components/icons';
import OrgUserSearch from './org-user-search.vue';
import OrgUserEdit from './org-user-edit.vue';
import {
  pageUsers,
  removeUsers,
  updateUserStatus,
  updateUserPassword,
  listUsers
} from '@/api/system/user';

const props = defineProps({
  /** 机构 id */
  organizationId: Number
});

/** 搜索栏实例 */
const searchRef = ref(null);

/** 表格实例 */
const tableRef = ref(null);

/** 表格列配置 */
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
    prop: 'username',
    label: 'Username',
    sortable: 'custom',
    minWidth: 110
  },
  {
    prop: 'nickname',
    label: 'Nickname',
    sortable: 'custom',
    minWidth: 110
  },
  {
    prop: 'sexName',
    label: 'Sex',
    sortable: 'custom',
    width: 90,
    align: 'center'
  },
  {
    columnKey: 'roles',
    label: 'Role',
    minWidth: 120,
    slot: 'roles',
    align: 'center',
    formatter: (row) => row.roles.map((d) => d.roleName).join(',')
  },
  {
    prop: 'createTime',
    label: 'Create time',
    sortable: 'custom',
    width: 180,
    align: 'center'
  },
  {
    prop: 'status',
    label: 'Status',
    width: 90,
    align: 'center',
    sortable: 'custom',
    slot: 'status',
    formatter: (row) => (row.status === 0 ? 'Normal' : 'Frozen')
  },
  {
    columnKey: 'action',
    label: 'Action',
    width: 128,
    align: 'center',
    slot: 'action',
    hideInPrint: true,
    hideInExport: true
  }
]);

/** 表格选中数据 */
const selections = ref([]);

/** 当前编辑数据 */
const current = ref(null);

/** 是否显示编辑弹窗 */
const showEdit = ref(false);

/** 表格数据源 */
const datasource = ({ pages, where, orders }) => {
  return pageUsers({
    ...where,
    ...orders,
    ...pages,
    organizationId: props.organizationId
  });
};

/** 搜索 */
const reload = (where) => {
  tableRef.value?.reload?.({ page: 1, where });
};

/** 打开编辑弹窗 */
const openEdit = (row) => {
  current.value = row ?? null;
  showEdit.value = true;
};

/** 删除 */
const remove = (row) => {
  const rows = row == null ? selections.value : [row];
  if (!rows.length) {
    EleMessage.error('Please select at least one data');
    return;
  }
  ElMessageBox.confirm(
    'Are you sure you want to delete “' +
    rows.map((d) => d.nickname).join(', ') +
    '”?',
    'System prompt',
    { type: 'warning', draggable: true }
  )
    .then(() => {
      const loading = EleMessage.loading({
        message: 'Requesting...',
        plain: true
      });
      removeUsers(rows.map((d) => d.userId))
        .then((msg) => {
          loading.close();
          EleMessage.success(msg);
          reload();
        })
        .catch((e) => {
          loading.close();
          EleMessage.error(e.message);
        });
    })
    .catch(() => { });
};

/** 修改用户状态 */
const editStatus = (checked, row) => {
  const status = checked ? 0 : 1;
  updateUserStatus(row.userId, status)
    .then((msg) => {
      row.status = status;
      EleMessage.success(msg);
    })
    .catch((e) => {
      EleMessage.error(e.message);
    });
};

/** 下拉菜单点击事件 */
const dropClick = (key, row) => {
  if (key === 'password') {
    ElMessageBox.prompt(`Please enter the new password for user "${row.nickname}":`, 'Reset password', {
      inputPattern: /^[\S]{5,18}$/,
      inputErrorMessage: 'Password must be 5-18 non-blank characters',
      draggable: true
    })
      .then(({ value }) => {
        const loading = EleMessage.loading({
          message: 'Requesting...',
          plain: true
        });
        updateUserPassword(row.userId, value)
          .then((msg) => {
            loading.close();
            EleMessage.success(msg);
          })
          .catch((e) => {
            loading.close();
            EleMessage.error(e.message);
          });
      })
      .catch(() => { });
  } else if (key === 'delete') {
    remove(row);
  }
};

// 监听机构 id 变化
watch(
  () => props.organizationId,
  () => {
    searchRef.value?.resetFields?.();
    reload({});
  }
);

/** 导出和打印全部数据的数据源 */
const exportSource = ({ where, orders }) => {
  return listUsers({
    ...where,
    ...orders,
    organizationId: props.organizationId
  });
};
</script>
