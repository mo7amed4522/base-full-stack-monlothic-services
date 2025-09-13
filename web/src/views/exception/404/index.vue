<template>
  <ele-page>
    <el-result title="404">
      <template #icon>
        <div style="width: 250px; height: 295px; margin: 20px 0 10px 0">
          <icon-svg />
        </div>
      </template>
      <template #sub-title>
        <ele-text type="placeholder">Sorry, the page you visited does not exist.</ele-text>
      </template>
      <template #extra>
        <router-link to="/" style="display: inline-flex; text-decoration: none">
          <el-button type="primary">Return to home</el-button>
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
  name: 'Exception404'
};
</script>
