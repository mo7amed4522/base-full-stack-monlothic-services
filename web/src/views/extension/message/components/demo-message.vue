<template>
  <ele-card header="Message demo">
    <option-item label="Theme">
      <el-radio-group v-model="theme">
        <el-radio :value="0" label="Default" />
        <el-radio :value="1" label="Simple" />
        <el-radio :value="2" label="Colorful" />
        <el-radio :value="3" label="Colorful simple" />
      </el-radio-group>
    </option-item>
    <option-item label="Limit in body" style="margin-top: 16px">
      <el-switch v-model="inner" size="small" />
    </option-item>
    <option-item label="Type">
      <el-button @click="showMessage('success')">Success</el-button>
      <el-button @click="showMessage('warning')">Warning</el-button>
      <el-button @click="showMessage('error')">Error</el-button>
      <el-button @click="showMessage('info')">Info</el-button>
    </option-item>
    <option-item label="Loading">
      <el-button @click="showLoading()">Open</el-button>
      <el-button @click="showLoading(true)">With mask</el-button>
      <el-button @click="showLoading(true, true)">Centered</el-button>
    </option-item>
    <option-item label="Reactive">
      <el-button @click="showReactiveLoading">Open</el-button>
    </option-item>
  </ele-card>
</template>

<script setup>
import { ref } from 'vue';
import { useMessage, useMessageBox, random } from 'ele-admin-plus/es';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';

const message = useMessage();
const messageBox = useMessageBox();

/** Theme Style */
const theme = ref(0);

/** Limit in body */
const inner = ref(false);

/** Show message */
const showMessage = (type) => {
  const messages = {
    success: 'This is a successful tip',
    warning: 'This is a warning tip',
    error: 'This is an error tip',
    info: 'This is a normal tip'
  };
  message({
    type,
    message: messages[type],
    plain: theme.value === 1,
    original: theme.value === 3 ? 'plain' : theme.value === 2,
    inner: inner.value,
    showClose: true
  });
};

/** Show loading */
const showLoading = (mask, centered) => {
  const loading = message.loading({
    message: 'Loading...',
    plain: theme.value === 1,
    original: theme.value === 3 ? 'plain' : theme.value === 2,
    inner: inner.value,
    showClose: true,
    mask,
    centered
  });
  setTimeout(() => {
    loading.close();
  }, 3000);
};

/** Loading progress */
const progress = ref(0);

/** Show reactive loading */
const showReactiveLoading = () => {
  progress.value = 0;
  const loading = message.loading({
    message: () =>
      `Loading...${String(progress.value).padStart(2, '0')}%`,
    plain: true,
    inner: inner.value,
    mask: true,
    centered: true
  });
  const timer = setInterval(() => {
    const num = progress.value + random(0, 5);
    progress.value = num > 100 ? 100 : num;
    if (progress.value == 100) {
      loading.close();
      clearInterval(timer);
      messageBox
        .alert(
          'Loading complete, a total of 80 pieces of data were added, 20 pieces of data were updated, and 0 pieces of data failed.',
          'Import successful',
          {
            type: 'success',
            confirmButtonText: 'Complete',
            inner: inner.value,
            draggable: true
          }
        )
        .then(() => {
          console.log('Click complete');
        })
        .catch(() => { });
    }
  }, 100);
};
</script>
