<template>
  <ele-modal form :width="550" :model-value="modelValue" :title="isUpdate ? 'Edit template' : 'Add template'"
    @update:modelValue="updateModelValue">
    <el-form ref="formRef" :model="form" :rules="rules" label-width="120px" @submit.prevent="">
      <el-form-item label="Application" prop="appId">
        <el-select v-model="form.appId" :empty-values="[null, undefined]" placeholder="Select application"
          class="ele-fluid">
          <el-option v-for="item in appList" :key="item.id" :label="item.appName" :value="item.id" />
        </el-select>
      </el-form-item>
      <el-form-item label="Domain" prop="domainId">
        <el-select v-model="form.domainId" :empty-values="[null, undefined]" placeholder="Select domain"
          class="ele-fluid">
          <el-option v-for="item in domainList" :key="item.id" :label="item.domain" :value="item.id" />
        </el-select>
      </el-form-item>
      <el-form-item label="Template" prop="templateId">
        <el-select v-model="form.templateId" :empty-values="[null, undefined]" placeholder="Select template"
          class="ele-fluid">
          <el-option v-for="item in templateList" :key="item.id" :label="item.name" :value="item.id" />
        </el-select>
      </el-form-item>
      <el-form-item label="Link name" prop="name">
        <el-input v-model="form.name" placeholder="Input link name" clearable />
      </el-form-item>
      <el-form-item label="ios jump url" prop="iosRedirectUrl">
        <el-input v-model="form.iosRedirectUrl" placeholder="Input ios jump url" clearable />
      </el-form-item>
      <el-form-item label="Package download address" prop="url">
        <el-input v-model="form.url" placeholder="Input package download address" clearable />
      </el-form-item>
      <el-form-item label="Allow modify url" prop="isAllowEditUrl">
        <el-radio-group v-model="form.isAllowEditUrl">
          <el-radio :value="1" label="Allow"></el-radio>
          <el-radio :value="-1" label="Not allow"></el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="Remark" prop="remark">
        <el-input v-model="form.remark" type="textarea" placeholder="Input remark" clearable />
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

import { createCPromotionLinkApi, updateCPromotionLinkApi } from '@/api/app';

const emit = defineEmits(['done', 'update:modelValue']);

const props = defineProps({
  /** Is the pop-up window open? */
  modelValue: Boolean,
  /** Modify the back display data */
  data: Object,
  appList: {
    type: Array,
    default: () => []
  },
  domainList: {
    type: Array,
    default: () => []
  },
  templateList: {
    type: Array,
    default: () => []
  }
});

/** Is it a modification? */
const isUpdate = ref(false);

/** Submit status */
const loading = ref(false);

/** Form instance */
const formRef = ref(null);

/** Form data */
const [form, resetFields, assignFields] = useFormData({
  id: void 0,
  name: null,
  appId: null,
  templateId: null,
  url: null,
  isAllowEditUrl: 1,
  iosRedirectUrl: null,
  domainId: null,
  remark: null
});

/** Form validation rules */
const rules = reactive({
  name: [{ required: true, message: 'Please input link name', trigger: 'blur' }],
  appId: [{ required: true, message: 'Please select the application', trigger: 'change' }],
  templateId: [{ required: true, message: 'Please select the template', trigger: 'change' }],
  url: [{ required: true, message: 'Please input package download address', trigger: 'blur' }],
  isAllowEditUrl: [{ required: true, message: 'Please select whether to allow modification of url', trigger: 'blur' }],
  iosRedirectUrl: [{ required: true, message: 'Please input ios jump url', trigger: 'blur' }],
  domainId: [{ required: true, message: 'Please select the domain', trigger: 'change' }],
});

/** Save edit */
const save = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    loading.value = true;
    const saveOrUpdate = isUpdate.value ? updateCPromotionLinkApi : createCPromotionLinkApi;
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
