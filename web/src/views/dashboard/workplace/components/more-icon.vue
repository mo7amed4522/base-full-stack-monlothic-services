<template>
  <ele-dropdown :items="[
    {
      title: 'Edit',
      command: 'edit',
      icon: EditOutlined
    },
    {
      title: 'Delete',
      command: 'remove',
      icon: DeleteOutlined,
      danger: true
    }
  ]" :icon-props="{ size: 15 }" placement="bottom-end" style="cursor: pointer" :popper-options="{
      modifiers: [{ name: 'offset', options: { offset: [12, 12] } }]
    }" @command="onCommand">
    <ele-text type="placeholder" :icon="MoreOutlined" style="transform: scale(1.1); outline: none" />
  </ele-dropdown>
</template>

<script setup>
import {
  MoreOutlined,
  EditOutlined,
  DeleteOutlined
} from '@/components/icons';

const emit = defineEmits(['command']);

const onCommand = (command) => {
  emit('command', command);
};
</script>
