<template>
  <ele-card header="MessageBox">
    <option-item label="Restricted to the main area" :responsive="false" style="margin-bottom: 20px">
      <el-switch v-model="inner" size="small" />
    </option-item>
    <div>
      <el-button class="ele-btn-icon" @click="openConfirm">Confirm</el-button>
      <el-button class="ele-btn-icon" @click="openPrompt">Prompt</el-button>
      <el-button class="ele-btn-icon" @click="openAlert">Alert</el-button>
    </div>
    <div style="margin-top: 20px">
      <el-button class="ele-btn-icon" @click="openSuccess">Succes </el-button>
      <el-button class="ele-btn-icon" @click="openWarning">Warning</el-button>
      <el-button class="ele-btn-icon" @click="openError">Error</el-button>
      <el-button class="ele-btn-icon" @click="openInfo">Info</el-button>
    </div>
  </ele-card>
</template>

<script setup>
import { ref } from 'vue';
import { useMessageBox } from 'ele-admin-plus/es';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';

const messageBox = useMessageBox();

const inner = ref(false);

const openConfirm = () => {
  messageBox
    .confirm('Are you sure you want to delete? After deletion, it will be kept in the recycle bin for 30 days', 'Delete confirmation', {
      type: 'error',
      inner: inner.value,
      draggable: true
    })
    .then(() => {
      console.log('Clicked OK');
    })
    .catch(() => { });
};

const openPrompt = () => {
  messageBox
    .prompt('Please enter the new password:', 'Reset password', {
      inputPattern: /^[\S]{5,18}$/,
      inputErrorMessage: 'Password must be 5-18 non-blank characters',
      inner: inner.value,
      draggable: true
    })
    .then(({ value }) => {
      console.log(value);
    })
    .catch(() => { });
};

const openAlert = () => {
  messageBox
    .alert('Content content content content content content content', 'Title', {
      inner: inner.value,
      draggable: true
    })
    .then(() => {
      console.log('Clicked OK');
    })
    .catch(() => { });
};

const openSuccess = () => {
  messageBox
    .alert(
      'Import completed, a total of 80 pieces of data were added, 20 pieces of data were updated, and 0 pieces of data failed to import.',
      'Import successful',
      {
        type: 'success',
        confirmButtonText: 'Great',
        inner: inner.value,
        draggable: true
      }
    )
    .then(() => {
      console.log('Clicked Great');
    })
    .catch(() => { });
};

const openWarning = () => {
  messageBox
    .confirm(
      'Your professional version membership will expire in 6 days. Please renew in time to avoid affecting normal use!',
      'Expiration reminder',
      {
        type: 'warning',
        cancelButtonText: 'We will talk next time',
        confirmButtonText: '前往续费',
        closeOnClickModal: false,
        inner: inner.value,
        draggable: true
      }
    )
    .then(() => {
      console.log('Click to renew');
    })
    .catch(() => { });
};

const openError = () => {
  messageBox
    .alert(
      'Your account has violated regulations and has been forced offline and frozen by the system. Please contact the administrator for processing.。',
      'Abnormal reminder',
      {
        type: 'error',
        confirmButtonText: 'View details',
        showClose: false,
        inner: inner.value,
        draggable: true
      }
    )
    .then(() => {
      console.log('Click to view details');
    })
    .catch(() => { });
};

const openInfo = () => {
  messageBox
    .alert(
      'This week, from 03:00 to 05:00, the website will be upgraded. During this period, the system will be unavailable. Please prepare in advance!',
      'System notice',
      {
        type: 'info',
        confirmButtonText: 'Know',
        closeOnClickModal: true,
        inner: inner.value,
        draggable: true
      }
    )
    .then(() => {
      console.log('Click to know');
    })
    .catch(() => { });
};
</script>
