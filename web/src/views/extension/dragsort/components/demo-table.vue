<template>
  <ele-card header="Table drag and drop sort">
    <el-row :gutter="16">
      <el-col v-for="(item, index) in taskList" :key="index" :md="8" :sm="24" :xs="24">
        <ele-table :border="true" class="demo-table">
          <thead>
            <tr>
              <th style="width: 40px"></th>
              <th>Task name</th>
              <th style="text-align: center; width: 80px">Status</th>
            </tr>
          </thead>
          <vue-draggable tag="tbody" item-key="id" :animation="300" :modelValue="item" :force-fallback="true"
            group="demoDragTable" handle=".demo-table-handle" :set-data="() => void 0"
            @update:modelValue="(value) => updateModelValue(value, index)">
            <template #item="{ element }">
              <tr>
                <td style="text-align: center; width: 40px">
                  <ele-text :icon="DragOutlined" type="placeholder" class="demo-table-handle" style="cursor: move" />
                </td>
                <td>{{ element.taskName }}</td>
                <td style="text-align: center; width: 80px">
                  <ele-text v-if="element.status === 0" type="warning">
                    Not started
                  </ele-text>
                  <ele-text v-else-if="element.status === 1" type="success">
                    In progress
                  </ele-text>
                  <ele-text v-else-if="element.status === 2" type="info">
                    Completed
                  </ele-text>
                </td>
              </tr>
            </template>
            <template #footer v-if="!item.length">
              <tr>
                <td colspan="3" style="text-align: center; background: none">
                  <ele-text type="placeholder">No data</ele-text>
                </td>
              </tr>
            </template>
          </vue-draggable>
        </ele-table>
      </el-col>
    </el-row>
  </ele-card>
</template>

<script setup>
import { ref } from 'vue';
import VueDraggable from 'vuedraggable';
import { DragOutlined } from '@/components/icons';

/** data */
const taskList = ref([]);

/** update data */
const updateModelValue = (value, index) => {
  taskList.value[index] = value;
};

/** handle data, 6 data per table */
const temp = [];
for (let i = 0; i < 18; i++) {
  const index = Number.parseInt(String(i / 6));
  if (temp[index] == null) {
    temp[index] = [];
  }
  temp[index].push({
    id: i,
    taskName: 'Task ' + (i + 1),
    status: 0
  });
}
taskList.value = temp;
</script>

<style lang="scss" scoped>
.demo-table {
  margin-bottom: 12px;
  table-layout: fixed;

  tr.sortable-ghost {
    opacity: 0;
  }

  tr.sortable-fallback {
    opacity: 1 !important;
    display: table !important;

    td {
      background: var(--el-color-primary-light-8);
      border-radius: 0 !important;
    }
  }
}
</style>
