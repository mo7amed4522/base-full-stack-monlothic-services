<template>
  <ele-modal form :width="550" :model-value="modelValue" :title="isView ? 'View subdomain' : 'Add subdomain'"
    @update:modelValue="updateModelValue">
    <el-form ref="formRef" :model="form" :rules="rules" label-width="100px" @submit.prevent="">
      <div class="form-notice">
        <el-icon class="warning-icon">
          <QuestionCircleOutlined />
        </el-icon>
        <span class="notice-text">Note: Prefixes should not contain uppercase letters, as this may cause the domain to
          be rejected!</span>
      </div>
      <el-form-item label="Subdomain prefix" prop="subDomains" required>
        <div class="subdomain-form">
          <div class="subdomain-list">
            <div v-for="(item, index) in form.subDomains" :key="index" class="subdomain-item">
              <el-input v-model="form.subDomains[index]" placeholder="Please enter" :rules="subdomainRules"
                @blur="validateSubdomain(index)" :disabled="props.isView" />
              <div class="item-actions" v-if="!props.isView">
                <el-button v-if="form.subDomains.length < 3" type="primary" :icon="Plus" circle size="small"
                  @click="addSubdomain" title="Add" />
                <el-button v-if="form.subDomains.length > 1" type="danger" :icon="Delete" circle size="small"
                  @click="removeSubdomain(index)" title="Delete" />
              </div>
            </div>
          </div>

          <div v-if="form.subDomains.length === 0 && !isView" class="empty-state">
            <el-button type="primary" :icon="Plus" @click="addSubdomain">Add subdomain prefix</el-button>
          </div>

          <!-- Empty state in view mode -->
          <div v-if="form.subDomains.length === 0 && isView" class="empty-state">
            <span class="no-data">No subdomain prefix</span>
          </div>
        </div>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="updateModelValue(false)">{{ isView ? 'Close' : 'Cancel' }}</el-button>
      <el-button v-if="!isView" type="primary" :loading="loading" @click="save">Save</el-button>
    </template>
  </ele-modal>
</template>

<script setup>
import { ref, reactive, watch } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';
import { QuestionCircleOutlined } from '@/components/icons/index.js';
import { createCSubDominManageApi } from '@/api/app';
import { Plus, Delete } from '@element-plus/icons-vue';

const emit = defineEmits(['done', 'update:modelValue']);

const props = defineProps({
  /** Whether the popup is open */
  modelValue: Boolean,
  /** Modify the data to be echoed */
  data: Object,
  /** Whether it is view mode */
  isView: Boolean
});

/** Submission status */
const loading = ref(false);

/** Form instance */
const formRef = ref(null);

/** Form data */
const [form, resetFields, assignFields] = useFormData({
  cDominId: null,
  subDomains: []
});

/** Subdomain validation rules */
const subdomainRules = [
  { required: true, message: 'Please enter subdomain prefix', trigger: 'blur' },
  {
    pattern: /^[a-z0-9-]+$/,
    message: 'Subdomain prefix can only contain lowercase letters, numbers, and hyphens',
    trigger: 'blur'
  },
  {
    validator: (rule, value, callback) => {
      if (value && value.length < 2) {
        callback(new Error('Subdomain prefix must be at least 2 characters'));
      } else if (value && value.length > 63) {
        callback(new Error('Subdomain prefix cannot exceed 63 characters'));
      } else {
        callback();
      }
    },
    trigger: 'blur'
  }
];

/** Form validation rules */
const rules = reactive({
  subDomains: [
    {
      validator: (rule, value, callback) => {
        if (!value || value.length === 0) {
          callback(new Error('Please add at least one subdomain prefix'));
          return;
        }

        // Check for duplicates
        const uniqueValues = new Set(value.filter((item) => item && item.trim()));
        if (uniqueValues.size !== value.filter((item) => item && item.trim()).length) {
          callback(new Error('Subdomain prefix cannot be repeated'));
          return;
        }

        // Check each subdomain for validity
        for (let i = 0; i < value.length; i++) {
          const item = value[i];
          if (!item || !item.trim()) {
            callback(new Error(`Subdomain prefix in the ${i + 1} position cannot be empty`));
            return;
          }

          if (!/^[a-z0-9-]+$/.test(item)) {
            callback(new Error(`Subdomain prefix in the ${i + 1} position can only contain lowercase letters, numbers, and hyphens`));
            return;
          }

          if (item.length < 2) {
            callback(new Error(`Subdomain prefix in the ${i + 1} position must be at least 2 characters`));
            return;
          }

          if (item.length > 63) {
            callback(new Error(`Subdomain prefix in the ${i + 1} position cannot exceed 63 characters`));
            return;
          }
        }

        callback();
      },
      trigger: 'blur'
    }
  ]
});

/** Add subdomain */
const addSubdomain = () => {
  form.subDomains.push('');
};

/** Delete subdomain */
const removeSubdomain = (index) => {
  form.subDomains.splice(index, 1);
};

/** Validate single subdomain */
const validateSubdomain = (index) => {
  // View mode does not perform validation and conversion
  if (props.isView) {
    return;
  }

  const value = form.subDomains[index];
  if (!value || !value.trim()) {
    return;
  }

  // Convert to lowercase automatically
  if (value !== value.toLowerCase()) {
    form.subDomains[index] = value.toLowerCase();
  }
};

/** Save edit */
const save = () => {
  // View mode does not perform save operation
  if (props.isView) {
    return;
  }

  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }

    // Filter out empty values and convert to lowercase
    const cleanSubDomains = form.subDomains.filter((item) => item && item.trim()).map((item) => item.toLowerCase().trim());

    const submitData = {
      ...form,
      subDomains: cleanSubDomains
    };

    loading.value = true;
    createCSubDominManageApi(submitData)
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
      assignFields({
        cDominId: props.data.id,
        subDomains: props.data.subDomainList.length > 0 ? props.data.subDomainList.map((item) => item.subDomain) : ['']
      });
    } else {
      resetFields();
      formRef.value?.clearValidate?.();
    }
  }
);
</script>

<style scoped>
.subdomain-form {
  width: 100%;
}

.form-notice {
  display: flex;
  align-items: center;
  margin-bottom: 16px;
  padding: 8px 12px;
  background-color: #fff2e8;
  border: 1px solid #ffd591;
  border-radius: 4px;
}

.warning-icon {
  color: #fa8c16;
  margin-right: 8px;
  font-size: 14px;
}

.notice-text {
  color: #d46b08;
  font-size: 12px;
}

.subdomain-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.subdomain-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.subdomain-item .el-input {
  flex: 1;
}

.item-actions {
  display: flex;
  gap: 4px;
}

.empty-state {
  display: flex;
  justify-content: center;
  padding: 20px 0;
}

.no-data {
  color: #909399;
  font-size: 14px;
}
</style>
