<!-- Checkbox group -->
<template>
  <div class="print-checkbox-group">
    <div v-for="item in items" :key="item.value" :class="[
      'print-checkbox',
      { 'is-checked': modelValue.includes(item.value) }
    ]" @click="onItemClick(item)">
      <div class="print-checkbox-icon"></div>
      <div>{{ item.label }}</div>
    </div>
  </div>
</template>

<script setup>
const emit = defineEmits(['update:modelValue']);

const props = defineProps({
  /** Checkbox group value */
  modelValue: Array,
  /** Checkbox group options */
  items: Array
});

/** Update value */
const updateModelValue = (value) => {
  emit('update:modelValue', value);
};

/** Checkbox item click event */
const onItemClick = (item) => {
  // Uncheck
  if (props.modelValue.includes(item.value)) {
    const temp = [...props.modelValue];
    temp.splice(temp.indexOf(item.value), 1);
    updateModelValue(temp);
    return;
  }
  // Check
  updateModelValue([...props.modelValue, item.value]);
};
</script>

<style lang="scss" scoped>
.print-checkbox-group {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
}

.print-checkbox {
  display: flex;
  align-items: center;
  color: #000;
  font-size: 15px;
  margin-right: 24px;
  cursor: pointer;

  .print-checkbox-icon {
    width: 15px;
    height: 15px;
    flex-shrink: 0;
    border: 1px solid #000;
    margin-right: 8px;
  }

  &.is-checked .print-checkbox-icon::before {
    content: '';
    width: 11px;
    height: 5px;
    display: inline-block;
    border: 2px solid transparent;
    border-left-color: #000;
    border-bottom-color: #000;
    transform: rotate(-45deg) translate(6px, -5px);
  }
}
</style>
