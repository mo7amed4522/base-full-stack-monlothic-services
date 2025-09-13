<template>
  <ele-card header="Basic usage" :body-style="{ overflow: 'hidden', padding: '8px 13px 13px 13px' }">
    <option-item label="Show collapse button" :responsive="false" style="margin: 0">
      <el-radio-group v-model="allowCollapse">
        <el-radio :value="true" label="Yes" />
        <el-radio :value="false" label="No" />
      </el-radio-group>
    </option-item>
    <option-item label="Freeze size" :responsive="false" style="margin: 0">
      <el-radio-group v-model="resizable">
        <el-radio :value="true" label="Yes" />
        <el-radio :value="false" label="No" />
      </el-radio-group>
    </option-item>
    <option-item label="Vertical layout" :responsive="false" style="margin: 0">
      <el-radio-group v-model="vertical">
        <el-radio :value="true" label="Yes" />
        <el-radio :value="false" label="No" />
      </el-radio-group>
    </option-item>
    <option-item label="Reverse layout" :responsive="false" style="margin: 0">
      <el-radio-group v-model="reverse">
        <el-radio :value="true" label="Yes" />
        <el-radio :value="false" label="No" />
      </el-radio-group>
    </option-item>
    <ele-split-panel space="0px" size="160px" :allow-collapse="allowCollapse" :resizable="resizable"
      :vertical="vertical" :reverse="reverse" :custom-style="{
        background: 'rgba(185, 182, 229, .4)',
        overflow: 'hidden',
        border: 'none'
      }" :body-style="{
        background: 'rgba(125, 226, 252, .4)',
        overflow: 'hidden'
      }" style="height: 360px; margin-top: 12px">
      <div>Sidebar</div>
      <template #body>
        <div>Content</div>
      </template>
    </ele-split-panel>
  </ele-card>
  <ele-card header="Compose" :body-style="{ padding: '13px' }">
    <div style="margin: 0 0 8px 0">First left to right then top to bottom</div>
    <ele-split-panel space="0px" :min-size="40" :max-size="-40" :resizable="true" :custom-style="{
      background: 'rgba(185, 182, 229, .4)',
      overflow: 'hidden',
      border: 'none'
    }" :body-style="{ overflow: 'hidden' }" :responsive="false" style="height: 400px">
      <div>Sidebar</div>
      <template #body>
        <ele-split-panel space="0px" :min-size="40" :max-size="-40" :vertical="true" :resizable="true" :custom-style="{
          background: 'rgba(171, 199, 255, .5)',
          overflow: 'hidden',
          border: 'none'
        }" :body-style="{
            background: 'rgba(125, 226, 252, .4)',
            overflow: 'hidden'
          }" :responsive="false" style="height: 400px">
          <div>Top bar</div>
          <template #body>
            <div>Content</div>
          </template>
        </ele-split-panel>
      </template>
    </ele-split-panel>
    <div style="margin: 16px 0 8px 0">First up and down, then left and right</div>
    <ele-split-panel space="0px" size="120px" :min-size="40" :max-size="-40" :vertical="true" :resizable="true"
      :custom-style="{
        background: 'rgba(185, 182, 229, .4)',
        overflow: 'hidden',
        border: 'none'
      }" :body-style="{ overflow: 'hidden' }" :responsive="false" style="height: 400px">
      <div>Top bar</div>
      <template #body>
        <ele-split-panel space="0px" :min-size="40" :max-size="-40" :resizable="true" :custom-style="{
          background: 'rgba(171, 199, 255, .5)',
          overflow: 'hidden',
          border: 'none'
        }" :body-style="{
            background: 'rgba(125, 226, 252, .4)',
            overflow: 'hidden'
          }" :responsive="false" style="height: 100%">
          <div>Sidebar</div>
          <template #body>
            <div>Content</div>
          </template>
        </ele-split-panel>
      </template>
    </ele-split-panel>
  </ele-card>
  <ele-card header="Can be freely combined" :body-style="{ overflow: 'hidden', padding: '13px' }">
    <ele-split-panel space="0px" size="80px" :min-size="40" :max-size="0.5" :vertical="true" :resizable="true"
      :allow-collapse="true" :custom-style="{
        background: 'rgba(185, 182, 229, .4)',
        overflow: 'hidden',
        border: 'none'
      }" :body-style="{ overflow: 'visible' }" :responsive="false" style="height: 400px">
      <div>Top bar</div>
      <template #body>
        <ele-split-panel space="0px" size="120px" :min-size="40" :max-size="280" :resizable="true"
          :allow-collapse="true" :custom-style="{
            background: 'rgba(171, 199, 255, .5)',
            overflow: 'hidden',
            border: 'none'
          }" :body-style="{ overflow: 'visible' }" :responsive="false" style="height: 100%">
          <div>Sidebar</div>
          <template #body>
            <ele-split-panel space="0px" size="120px" :min-size="40" :max-size="280" :resizable="true"
              :allow-collapse="true" :reverse="true" :custom-style="{
                background: 'rgba(171, 199, 255, .5)',
                overflow: 'hidden',
                border: 'none'
              }" :body-style="{
                background: 'rgba(125, 226, 252, .4)',
                overflow: 'hidden'
              }" :responsive="false" style="height: 100%">
              <div>Right sidebar</div>
              <template #body>
                <div>Content</div>
              </template>
            </ele-split-panel>
          </template>
        </ele-split-panel>
      </template>
    </ele-split-panel>
  </ele-card>
  <ele-card header="Actual case">
    <el-button type="primary" @click="openTable">Combined table case</el-button>
  </ele-card>
</template>

<script setup>
import { ref } from 'vue';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';

const emit = defineEmits(['change']);

/** Whether to display the collapse button */
const allowCollapse = ref(true);

/** Whether to support free stretching */
const resizable = ref(false);

/** Whether to use vertical layout mode */
const vertical = ref(false);

/** Whether to reverse the layout direction */
const reverse = ref(false);

/** Switch to combined table case */
const openTable = () => {
  emit('change', 1);
};
</script>
