<template>
  <ele-card header="Open multiple dialogs">
    <option-item label="Limit in the main area" :responsive="false" style="margin-bottom: 20px">
      <el-switch v-model="inner" size="small" />
    </option-item>
    <div>
      <el-button class="ele-btn-icon" @click="openDialog1">
        Open dialog1
      </el-button>
      <el-button class="ele-btn-icon" @click="openDialog2">
        Open dialog2
      </el-button>
      <el-button class="ele-btn-icon" @click="openDialog3">
        Open dialog3
      </el-button>
    </div>
    <ele-text block type="secondary" style="margin-top: 20px">
      When opening multiple dialogs, clicking will automatically top
    </ele-text>
  </ele-card>
  <ele-modal :inner="inner" :width="460" title="Dialog1" v-model="visible1" :resizable="true" :maxable="true"
    :multiple="true" :destroy-on-close="false" :move-out="['right', 'bottom']" position="center">
    <div style="padding: 40px 0">Dialog1</div>
    <template #footer>
      <el-button @click="visible1 = false">Cancel</el-button>
      <el-button type="primary">OK</el-button>
    </template>
  </ele-modal>
  <ele-modal :inner="inner" :width="460" title="Dialog2" v-model="visible2" :resizable="true" :maxable="true"
    :multiple="true" :destroy-on-close="false" position="leftBottom">
    <div style="padding: 40px 0">Dialog2</div>
    <template #footer>
      <el-button @click="visible2 = false">Cancel</el-button>
      <el-button type="primary">OK</el-button>
    </template>
  </ele-modal>
  <ele-modal :inner="inner" :width="360" title="Dialog3" v-model="visible3" :resizable="true" :maxable="true"
    :multiple="true" :destroy-on-close="false" position="rightTop" :style="{
      borderRadius: '8px',
      backdropFilter: 'blur(4px)',
      background: 'hsla(0, 0%, 26%, 0.4)',
      '--ele-modal-icon-color': 'rgba(255, 255, 255, 0.8)',
      '--ele-modal-icon-hover-color': 'rgba(255, 255, 255, 1)',
      '--ele-modal-icon-hover-bg': 'rgba(255, 255, 255, .15)',
      '--ele-modal-header-border': '1px solid rgba(255, 255, 255, 0.1)',
      '--ele-modal-header-padding': '14px 20px',
      '--ele-modal-fullscreen-border': 'none',
      '--ele-scrollbar-color': 'rgba(255, 255, 255, .36)',
      '--ele-scrollbar-hover-color': 'rgba(255, 255, 255, .58)'
    }" :title-style="{ color: '#fff' }" :body-style="{ color: '#fff' }">
    <div style="padding: 40px 0">弹窗3</div>
  </ele-modal>
</template>

<script setup>
import { ref } from 'vue';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';

/** Dialog1 Is it open? */
const visible1 = ref(false);

const visible2 = ref(false);

const visible3 = ref(false);

/** Open Dialog1 */
const openDialog1 = () => {
  visible1.value = true;
};

const openDialog2 = () => {
  visible2.value = true;
};

const openDialog3 = () => {
  visible3.value = true;
};

/** Limit in the main area */
const inner = ref(false);
</script>
