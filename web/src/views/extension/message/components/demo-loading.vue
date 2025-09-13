<template>
  <ele-card header="Loading demo">
    <option-item label="Status" :responsive="false">
      <el-switch v-model="loading" size="small" />
    </option-item>
    <option-item label="Size" :responsive="false" style="margin: 6px 0 0 0">
      <el-radio-group v-model="loadingSize" :disabled="loadingType === 'circle'">
        <el-radio value="small" label="Small" />
        <el-radio value="default" label="Default" />
        <el-radio value="large" label="Large" />
      </el-radio-group>
    </option-item>
    <option-item label="Text" :responsive="false" style="margin: 2px 0 0 0">
      <el-radio-group v-model="loadingText" @change="onLoadingTextChange">
        <el-radio value="" label="No text" />
        <el-radio value="Loading..." label="With text" />
      </el-radio-group>
    </option-item>
    <option-item label="Type" :responsive="false" style="margin: 2px 0 0 0">
      <el-radio-group v-model="loadingType">
        <el-radio value="dot" label="Dot" />
        <el-radio value="circle" label="Circle" />
      </el-radio-group>
    </option-item>
    <ele-loading :key="loadingKey" :loading="loading" :size="loadingSize" :type="loadingType" :text="loadingText"
      :blur="true" style="padding: 20px 16px; max-width: 520px; position: relative">
      <div style="font-size: 17px; max-height: 100%">Alert message title</div>
      <div style="word-break: break-all; margin-top: 4px; max-height: 100%">
        <span>Further details about the context of this alert.</span>
        <span>Further details about the context of this alert.</span>
        <span>Further details about the context of this alert.</span>
        <span>Further details about the context of this alert.</span>
        <span>Further details about the context of this alert.</span>
        <span>Further details about the context of this alert.</span>
        <span>Further details about the context of this alert.</span>
        <span>Further details about the context of this alert.</span>
      </div>
    </ele-loading>
  </ele-card>
</template>

<script setup>
import { ref } from 'vue';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';

/** Loading status */
const loading = ref(false);

/** Loading size */
const loadingSize = ref('default');

/** Loading type */
const loadingType = ref('dot');

/** Loading text */
const loadingText = ref('');

/** Key for re-rendering loading */
const loadingKey = ref(0);

/** Loading type change event */
const onLoadingTextChange = () => {
  if (loadingType.value === 'circle' && loading.value) {
    loadingKey.value++;
  }
};
</script>
