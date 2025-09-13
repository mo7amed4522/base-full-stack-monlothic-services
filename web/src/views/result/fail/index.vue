<template>
  <ele-page :multi-card="false">
    <ele-card :body-style="{ paddingLeft: '12px', paddingRight: '12px' }">
      <el-result icon="error" title="提交失败">
        <template #sub-title>
          <ele-text type="placeholder">
            Please check and modify the following information, then resubmit.
          </ele-text>
        </template>
      </el-result>
      <div style="text-align: center">
        <el-button type="primary">Return to modification</el-button>
        <el-button>Resubmit</el-button>
      </div>
      <div style="
          max-width: 760px;
          padding: 18px 24px;
          background: hsla(0, 0%, 60%, 0.1);
          border-radius: var(--el-border-radius-base);
          margin: 26px auto;
        ">
        <div>
          The following errors have been found in your submission:
        </div>
        <div style="margin-top: 12px">
          <ele-text type="danger" :icon="CloseCircleOutlined" tag="span" />
          <span>&nbsp;Your account has been frozen&nbsp;</span>
          <el-link type="primary" :underline="false" style="display: inline; vertical-align: 0">
            <span>Unfreeze now</span>
            <el-icon :size="14">
              <arrow-right />
            </el-icon>
          </el-link>
        </div>
        <div style="margin-top: 12px">
          <ele-text type="danger" :icon="CloseCircleOutlined" tag="span" />
          <span>&nbsp;Your account does not have the qualification to apply&nbsp;</span>
          <el-link type="primary" :underline="false" style="display: inline; vertical-align: 0">
            <span>Upgrade now</span>
            <el-icon :size="14">
              <arrow-right />
            </el-icon>
          </el-link>
        </div>
      </div>
    </ele-card>
  </ele-page>
</template>

<script setup>
import { CloseCircleOutlined, ArrowRight } from '@/components/icons';
</script>

<script>
export default {
  name: 'ResultFail'
};
</script>
