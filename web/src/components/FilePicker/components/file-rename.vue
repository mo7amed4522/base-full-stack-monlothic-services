<template>
  <EleModal :form="true" :width="460" title="Rename" :zIndex="baseIndex" :appendToBody="false" v-bind="modalProps || {}"
    v-model="visible" @open="handleOpen" @closed="handleClosed">
    <ElForm ref="formRef" :model="form" labelWidth="82px" @submit.prevent="">
      <ElFormItem :label="`${data && data.isDirectory ? 'Group' : 'File'} Name`" prop="name" :rules="[
        {
          required: true,
          message: `Please enter${data && data.isDirectory ? 'Group' : 'File'} Name`,
          type: 'string',
          trigger: 'blur'
        }
      ]">
        <ElInput :maxlength="20" :clearable="true" v-model="form.name"
          :placeholder="`Please enter ${data && data.isDirectory ? 'Group' : 'File'} name`" />
      </ElFormItem>
    </ElForm>
    <template #footer>
      <ElButton @click="handleCancel">Cancel</ElButton>
      <ElButton type="primary" :loading="loading" @click="save">
        keep
      </ElButton>
    </template>
  </EleModal>
</template>

<script setup>
import { ref } from 'vue';
import { useFormData } from '@/utils/use-form-data';
import { updateUserFile } from '@/api/system/user-file';

const props = defineProps({
  /** File data */
  data: Object,
  /** Modal parameters */
  modalProps: Object,
  /** Base index */
  baseIndex: Number,
  /** Message instance */
  messageIns: [Object, Function]
});

const emit = defineEmits(['done']);

/** Modal visibility */
const visible = defineModel({ type: Boolean });

/** Form instance */
const formRef = ref(null);

/** Submit status */
const loading = ref(false);

/** Form data */
const [form, resetFields] = useFormData({ name: '' });

/** Modal close event */
const handleCancel = () => {
  visible.value = false;
};

/** Save event */
const save = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    loading.value = true;
    updateUserFile({ ...form, id: props.data?.id })
      .then((msg) => {
        loading.value = false;
        props.messageIns?.success?.(msg);
        handleCancel();
        emit('done');
      })
      .catch((e) => {
        loading.value = false;
        props.messageIns?.error?.(e.message);
      });
  });
};

/** Modal close event */
const handleClosed = () => {
  resetFields();
  formRef.value?.clearValidate?.();
};

/** Modal open event */
const handleOpen = () => {
  if (props.data) {
    form.name = props.data.name;
  }
};
</script>
