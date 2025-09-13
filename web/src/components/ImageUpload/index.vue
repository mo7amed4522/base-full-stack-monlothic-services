<template>
  <CommonUpload
    v-model="images"
    :fileLimit="fileLimit"
    :readonly="readonly"
    :disabled="disabled"
    :preview="preview"
    :limit="limit"
    :multiple="multiple"
    :drag="drag"
    :accept="accept"
    :itemStyle="itemStyle"
    :buttonStyle="buttonStyle"
    :sortable="sortable"
    :imageProps="imageProps"
    :progressProps="progressProps"
    :previewProps="previewProps"
    :tools="tools"
    :listType="listType"
    :beforeUploadClick="beforeUploadClick"
    :beforeItemEdit="beforeItemEdit"
    :locale="locale"
    :type="type"
    :name="name"
  />
</template>

<script setup>
  import { ref, watch } from 'vue';
  import CommonUpload from '../CommonUpload/index.vue';

  defineOptions({ name: 'ImageUpload' });

  const props = defineProps({
    type: {
      type: Number,
      default: 1
    },
    name: {
      type: String,
      default: 'icon'
    },
    /** Binding value */
    modelValue: String,
    /** File size limit, unit MB */
    fileLimit: {
      type: Number,
      default: 100
    },
    /** Whether it is read-only */
    readonly: Boolean,
    /** Whether it is disabled */
    disabled: Boolean,
    /** Whether to support clicking to preview */
    preview: {
      type: Boolean,
      default: true
    },
    /** Maximum upload quantity */
    limit: Number,
    /** Whether to support multiple file selection */
    multiple: Boolean,
    /** Whether to enable drag-and-drop upload */
    drag: {
      type: Boolean,
      default: true
    },
    /** Accepted file types for upload */
    accept: {
      type: String,
      default: 'image/*'
    },
    /** Custom style for each item */
    itemStyle: Object,
    /** Custom style for upload button */
    buttonStyle: Object,
    /** Whether to enable drag-and-drop upload */
    sortable: {
      type: [Boolean, Object],
      default: () => {
        return { forceFallback: true };
      }
    },
    /** Custom image properties */
    imageProps: Object,
    /** Custom progress bar properties */
    progressProps: Object,
    /** Custom image preview properties */
    previewProps: Object,
    /** Whether to enable bottom preview and modification operation buttons */
    tools: {
      type: Boolean,
      default: true
    },
    /** List display style */
    listType: String,
    /** Upload button click hook */
    beforeUploadClick: Function,
    /** Modify button click hook */
    beforeItemEdit: Function,
    /** Internationalization */
    locale: Object
  });

  const emit = defineEmits(['update:modelValue', 'change']);

  /** Upload component data */
  const images = ref([]);

  /** Update binding value */
  const updateModelValue = (value) => {
    const oldValue = props.modelValue ?? '';
    const newValue = value ?? '';
    if (oldValue !== newValue) {
      emit('update:modelValue', newValue);
      emit('change', value);
    }
  };

  /** Clear upload component data */
  const clearData = () => {
    images.value = [];
  };

  /** Check if all uploads are complete */
  const isDone = () => {
    return !images.value.some((item) => item.status !== 'done');
  };

  /** Get upload data corresponding value */
  const getDataValue = (imagesValue) => {
    const urls = [];
    imagesValue.forEach((item) => {
      if (item.status === 'done' && (item.originalUrl || item.url) != null) {
        // The original URL is used first, if not, the processed URL is used
        urls.push(item.originalUrl || item.url);
      }
    });
    return urls;
  };

  /** Process the image URL to ensure it can be accessed correctly */
  const processImageUrl = (url) => {
    if (!url) return url;

    // If it is a complete URL (including http/https), return it directly
    if (url.startsWith('http://') || url.startsWith('https://')) {
      return url;
    }

    // If it is a relative path, convert it to an API path
    if (url.startsWith('uploads/')) {
      return `/api/${url}`;
    }

    return url;
  };

  /** Get upload data corresponding to the value */
  const getImagesValue = (dataValue, imagesValue) => {
    if (!dataValue) {
      return imagesValue.filter((item) => item.status !== 'done');
    }
    const result = [];
    if (props.limit === 1) {
      result.push({
        key: `0-${dataValue}`,
        originalUrl: dataValue, // Storing the original URL
        url: processImageUrl(dataValue), // Store the processed URL for display
        status: 'done'
      });
      return result;
    }
    try {
      const urls = JSON.parse(dataValue);
      imagesValue.forEach((item) => {
        if ((item.originalUrl && urls.includes(item.originalUrl)) || (item.url && urls.includes(item.url)) || item.status !== 'done') {
          result.push(item);
        }
      });
      urls.forEach((url, index) => {
        if (url && !result.some((item) => item.originalUrl === url || item.url === url)) {
          const key = `${index}-${url}`;
          const oldIndex = result.findIndex((item) => item.key === key);
          if (oldIndex !== -1) {
            result.splice(oldIndex, 1);
          }
          result.push({
            key,
            originalUrl: url, // Storing the original URL
            url: processImageUrl(url), // Store the processed URL for display
            status: 'done'
          });
        }
      });
      if (props.limit != null && result.length > props.limit) {
        const temp = result.filter((item) => item.status !== 'done');
        for (const t of temp) {
          const i = result.indexOf(t);
          result.splice(i, 1);
          if (result.length <= props.limit) {
            break;
          }
        }
      }
    } catch (e) {
      console.error(e);
      imagesValue.forEach((item) => {
        if (item.status !== 'done') {
          result.push(item);
        }
      });
    }
    return result;
  };

  /** Whether upload data has changed */
  const imagesIsChanged = (newImages, oldImages) => {
    if (newImages.length !== oldImages.length) {
      return true;
    }
    for (let i = 0; i < newImages.length; i++) {
      const newItem = newImages[i];
      const oldItem = oldImages[i];
      if (newItem.key !== oldItem.key || (newItem.originalUrl || newItem.url) !== (oldItem.originalUrl || oldItem.url) || newItem.status !== oldItem.status) {
        return true;
      }
    }
    return false;
  };

  /** Synchronize binding value */
  watch(
    images,
    (imagesValue) => {
      const dataValue = getDataValue(imagesValue);
      if (props.limit === 1) {
        updateModelValue(dataValue.join());
        return;
      }
      updateModelValue(dataValue.length ? JSON.stringify(dataValue) : '');
    },
    { deep: true }
  );

  watch(
    () => props.modelValue,
    (dataValue) => {
      const imagesValue = getImagesValue(dataValue, images.value);
      if (imagesIsChanged(imagesValue, images.value)) {
        images.value = imagesValue;
      }
    },
    { immediate: true }
  );

  defineExpose({ clearData, isDone });
</script>
