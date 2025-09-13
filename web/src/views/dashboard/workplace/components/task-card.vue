<!-- My Tasks -->
<template>
  <ele-card :header="title" :body-style="{ height: '370px', padding: '7px 8px 6px 8px' }">
    <template #extra>
      <more-icon @command="onCommand" />
    </template>
    <el-scrollbar :wrap-style="{ position: 'relative', zIndex: 1 }">
      <ele-table class="task-table" size="large">
        <thead style="position: sticky; top: 0; z-index: 2">
          <tr>
            <th style="position: sticky; left: 0; z-index: 1; width: 38px"></th>
            <th style="text-align: center; width: 78px">Priority</th>
            <th>Task Name</th>
            <th style="text-align: center; width: 80px">Status</th>
          </tr>
        </thead>
        <vue-draggable tag="tbody" item-key="id" v-model="taskList" handle=".sort-handle" :animation="300"
          :set-data="() => void 0" :force-fallback="true">
          <template #item="{ element }">
            <tr>
              <td :style="{
                textAlign: 'center',
                paddingLeft: 0,
                paddingRight: 0,
                position: 'sticky',
                left: 0,
                zIndex: 1,
                width: '38px'
              }">
                <ele-text :icon="DragOutlined" :icon-style="{ transform: 'scale(1.15)' }" type="placeholder"
                  class="sort-handle" />
              </td>
              <td style="text-align: center; width: 78px">
                <el-tag v-if="element.priority === 1" type="danger" :disable-transitions="true">
                  {{ element.priority }}
                </el-tag>
                <el-tag v-else-if="element.priority === 2" type="warning" :disable-transitions="true">
                  {{ element.priority }}
                </el-tag>
                <el-tag v-else :disable-transitions="true">
                  {{ element.priority }}
                </el-tag>
              </td>
              <td>
                <ele-ellipsis style="line-height: 20px">
                  <el-link type="primary" :underline="false">
                    {{ element.taskName }}
                  </el-link>
                </ele-ellipsis>
              </td>
              <td style="text-align: center; width: 80px">
                <ele-text v-if="element.status === 0" type="warning">
                  Not Started
                </ele-text>
                <ele-text v-else-if="element.status === 1" type="success">
                  In Progress
                </ele-text>
                <ele-text v-else-if="element.status === 2" type="info" deleted>
                  Completed
                </ele-text>
              </td>
            </tr>
          </template>
        </vue-draggable>
      </ele-table>
    </el-scrollbar>
  </ele-card>
</template>

<script setup>
import { ref } from 'vue';
import VueDraggable from 'vuedraggable';
import { DragOutlined } from '@/components/icons';
import MoreIcon from './more-icon.vue';

defineProps({
  title: String
});

const emit = defineEmits(['command']);

/** My Tasks Data */
const taskList = ref([]);

/** Query My Tasks */
const queryTaskList = () => {
  taskList.value = [
    {
      id: 1,
      priority: 1,
      taskName: 'Solve the bug of project one',
      status: 0
    },
    {
      id: 2,
      priority: 2,
      taskName: 'Solve the bug of project two',
      status: 0
    },
    {
      id: 3,
      priority: 2,
      taskName: 'Solve the bug of project three',
      status: 1
    },
    {
      id: 4,
      priority: 3,
      taskName: 'Solve the bug of project four',
      status: 1
    },
    {
      id: 5,
      priority: 3,
      taskName: 'Solve the bug of project five',
      status: 2
    },
    {
      id: 6,
      priority: 3,
      taskName: 'Solve the bug of project six',
      status: 2
    }
  ];
};

const onCommand = (command) => {
  emit('command', command);
};

queryTaskList();
</script>

<style lang="scss" scoped>
.task-table {
  table-layout: fixed;
  min-width: 300px;

  .sort-handle {
    cursor: move;
  }

  .el-tag {
    width: 20px;
    height: 20px;
    border-radius: 50%;
  }

  td,
  th {
    box-sizing: border-box;
  }

  tr.sortable-ghost {
    opacity: 0;
  }

  tr.sortable-fallback {
    opacity: 1 !important;
    display: table !important;
    table-layout: fixed !important;

    td {
      background: var(--el-color-primary-light-8);
    }
  }
}
</style>
