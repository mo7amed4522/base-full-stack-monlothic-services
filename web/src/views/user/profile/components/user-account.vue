<template>
  <div class="list-wrapper">
    <div class="list-item">
      <div class="list-item-body">
        <ele-text size="md">Security phone</ele-text>
        <ele-text type="placeholder">Binded phone: {{ data.phone }}</ele-text>
      </div>
      <el-link type="primary" :underline="false">Go to modify</el-link>
    </div>
    <el-divider style="margin: 0; opacity: 0.6" />
    <div class="list-item">
      <div class="list-item-body">
        <ele-text size="md">Security email</ele-text>
        <ele-text type="placeholder">
          Binded email: eleadmin@eclouds.com
        </ele-text>
      </div>
      <el-link type="primary" :underline="false">Go to modify</el-link>
    </div>
    <el-divider style="margin: 0; opacity: 0.6" />
    <div class="list-item">
      <div class="list-item-body">
        <ele-text size="md">Security question</ele-text>
        <ele-text type="placeholder">Not set security question</ele-text>
      </div>
      <el-link type="primary" :underline="false">Go to set</el-link>
    </div>
    <el-divider style="margin: 0; opacity: 0.6" />
    <div class="list-item">
      <el-icon class="list-item-icon">
        <QqFilled />
      </el-icon>
      <div class="list-item-body">
        <ele-text size="md">Security QQ</ele-text>
        <ele-text type="placeholder">Not binded QQ account</ele-text>
      </div>
      <el-link type="primary" :underline="false">Go to bind</el-link>
    </div>
    <el-divider style="margin: 0; opacity: 0.6" />
    <div class="list-item">
      <el-icon class="list-item-icon" style="background: #4daf29">
        <WechatFilled />
      </el-icon>
      <div class="list-item-body">
        <ele-text size="md">Security Wechat</ele-text>
        <ele-text type="placeholder">Not binded Wechat account</ele-text>
      </div>
      <el-link type="primary" :underline="false">Go to bind</el-link>
    </div>
    <el-divider style="margin: 0; opacity: 0.6" />
    <div class="list-item">
      <el-icon class="list-item-icon" style="background: #1476fe">
        <AlipayFilled />
      </el-icon>
      <div class="list-item-body">
        <ele-text size="md">Security Alipay</ele-text>
        <ele-text type="placeholder">Currently not bound to Alipay account</ele-text>
      </div>
      <el-link type="primary" :underline="false">Go to bind</el-link>
    </div>
  </div>
</template>

<script setup>
import { QqFilled, WechatFilled, AlipayFilled } from '@/components/icons';

defineProps({
  data: Object
});
</script>

<style lang="scss" scoped>
.list-wrapper {
  padding: 6px 14px;
  box-sizing: border-box;

  .list-item {
    display: flex;
    align-items: center;
    box-sizing: border-box;
    padding: 12px 10px;

    .list-item-body {
      flex: 1;

      .ele-text+.ele-text {
        margin-top: 2px;
      }
    }

    .list-item-icon {
      color: #fff;
      padding: 10px;
      font-size: 24px;
      border-radius: 50%;
      background: #3492ed;
      margin-right: 12px;
    }
  }
}
</style>
