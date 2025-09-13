<template>
  <ElFormItem v-bind="footerProps || {}">
    <template v-for="name in Object.keys(footerSlots || {}).filter((k) => k !== 'footer' && !!(footerSlots && footerSlots[k] && $slots[footerSlots[k]]))" #[name]="slotProps">
      <slot :name="footerSlots?.[name]" v-bind="slotProps || {}"></slot>
    </template>
    <div
      :style="{
        flex: 1,
        display: 'flex',
        alignItems: 'center',
        ...(footerStyle || {})
      }"
    >
      <slot name="footer">
        <ElButton type="primary" v-bind="submitButtonProps || {}" @click="submit">
          {{ submitText }}
        </ElButton>
        <ElButton v-bind="resetButtonProps || {}" @click="reset">
          {{ resetText }}
        </ElButton>
      </slot>
      <ElLink v-if="showSearchExpand" type="primary" :underline="false" style="margin-left: 12px" v-bind="searchExpandButtonProps || {}" @click="toggleSearchExpand">
        <template v-if="searchExpand">
          <span>{{ searchShrinkText }}</span>
          <ElIcon style="vertical-align: -1px">
            <ArrowUp />
          </ElIcon>
        </template>
        <template v-else>
          <span>{{ searchExpandText }}</span>
          <ElIcon style="vertical-align: -2px">
            <ArrowDown />
          </ElIcon>
        </template>
      </ElLink>
      <slot name="footerExtra"></slot>
    </div>
  </ElFormItem>
</template>

<script setup>
  import { ArrowDown, ArrowUp } from '@/components/icons';

  const props = defineProps({
    /** Footer ElFormItem attributes */
    footerProps: Object,
    /** Footer ElFormItem slots */
    footerSlots: Object,
    /** Footer style */
    footerStyle: Object,
    /** Submit button text */
    submitText: String,
    /** Reset button text */
    resetText: String,
    /** Submit button attributes */
    submitButtonProps: Object,
    /** Reset button attributes */
    resetButtonProps: Object,
    /** Whether to display the expand/collapse button in the footer */
    showSearchExpand: Boolean,
    /** Whether the search form is expanded */
    searchExpand: Boolean,
    /** Expand and collapse button attributes */
    searchExpandButtonProps: Object,
    /** Expand button text */
    searchExpandText: String,
    /** Collapse button text */
    searchShrinkText: String
  });

  const emit = defineEmits(['submit', 'reset', 'updateSearchExpand']);

  /** Submit */
  const submit = () => {
    emit('submit');
  };

  /** Reset */
  const reset = () => {
    emit('reset');
  };

  /** Toggle search form expand status */
  const toggleSearchExpand = () => {
    emit('updateSearchExpand', !props.searchExpand);
  };
</script>
