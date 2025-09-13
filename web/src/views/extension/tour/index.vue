<template>
  <ele-page>
    <ele-card header="Basic usage">
      <div>
        <el-button type="primary" @click="onStart1">Start tour</el-button>
      </div>
      <div style="margin-top: 20px">
        <el-button ref="uploadRef1">Upload</el-button>
        <el-button ref="saveRef1" type="primary">Save</el-button>
        <el-button ref="moreRef1">More</el-button>
      </div>
      <ele-tour v-model="current1" :steps="steps1" />
    </ele-card>
    <ele-card header="Without mask layer">
      <div>
        <el-button type="primary" @click="onStart2">Start tour</el-button>
      </div>
      <div style="margin-top: 20px">
        <el-button ref="uploadRef2">Upload</el-button>
        <el-button ref="saveRef2" type="primary">Save</el-button>
        <el-button ref="moreRef2">More</el-button>
      </div>
      <ele-tour v-model="current2" :steps="steps2" :mask="false" />
    </ele-card>
    <ele-card header="Multiple forms">
      <div>
        <el-button type="primary" @click="onStart3">Start tour</el-button>
      </div>
      <div style="margin-top: 20px">
        <el-button ref="uploadRef3">Upload</el-button>
        <el-button ref="saveRef3" type="primary">Save</el-button>
        <el-button ref="moreRef3">More</el-button>
      </div>
      <ele-tour v-model="current3" :steps="steps3">
        <template #text="{ step, current }">
          <template v-if="current === 0">
            <div style="margin-bottom: 10px">
              <img src="https://gw.alipayobjects.com/mdn/rms_08e378/afts/img/A*P0S-QIRUbsUAAAAAAAAAAABkARQnAQ"
                style="height: 184px; width: 100%; object-fit: cover" />
            </div>
            <div>{{ step.description }}</div>
          </template>
        </template>
      </ele-tour>
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref } from 'vue';

/** 当前步骤 */
const current1 = ref();

/** 按钮 */
const uploadRef1 = ref(null);
const saveRef1 = ref(null);
const moreRef1 = ref(null);

/** 步骤 */
const steps1 = ref([
  {
    target: () => uploadRef1.value?.$el,
    title: 'Upload',
    description: 'Click this button to upload a file.'
  },
  {
    target: () => saveRef1.value?.$el,
    title: 'Save',
    description: 'Click this button to save the data.'
  },
  {
    target: () => moreRef1.value?.$el,
    title: 'More',
    description: 'Hover this button to show more operations.'
  }
]);

/** 开始引导 */
const onStart1 = () => {
  current1.value = 0;
};

/** 当前步骤 */
const current2 = ref();

/** 按钮 */
const uploadRef2 = ref(null);
const saveRef2 = ref(null);
const moreRef2 = ref(null);

/** 步骤 */
const steps2 = ref([
  {
    target: () => uploadRef2.value?.$el,
    title: 'Upload',
    description: 'Click this button to upload a file.',
    popoverProps: { placement: 'top-start' }
  },
  {
    target: () => saveRef2.value?.$el,
    title: 'Save',
    description: 'Click this button to save the data.',
    popoverProps: { placement: 'bottom' }
  },
  {
    target: () => moreRef2.value?.$el,
    title: 'More',
    description: 'Hover this button to show more operations.',
    popoverProps: { placement: 'top-end' }
  }
]);

/** 开始引导 */
const onStart2 = () => {
  current2.value = 0;
};

/** 当前步骤 */
const current3 = ref();

/** 按钮 */
const uploadRef3 = ref(null);
const saveRef3 = ref(null);
const moreRef3 = ref(null);

/** 步骤 */
const steps3 = ref([
  {
    title: 'Welcome to EleAdminPlus',
    description:
      'Below will introduce some common functions. If you have already introduced them, you can click skip to end the guide.'
  },
  {
    target: () => uploadRef3.value?.$el,
    title: 'Upload',
    description: 'Click this button to upload a file.'
  },
  {
    target: () => saveRef3.value?.$el,
    title: 'Save',
    description: 'Click this button to save the data.',
    mask: false
  },
  {
    target: () => moreRef3.value?.$el,
    title: 'More',
    description: 'Hover this button to show more operations.'
  }
]);

/** 开始引导 */
const onStart3 = () => {
  current3.value = 0;
};
</script>

<script>
export default {
  name: 'ExtensionTour'
};
</script>
