<template>
  <ele-modal form :width="460" title="New folder" :model-value="modelValue" @update:modelValue="updateModelValue">
    <el-form ref="formRef" :model="form" :rules="rules" label-width="100px" @submit.prevent="">
      <el-form-item label="Folder name" prop="name">
        <el-input clearable :maxlength="20" v-model="form.name" placeholder="Please enter the folder name" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="updateModelValue(false)">Cancel</el-button>
      <el-button type="primary" :loading="loading" @click="save">
        Save
      </el-button>
    </template>
  </ele-modal>
</template>

<script setup>
import { ref, reactive, watch } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';
import { addUserFile } from '@/api/system/user-file';

const emit = defineEmits(['done', 'update:modelValue']);

const props = defineProps({
  /** Is the pop-up window open? */
  modelValue: Boolean,
  /** Parent folder id */
  parentId: Number
});

/** Form instance */
const formRef = ref(null);

/** Submit status */
const loading = ref(false);

/** Form data */
const [form, resetFields] = useFormData({ name: '' });

/** Form validation rules */
const rules = reactive({
  name: [
    {
      required: true,
      message: 'Please enter the folder name',
      type: 'string',
      trigger: 'blur'
    }
  ]
});

/** Save edit */
const save = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    loading.value = true;
    addUserFile({
      ...form,
      parentId: props.parentId,
      isDirectory: 1
    })
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
    if (!modelValue) {
      resetFields();
      formRef.value?.clearValidate?.();
    }
  }
);
</script>
