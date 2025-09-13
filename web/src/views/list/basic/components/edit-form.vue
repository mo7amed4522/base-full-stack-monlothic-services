<!-- 添加和修改的表单 -->
<template>
  <el-form ref="formRef" :model="form" :rules="rules" label-width="80px" style="max-width: 600px; margin: 10px auto"
    @submit.prevent="">
    <el-form-item label="User Account" prop="username">
      <el-input clearable :maxlength="20" v-model="form.username" placeholder="Please enter your user account" />
    </el-form-item>
    <el-form-item label="User Name" prop="nickname">
      <el-input clearable :maxlength="20" v-model="form.nickname" placeholder="Please enter your user name" />
    </el-form-item>
    <el-form-item label="Sex" prop="sex">
      <dict-data code="sex" v-model="form.sex" placeholder="Please select your sex" />
    </el-form-item>
    <el-form-item label="Role" prop="roles">
      <role-select v-model="form.roles" />
    </el-form-item>
    <el-form-item label="Email" prop="email">
      <el-input clearable :maxlength="100" v-model="form.email" placeholder="Please enter your email" />
    </el-form-item>
    <el-form-item label="Phone" prop="phone">
      <el-input clearable :maxlength="11" v-model="form.phone" placeholder="Please enter your phone number" />
    </el-form-item>
    <el-form-item label="Birthday" prop="birthday">
      <el-date-picker v-model="form.birthday" value-format="YYYY-MM-DD" placeholder="Please select your birthday"
        class="ele-fluid" />
    </el-form-item>
    <el-form-item v-if="!isUpdate" label="Password" prop="password">
      <el-input show-password type="password" :maxlength="20" v-model="form.password"
        placeholder="Please enter your login password" />
    </el-form-item>
    <el-form-item label="Introduction" prop="introduction">
      <el-input type="textarea" :rows="4" v-model="form.introduction" :maxlength="200"
        placeholder="Please enter your introduction" />
    </el-form-item>
    <el-form-item>
      <el-button @click="onClose">Close</el-button>
      <el-button type="primary" :loading="loading" @click="save">Save</el-button>
    </el-form-item>
  </el-form>
</template>

<script setup>
import { ref, reactive, watch } from 'vue';
import { useRouter } from 'vue-router';
import { EleMessage, emailReg, phoneReg } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';
import { usePageTab } from '@/utils/use-page-tab';
import { addUser, updateUser, checkExistence } from '@/api/system/user';

const props = defineProps({
  /** 修改回显的数据 */
  data: Object
});

const { push } = useRouter();
const { routeTabKey, removePageTab, reloadPageTab } = usePageTab();

/** 是否是修改 */
const isUpdate = ref(false);

/** 提交状态 */
const loading = ref(false);

/** 表单实例 */
const formRef = ref(null);

/** 表单数据 */
const [form, resetFields, assignFields] = useFormData({
  userId: void 0,
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
      type: 'string',
      trigger: 'blur',
      validator: (_rule, value, callback) => {
        if (!value) {
          return callback(new Error('Please enter your user account'));
        }
        checkExistence('username', value, props.data?.userId)
          .then(() => {
            callback(new Error('Account already exists'));
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
      message: 'Please enter your user name',
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
      message: 'Email format is incorrect',
      type: 'string',
      trigger: 'blur'
    }
  ],
  password: [
    {
      required: true,
      type: 'string',
      trigger: 'blur',
      validator: (_rule, value, callback) => {
        if (isUpdate.value || /^[\S]{5,18}$/.test(value)) {
          return callback();
        }
        callback(new Error('Password must be 5-18 non-blank characters'));
      }
    }
  ],
  phone: [
    {
      pattern: phoneReg,
      message: 'Phone number format is incorrect',
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
        onDone();
      })
      .catch((e) => {
        loading.value = false;
        EleMessage.error(e.message);
      });
  });
};

/** 关闭当前页面并跳转到列表页面 */
const onClose = () => {
  removePageTab({ key: routeTabKey });
  push('/list/basic');
};

/** 关闭当前页面并刷新列表页面 */
const onDone = () => {
  removePageTab({ key: routeTabKey });
  reloadPageTab({ fullPath: '/list/basic' });
};

watch(
  () => props.data,
  () => {
    if (props.data) {
      assignFields({
        ...props.data,
        password: ''
      });
      isUpdate.value = true;
    } else {
      isUpdate.value = false;
      resetFields();
      formRef.value?.clearValidate?.();
    }
  },
  { immediate: true }
);
</script>
