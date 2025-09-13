<template>
  <div>
    <ele-card style="border-radius: 0">
      <ele-text type="heading" size="lg">Basic Form</ele-text>
      <ele-text type="placeholder" style="margin-top: 6px">
        Form pages are used to collect or verify information, and are common in scenarios with few data items.
      </ele-text>
    </ele-card>
    <ele-page :multi-card="false">
      <ele-card>
        <el-form status-icon ref="formRef" :model="form" :rules="rules" label-width="80px"
          style="max-width: 700px; margin: 10px auto" @submit.prevent="">
          <el-form-item label="Title" prop="title">
            <el-input clearable v-model="form.title" placeholder="Please enter a title" />
          </el-form-item>
          <el-form-item label="Date Range" prop="datetime">
            <el-date-picker unlink-panels type="daterange" v-model="form.datetime" range-separator="-"
              format="YYYY-MM-DD" value-format="YYYY-MM-DD" start-placeholder="Start date" end-placeholder="End date"
              class="ele-fluid" />
          </el-form-item>
          <el-form-item label="Goal Description" prop="goal">
            <el-input show-word-limit maxlength="800" :rows="4" type="textarea" v-model="form.goal"
              placeholder="Please enter a goal description" />
          </el-form-item>
          <el-form-item label="Standard" prop="standard">
            <el-input :rows="4" type="textarea" v-model="form.standard" placeholder="Please enter a standard" />
          </el-form-item>
          <el-form-item label="Location" prop="address">
            <el-select clearable v-model="form.address" placeholder="Please select a location" class="ele-fluid"
              :popper-options="{ strategy: 'fixed' }">
              <el-option value="1" label="Location 1" />
              <el-option value="2" label="Location 2" />
              <el-option value="3" label="Location 3" />
            </el-select>
          </el-form-item>
          <el-form-item label="Invitees" prop="invites">
            <el-select multiple clearable v-model="form.invites" placeholder="Please select invitees" class="ele-fluid"
              :popper-options="{ strategy: 'fixed' }">
              <el-option-group label="Group 1">
                <el-option :value="1" label="SunSmile" />
                <el-option :value="2" label="Your name is good" />
                <el-option :value="3" label="The hope of the whole family" />
              </el-option-group>
              <el-option-group label="Group 2">
                <el-option :value="4" label="Jasmine" />
                <el-option :value="5" label="Cool man" />
              </el-option-group>
            </el-select>
          </el-form-item>
          <el-form-item label="Weight" prop="weight">
            <el-input-number :min="0" :max="100" v-model="form.weight" controls-position="right" class="ele-fluid"
              style="max-width: 160px" />
            <div style="padding-left: 8px">%</div>
          </el-form-item>
          <el-form-item label="Public Type" style="margin-bottom: 0">
            <el-radio-group v-model="form.publicType">
              <el-radio :value="1" label="Public" />
              <el-radio :value="2" label="Partial Public" />
              <el-radio :value="3" label="Not Public" />
            </el-radio-group>
          </el-form-item>
          <el-form-item v-if="form.publicType === 2" style="margin-bottom: 16px">
            <el-input clearable v-model="form.publicValue" placeholder="Please enter the public value"
              class="ele-fluid" />
          </el-form-item>
          <el-form-item>
            <ele-text type="placeholder" style="line-height: 1.6">
              The default share is to the customer and invitees.
            </ele-text>
          </el-form-item>
          <el-form-item>
            <el-button @click="finishPageTab()">Close</el-button>
            <el-button type="primary" :loading="loading" @click="submit">
              Submit
            </el-button>
          </el-form-item>
        </el-form>
      </ele-card>
    </ele-page>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';
import { usePageTab } from '@/utils/use-page-tab';

const { finishPageTab } = usePageTab();

/** 加载状态 */
const loading = ref(false);

/** 表单实例 */
const formRef = ref(null);

/** 表单数据 */
const [form, resetFields] = useFormData({
  title: '',
  datetime: void 0,
  goal: '',
  standard: '',
  address: void 0,
  invites: [],
  weight: 0,
  publicType: 1,
  publicValue: ''
});

/** 表单验证规则 */
const rules = reactive({
  title: [
    {
      required: true,
      message: 'Please enter a title',
      type: 'string',
      trigger: 'blur'
    }
  ],
  datetime: [
    {
      required: true,
      message: 'Please select the start and end date',
      type: 'array',
      trigger: 'blur'
    }
  ],
  goal: [
    {
      required: true,
      message: 'Please enter a goal description',
      type: 'string',
      trigger: 'blur'
    }
  ],
  standard: [
    {
      required: true,
      message: 'Please enter a standard',
      type: 'string',
      trigger: 'blur'
    }
  ],
  address: [
    {
      required: true,
      message: 'Please select a location',
      type: 'string',
      trigger: 'blur'
    }
  ],
  invites: [
    {
      required: true,
      message: 'Please select invitees',
      type: 'array',
      trigger: 'change'
    }
  ],
  weight: [
    {
      required: true,
      message: 'Please enter a weight',
      type: 'number',
      trigger: 'change'
    }
  ]
});

/** 提交 */
const submit = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    loading.value = true;
    setTimeout(() => {
      loading.value = false;
      resetFields();
      EleMessage.success('Submit success');
    }, 1500);
  });
};
</script>

<script>
export default {
  name: 'FormBasic'
};
</script>
