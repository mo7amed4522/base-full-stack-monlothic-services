<template>
  <ele-card header="Modify menu badge">
    <el-form label-width="80px" style="max-width: 360px" @submit.prevent="">
      <el-form-item label="Menu">
        <el-tree-select clearable :data="treeData" check-strictly default-expand-all placeholder="Please select menu"
          v-model="path" class="ele-fluid" :popper-options="{ strategy: 'fixed' }">
          <template #default="{ data }">
            <el-icon v-if="data.meta?.icon" style="margin-right: 4px; transform: translateY(-1px)">
              <component :is="data.meta?.icon" />
            </el-icon>
            <span>{{ data.label }}</span>
          </template>
          <template v-if="selectedIcon" #prefix>
            <el-icon color="var(--el-text-color-regular)" style="margin-right: 6px">
              <component :is="selectedIcon" />
            </el-icon>
          </template>
        </el-tree-select>
      </el-form-item>
      <el-form-item label="Badge Value">
        <el-input clearable v-model="badge" placeholder="Please enter a badge value" />
      </el-form-item>
      <el-form-item label="Badge Type">
        <el-select clearable v-model="type" placeholder="Please select a badge type" class="ele-fluid"
          :popper-options="{ strategy: 'fixed' }">
          <el-option label="primary" value="primary" />
          <el-option label="success" value="success" />
          <el-option label="warning" value="warning" />
          <el-option label="danger" value="danger" />
          <el-option label="info" value="info" />
        </el-select>
      </el-form-item>
      <el-form-item style="margin: 0">
        <el-button type="primary" @click="setBadge">renew</el-button>
      </el-form-item>
    </el-form>
  </ele-card>
</template>

<script setup>
import { ref, computed } from 'vue';
import { storeToRefs } from 'pinia';
import { EleMessage, mapTree, findTree } from 'ele-admin-plus/es';
import { useUserStore } from '@/store/modules/user';
import * as MenuIcons from '@/layout/menu-icons';

defineOptions({
  components: MenuIcons
});

/** Menu tree */
const userStore = useUserStore();
const { menus } = storeToRefs(userStore);

/** Menu tree dropdown data */
const treeData = computed(() => {
  return mapTree(menus.value, (m) => {
    return {
      ...m,
      value: m.path,
      label: m.meta?.title
    };
  });
});

/** Selected menu */
const path = ref();

/** Badge value */
const badge = ref();

/** Badge type */
const type = ref();

/** Select the menu icon */
const selectedIcon = computed(() => {
  if (!path.value) {
    return;
  }
  return findTree(treeData.value, (d) => d.path == path.value)?.meta?.icon;
});

/** Modify the menu badge data */
const setBadge = () => {
  if (!path.value) {
    EleMessage.error('Please select a menu');
    return;
  }
  userStore.setMenuBadge(
    path.value,
    !badge.value || isNaN(Number(badge.value))
      ? badge.value
      : Number(badge.value),
    type.value
  );
};
</script>
