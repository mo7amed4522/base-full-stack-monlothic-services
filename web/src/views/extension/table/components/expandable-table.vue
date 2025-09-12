<template>
  <div>
    <ele-pro-table row-key="userId" :columns="columns" :datasource="datasource" :show-overflow-tooltip="true"
      :export-config="{ fileName: 'User data', datasource: exportSource }" :print-config="{ datasource: exportSource }">
      <template #expand="{ row }">
        <div style="padding: 12px 16px 12px 32px">
          <div>
            <span>Phone Number：</span>
            <span>{{ row.phone }}</span>
          </div>
          <div style="margin-top: 12px">
            <span>Organization：</span>
            <span>{{ row.organizationName }}</span>
          </div>
          <div style="margin-top: 12px">
            <span>Role：</span>
            <span>
              <el-tag v-for="item in row.roles" :key="item.roleId" size="small" :disable-transitions="true">
                {{ item.roleName }}
              </el-tag>
            </span>
          </div>
          <div style="margin-top: 12px">
            <span>Condition&emsp;&emsp;state: </span>
            <span>
              <ele-dot v-if="row.status === 0" text="Normal" size="8px" />
              <ele-dot v-else-if="row.status === 1" text="Frozen" type="danger" :ripple="false" size="8px" />
            </span>
          </div>
        </div>
      </template>
    </ele-pro-table>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { pageUsers, listUsers } from '@/api/system/user';

/** 表格列配置 */
const columns = ref([
  {
    type: 'expand',
    columnKey: 'expand',
    width: 50,
    align: 'center',
    slot: 'expand',
    formatter: (row) => {
      return [
        `Phone Number：${row.phone}`,
        `Organization：${row.organizationName}`,
        `Role：${row.roles.map((r) => r.roleName).join()}`,
        `Condition&emsp;&emsp;state: ${row.status == 0 ? 'Normal' : 'Frozen'}`
      ].join('\n');
    }
  },
  {
    type: 'index',
    columnKey: 'index',
    width: 50,
    align: 'center'
  },
  {
    prop: 'username',
    label: 'User account',
    sortable: 'custom',
    minWidth: 140
  },
  {
    prop: 'nickname',
    label: 'User name',
    sortable: 'custom',
    minWidth: 140
  },
  {
    prop: 'createTime',
    label: 'Create time',
    sortable: 'custom',
    minWidth: 180,
    align: 'center'
  }
]);

/** 表格数据源 */
const datasource = ({ pages, orders }) => {
  return pageUsers({ ...orders, ...pages });
};

/** 导出和打印全部数据的数据源 */
const exportSource = ({ orders }) => {
  return listUsers({ ...orders });
};
</script>
