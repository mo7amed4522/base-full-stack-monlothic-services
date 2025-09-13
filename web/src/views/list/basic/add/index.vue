<template>
  <div>
    <ele-card style="border-radius: 0">
      <ele-text type="heading" size="lg">Add user</ele-text>
      <ele-text type="placeholder" style="margin-top: 6px">
        Fill in the user's basic information and click the save button
      </ele-text>
    </ele-card>
    <ele-page>
      <ele-card>
        <edit-form />
      </ele-card>
    </ele-page>
  </div>
</template>

<script setup>
import EditForm from '../components/edit-form.vue';
</script>

<script>
export default {
  name: 'ListBasicAdd'
};
</script>
