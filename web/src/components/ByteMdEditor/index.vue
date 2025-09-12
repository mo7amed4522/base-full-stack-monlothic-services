<!-- markdown editor -->
<template>
  <div ref="rootRef" class="ele-bytemd-wrap"></div>
</template>

<script setup>
  import { Editor } from 'bytemd';
  import 'bytemd/dist/index.min.css';
  import { onMounted, ref, watch } from 'vue';

  const props = defineProps({
    /** Binding Value */
    modelValue: {
      type: String,
      required: true
    },
    /** Editor Configuration */
    config: Object,
    /** Editor Height */
    height: String,
    /** Fullscreen Z-Index */
    fullIndex: {
      type: Number,
      default: 999
    }
  });

  const emit = defineEmits(['update:modelValue', 'change']);

  /** ByteMD Editor Instance */
  let editor = null;

  /** Root Node */
  const rootRef = ref(null);

  // Render Editor
  onMounted(() => {
    const ins = new Editor({
      target: rootRef.value,
      props: Object.assign({}, props.config, { value: props.modelValue })
    });
    ins.$on('change', (e) => {
      emit('update:modelValue', e.detail.value);
      emit('change', e.detail.value);
    });
    editor = ins;
  });

  // Update Configuration
  watch(
    () => props.config,
    (config) => {
      const option = Object.assign({}, config);
      Object.keys(option).forEach((key) => {
        if (typeof option[key] === 'undefined') {
          delete option[key];
        }
      });
      editor?.$set?.(option);
    },
    { deep: true }
  );

  watch(
    () => props.modelValue,
    (value) => {
      editor?.$set?.({ value });
    }
  );

  defineExpose({ editor });
</script>

<style lang="scss" scoped>
  .ele-bytemd-wrap {
    width: 100%;
    line-height: initial;
  }

  /* Editor Height */
  .ele-bytemd-wrap :deep(.bytemd) {
    height: v-bind(height);

    /* Fullscreen Z-Index */
    &.bytemd-fullscreen {
      z-index: v-bind(fullIndex);
    }

    /* Remove Default Max-width Limit */
    .CodeMirror .CodeMirror-lines {
      max-width: 100%;
    }

    pre.CodeMirror-line,
    pre.CodeMirror-line-like {
      padding: 0 24px;
    }

    .markdown-body {
      max-width: 100%;
      padding: 16px 24px;
    }

    /* Remove Github Icon */
    .bytemd-toolbar-right > .bytemd-toolbar-icon:last-child {
      display: none;
    }
  }
</style>
