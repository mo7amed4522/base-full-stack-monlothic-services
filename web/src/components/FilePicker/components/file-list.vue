<template>
  <div class="file-picker-main">
    <div class="file-picker-body">
      <div class="file-picker-toolbar">
        <ElUpload action="" :accept="accept" :showFileList="false" :beforeUpload="handleUpload">
          <ElButton type="primary" class="ele-btn-icon" :icon="UploadOutlined">
            Upload
          </ElButton>
        </ElUpload>
        <div class="file-picker-search">
          <ElInput :clearable="true" v-model="searchKeyword" placeholder="Search by file name" @clear="handleSearch"
            @change="handleSearch" />
          <ElButton type="primary" @click="handleSearch">Search</ElButton>
        </div>
        <EleSegmented v-model="isGridMode" :items="modeSegmentedItems" />
      </div>
      <template v-if="fileData.length">
        <div class="file-picker-file-list" @scroll="handleFileListScroll">
          <EleFileList :boxChoose="true" :icons="localIcons" :smallIcons="localSmallIcons"
            :contextMenuProps="contextMenuProps" v-bind="fileListProps || {}" :data="fileData" :grid="isGridMode === 1"
            :selectionType="limit === 1 ? 'radio' : 'checkbox'" :selections="fileSelections"
            v-model:current="fileCurrent" :contextMenus="fileContextMenus" :class="[{ 'is-ping-top': isPingTop }]"
            @itemClick="handleFileItemClick" @itemContextMenu="handleFileCtxMenuClick"
            @update:selections="updateSelections" @itemContextOpen="handleFileCtxMenuOpen" />
        </div>
        <ElePagination size="small" :teleported="false" :pageSizes="[18, 24, 30, 36, 42]"
          layout="total,prev,pager,next,sizes" v-bind="paginationProps || {}" :currentPage="currentPage"
          :pageSize="pageSize" :total="total" @update:currentPage="handleCurrentPageChange"
          @update:pageSize="handlePageSizeChange" />
      </template>
      <ElEmpty v-else :imageSize="80" description="No data" v-bind="emptyProps || {}" />
    </div>
    <FileSelections ref="fileSelectionsRef" :fileSelections="limit === 1 ? (fileCurrent ? [fileCurrent] : []) : fileSelections
      " :limit="limit" :selectionListProps="selectionListProps" :baseIndex="baseIndex"
      @clearSelections="clearSelections" @removeItem="removeItem" />
  </div>
</template>

<script setup>
import { ref, reactive, watch, markRaw } from 'vue';
import {
  localIcons,
  localSmallIcons
} from 'ele-admin-plus/es/ele-file-list/icons';
import {
  UploadOutlined,
  MenuOutlined,
  AppstoreOutlined,
  EditOutlined,
  DragOutlined,
  DeleteOutlined
} from '@/components/icons';
import { uploadFile } from '@/api/system/file';
import { pageUserFiles } from '@/api/system/user-file';
import FileSelections from './file-selections.vue';

const props = defineProps({
  /** Maximum number of selectable files */
  limit: Number,
  /** File size limit, unit MB */
  fileLimit: Number,
  /** Accepted file types for upload */
  accept: String,
  /** Interface query parameters */
  params: Object,
  /** File list custom attributes */
  fileListProps: Object,
  /** Selected list custom attributes */
  selectionListProps: Object,
  /** Empty component attributes */
  emptyProps: Object,
  /** Pagination component attributes */
  paginationProps: Object,
  /** Base index for setting z-index */
  baseIndex: Number,
  /** Message prompt component */
  messageIns: [Object, Function]
});

const emit = defineEmits([
  'queryStart',
  'queryDone',
  'renameFile',
  'moveFile',
  'removeFile',
  'fileItemContextOpen'
]);

/** Current group id */
const fileParentId = ref();

/** File data */
const fileData = ref([]);

/** Selected file data */
const fileSelections = ref([]);

/** Single selected file data */
const fileCurrent = ref();

/** Whether grid mode is enabled */
const isGridMode = ref(1);

/** Search keyword */
const searchKeyword = ref('');

/** Current page number */
const currentPage = ref(1);

/** Number of items per page */
const pageSize = ref(24);

/** Total number of items */
const total = ref(0);

/** Whether the file list header is fixed */
const isPingTop = ref(false);

/** View mode segmented data */
const modeSegmentedItems = [
  { icon: MenuOutlined, value: 0, iconStyle: { transform: 'scale(0.9)' } },
  { icon: AppstoreOutlined, value: 1 }
];

/** Selected list component */
const fileSelectionsRef = ref();

/** File list right-click menu attributes */
const contextMenuProps = reactive({
  menuStyle: { minWidth: '120px' },
  iconProps: { size: 15 },
  popperOptions: { strategy: 'fixed' },
  zIndex: props.baseIndex
});

/** Check the maximum number of selections */
const checkLimit = (selections, isAdd) => {
  if (
    props.limit &&
    props.limit > 1 &&
    (isAdd
      ? selections.length >= props.limit
      : selections.length > props.limit)
  ) {
    props.messageIns?.error?.(`You can only choose ${props.limit} indivual`);
    return false;
  }
};

/** Update selected data */
const updateSelections = (selections) => {
  if (checkLimit(selections) !== false) {
    fileSelections.value = selections;
    return;
  }
  if (props.limit && props.limit > 1) {
    fileSelections.value = selections.slice(0, props.limit);
  }
};

/** Clear selected */
const clearSelections = () => {
  fileSelections.value = [];
  fileCurrent.value = void 0;
};

/** Get selected data */
const getSelections = () => {
  if (fileSelectionsRef.value) {
    const selections = fileSelectionsRef.value.getSelections() || [];
    return selections.map((d) => d.userFile);
  }
  return [];
};

/** Remove selected */
const removeItem = (item) => {
  if (props.limit === 1) {
    clearSelections();
    return;
  }
  const index = fileSelections.value.findIndex((d) => d.key === item.key);
  if (index !== -1) {
    fileSelections.value.splice(index, 1);
  }
};

/** Search */
const handleSearch = () => {
  queryData({ page: 1 });
};

/** Switch current page */
const handleCurrentPageChange = (page) => {
  if (currentPage.value !== page) {
    currentPage.value = page;
    queryData();
  }
};

/** Switch page size */
const handlePageSizeChange = (limit) => {
  if (pageSize.value !== limit) {
    currentPage.value = 1;
    pageSize.value = limit;
    const maxPage = Math.ceil(total.value / limit);
    if (maxPage && currentPage.value > maxPage) {
      currentPage.value = maxPage;
    }
    queryData();
  }
};

/** File list click event */
const handleFileItemClick = (item) => {
  if (props.limit === 1) {
    fileCurrent.value = item;
    return;
  }
  const index = fileSelections.value.findIndex((d) => d.key === item.key);
  if (index !== -1) {
    fileSelections.value.splice(index, 1);
    return;
  }
  if (checkLimit(fileSelections.value, true) !== false) {
    fileSelections.value.push(item);
  }
};

/** File list right-click menu click event */
const handleFileCtxMenuClick = (option) => {
  const { key, item } = option;
  const userFileItem = item.userFile;
  if (key === 'preview') {
    if (!item.thumbnail) {
      window.open(item.url);
      return;
    }
    const data = fileData.value.filter((d) => !!d.thumbnail);
    const urls = data.map((d) => d.url);
    const index = data.indexOf(item);
    if (index !== -1) {
      fileSelectionsRef.value &&
        fileSelectionsRef.value.openImagePreview(urls, index);
    }
  } else if (key === 'rename') {
    emit('renameFile', userFileItem, true);
  } else if (key === 'move') {
    emit('moveFile', userFileItem, true);
  } else if (key === 'remove') {
    emit('removeFile', userFileItem, true);
  }
};

/** File list right-click menu data */
const fileContextMenus = (item) => {
  const menus = [
    {
      title: 'Rename',
      command: 'rename',
      icon: markRaw(EditOutlined)
    },
    {
      title: 'Move',
      command: 'move',
      icon: markRaw(DragOutlined)
    },
    {
      title: 'Delete',
      command: 'remove',
      icon: markRaw(DeleteOutlined),
      divided: true,
      danger: true
    }
  ];
  if (item.thumbnail) {
    menus[0].divided = true;
    menus.unshift({ title: 'Preview', command: 'preview' });
  } else {
    menus[0].divided = true;
    menus.unshift({ title: 'Open', command: 'preview' });
  }
  return menus;
};

/** File list right-click menu open event */
const handleFileCtxMenuOpen = () => {
  emit('fileItemContextOpen');
};

/** Check selected file */
const checkFile = (file) => {
  if (!file) {
    return;
  }
  if (props.accept === 'image/*') {
    if (!file.type.startsWith('image')) {
      props.messageIns?.error?.('Only pictures can be selected');
      return;
    }
  } else if (props.accept === '.xls,.xlsx') {
    if (
      ![
        'application/vnd.ms-excel',
        'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
      ].includes(file.type)
    ) {
      props.messageIns?.error?.('Only excel files can be selected');
      return;
    }
  }
  if (props.fileLimit && file.size / 1024 / 1024 > props.fileLimit) {
    props.messageIns?.error?.(`Size cannot exceed ${props.fileLimit}MB`);
    return;
  }
  return true;
};

/** File upload event */
const handleUpload = (file) => {
  if (checkFile(file)) {
    const loading = props.messageIns?.loading?.({
      message: 'Uploading...',
      plain: true,
      mask: true
    });
    uploadFile(file)
      .then(() => {
        loading?.close?.();
        props.messageIns?.success?.('Upload successful');
        queryData();
      })
      .catch((e) => {
        loading?.close?.();
        props.messageIns?.error?.(e.message);
      });
  }
  return false;
};

/** File list scroll event */
const handleFileListScroll = (e) => {
  const wrapEl = e.currentTarget;
  const scrollTop = wrapEl.scrollTop;
  isPingTop.value = scrollTop > 1;
};

/** Trigger file data request completion */
const handleQueryDone = () => {
  emit('queryDone');
};

/** Judge whether it is a picture file */
const isImageFile = (item) => {
  return (
    typeof item.contentType === 'string' &&
    item.contentType.startsWith('image/') &&
    item.url
  );
};

/** Format file size */
const formatLength = (length) => {
  if (length == null) {
    return '-';
  }
  if (length < 1024) {
    return length + 'B';
  } else if (length < 1024 * 1024) {
    return (length / 1024).toFixed(1) + 'KB';
  } else if (length < 1024 * 1024 * 1024) {
    return (length / 1024 / 1024).toFixed(1) + 'M';
  } else {
    return (length / 1024 / 1024 / 1024).toFixed(1) + 'G';
  }
};

/** Query file data */
const queryData = (params) => {
  emit('queryStart');
  if (params) {
    if (params.page != null) {
      currentPage.value = params.page;
    }
    if (params.parentId != null) {
      fileParentId.value = params.parentId;
    }
    if (params.name != null) {
      searchKeyword.value = params.name;
    }
  }
  const parentId = fileParentId.value;
  pageUserFiles({
    page: currentPage.value,
    limit: pageSize.value,
    parentId: parentId === -1 ? void 0 : parentId,
    name: searchKeyword.value,
    isDirectory: 0,
    ...(props.params || {})
  })
    .then((result) => {
      if (!result?.list?.length && result?.count) {
        const maxPage = Math.ceil(result.count / pageSize.value);
        if (maxPage && currentPage.value > maxPage) {
          queryData({ page: maxPage });
          return;
        }
      }
      total.value = result?.count || 0;
      fileData.value =
        result?.list?.map?.((d) => {
          return {
            key: d.id,
            name: d.name,
            url: d.url,
            thumbnail: isImageFile(d) ? d.url : void 0,
            length: formatLength(d.length),
            updateTime: d.updateTime,
            isDirectory: d.isDirectory === 1,
            userFile: d
          };
        }) || [];
      handleQueryDone();
    })
    .catch((e) => {
      fileData.value = [];
      props.messageIns?.error?.(e.message);
      handleQueryDone();
    });
};

watch(
  () => props.limit,
  () => {
    clearSelections();
  }
);

watch(
  () => props.baseIndex,
  (baseIndex) => {
    contextMenuProps.zIndex = baseIndex;
  }
);

defineExpose({ queryData, clearSelections, getSelections });
</script>
