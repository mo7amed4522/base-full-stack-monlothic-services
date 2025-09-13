<template>
  <ele-modal form :width="550" :model-value="modelValue" :title="isUpdate ? 'Update template' : 'Add template'"
    @update:modelValue="updateModelValue">
    <el-form ref="formRef" :model="form" :rules="rules" label-width="80px" @submit.prevent="">
      <el-form-item label="Template name" prop="name">
        <el-input v-model="form.name" placeholder="Please enter a template name" clearable />
      </el-form-item>
      <el-form-item label="Template code" prop="code">
        <el-input v-model="form.code" placeholder="Please enter a template code" clearable />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="updateModelValue(false)">Cancel</el-button>
      <el-button type="primary" :loading="loading" @click="save">Save</el-button>
    </template>
  </ele-modal>
</template>

<script setup>
import { ref, reactive, watch } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';

import { createCPageTemplateApi, updateCPageTemplateApi } from '@/api/app';

const emit = defineEmits(['done', 'update:modelValue']);

const props = defineProps({
  /** Is the pop-up window open? */
  modelValue: Boolean,
  /** Template data for editing */
  data: Object
});

/** Is it an update operation? */
const isUpdate = ref(false);

/** Is the form loading? */
const loading = ref(false);

/** Form instance */
const formRef = ref(null);

/** Form data */
const [form, resetFields, assignFields] = useFormData({
  id: void 0,
  name: '',
  code: ''
});

/** Form validation rules */
const rules = reactive({
  name: [{ required: true, message: 'Please enter a template name', trigger: 'blur' }],
  code: [{ required: true, message: 'Please enter a template code', trigger: 'blur' }]
});

/** Save edit */
const save = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    loading.value = true;
    const saveOrUpdate = isUpdate.value ? updateCPageTemplateApi : createCPageTemplateApi;
    saveOrUpdate({ ...form })
      .then((msg) => {
        loading.value = false;
        EleMessage.success(msg);
        updateModelValue(false);
        emit('done');
      })
      .catch((e) => {
        loading.value = false;
        EleMessage.error(e.message);
      });
  });
};

/** Update modelValue */
const updateModelValue = (value) => {
  emit('update:modelValue', value);
};

watch(
  () => props.modelValue,
  (modelValue) => {
    if (modelValue) {
      if (props.data) {
        assignFields({
          ...props.data
        });
        isUpdate.value = true;
      } else {
        isUpdate.value = false;
      }
    } else {
      resetFields();
      formRef.value?.clearValidate?.();
    }
  }
);
</script>
