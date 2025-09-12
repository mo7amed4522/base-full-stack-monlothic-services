<!-- 用户编辑弹窗 -->
<template>
  <ele-modal form :width="640" :model-value="modelValue" :title="isUpdate ? 'Modify user' : 'Create new user'"
    @update:modelValue="updateModelValue">
    <el-form ref="formRef" :model="form" :rules="rules" label-width="80px" @submit.prevent="">
      <el-row :gutter="16">
        <el-col :sm="12" :xs="24">
          <el-form-item label="Organization"></el-form-item>
          <el-form-item label="User account" prop="username">
            <el-input clearable :maxlength="20" v-model="form.username" placeholder="Please enter the user account" />
          </el-form-item>
          <el-form-item label="Username" prop="nickname">
            <el-input clearable :maxlength="20" v-model="form.nickname" placeholder="Please enter the username" />
          </el-form-item>
          <el-form-item label="Sex" prop="sex">
            <dict-data code="sex" v-model="form.sex" placeholder="Please select the sex" />
          </el-form-item>
          <el-form-item label="Role" prop="roles"></el-form-item>
        </el-col>
        <el-col :sm="12" :xs="24">
          <el-form-item label="Phone" prop="phone">
            <el-input clearable :maxlength="11" v-model="form.phone" placeholder="Please enter the phone number" />
          </el-form-item>
          <el-form-item label="Email" prop="email">
            <el-input clearable :maxlength="100" v-model="form.email" placeholder="Please enter the email" />
          </el-form-item>
          <el-form-item v-if="!isUpdate" label="Login password" prop="password">
            <el-input show-password type="password" :maxlength="20" v-model="form.password"
              placeholder="Please enter the login password" />
          </el-form-item>
          <el-form-item label="Personal introduction">
            <el-input type="textarea" :rows="3" :maxlength="200" v-model="form.introduction"
              placeholder="Please enter the personal introduction" />
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
    <template #footer>
      <el-button @click="updateModelValue(false)">Cancel</el-button>
      <el-button type="primary" :loading="loading" @click="save">Save</el-button>
    </template>
  </ele-modal>
</template>

<script setup>
import { ref, reactive, watch } from 'vue';
import { EleMessage, emailReg, phoneReg } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';

import { addUser, updateUser, checkExistence } from '@/api/system/user';

const emit = defineEmits(['done', 'update:modelValue']);

const props = defineProps({
  /** 弹窗是否打开 */
  modelValue: Boolean,
  /** 修改回显的数据 */
  data: Object,
  /** 添加时机构id */
  organizationId: Number
});

/** 是否是修改 */
const isUpdate = ref(false);

/** 提交状态 */
const loading = ref(false);

/** 表单实例 */
const formRef = ref(null);

/** 表单数据 */
const [form, resetFields, assignFields] = useFormData({
  userId: void 0,
  organizationId: void 0,
  username: '',
  nickname: '',
  sex: void 0,
  roles: [],
  email: '',
  phone: '',
  password: '',
  introduction: ''
});

/** 表单验证规则 */
const rules = reactive({
  username: [
    {
      required: true,
      message: 'Please enter your user account',
      type: 'string',
      trigger: 'blur'
    },
    {
      min: 4,
      message: 'The account length must be at least 4 characters',
      type: 'string',
      trigger: 'blur'
    },
    {
      type: 'string',
      trigger: 'blur',
      validator: (_rule, value, callback) => {
        checkExistence('username', value, props.data?.userId)
          .then(() => {
            callback(new Error('The account already exists'));
          })
          .catch(() => {
            callback();
          });
      }
    }
  ],
  nickname: [
    {
      required: true,
      message: 'Please enter your username',
      type: 'string',
      trigger: 'blur'
    }
  ],
  sex: [
    {
      required: true,
      message: 'Please select your sex',
      type: 'string',
      trigger: 'blur'
    }
  ],
  roles: [
    {
      required: true,
      message: 'Please select your role',
      type: 'array',
      trigger: 'blur'
    }
  ],
  email: [
    {
      pattern: emailReg,
      message: 'The email format is incorrect',
      type: 'string',
      trigger: 'blur'
    }
  ],
  password: [
    {
      required: true,
      message: 'Please enter your login password',
      type: 'string',
      trigger: 'blur'
    },
    {
      type: 'string',
      trigger: 'blur',
      validator: (_rule, value, callback) => {
        if (isUpdate.value || /^[\S]{5,18}$/.test(value)) {
          return callback();
        }
        callback(new Error('The password must be 5-18 non-blank characters'));
      }
    }
  ],
  phone: [
    {
      pattern: phoneReg,
      message: 'The phone format is incorrect',
      type: 'string',
      trigger: 'blur'
    }
  ]
});

/** 保存编辑 */
const save = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    loading.value = true;
    const saveOrUpdate = isUpdate.value ? updateUser : addUser;
    saveOrUpdate(form)
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
        assignFields({ ...props.data, password: '' });
        isUpdate.value = true;
      } else {
        form.organizationId = props.organizationId;
        isUpdate.value = false;
      }
    } else {
      resetFields();
      formRef.value?.clearValidate?.();
    }
  }
);
</script>
