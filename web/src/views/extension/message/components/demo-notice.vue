<template>
  <ele-card header="Notice demo">
    <option-item label="Usage">
      <el-button @click="showNotification()" class="ele-btn-icon">
        Open notice
      </el-button>
      <el-button @click="showNotification('success')" class="ele-btn-icon">
        With icon
      </el-button>
      <el-button @click="showNotification('info', 0)" class="ele-btn-icon">
        Do not close automatically
      </el-button>
    </option-item>
    <option-item label="Type">
      <el-button @click="showNotification('success')">Success</el-button>
      <el-button @click="showNotification('warning')">Warning</el-button>
      <el-button @click="showNotification('error')">Error</el-button>
      <el-button @click="showNotification('info')">Info</el-button>
    </option-item>
    <option-item label="Position">
      <el-button @click="showNotification('info', null, 'top-left')" class="ele-btn-icon">
        TopLeft
      </el-button>
      <el-button @click="showNotification('info', null, 'bottom-left')" class="ele-btn-icon">
        BottomLeft
      </el-button>
      <el-button @click="showNotification('info', null, 'bottom-right')" class="ele-btn-icon">
        BottomRight
      </el-button>
    </option-item>
  </ele-card>
</template>

<script setup>
import { ElNotification } from 'element-plus/es';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';

/** Open notice */
const showNotification = (type, duration, position) => {
  ElNotification({
    title: 'Notification Title',
    message:
      'This is the content of the notification. This is the content of the notification. ',
    type: type || '',
    duration: duration ?? 4500,
    position: position || 'top-right'
  });
};
</script>
