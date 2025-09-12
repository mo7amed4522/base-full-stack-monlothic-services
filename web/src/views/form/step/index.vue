<template>
  <div>
    <ele-card style="border-radius: 0">
      <ele-text type="heading" size="lg">Transfer</ele-text>
      <ele-text type="placeholder" style="margin-top: 6px">
        Break a lengthy or unfamiliar form task into multiple steps to guide users through it.
      </ele-text>
    </ele-card>
    <ele-page>
      <ele-card class="step-card">
        <div style="max-width: 800px; margin: 0 auto; padding: 10px 0 16px 0">
          <ele-steps :active="active" finish-status="success" type="inline" :items="[
            { title: 'Transfer information', description: 'Fill in the transfer information' },
            { title: 'Confirm transfer information', description: 'Confirm the transfer information' },
            { title: 'Transfer successful', description: 'Transfer successful' }
          ]" style="margin: 0 0 30px 0" />
          <step-edit v-if="active === 0" @done="onDone" />
          <step-confirm v-if="active === 1" :data="form" @done="onNext" @back="onBack" />
          <step-success v-if="active === 2" :data="form" @back="onBack" />
        </div>
        <template v-if="active === 0">
          <el-divider style="margin: 0 0 26px 0; opacity: 0.6" />
          <ele-alert :closable="false">
            <ele-text size="lg" style="margin-bottom: 12px">Transfer notice</ele-text>
            <ele-text size="md" style="margin-bottom: 8px">
              Transfer to Alipay
            </ele-text>
            <ele-text style="margin-bottom: 12px">
              If you need, you can put some common product questions here. If you need,
              you can put some common product questions here. If you need, you can put some common product questions
              here.
            </ele-text>
            <ele-text size="md" style="margin-bottom: 8px">Transfer to WeChat</ele-text>
            <ele-text style="margin-bottom: 8px">
              If you need, you can put some common product questions here. If you need,
              you can put some common product questions here. If you need, you can put some common product questions
              here.
            </ele-text>
          </ele-alert>
        </template>
      </ele-card>
    </ele-page>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import StepEdit from './components/step-edit.vue';
import StepConfirm from './components/step-confirm.vue';
import StepSuccess from './components/step-success.vue';

/** 选中步骤 */
const active = ref(0);

/** 表单数据 */
const form = reactive({});

/** 表单编辑完成事件 */
const onDone = (data) => {
  Object.assign(form, data);
  active.value = 1;
};

/** 下一步 */
const onNext = () => {
  active.value = 2;
};

/** 上一步 */
const onBack = () => {
  active.value = 0;
};
</script>

<script>
export default {
  name: 'FormStep'
};
</script>

<style lang="scss" scoped>
@media screen and (max-width: 768px) {
  .step-card :deep(.ele-card-body) {
    padding: 4px 8px 16px 8px;

    .el-form {
      padding-right: 8px !important;
    }

    .el-result {
      padding: 0;

      .el-result__extra {
        max-width: 100%;
      }
    }
  }
}
</style>
