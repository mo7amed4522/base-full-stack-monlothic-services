<template>
  <ele-page>
    <ele-card header="Basic Usage">
      <ele-steps :active="active" finish-status="success" :items="[
        { title: 'Step 1', description: 'Fill in transfer information' },
        { title: 'Step 2', description: 'Confirm transfer information' },
        { title: 'Step 3', description: 'Transfer successful' }
      ]" :align-center="!!alignCenter" />
      <div style="margin-top: 28px; display: flex; align-items: center">
        <el-button @click="onPrev">Previous</el-button>
        <el-button type="primary" @click="onNext">Next</el-button>
        <div style="margin-left: 12px">
          <el-checkbox v-model="alignCenter" :true-value="1" :false-value="0" label="Center" />
        </div>
      </div>
    </ele-card>
    <ele-card header="Compact Version">
      <ele-steps :active="active" finish-status="success" type="inline" :items="[
        { title: 'Step 1', description: 'Fill in transfer information' },
        { title: 'Step 2', description: 'Confirm transfer information' },
        { title: 'Step 3', description: 'Transfer successful' }
      ]" />
    </ele-card>
    <ele-card header="Vertical Direction">
      <div style="height: 320px">
        <ele-steps :active="active" finish-status="success" direction="vertical" :items="[
          { title: 'Step 1', description: 'Fill in transfer information' },
          { title: 'Step 2', description: 'Confirm transfer information' },
          { title: 'Step 3', description: 'Transfer successful' }
        ]" />
      </div>
    </ele-card>
    <ele-card header="Icon Compact Version">
      <div style="margin-bottom: 28px">
        <el-button @click="onPrev">Previous</el-button>
        <el-button type="primary" @click="onNext">Next</el-button>
      </div>
      <ele-steps :active="active" type="inline" :items="[
        { title: 'Step 1', icon: UserFilled },
        { title: 'Step 2', icon: Checked },
        { title: 'Step 3', icon: SuccessFilled }
      ]" :style="{
          '--ele-step-icon-size': '24px',
          '--ele-step-line-height': '24px'
        }" />
    </ele-card>
    <ele-card header="Icon Center Version">
      <ele-steps :active="active" :align-center="true" :items="[
        { title: 'Step 1', icon: UserFilled },
        { title: 'Step 2', icon: Checked },
        { title: 'Step 3', icon: SuccessFilled }
      ]" :style="{
          '--ele-step-icon-size': '24px',
          '--ele-step-line-height': '24px'
        }" />
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref } from 'vue';
import { UserFilled, Checked, SuccessFilled } from '@element-plus/icons-vue';

/** Active Step */
const active = ref(0);

/** Center */
const alignCenter = ref(0);

/** Previous */
const onPrev = () => {
  if (active.value < 1) {
    active.value = 3;
    return;
  }
  active.value = active.value - 1;
};

/** Next */
const onNext = () => {
  if (active.value > 2) {
    active.value = 0;
    return;
  }
  active.value = active.value + 1;
};
</script>

<script>
export default {
  name: 'ExtensionSteps'
};
</script>
