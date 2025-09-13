<template>
  <div class="file-picker-right">
    <div class="file-picker-right-header">
      <div class="file-picker-right-title">
        <span>Selected </span>
        <span>{{ fileSelections.length }}</span>
        <template v-if="limit">
          <span> / </span>
          <span>{{ limit }}</span>
        </template>
        <span v-else> 个</span>
      </div>
      <ElLink type="danger" :underline="false" class="file-picker-right-clear" @click="clearSelections">
        Clear
      </ElLink>
    </div>
    <EleUploadList :limit="limit" :previewProps="previewProps" v-bind="selectionListProps || {}" ref="uploadListRef"
      v-model="images" :buttonStyle="false" :sortable="!mobileDevice" @remove="handleRemove" />
  </div>
</template>

<script setup>
import { ref, reactive, watch } from 'vue';
import { valueIsChanged } from 'ele-admin-plus/es/ele-basic-select/util';
import { useMobileDevice } from '@/utils/use-mobile';

const props = defineProps({
  /** Selected file data */
  fileSelections: Array,
  /** Maximum selection count */
  limit: Number,
  /** File list custom attributes */
  selectionListProps: Object,
  /** Uniformly set the layer */
  baseIndex: Number
});

const emit = defineEmits(['clearSelections', 'removeItem']);

/** Touch device disables drag sorting */
const { mobileDevice } = useMobileDevice();

/** Image upload list component */
const uploadListRef = ref(null);

/** Image data */
const images = ref([]);

/** Image preview component attributes */
const previewProps = reactive({
  zIndex: props.baseIndex
});

/** Delete event */
const handleRemove = (item) => {
  const fileItem = props.fileSelections.find((d) => d.key === item.key);
  if (fileItem) {
    emit('removeItem', fileItem);
  }
};

/** Clear selected data */
const clearSelections = () => {
  emit('clearSelections');
};

/** Get selected data */
const getSelections = () => {
  const result = [];
  images.value.forEach((d) => {
    const item = props.fileSelections.find((t) => t.key === d.key);
    if (item) {
      result.push(item);
    }
  });
  return result;
};

/** Open preview image popup */
const openImagePreview = (urls, index) => {
  uploadListRef.value && uploadListRef.value.openImagePreview(urls, index);
};

watch(
  () => props.fileSelections,
  (selections) => {
    if (!selections || !selections.length) {
      if (images.value.length) {
        images.value = [];
      }
      return;
    }
    const temp = images.value.filter((d) =>
      selections.some((t) => t.key === d.key)
    );
    selections.forEach((d) => {
      if (!temp.some((t) => t.key === d.key)) {
        temp.push({
          key: d.key,
          name: d.name,
          url: d.thumbnail,
          status: 'done'
        });
      }
    });
    if (
      valueIsChanged(
        images.value.map((d) => d.key),
        temp.map((d) => d.key),
        true
      )
    ) {
      images.value = temp;
    }
  },
  {
    immediate: true,
    deep: true
  }
);

watch(
  () => props.baseIndex,
  (baseIndex) => {
    previewProps.zIndex = baseIndex;
  }
);

defineExpose({ getSelections, openImagePreview });
</script>
