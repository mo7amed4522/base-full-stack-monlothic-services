<template>
  <el-form label-width="60px" size="small" class="demo-search" @submit.prevent="">
    <el-row>
      <el-col :lg="8" :md="8" :sm="24" :xs="24">
        <el-form-item label="Username">
          <el-input clearable v-model.trim="form.username" placeholder="Username" />
        </el-form-item>
      </el-col>
      <el-col :lg="8" :md="8" :sm="24" :xs="24">
        <el-form-item label="Sex">
          <dict-data code="sex" v-model="form.sex" placeholder="Please select" :teleported="false" />
        </el-form-item>
      </el-col>
      <el-col :lg="8" :md="8" :sm="24" :xs="24">
        <el-form-item label-width="12px">
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
  sex: ''
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

<style lang="scss" scoped>
.demo-search :deep(.el-form-item) {
  margin-bottom: 8px;
}
</style>
