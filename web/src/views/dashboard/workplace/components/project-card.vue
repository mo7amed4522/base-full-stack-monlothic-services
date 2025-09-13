<!-- Project Progress -->
<template>
  <ele-card :header="title" :body-style="{ padding: '10px', height: '370px' }">
    <template #extra>
      <more-icon @command="onCommand" />
    </template>
    <ele-pro-table :height="352" row-key="id" :columns="columns" :datasource="projectList" :show-overflow-tooltip="true"
      highlight-current-row :pagination="false" :toolbar="false" :bottom-line="false" size="large"
      class="project-table">
      <template #projectName="{ row }">
        <el-link type="primary" :underline="false">
          {{ row.projectName }}
        </el-link>
      </template>
      <template #status="{ row }">
        <ele-text v-if="row.status === 0" type="success">In Progress</ele-text>
        <ele-text v-else-if="row.status === 1" type="danger">Delayed</ele-text>
        <ele-text v-else-if="row.status === 2" type="warning">
          Not Started
        </ele-text>
        <ele-text v-else-if="row.status === 3" type="info">Completed</ele-text>
      </template>
      <template #progress="{ row }">
        <el-progress :percentage="row.progress" />
      </template>
    </ele-pro-table>
  </ele-card>
</template>

<script setup>
import { ref } from 'vue';
import MoreIcon from './more-icon.vue';

defineProps({
  title: String
});

const emit = defineEmits(['command']);

/** Table column configuration */
const columns = ref([
  {
    type: 'index',
    columnKey: 'index',
    width: 50,
    align: 'center',
    fixed: 'left'
  },
  {
    prop: 'projectName',
    label: 'Project Name',
    slot: 'projectName',
    minWidth: 110
  },
  {
    prop: 'startDate',
    label: 'Start Date',
    align: 'center',
    minWidth: 110
  },
  {
    prop: 'endDate',
    label: 'End Date',
    align: 'center',
    minWidth: 110
  },
  {
    prop: 'status',
    label: 'Status',
    slot: 'status',
    align: 'center',
    width: 90
  },
  {
    prop: 'progress',
    label: 'Progress',
    width: 180,
    align: 'center',
    slot: 'progress',
    showOverflowTooltip: false
  }
]);

/** Project progress data */
const projectList = ref([]);

/** Query project progress */
const queryProjectList = () => {
  projectList.value = [
    {
      id: 1,
      projectName: 'Project 000000001',
      status: 0,
      startDate: '2020-03-01',
      endDate: '2020-06-01',
      progress: 30
    },
    {
      id: 2,
      projectName: 'Project 000000002',
      status: 0,
      startDate: '2020-03-01',
      endDate: '2020-08-01',
      progress: 10
    },
    {
      id: 3,
      projectName: 'Project 000000003',
      status: 1,
      startDate: '2020-01-01',
      endDate: '2020-05-01',
      progress: 60
    },
    {
      id: 4,
      projectName: 'Project 000000004',
      status: 1,
      startDate: '2020-06-01',
      endDate: '2020-10-01',
      progress: 0
    },
    {
      id: 5,
      projectName: 'Project 000000005',
      status: 2,
      startDate: '2020-03-01',
      endDate: '2020-07-01',
      progress: 100
    },
    {
      id: 6,
      projectName: 'Project 000000006',
      status: 3,
      startDate: '2020-02-01',
      endDate: '2020-06-01',
      progress: 100
    }
  ];
};

const onCommand = (command) => {
  emit('command', command);
};

queryProjectList();
</script>

<style lang="scss" scoped>
.project-table :deep(.el-progress__text) {
  font-size: 12px !important;
}
</style>
