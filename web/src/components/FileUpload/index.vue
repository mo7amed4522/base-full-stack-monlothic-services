<template>
  <CommonUpload v-model="files" :fileLimit="fileLimit" :readonly="readonly" :disabled="disabled" :preview="preview"
    :limit="limit" :multiple="multiple" :drag="drag" :accept="accept" :itemStyle="itemStyle" :buttonStyle="buttonStyle"
    :sortable="sortable" :imageProps="imageProps" :progressProps="progressProps" :previewProps="previewProps"
    :tools="tools" :listType="listType" :beforeUploadClick="beforeUploadClick" :beforeItemEdit="beforeItemEdit"
    :locale="locale" @itemClick="handleItemClick" @preview="handleItemPreview" />
</template>

<script setup>
import { ref, watch } from 'vue';
import CommonUpload from '../CommonUpload/index.vue';

defineOptions({ name: 'FileUpload' });

const props = defineProps({
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
  /** Whether to support clicking preview */
  preview: {
    type: Boolean,
    default: true
  },
  /** Maximum upload quantity */
  limit: Number,
  /** Whether to support multiple file selection */
  multiple: Boolean,
  /** Whether to enable drag and drop upload */
  drag: {
    type: Boolean,
    default: true
  },
  /** Accept upload file type */
  accept: {
    type: String,
    default: ''
  },
  /** Custom style */
  itemStyle: Object,
  /** Custom upload button style */
  buttonStyle: Object,
  /** Whether to enable drag and drop sort */
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
  listType: {
    type: String,
    default: 'file'
  },
  /** Upload button click hook */
  beforeUploadClick: Function,
  /** Modify button click hook */
  beforeItemEdit: Function,
  /** Internationalization */
  locale: Object
});

const emit = defineEmits(['update:modelValue', 'change']);

/** Upload component data */
const files = ref([]);

/** Whether it is an image */
const isImage = (url) => {
  const parts = url ? url.split('.') : [];
  const suffix =
    parts.length <= 1 || (parts[0] === '' && parts.length === 2)
      ? ''
      : parts.pop()?.toLowerCase?.();
  return suffix && ['png', 'jpg', 'jpeg', 'gif', 'svg'].includes(suffix);
};

/** Get upload component data address */
const getItemFileUrl = (item) => {
  return item.fileUrl;
};

/** Click event */
const handleItemClick = (item) => {
  handleItemPreview(item);
};

/** Preview event */
const handleItemPreview = (item) => {
  if (!item.url) {
    const fileUrl = getItemFileUrl(item);
    if (fileUrl) {
      window.open(fileUrl);
    } else if (item.file && !item.file.type.startsWith('image')) {
      window.open(URL.createObjectURL(item.file));
    }
  }
};

/** Update binding value */
const updateModelValue = (value) => {
  if (props.modelValue !== value) {
    emit('update:modelValue', value);
    emit('change', value);
  }
};

/** Clear upload component data */
const clearData = () => {
  files.value = [];
};

/** Whether all uploads are complete */
const isDone = () => {
  return !files.value.some((d) => d.status !== 'done');
};

/** Get upload data corresponding value */
const getDataValue = (filesValue) => {
  const result = [];
  filesValue.forEach((item) => {
    const fileUrl = getItemFileUrl(item);
    if (item.status === 'done' && fileUrl != null) {
      result.push({ url: fileUrl, name: item.name });
    }
  });
  return result;
};

/** Get upload data corresponding to value */
const getFilesValue = (dataValue, filesValue) => {
  if (!dataValue) {
    return filesValue.filter((item) => item.status !== 'done');
  }
  const result = [];
  if (props.limit === 1) {
    try {
      const temp = JSON.parse(dataValue);
      if (temp) {
        const url = temp.url ?? '';
        result.push({
          key: `0-${url}`,
          name: temp.name ?? '',
          url: isImage(url) ? url : void 0,
          fileUrl: url,
          status: 'done'
        });
        return result;
      }
    } catch (e) {
      console.error(e);
    }
    const temp = filesValue.find((item) => item.status !== 'done');
    if (temp != null) {
      result.push(temp);
    }
    return result;
  }
  try {
    const temps = JSON.parse(dataValue);
    filesValue.forEach((item) => {
      const fileUrl = getItemFileUrl(item);
      if (
        (fileUrl && temps.some((temp) => temp.url === fileUrl)) ||
        item.status !== 'done'
      ) {
        result.push(item);
      }
    });
    temps.forEach((temp, index) => {
      const url = temp?.url ?? '';
      if (url && !result.some((item) => getItemFileUrl(item) === url)) {
        const key = `${index}-${url}`;
        const oldIndex = result.findIndex((item) => item.key === key);
        if (oldIndex !== -1) {
          result.splice(oldIndex, 1);
        }
        result.push({
          key,
          name: temp.name ?? '',
          url: isImage(url) ? url : void 0,
          fileUrl: url,
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
    filesValue.forEach((item) => {
      if (item.status !== 'done') {
        result.push(item);
      }
    });
  }
  return result;
};

/** Whether upload data has changed */
const filesIsChanged = (newFiles, oldFiles) => {
  if (newFiles.length !== oldFiles.length) {
    return true;
  }
  for (let i = 0; i < newFiles.length; i++) {
    const newItem = newFiles[i];
    const oldItem = oldFiles[i];
    if (
      newItem.key !== oldItem.key ||
      newItem.name !== oldItem.name ||
      getItemFileUrl(newItem) !== getItemFileUrl(oldItem) ||
      newItem.status !== oldItem.status
    ) {
      return true;
    }
  }
  return false;
};

/** Synchronize binding value */
watch(
  files,
  (filesValue) => {
    const dataValue = getDataValue(filesValue);
    if (props.limit === 1) {
      updateModelValue(dataValue.length ? JSON.stringify(dataValue[0]) : '');
      return;
    }
    updateModelValue(dataValue.length ? JSON.stringify(dataValue) : '');
  },
  { deep: true }
);

watch(
  () => props.modelValue,
  (dataValue) => {
    const filesValue = getFilesValue(dataValue, files.value);
    if (filesIsChanged(filesValue, files.value)) {
      files.value = filesValue;
    }
  },
  { immediate: true }
);

defineExpose({ clearData, isDone });
</script>
