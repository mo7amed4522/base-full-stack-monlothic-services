<!-- Change password pop-up window -->
<template>
  <ele-modal form :width="420" title="Change password" :append-to-body="true" :model-value="modelValue"
    @update:modelValue="updateModelValue" @closed="onCancel">
    <el-form ref="formRef" :model="form" :rules="rules" label-width="80px" @submit.prevent="">
      <el-form-item label="Old password" prop="password">
        <el-input show-password type="password" :maxlength="20" v-model="form.password"
          placeholder="Please enter old password" />
      </el-form-item>
      <el-form-item label="New Password" prop="newPassword">
        <el-input show-password type="password" :maxlength="20" v-model="form.newPassword"
          placeholder="Please enter your new password" />
      </el-form-item>
      <el-form-item label="Confirm Password" prop="repetPassword">
        <el-input show-password type="password" :maxlength="20" v-model="form.repetPassword"
          placeholder="Please enter your new password again" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="updateModelValue(false)">Cancel</el-button>
      <el-button type="primary" :loading="loading" @click="onOk">Save</el-button>
    </template>
  </ele-modal>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';
import { changePasswordApi } from '@/api/layout';

const emit = defineEmits(['update:modelValue']);

defineProps({
  modelValue: Boolean
});

/** Submit loading */
const loading = ref(false);

/** Form instance */
const formRef = ref(null);

/** Form data */
const [form, resetFields] = useFormData({
  password: '',
  newPassword: '',
  repetPassword: ''
});

/** Form validation rules */
const rules = reactive({
  password: [
    {
      required: true,
      message: 'Please enter old password',
      type: 'string',
      trigger: 'blur'
    }
  ],
  newPassword: [
    {
      required: true,
      message: 'Please enter new password',
      type: 'string',
      trigger: 'blur'
    },
    {
      type: 'string',
      trigger: 'blur',
      pattern: /^[\S]{5,18}$/,
      message: 'Password must be 5-18 non-blank characters'
    }
  ],
  repetPassword: [
    {
      required: true,
      message: 'Please enter new password again',
      type: 'string',
      trigger: 'blur'
    },
    {
      type: 'string',
      trigger: 'blur',
      validator: (_rule, value, callback) => {
        if (value && value !== form.newPassword) {
          return callback(new Error('Passwords do not match'));
        }
        callback();
      }
    }
  ]
});

/** Update modelValue */
const updateModelValue = (value) => {
  emit('update:modelValue', value);
};

/** Save changes */
const onOk = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    loading.value = true;
    changePasswordApi(form)
      .then((msg) => {
        loading.value = false;
        EleMessage.success(msg);
        updateModelValue(false);
      })
      .catch((e) => {
        loading.value = false;
        EleMessage.error(e.message);
      });
  });
};

/** Close callback */
const onCancel = () => {
  resetFields();
  formRef.value?.clearValidate?.();
  loading.value = false;
};
</script>
