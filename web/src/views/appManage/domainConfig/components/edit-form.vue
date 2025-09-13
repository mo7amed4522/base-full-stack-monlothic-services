<template>
  <ele-modal form :width="550" :model-value="modelValue" :title="isUpdate ? 'Domain configuration' : 'Add domain configuration'" @update:modelValue="updateModelValue">
    <el-form ref="formRef" :model="form" :rules="rules" label-width="120px" @submit.prevent="">
      <el-form-item label="Domain name" prop="domain">
        <el-input v-model="form.domain" placeholder="Please enter the domain name" clearable />
      </el-form-item>
      <el-form-item label="Resolve IP" prop="resolveIp">
        <el-input v-model="form.resolveIp" placeholder="Please enter the resolve IP" clearable />
      </el-form-item>
      <el-form-item label="Access type" prop="accessType">
        <el-select v-model="form.accessType" placeholder="Please select the access type">
          <el-option label="Domain hosting" :value="1" />
        </el-select>
      </el-form-item>
      <el-form-item label="Status" prop="status">
        <el-select v-model="form.status" placeholder="Please select the status">
          <el-option label="Matching" :value="1" />
          <el-option label="Correct match" :value="2" />
          <el-option label="Match failed" :value="3" />
        </el-select>
      </el-form-item>
      <el-form-item label="Visible" prop="visible">
        <el-radio-group v-model="form.visible">
          <el-radio :value="1">Visible</el-radio>
          <el-radio :value="-1">Hidden</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="Remark" prop="remark">
        <el-input v-model="form.remark" type="textarea" placeholder="Please enter the remark" clearable />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="updateModelValue(false)">Cancel</el-button>
      <el-button type="primary" :loading="loading" @click="save">keep</el-button>
    </template>
  </ele-modal>
</template>

<script setup>
  import { ref, reactive, watch } from 'vue';
  import { EleMessage } from 'ele-admin-plus/es';
  import { useFormData } from '@/utils/use-form-data';
  import { createCDominManageApi } from '@/api/app';

  const emit = defineEmits(['done', 'update:modelValue']);

  const props = defineProps({
    /** Is the pop-up window open? */
    modelValue: Boolean,
    /** Modify the display data */
    data: Object
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
    domain: '',
    status: 1,
    visible: 1,
    accessType: 1,
    resolveIp: '',
    domainSource: 1,
    remark: ''
  });

  /** Form validation rules */
  const rules = reactive({
    domain: [{ required: true, message: 'Please enter the domain name', trigger: 'blur' }],
    resolveIp: [{ required: true, message: 'Please enter the resolve IP', trigger: 'blur' }],
    accessType: [{ required: true, message: 'Please select the access type', trigger: 'blur' }],
    status: [{ required: true, message: 'Please select the status', trigger: 'blur' }],
    visible: [{ required: true, message: 'Please select the visible status', trigger: 'blur' }],
    remark: [{ required: true, message: 'Please enter the remark', trigger: 'blur' }]
  });

  /** Save edit */
  const save = () => {
    formRef.value?.validate?.((valid) => {
      if (!valid) {
        return;
      }
      loading.value = true;
      const saveOrUpdate = isUpdate.value ? createCDominManageApi : createCDominManageApi;
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
