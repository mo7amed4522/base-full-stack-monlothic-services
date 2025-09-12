<template>
  <template v-if="!grid">
    <template v-for="item in items">
      <ProFormItem
        v-if="isShowItem(item, model, formItems || items, searchExpand)"
        :key="item.key ?? item.prop"
        :item="item"
        :model="model"
        :rules="rules"
        :formItems="formItems"
        :searchExpand="searchExpand"
        @updateItemValue="updateItemValue"
      >
        <template v-for="name in Object.keys($slots).filter((k) => !slotExcludes.includes(k))" #[name]="slotProps">
          <slot :name="name" v-bind="slotProps || {}"></slot>
        </template>
      </ProFormItem>
    </template>
    <slot name="contentExtra"></slot>
  </template>
  <ElRow v-else v-bind="rowProps || {}">
    <template v-for="item in items">
      <ElCol
        v-if="isShowItem(item, model, formItems || items, searchExpand)"
        :key="item.key ?? item.prop"
        v-bind="{
          ...(grid === true ? { span: 12 } : grid),
          ...(item.colProps || {})
        }"
      >
        <ProFormItem :item="item" :model="model" :rules="rules" :formItems="formItems" :searchExpand="searchExpand" @updateItemValue="updateItemValue">
          <template v-for="name in Object.keys($slots).filter((k) => !slotExcludes.includes(k))" #[name]="slotProps">
            <slot :name="name" v-bind="slotProps || {}"></slot>
          </template>
        </ProFormItem>
      </ElCol>
    </template>
    <ElCol v-if="$slots.contentExtra" v-bind="contentExtraColProps || {}">
      <slot name="contentExtra"></slot>
    </ElCol>
  </ElRow>
</template>

<script setup>
  import { isShowItem } from '../util';
  import ProFormItem from './pro-form-item.vue';
  const slotExcludes = ['default', 'contentExtra'];

  defineProps({
    /** Form data */
    model: {
      type: Object,
      required: true
    },
    /** Validation rules */
    rules: Object,
    /** Form items */
    items: {
      type: Array,
      required: true
    },
    /** Whether to use grid layout */
    grid: [Boolean, Object],
    /** ElRow attributes */
    rowProps: Object,
    /** Additional ElCol attributes */
    contentExtraColProps: Object,
    /** All form items */
    formItems: Array,
    /** Whether the search form is expanded */
    searchExpand: Boolean
  });

  const emit = defineEmits(['updateItemValue']);

  /** Update value */
  const updateItemValue = (prop, value) => {
    emit('updateItemValue', prop, value);
  };
</script>
