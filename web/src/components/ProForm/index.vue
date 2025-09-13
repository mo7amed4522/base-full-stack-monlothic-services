<!-- 高级表单 -->
<template>
  <ElForm
    ref="formRef"
    :model="model"
    :labelPosition="labelPosition"
    :labelWidth="labelWidth"
    :statusIcon="statusIcon"
    :validateOnRuleChange="validateOnRuleChange"
    :size="size"
    :disabled="disabled"
    :scrollToError="scrollToError"
    @submit.prevent=""
  >
    <slot name="topExtra"></slot>
    <ProFormContent :model="model" :rules="rules" :items="items" :grid="grid" :rowProps="rowProps" :contentExtraColProps="footerColProps" :searchExpand="searchExpand" @updateItemValue="updateValue">
      <template v-for="name in Object.keys($slots).filter((k) => !slotExcludes.includes(k))" #[name]="slotProps">
        <slot :name="name" v-bind="slotProps || {}"></slot>
      </template>
      <template v-if="footer" #contentExtra>
        <ProFormFooter
          :footerProps="footerProps"
          :footerSlots="footerSlots"
          :footerStyle="footerStyle"
          :submitText="submitText"
          :resetText="resetText"
          :submitButtonProps="submitButtonProps"
          :resetButtonProps="resetButtonProps"
          :showSearchExpand="showSearchExpand"
          :searchExpand="searchExpand"
          :searchExpandButtonProps="searchExpandButtonProps"
          :searchExpandText="searchExpandText"
          :searchShrinkText="searchShrinkText"
          @submit="submit"
          @reset="reset"
          @updateSearchExpand="updateSearchExpand"
        >
          <template v-for="name in Object.keys($slots).filter((k) => !fSlotExcludes.includes(k))" #[name]="slotProps">
            <slot :name="name" v-bind="slotProps || {}"></slot>
          </template>
        </ProFormFooter>
      </template>
    </ProFormContent>
    <slot name="bottomExtra"></slot>
  </ElForm>
</template>

<script setup>
  import { ref, nextTick } from 'vue';
  import ProFormContent from './components/pro-form-content.vue';
  import ProFormFooter from './components/pro-form-footer.vue';
  const fSlotExcludes = ['default', 'topExtra', 'bottomExtra', 'contentExtra'];
  const slotExcludes = [...fSlotExcludes, 'footer', 'footerExtra'];

  defineOptions({ name: 'ProForm' });

  const props = defineProps({
    /** Form Data */
    model: {
      type: Object,
      required: true
    },
    /** Validation Rules */
    rules: Object,
    /** Positioning of form field labels */
    labelPosition: String,
    /** Label length */
    labelWidth: {
      type: [String, Number],
      default: '80px'
    },
    /** Whether to display the verification result icon */
    statusIcon: Boolean,
    /** Whether to trigger a validation immediately after the rules property changes */
    validateOnRuleChange: Boolean,
    /** size */
    size: String,
    /** Disable */
    disabled: Boolean,
    /** Scroll to the first error field when validation fails */
    scrollToError: Boolean,
    /** Form Items */
    items: {
      type: Array,
      required: true
    },
    /** Whether to grid layout */
    grid: [Boolean, Object],
    /** ElRow Properties */
    rowProps: Object,
    /** Is a bottom bar required? */
    footer: Boolean,
    /** Bottom bar ElFormItem properties */
    footerProps: Object,
    /** Bottom bar ElFormItem slot */
    footerSlots: Object,
    /** Bottom bar ElCol properties */
    footerColProps: {
      type: Object,
      default: () => {
        return { span: 24 };
      }
    },
    /** Bottom bar style */
    footerStyle: Object,
    /** Submit button text */
    submitText: {
      type: String,
      default: 'submit'
    },
    /** Reset button text */
    resetText: {
      type: String,
      default: 'Reset'
    },
    /** Submit Button Properties */
    submitButtonProps: Object,
    /** Reset Button Properties */
    resetButtonProps: Object,
    /** Whether to display the form expand and collapse buttons in the bottom bar */
    showSearchExpand: Boolean,
    /** Expand and collapse button properties */
    searchExpandButtonProps: Object,
    /** Expand button text */
    searchExpandText: {
      type: String,
      default: 'Expand'
    },
    /** Collapse button text */
    searchShrinkText: {
      type: String,
      default: 'Close'
    }
  });

  const emit = defineEmits(['updateValue', 'submit', 'reset']);

  /** Search form expanded state */
  const searchExpand = defineModel('searchExpand', { type: Boolean });

  /** Form Example */
  const formRef = ref(null);

  /** Update value */
  const updateValue = (prop, value) => {
    emit('updateValue', prop, value);
  };

  /** Update the search form expanded state */
  const updateSearchExpand = (expand) => {
    searchExpand.value = expand;
  };

  /** submit */
  const submit = () => {
    formRef.value?.validate?.((valid) => {
      if (valid) {
        emit('submit', props.model);
      }
    });
  };

  /** Reset */
  const reset = () => {
    emit('reset');
    clearValidate();
    nextTick(() => {
      clearValidate();
      nextTick(() => {
        clearValidate();
      });
    });
  };

  /** Validation Form */
  const validate = (callback) => {
    if (!formRef.value) {
      throw new Error('formRef is null');
    }
    return formRef.value.validate(callback);
  };

  /** Validate a form field */
  const validateField = (props, callback) => {
    if (!formRef.value) {
      throw new Error('formRef is null');
    }
    return formRef.value.validateField(props, callback);
  };

  /** Reset Form */
  const resetFields = (props) => {
    if (!formRef.value) {
      throw new Error('formRef is null');
    }
    return formRef.value.resetFields(props);
  };

  /** Scroll to the specified field */
  const scrollToField = (prop) => {
    if (!formRef.value) {
      throw new Error('formRef is null');
    }
    return formRef.value.scrollToField(prop);
  };

  /** Clean up form validation information */
  const clearValidate = (props) => {
    if (!formRef.value) {
      throw new Error('formRef is null');
    }
    return formRef.value.clearValidate(props);
  };

  defineExpose({
    formRef,
    validate,
    validateField,
    resetFields,
    scrollToField,
    clearValidate
  });
</script>
