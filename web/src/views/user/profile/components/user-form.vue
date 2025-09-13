<template>
  <el-form ref="formRef" :model="form" :rules="rules" label-width="100px"
    style="max-width: 580px; padding: 34px 16px 12px 0; box-sizing: border-box" @submit.prevent="">
    <el-form-item label="Nick name" prop="nickname">
      <el-input clearable :maxlength="20" v-model="form.nickname" placeholder="Please input nick name" />
    </el-form-item>
    <el-form-item label="Sex" prop="sex">
      <el-select clearable v-model="form.sex" placeholder="Please select sex" class="ele-fluid">
        <el-option value="1" label="Male" />
        <el-option value="2" label="Female" />
      </el-select>
    </el-form-item>
    <el-form-item label="Email" prop="email">
      <el-input clearable :maxlength="100" v-model="form.email" placeholder="Please input email" />
    </el-form-item>
    <el-form-item label="Personal introduction">
      <el-input type="textarea" :rows="4" :maxlength="200" v-model="form.introduction"
        placeholder="Please input personal introduction" />
    </el-form-item>
    <el-form-item label="Address">
      <el-input clearable :maxlength="200" v-model="form.address" placeholder="Please input address" />
    </el-form-item>
    <el-form-item label="Contact phone:">
      <div class="form-tell">
        <div class="form-tell-prefix">
          <el-input v-model="form.tellPre" :maxlength="4" />
        </div>
        <div class="form-tell-body">
          <el-input clearable :maxlength="10" v-model="form.tell" placeholder="Please input contact phone" />
        </div>
      </div>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" :loading="loading" @click="save">
        {{ loading ? 'Saving...' : 'Save changes' }}
      </el-button>
    </el-form-item>
  </el-form>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';
import { updateUserInfo } from '@/api/layout';

const props = defineProps({
  data: Object
});

const emit = defineEmits(['done']);

/** 提交状态 */
const loading = ref(false);

/** 表单实例 */
const formRef = ref(null);

/** 表单数据 */
const [form, _resetFields, assignFields] = useFormData({
  nickname: '',
  sex: void 0,
  email: '',
  introduction: '',
  address: '',
  tellPre: '0752',
  tell: ''
});

/** 表单验证规则 */
const rules = reactive({
  nickname: [
    {
      required: true,
      message: 'Please input nick name',
      type: 'string',
      trigger: 'blur'
    }
  ],
  sex: [
    {
      required: true,
      message: 'Please select sex',
      type: 'string',
      trigger: 'blur'
    }
  ],
  email: [
    {
      required: true,
      message: 'Please input email',
      type: 'string',
      trigger: 'blur'
    }
  ]
});

/** 保存更改 */
const save = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    loading.value = true;
    updateUserInfo(form)
      .then((data) => {
        loading.value = false;
        EleMessage.success('Save success');
        emit('done', data);
      })
      .catch((e) => {
        loading.value = false;
        EleMessage.error(e.message);
      });
  });
};

// 回显当前登录用户信息
assignFields({
  ...props.data,
  email: props.data.email || 'eleadmin@eclouds.com',
  tellPre: props.data.tellPre || '0752'
});
</script>

<style lang="scss" scoped>
.form-tell {
  display: flex;
  align-items: center;
  width: 100%;

  .form-tell-prefix {
    width: 65px;
  }

  .form-tell-body {
    flex: 1;
    padding-left: 12px;
    box-sizing: border-box;
  }
}
</style>
