<template>
  <ele-card header="Basic Usage">
    <div style="max-width: 260px">
      <ele-table-select clearable placeholder="Please select" value-key="userId" label-key="username"
        v-model="selectedValue" :table-props="tableProps" :popper-width="580" :popper-options="{ strategy: 'fixed' }">
        <!-- 角色列 -->
        <template #roles="{ row }">
          <el-tag v-for="item in row.roles" :key="item.roleId" size="small" :disable-transitions="true">
            {{ item.roleName }}
          </el-tag>
        </template>
      </ele-table-select>
    </div>
    <div style="margin-top: 14px">
      <el-button type="primary" @click="setValue">Set Value</el-button>
    </div>
  </ele-card>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { listUsers } from '@/api/system/user';

/** 表格下拉选中值 */
const selectedValue = ref();

/** 表格配置 */
const tableProps = reactive({
  datasource: [],
  columns: [
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
      sortable: 'custom'
    },
    {
      prop: 'nickname',
      label: 'Nickname',
      sortable: 'custom',
      slot: 'nickname'
    },
    {
      prop: 'sexName',
      label: 'Sex',
      sortable: 'custom',
      width: 80
    },
    {
      columnKey: 'roles',
      label: 'Roles',
      slot: 'roles'
    }
  ],
  showOverflowTooltip: true,
  highlightCurrentRow: true,
  toolbar: false,
  pagination: {
    pageSize: 6,
    layout: 'total, prev, pager, next, jumper',
    style: { padding: '0px' }
  },
  rowStyle: { cursor: 'pointer' }
});

/** 查询表格数据 */
listUsers().then((data) => {
  tableProps.datasource = data;
});

/** 回显数据 */
const setValue = () => {
  selectedValue.value = 40;
};
</script>
