<template>
  <EleModal :form="true" :width="460" title="Add Group" :zIndex="baseIndex" :appendToBody="false"
    v-bind="modalProps || {}" v-model="visible" @open="handleOpen" @closed="handleClosed">
    <ElForm ref="formRef" :model="form" labelWidth="82px" @submit.prevent="">
      <ElFormItem label="Parent Group" prop="parentId">
        <ElTreeSelect v-model="form.parentId" :data="groupData.filter((d) => d.id !== 0 && d.id !== -1)" nodeKey="id"
          :props="{ label: 'name' }" :clearable="true" :checkStrictly="true" :defaultExpandAll="true"
          placeholder="Please select the parent group" :popperOptions="{ strategy: 'fixed' }" :teleported="false"
          class="ele-fluid">
          <template #default="{ data }">
            <img src="/ele-file-list/ic_file_folder.png" class="file-picker-tree-icon" />
            <span>{{ data.name }}</span>
          </template>
          <template v-if="form.parentId" #prefix>
            <img src="/ele-file-list/ic_file_folder.png" class="file-picker-tree-icon" />
          </template>
        </ElTreeSelect>
      </ElFormItem>
      <ElFormItem label="Group Name" prop="name" :rules="[
        {
          required: true,
          message: 'Please enter a group name',
          type: 'string',
          trigger: 'blur'
        }
      ]">
        <ElInput :maxlength="20" :clearable="true" v-model="form.name" placeholder="Please enter a group name" />
      </ElFormItem>
    </ElForm>
    <template #footer>
      <ElButton @click="handleCancel">Cancel</ElButton>
      <ElButton type="primary" :loading="loading" @click="save">
        Save
      </ElButton>
    </template>
  </EleModal>
</template>

<script setup>
import { ref } from 'vue';
import { useFormData } from '@/utils/use-form-data';
import { addUserFile } from '@/api/system/user-file';

const props = defineProps({
  /** Parent group */
  parentId: Number,
  /** Group data */
  groupData: Array,
  /** Modal parameters */
  modalProps: Object,
  /** Uniformly set the layer */
  baseIndex: Number,
  /** Message prompt component */
  messageIns: [Object, Function]
});

const emit = defineEmits(['done']);

/** Modal is open */
const visible = defineModel({ type: Boolean });

/** Form instance */
const formRef = ref(null);

/** Submission status */
const loading = ref(false);

/** Form data */
const [form, resetFields] = useFormData({
  name: '',
  parentId: void 0
});

/** Close modal */
const handleCancel = () => {
  visible.value = false;
};

/** Save edit */
const save = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    loading.value = true;
    addUserFile({ ...form, isDirectory: 1 })
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

/** Modal closed event */
const handleClosed = () => {
  resetFields();
  formRef.value?.clearValidate?.();
};

/** Modal open event */
const handleOpen = () => {
  if (props.parentId != null) {
    form.parentId = props.parentId;
  }
};
</script>
