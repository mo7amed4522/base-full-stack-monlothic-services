<template>
  <el-row :gutter="16">
    <el-col :md="8" :sm="24" :xs="24">
      <ele-card header="List drag and drop sort">
        <vue-draggable v-model="list" item-key="id" :animation="300" :set-data="() => void 0" class="demo-list">
          <template #item="{ element }">
            <div class="demo-list-item">{{ element.name }}</div>
          </template>
        </vue-draggable>
      </ele-card>
    </el-col>
    <el-col :md="16" :sm="24" :xs="24">
      <ele-card header="List drag and drop sort">
        <div style="display: flex; align-items: flex-start">
          <vue-draggable v-model="list1" item-key="id" :animation="300" group="demoDragList" :set-data="() => void 0"
            class="demo-list">
            <template #item="{ element }">
              <div class="demo-list-item">{{ element.name }}</div>
            </template>
          </vue-draggable>
          <vue-draggable v-model="list2" item-key="id" :animation="300" group="demoDragList" :set-data="() => void 0"
            class="demo-list" style="margin-left: 12px">
            <template #item="{ element }">
              <div class="demo-list-item">{{ element.name }}</div>
            </template>
          </vue-draggable>
        </div>
      </ele-card>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref } from 'vue';
import VueDraggable from 'vuedraggable';

/** List data */
const list = ref([
  { id: 1, name: 'List item 0000001' },
  { id: 2, name: 'List item 0000002' },
  { id: 3, name: 'List item 0000003' },
  { id: 4, name: 'List item 0000004' },
  { id: 5, name: 'List item 0000005' }
]);

/** List 1 data */
const list1 = ref([
  { id: 1, name: 'List item 0000001' },
  { id: 2, name: 'List item 0000002' },
  { id: 3, name: 'List item 0000003' },
  { id: 4, name: 'List item 0000004' },
  { id: 5, name: 'List item 0000005' }
]);

/** List 2 data */
const list2 = ref([
  { id: 6, name: 'List item 0000006' },
  { id: 7, name: 'List item 0000007' },
  { id: 8, name: 'List item 0000008' },
  { id: 9, name: 'List item 0000009' },
  { id: 10, name: 'List item 0000010' }
]);
</script>

<style lang="scss" scoped>
.demo-list {
  width: 230px;
  margin: 0 auto;
  max-width: 100%;
  min-height: 94px;
  border-radius: 4px;
  border: 1px solid hsla(0, 0%, 60%, 0.2);
  box-sizing: border-box;
  padding: 8px 8px 0 8px;
}

.demo-list-item {
  color: #fff;
  background: #1677ff;
  border-radius: 4px;
  padding: 8px 12px;
  margin-bottom: 8px;
  cursor: move;

  &.sortable-ghost {
    opacity: 0;
  }
}
</style>
