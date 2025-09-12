<template>
  <ele-page>
    <el-result title="500">
      <template #icon>
        <div style="width: 250px; height: 295px; margin: 20px 0 10px 0">
          <icon-svg />
        </div>
      </template>
      <template #sub-title>
        <ele-text type="placeholder">Sorry, there was a server error.</ele-text>
      </template>
      <template #extra>
        <router-link to="/" style="display: inline-flex; text-decoration: none">
          <el-button type="primary">Return to Home</el-button>
        </router-link>
      </template>
    </el-result>
  </ele-page>
</template>

<script setup>
import IconSvg from './components/icon-svg.vue';
</script>

<script>
export default {
  name: 'Exception500'
};
</script>
