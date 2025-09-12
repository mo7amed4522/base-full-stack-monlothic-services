<!-- 搜索表单 -->
<template>
  <el-form label-width="72px" @keyup.enter="search" @submit.prevent="">
    <el-row :gutter="8">
      <el-col :lg="6" :md="12" :sm="12" :xs="24">
        <el-form-item label="User account">
          <el-input clearable v-model.trim="form.username" placeholder="Please enter" />
        </el-form-item>
      </el-col>
      <el-col :lg="6" :md="12" :sm="12" :xs="24">
        <el-form-item label="User name">
          <el-input clearable v-model.trim="form.nickname" placeholder="Please enter" />
        </el-form-item>
      </el-col>
      <el-col :lg="6" :md="12" :sm="12" :xs="24">
        <el-form-item label="User sex">
          <dict-data code="sex" v-model="form.sex" placeholder="Please select" />
        </el-form-item>
      </el-col>
      <el-col :lg="6" :md="12" :sm="12" :xs="24">
        <el-form-item label-width="16px">
          <el-button type="primary" @click="search">Search</el-button>
          <el-button @click="reset">Reset</el-button>
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>
</template>

<script setup>
import { useFormData } from '@/utils/use-form-data';

const emit = defineEmits(['search']);

/** 表单数据 */
const [form, resetFields] = useFormData({
  username: '',
  nickname: '',
  sex: void 0
});

/** 搜索 */
const search = () => {
  emit('search', { ...form });
};

/**  重置 */
const reset = () => {
  resetFields();
  search();
};

defineExpose({ resetFields });
</script>
