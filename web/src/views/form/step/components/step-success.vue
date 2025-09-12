<template>
  <el-result icon="success" title="Transfer success">
    <template #sub-title>
      <ele-text type="placeholder">Expected to arrive within two hours</ele-text>
    </template>
    <template #extra>
      <div style="margin-bottom: 24px">
        <el-button type="primary" @click="back">Transfer again</el-button>
        <el-button>View bill</el-button>
      </div>
      <el-descriptions :border="true" :column="1" size="large" class="detail-table"
        style="width: 680px; max-width: 100%">
        <el-descriptions-item label="Payment account">
          <div>{{ data.account }}</div>
        </el-descriptions-item>
        <el-descriptions-item label="Receiving account">
          <div>{{ data.receiver }}</div>
        </el-descriptions-item>
        <el-descriptions-item label="Payee name">
          <div>{{ data.name }}</div>
        </el-descriptions-item>
        <el-descriptions-item label="Transfer amount">
          <div style="display: flex; align-items: flex-end">
            <ele-text size="xl" style="line-height: 1">
              {{ data.amount }}&nbsp;
            </ele-text>
            <div style="line-height: 1.15">&nbsp;元</div>
          </div>
        </el-descriptions-item>
      </el-descriptions>
    </template>
  </el-result>
</template>

<script setup>
defineProps({
  data: Object
});

const emit = defineEmits(['back']);

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
