<template>
  <ele-card header="Tab operation">
    <div style="display: flex; align-items: center; padding-bottom: 8px">
      <div style="flex-shrink: 0; padding-right: 12px">Operation specified tab:</div>
      <div style="flex: 1; max-width: 160px">
        <el-select clearable v-model="tabPath" placeholder="Please select tab" class="ele-fluid"
          :popper-options="{ strategy: 'fixed' }">
          <el-option v-for="(item, index) in tabs" :key="item.key" :label="item.title"
            :value="item.fullPath || index" />
        </el-select>
      </div>
    </div>
    <div style="display: flex; flex-wrap: wrap; width: 460px; max-width: 100%">
      <el-button type="primary" style="margin: 16px 12px 0 0" @click="onUpdate2">
        Update tab title
      </el-button>
      <el-button type="primary" style="margin: 16px 12px 0 0" @click="onUpdate3">
        Update tab to non-closable
      </el-button>
      <el-button type="primary" style="margin: 16px 12px 0 0" @click="onClose2">
        Close tab
      </el-button>
      <el-button type="primary" style="margin: 16px 12px 0 0" @click="onClose3">
        Close current tab and switch to this tab
      </el-button>
      <el-button type="primary" style="margin: 16px 12px 0 0" @click="onCloseAndRefresh">
        Close current tab and refresh this tab
      </el-button>
    </div>
    <div style="margin: 42px 0 16px 0">Operation current tab:</div>
    <div style="display: flex; flex-wrap: wrap">
      <el-button type="primary" style="margin: 0 12px 16px 0" @click="onRefresh">
        Refresh current tab
      </el-button>
      <el-button type="primary" style="margin: 0 12px 16px 0" @click="onUpdate">
        Update current tab title
      </el-button>
      <el-button type="primary" style="margin: 0 12px 16px 0" @click="onClose">
        Close current tab
      </el-button>
      <el-button type="primary" style="margin: 0 12px 16px 0" @click="onCloseAll">
        Close all tabs
      </el-button>
    </div>
  </ele-card>
</template>

<script setup>
import { ref, unref } from 'vue';
import { storeToRefs } from 'pinia';
import { useRouter } from 'vue-router';
import { EleMessage } from 'ele-admin-plus/es';
import { useThemeStore } from '@/store/modules/theme';
import { usePageTab } from '@/utils/use-page-tab';

const {
  reloadPageTab,
  setPageTabTitle,
  finishPageTab,
  removeAllPageTab,
  getRouteTabKey,
  setPageTab,
  removePageTab,
  routeTabKey
} = usePageTab();

const { currentRoute, push } = useRouter();

/** Currently open tab */
const themeStore = useThemeStore();
const { tabs } = storeToRefs(themeStore);

/** Refresh current tab */
const onRefresh = () => {
  reloadPageTab();
};

/** Update current tab title */
const onUpdate = () => {
  setPageTabTitle('New title');
};

/** Close current tab */
const onClose = () => {
  finishPageTab();
};

/** Close all tabs */
const onCloseAll = () => {
  removeAllPageTab({ active: getRouteTabKey() });
};

/** Dropdown select */
const tabPath = ref('');

/** Update specified tab title */
const onUpdate2 = () => {
  if (
    !tabPath.value ||
    !tabs.value.some((t) => t.fullPath === tabPath.value)
  ) {
    tabPath.value = '';
    EleMessage.error('Please select a tab');
    return;
  }
  setPageTab({
    path: tabPath.value,
    title: 'New title'
  });
};

/** Update specified tab to non-closable */
const onUpdate3 = () => {
  if (
    !tabPath.value ||
    !tabs.value.some((t) => t.fullPath === tabPath.value)
  ) {
    tabPath.value = '';
    EleMessage.error('Please select a tab');
    return;
  }
  setPageTab({
    path: tabPath.value,
    closable: false
  });
};

/** Close specified tab */
const onClose2 = () => {
  if (!tabPath.value) {
    EleMessage.error('Please select a tab');
    return;
  }
  const t = tabs.value.find((t) => t.fullPath === tabPath.value);
  if (!t) {
    tabPath.value = '';
    EleMessage.error('Please select a tab');
    return;
  }
  if (t.path === unref(currentRoute).path) {
    onClose();
    return;
  }
  removePageTab({
    key: tabPath.value,
    active: getRouteTabKey()
  });
};

/** Close current tab and switch to specified tab */
const onClose3 = () => {
  if (!tabPath.value) {
    EleMessage.error('Please select a tab');
    return;
  }
  const t = tabs.value.find((t) => t.fullPath === tabPath.value);
  if (!t) {
    tabPath.value = '';
    EleMessage.error('Please select a tab');
    return;
  }
  if (!tabs.value.find((t) => t.key === routeTabKey)?.closable) {
    EleMessage.error('Current tab is not closable');
    return;
  }
  if (t.path === routeTabKey) {
    onRefresh();
    return;
  }
  removePageTab({ key: routeTabKey });
  push(tabPath.value);
};

/** Close current tab and switch to specified tab and refresh specified tab */
const onCloseAndRefresh = () => {
  if (!tabPath.value) {
    EleMessage.error('Please select a tab');
    return;
  }
  const t = tabs.value.find((t) => t.fullPath === tabPath.value);
  if (!t) {
    tabPath.value = '';
    EleMessage.error('Please select a tab');
    return;
  }
  if (!tabs.value.find((t) => t.key === routeTabKey)?.closable) {
    EleMessage.error('Current tab is not closable');
    return;
  }
  if (t.path === routeTabKey) {
    onRefresh();
    return;
  }
  removePageTab({ key: routeTabKey });
  reloadPageTab({ fullPath: tabPath.value });
};
</script>
