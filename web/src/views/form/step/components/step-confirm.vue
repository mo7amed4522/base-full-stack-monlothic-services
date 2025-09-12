<template>
  <div>
    <ele-alert show-icon
      title="After transfer, the money will be directly into the other account and cannot be returned."
      style="margin-bottom: 18px" />
    <el-descriptions :border="true" :column="1" size="large" class="detail-table">
      <el-descriptions-item label="Payment Account">
        <div>{{ data.account }}</div>
      </el-descriptions-item>
      <el-descriptions-item label="Receiving Account">
        <div>{{ data.receiver }}</div>
      </el-descriptions-item>
      <el-descriptions-item label="Payee Name">
        <div>{{ data.name }}</div>
      </el-descriptions-item>
      <el-descriptions-item label="Transfer Amount">
        <div style="display: flex; align-items: flex-end">
          <ele-text size="xl" style="line-height: 1">
            {{ data.amount }}&nbsp;
          </ele-text>
          <div style="line-height: 1.15">&nbsp;CNY</div>
        </div>
      </el-descriptions-item>
      <el-descriptions-item label="Payment Password">
        <div style="max-width: 220px">
          <el-input type="password" v-model="form.password" placeholder="Please input payment password" show-password />
        </div>
      </el-descriptions-item>
    </el-descriptions>
    <div style="margin-top: 22px; padding-left: 140px">
      <el-button type="primary" :loading="loading" @click="submit">
        Next Step
      </el-button>
      <el-button @click="back">Previous Step</el-button>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';

defineProps({
  data: Object
});

const emit = defineEmits(['done', 'back']);

/** 提交状态 */
const loading = ref(false);

/** 表单数据 */
const form = reactive({
  password: '123456'
});

/** 下一步 */
const submit = () => {
  if (!form.password) {
    EleMessage.error('Please input payment password');
    return;
  }
  loading.value = true;
  setTimeout(() => {
    loading.value = false;
    emit('done');
  }, 300);
};

/** 上一步 */
const back = () => {
  emit('back');
};
</script>

<style lang="scss" scoped>
.detail-table :deep(.el-descriptions__label) {
  width: 120px;
  text-align: right;
  font-weight: normal;
}
</style>
