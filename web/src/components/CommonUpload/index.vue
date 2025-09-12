<template>
  <EleUploadList v-model="images" :readonly="readonly" :disabled="disabled" :preview="preview" :limit="limit"
    :multiple="multiple" :drag="drag" :accept="accept" :itemStyle="itemStyle" :buttonStyle="buttonStyle"
    :sortable="sortable" :imageProps="imageProps" :progressProps="progressProps" :previewProps="previewProps"
    :tools="tools" :listType="listType" :beforeUploadClick="beforeUploadClick" :beforeItemEdit="beforeItemEdit"
    :locale="locale" @upload="handleItemUpload" @retry="(item) => handleItemUpload(item, true)"
    @remove="handleItemRemove" @editUpload="handleItemEditUpload" @itemClick="handleItemClick"
    @preview="handleItemPreview">
    <template v-for="name in Object.keys($slots)" #[name]="slotProps">
      <slot :name="name" v-bind="slotProps || {}"></slot>
    </template>
  </EleUploadList>
</template>

<script setup>
import { EleMessage } from 'ele-admin-plus/es';
import { uploadFile } from '@/api/system/file';

defineOptions({ name: 'CommonUpload' });

const props = defineProps({
  /** File size limit, in MB */
  fileLimit: {
    type: Number,
    default: 100
  },
  /** Whether it is read-only */
  readonly: Boolean,
  /** Whether it is disabled */
  disabled: Boolean,
  /** Whether to support previewing */
  preview: {
    type: Boolean,
    default: true
  },
  /** Maximum upload quantity */
  limit: Number,
  /** Whether to support multiple file uploads */
  multiple: Boolean,
  /** Whether to enable drag-and-drop upload */
  drag: {
    type: Boolean,
    default: true
  },
  /** Accepted file types for upload */
  accept: {
    type: String,
    default: ''
  },
  /** Custom styles for uploaded items */
  itemStyle: Object,
  /** Custom styles for the upload button */
  buttonStyle: Object,
  /** Whether to enable drag-and-drop upload */
  sortable: {
    type: [Boolean, Object],
    default: () => {
      return { forceFallback: true };
    }
  },
  /** Custom attributes for uploaded items */
  imageProps: Object,
  /** Custom attributes for progress bars */
  progressProps: Object,
  /** Custom attributes for image previews */
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

const emit = defineEmits(['itemClick', 'preview']);

/** Uploaded data */
const images = defineModel({
  type: Array,
  default: () => []
});

/** Verify the selected files */
const checkFile = (file) => {
  if (!file) {
    return;
  }
  if (props.accept === 'image/*') {
    if (!file.type.startsWith('image')) {
      EleMessage.error('Only image files are allowed');
      return;
    }
  } else if (props.accept === '.xls,.xlsx') {
    if (
      ![
        'application/vnd.ms-excel',
        'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
      ].includes(file.type)
    ) {
      EleMessage.error('Only Excel files are allowed');
      return;
    }
  }
  if (props.fileLimit && file.size / 1024 / 1024 > props.fileLimit) {
    EleMessage.error(`File size cannot exceed ${props.fileLimit}MB`);
    return;
  }
  return true;
};

/** Process the image URL to ensure it can be accessed correctly */
const processImageUrl = (url) => {
  if (!url) return url;

  // If it is a complete URL (including http/https), return directly
  if (url.startsWith('http://') || url.startsWith('https://')) {
    return url;
  }

  // If it is a relative path, convert it to the API path
  if (url.startsWith('uploads/')) {
    return `/api/${url}`;
  }

  return url;
};

/** Upload event */
const handleItemUpload = (data, retry) => {
  if (!data.file || !checkFile(data.file)) {
    return;
  }
  if (!retry) {
    images.value.push({ ...data });
  }
  const item = images.value.find((t) => t.key === data.key);
  if (!item) {
    return;
  }
  item.status = 'uploading';
  item.progress = 0;
  uploadFile(data.file, {
    onUploadProgress: (e) => {
      if (e.total != null && item.status !== 'done') {
        item.progress = (e.loaded / e.total) * 100;
      }
    }
  })
    .then((res) => {
      item.progress = 100;
      item.status = 'done';
      item.originalUrl = res.file.url; // Stores the original URL (for submission)
      item.url = processImageUrl(res.file.url); // Stores the processed URL (for display）
    })
    .catch((e) => {
      item.status = 'exception';
      EleMessage.error(e.message);
    });
};

/** Modify event */
const handleItemEditUpload = ({ item, newItem }) => {
  if (!checkFile(newItem.file)) {
    return;
  }
  const oldItem = images.value.find((t) => t.key === item.key);
  if (oldItem) {
    oldItem.url = void 0;
    oldItem.name = newItem.name;
    oldItem.file = newItem.file;
    oldItem.progress = 0;
    oldItem.status = void 0;
    handleItemUpload(oldItem, true);
  }
};

/** Delete event */
const handleItemRemove = (item) => {
  images.value.splice(images.value.indexOf(item), 1);
};

/** Click event */
const handleItemClick = (item) => {
  emit('itemClick', item);
};

/** Preview event */
const handleItemPreview = (item) => {
  emit('preview', item);
};
</script>
