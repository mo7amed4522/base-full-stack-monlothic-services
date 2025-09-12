<!-- Latest Activities -->
<template>
  <ele-card :header="title" :body-style="{ padding: '6px 0', height: '370px' }">
    <template #extra>
      <more-icon @command="onCommand" />
    </template>
    <el-scrollbar :view-style="{ padding: '20px 20px 0 20px' }">
      <el-timeline :reverse="false" class="demo-timeline">
        <el-timeline-item v-for="item in activities" :key="item.id" :timestamp="item.time" :type="item.type"
          :hollow="true">
          {{ item.title }}
        </el-timeline-item>
      </el-timeline>
    </el-scrollbar>
  </ele-card>
</template>

<script setup>
import { ref } from 'vue';
import MoreIcon from './more-icon.vue';

defineProps({
  title: String
});

const emit = defineEmits(['command']);

/** Latest Activities Data */
const activities = ref([]);

/** Query Latest Activities */
const queryActivities = () => {
  activities.value = [
    {
      id: 1,
      title: 'User SunSmile solved the bug of login prompt operation failed',
      time: '20:30'
    },
    {
      id: 2,
      title: 'User Jasmine solved the bug of button color not match the design',
      time: '19:30'
    },
    {
      id: 3,
      title: 'Project Manager assigned the task to solve the bug of project one',
      time: '18:30',
      type: 'primary'
    },
    {
      id: 4,
      title: 'Project Manager assigned the task to solve the bug of project two',
      time: '17:30',
      type: 'primary'
    },
    {
      id: 5,
      title: 'Project Manager assigned the task to solve the bug of project three',
      time: '16:30',
      type: 'primary'
    },
    {
      id: 6,
      title: 'The project manager assigned the task to fix the bug in project 4.',
      time: '15:30'
    },
    {
      id: 7,
      title: 'The project manager assigned the task to fix the bug in project 5.',
      time: '14:30'
    },
    {
      id: 8,
      title: 'The project manager assigned the task to fix the bug in project 6',
      time: '12:30'
    },
    {
      id: 9,
      title: 'The project manager assigned the task to fix the bug in project 7.',
      time: '11:30',
      type: 'primary'
    },
    {
      id: 10,
      title: 'The project manager assigned the task to fix the bug in project 8.',
      time: '10:30'
    },
    {
      id: 11,
      title: 'The project manager assigned the task to fix the bug in project 9.',
      time: '09:30',
      type: 'success'
    },
    {
      id: 12,
      title: 'The project manager assigned the task to fix the bug in project 10.',
      time: '08:30',
      type: 'danger'
    }
  ];
};

const onCommand = (command) => {
  emit('command', command);
};

queryActivities();
</script>

<style lang="scss" scoped>
/* Timeline */
.demo-timeline {
  padding-left: 0;

  :deep(.el-timeline-item__wrapper) {
    display: flex;

    .el-timeline-item__timestamp {
      order: 0;
      flex-shrink: 0;
      margin: 0 16px 0 0;
      height: 22px;
      line-height: 22px;
      font-size: 14px;
    }

    .el-timeline-item__content {
      order: 1;
      flex: 1;
    }
  }

  :deep(.el-timeline-item__node) {
    top: 3px;
    --el-color-white: var(--el-bg-color-overlay);
  }

  :deep(.el-timeline-item__tail) {
    top: 3px;
  }
}
</style>
