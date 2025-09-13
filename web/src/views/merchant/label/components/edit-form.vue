<template>
  <ele-modal form :width="550" :model-value="modelValue" :title="isUpdate ? t('common.edit') : t('common.add')"
    @update:modelValue="updateModelValue">
    <el-form ref="formRef" :model="form" :rules="rules" label-width="120px" @submit.prevent="">
      <el-form-item label="{{ t('common.remark') }}" prop="remark">
        <el-input v-model="form.remark" placeholder="{{ t('common.remark') }}" clearable />
      </el-form-item>
      <el-form-item label="{{ t('common.status') }}" prop="valid">
        <el-switch v-model="form.valid" :active-value="1" :inactive-value="0" active-text="{{ t('common.enable') }}"
          inactive-text="{{ t('common.disable') }}" inline-prompt />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="updateModelValue(false)">{{ t('common.cancel') }}</el-button>
      <el-button type="primary" :loading="loading" @click="save">{{ t('common.save') }}</el-button>
    </template>
  </ele-modal>
</template>

<script setup>
import { ref, reactive, watch } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';
import { createMUserLabelApi, updateMUserLabelApi } from '@/api/merchant';

const emit = defineEmits(['done', 'update:modelValue']);

const props = defineProps({
  /** 弹窗是否打开 */
  modelValue: Boolean,
  /** 修改回显的数据 */
  data: Object
});

/** 是否是修改 */
const isUpdate = ref(false);

/** 提交状态 */
const loading = ref(false);

/** 表单实例 */
const formRef = ref(null);

/** 表单数据 */
const [form, resetFields, assignFields] = useFormData({
  id: void 0,
  remark: '',
  valid: 1
});

/** 表单验证规则 */
const rules = reactive({
  remark: [{ required: true, message: 'Please enter a label name', trigger: 'blur' }],
  valid: [{ required: true, message: 'Please select status', trigger: ['change', 'blur'] }]
});

/** 保存编辑 */
const save = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    loading.value = true;
    const saveOrUpdate = isUpdate.value ? updateMUserLabelApi : createMUserLabelApi;
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

/** 更新modelValue */
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
