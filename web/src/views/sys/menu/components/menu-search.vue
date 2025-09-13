<!-- 搜索表单 -->
<template>
  <ele-card :body-style="{ paddingBottom: '2px' }">
    <el-form label-width="72px" @keyup.enter="search" @submit.prevent="">
      <el-row :gutter="8">
        <el-col :lg="6" :md="12" :sm="12" :xs="24">
          <el-form-item label="Menu name">
            <el-input clearable v-model.trim="form.title" placeholder="Please input" />
          </el-form-item>
        </el-col>
        <el-col :lg="6" :md="12" :sm="12" :xs="24">
          <el-form-item label="Menu address">
            <el-input clearable v-model.trim="form.path" placeholder="Please input" />
          </el-form-item>
        </el-col>
        <el-col :lg="6" :md="12" :sm="12" :xs="24">
          <el-form-item label="Permission identifier">
            <el-input clearable v-model.trim="form.authority" placeholder="Please input" />
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
  </ele-card>
</template>

<script setup>
import { useFormData } from '@/utils/use-form-data';

const emit = defineEmits(['search']);

/** 表单数据 */
const [form, resetFields] = useFormData({
  title: '',
  path: '',
  authority: ''
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
</script>
