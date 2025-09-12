<template>
  <el-dialog v-model="visible" :title="feedback ? t('feedbacks.editFeedback') : t('feedbacks.addFeedback')"
    width="600px" :close-on-click-modal="false">
    <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
      <el-form-item :label="t('feedbacks.title')" prop="title">
        <el-input v-model="form.title" :placeholder="t('feedbacks.titlePlaceholder')" />
      </el-form-item>
      <el-form-item :label="t('feedbacks.message')" prop="message">
        <el-input v-model="form.message" type="textarea" :rows="4" :placeholder="t('feedbacks.messagePlaceholder')" />
      </el-form-item>
      <el-form-item :label="t('feedbacks.type')" prop="type">
        <el-select v-model="form.type" :placeholder="t('feedbacks.selectType')">
          <el-option v-for="type in feedbackTypes" :key="type.value" :label="type.label" :value="type.value" />
        </el-select>
      </el-form-item>
      <el-form-item :label="t('feedbacks.status')" prop="status">
        <el-select v-model="form.status" :placeholder="t('feedbacks.selectStatus')">
          <el-option v-for="status in statusOptions" :key="status.value" :label="status.label" :value="status.value" />
        </el-select>
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="ele-form-actions">
        <el-button @click="handleCancel">
          <template #icon>
            <ele-icon :size="14">
              <close-outlined />
            </ele-icon>
          </template>
          {{ t('common.cancel') }}
        </el-button>
        <el-button type="primary" @click="handleSubmit" :loading="loading">
          <template #icon>
            <ele-icon :size="14">
              <check-outlined />
            </ele-icon>
          </template>
          {{ t('common.save') }}
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, computed, watch } from 'vue';
import { useI18n } from 'vue-i18n';
import { EleMessage } from 'ele-admin-plus/es';
import { CloseOutlined, CheckOutlined } from '@/components/icons';

const { t } = useI18n();

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  feedback: {
    type: Object,
    default: null
  }
});

const emit = defineEmits(['update:modelValue', 'refresh']);

const formRef = ref();
const loading = ref(false);

const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
});

const form = ref({
  title: '',
  message: '',
  type: '',
  status: 'unresolved'
});

const feedbackTypes = [
  { label: t('feedbacks.types.bug'), value: 'bug' },
  { label: t('feedbacks.types.feature'), value: 'feature' },
  { label: t('feedbacks.types.improvement'), value: 'improvement' },
  { label: t('feedbacks.types.other'), value: 'other' }
];

const statusOptions = [
  { label: t('feedbacks.unresolved'), value: 'unresolved' },
  { label: t('feedbacks.resolved'), value: 'resolved' }
];

const rules = {
  title: [
    { required: true, message: t('feedbacks.titleRequired'), trigger: 'blur' }
  ],
  message: [
    { required: true, message: t('feedbacks.messageRequired'), trigger: 'blur' }
  ],
  type: [
    { required: true, message: t('feedbacks.typeRequired'), trigger: 'change' }
  ]
};

watch(
  () => props.feedback,
  (newFeedback) => {
    if (newFeedback) {
      form.value = {
        title: newFeedback.title || '',
        message: newFeedback.message || '',
        type: newFeedback.type || '',
        status: newFeedback.status || 'unresolved'
      };
    } else {
      form.value = {
        title: '',
        message: '',
        type: '',
        status: 'unresolved'
      };
    }
  },
  { immediate: true }
);

const handleCancel = () => {
  visible.value = false;
  formRef.value?.resetFields();
};

const handleSubmit = async () => {
  try {
    await formRef.value?.validate();
    loading.value = true;
    await new Promise(resolve => setTimeout(resolve, 1000));

    EleMessage.success(props.feedback ? t('feedbacks.updateSuccess') : t('feedbacks.createSuccess'));
    visible.value = false;
    emit('refresh');
  } catch (error) {
    console.error('Validation failed:', error);
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
.el-form {
  padding: 20px 0;
}

.ele-form-actions {
  display: flex;
  gap: 8px;
  justify-content: flex-end;
}
</style>
