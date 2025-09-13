<!-- File picker -->
<template>
  <EleModal :width="988" :title="title" position="center" :zIndex="baseIndex" v-bind="modalProps || {}"
    v-model="visible" class="file-picker-modal" @open="handleOpen" @close="handleClose">
    <EleSplitPanel ref="splitRef" space="0px" size="186px" :flexTable="true" :allowCollapse="mobile"
      :customStyle="{ borderWidth: '0 1px 0 0' }" class="file-picker-wrapper">
      <FileGroup ref="fileGroupRef" :groupData="groupData" @groupSelect="handleGroupSelect"
        @addBtnClick="openGroupAdd()" @moreIconClick="openCtxMenuDropdown" @hideGroupCtxMenu="hideGroupContextmenu()" />
      <template #body>
        <FileList ref="fileListRef" :limit="limit" :fileLimit="fileLimit" :accept="accept" :params="params"
          :fileListProps="fileListProps" :selectionListProps="selectionListProps" :emptyProps="emptyProps"
          :paginationProps="paginationProps" :baseIndex="componentIndex" :messageIns="messageIns"
          @queryStart="showLoading" @queryDone="hideLoading" @renameFile="openRenameItem" @moveFile="openMoveItem"
          @removeFile="removeItem" @fileItemContextOpen="hideGroupContextmenu" />
      </template>
    </EleSplitPanel>
    <EleLoading :loading="loading" class="file-picker-loading" />
    <!-- Add group popup -->
    <FileGroupAdd v-model="showGroupAdd" :parentId="groupParentId" :groupData="groupData"
      :modalProps="groupAddModalProps" :baseIndex="componentIndex" :messageIns="messageIns" @done="handleReloadData" />
    <!-- Group or file rename popup -->
    <FileRename v-model="showRename" :data="renameData" :modalProps="renameModalProps" :baseIndex="componentIndex"
      :messageIns="messageIns" @done="handleReloadData" />
    <!-- File move popup -->
    <FileMove v-model="showMove" :data="moveData" :groupData="groupData" :modalProps="moveModalProps"
      :baseIndex="componentIndex" :messageIns="messageIns" @done="handleReloadData" />
    <!-- Group item right-click menu -->
    <EleDropdown :triggerKeys="[]" :persistent="false" :zIndex="componentIndex" placement="bottom-end"
      :iconProps="{ size: 15 }" :popperOptions="{
        strategy: 'fixed',
        modifiers: [{ name: 'offset', options: { offset: [6, 6] } }]
      }" v-bind="groupDropdownProps || {}" ref="ctxMenuDropdownRef" trigger="click" componentType="pro"
      :preventContextmenu="true" :virtualTriggering="true" :virtualRef="ctxMenuDropdownVirtualRef"
      :items="ctxMenuDropdownItems" @command="handleItemCommand" />
    <!-- Message prompt container -->
    <div ref="messageWrapRef" class="ele-message-wrapper" :style="{ position: 'fixed', zIndex: messageIndex }"></div>
    <div ref="messageBoxWrapRef" class="ele-message-box-wrapper" :style="{ position: 'fixed', zIndex: messageIndex }">
    </div>
    <template #footer>
      <ElButton @click="handleCancel">Cancel</ElButton>
      <ElButton type="primary" @click="handleConfirm">Confirm</ElButton>
    </template>
  </EleModal>
</template>

<script setup>
import { ref, computed, nextTick, markRaw } from 'vue';
import {
  useMessage,
  useMessageBox,
  toTree,
  findTree
} from 'ele-admin-plus/es';
import {
  EditOutlined,
  FolderOutlined,
  DeleteOutlined
} from '@/components/icons';
import { useMobile } from '@/utils/use-mobile';
import FileGroup from './components/file-group.vue';
import FileList from './components/file-list.vue';
import FileGroupAdd from './components/file-group-add.vue';
import FileRename from './components/file-rename.vue';
import FileMove from './components/file-move.vue';
import { listUserFiles, removeUserFile } from '@/api/system/user-file';

defineOptions({ name: 'FilePicker' });

const props = defineProps({
  /** Modal title */
  title: {
    type: String,
    default: 'File picker'
  },
  /** Modal parameters */
  modalProps: Object,
  /** Maximum number of files to select */
  limit: Number,
  /** File size limit, unit MB */
  fileLimit: {
    type: Number,
    default: 100
  },
  /** Accepted file types for upload */
  accept: String,
  /** API query parameters */
  params: Object,
  /** Group dropdown menu properties */
  groupDropdownProps: Object,
  /** Add group popup parameters */
  groupAddModalProps: Object,
  /** Rename popup parameters */
  renameModalProps: Object,
  /** Move group popup parameters */
  moveModalProps: Object,
  /** File list custom attributes */
  fileListProps: Object,
  /** Selected list custom attributes */
  selectionListProps: Object,
  /** Empty component attributes */
  emptyProps: Object,
  /** Pagination component attributes */
  paginationProps: Object,
  /** Set the layer uniformly */
  baseIndex: Number
});

const emit = defineEmits(['done', 'close']);

/** Modal visibility */
const visible = defineModel({ type: Boolean });

/** Whether it is a mobile device */
const { mobile } = useMobile();

/** Split panel component */
const splitRef = ref(null);

/** File group component */
const fileGroupRef = ref(null);

/** File list component */
const fileListRef = ref(null);

/** Data request status */
const loading = ref(false);

/** Group data */
const groupData = ref([]);

/** Selected group data */
const groupSelected = ref(null);

/** Whether to display the rename popup */
const showRename = ref(false);

/** Current rename data */
const renameData = ref(null);

/** Whether to display the add group popup */
const showGroupAdd = ref(false);

/** Add group parent id */
const groupParentId = ref();

/** Whether to display the move popup */
const showMove = ref(false);

/** Current mobile data */
const moveData = ref(null);

/** Whether the current edit is file operation */
const isEditFile = ref(false);

/** Right-click menu component */
const ctxMenuDropdownRef = ref();

/** Right-click menu virtual trigger node */
const ctxMenuDropdownVirtualRef = ref();

/** Right-click menu item data */
const ctxMenuDropdownItems = ref([]);

/** Current open right-click menu corresponding data */
let ctxMenuFileItem = null;

/** Component unified layer */
const componentIndex = computed(() => {
  if (props.baseIndex == null) {
    return;
  }
  return props.baseIndex + 100;
});

/** Message prompt component layer */
const messageIndex = computed(() => {
  if (componentIndex.value == null) {
    return 3000;
  }
  return componentIndex.value + 100;
});

/** Message prompt container */
const messageWrapRef = ref();

/** Message popup box container */
const messageBoxWrapRef = ref();

/** Message prompt box */
const messageIns = useMessage({
  inner: true,
  appendTo: messageWrapRef
});

/** Message popup box */
const messageBoxIns = useMessageBox({
  appendTo: messageBoxWrapRef
});

/** Get group right-click menu item data */
const getGroupCtxMenuItems = (_item) => {
  return [
    { title: 'Rename', command: 'rename', icon: markRaw(EditOutlined) },
    {
      title: 'Add group',
      command: 'add',
      icon: markRaw(FolderOutlined),
      iconStyle: { transform: 'scale(0.92)' }
    },
    {
      title: 'Delete group',
      command: 'remove',
      icon: markRaw(DeleteOutlined),
      divided: true,
      danger: true
    }
  ];
};

/** Close group right-click menu */
const hideGroupContextmenu = () => {
  ctxMenuDropdownRef.value && ctxMenuDropdownRef.value.handleClose();
};

/** Open group right-click menu */
const openCtxMenuDropdown = (triggerEl, item) => {
  if (
    triggerEl == null ||
    ctxMenuDropdownVirtualRef.value === triggerEl ||
    (ctxMenuFileItem != null && ctxMenuFileItem === item)
  ) {
    return;
  }
  hideGroupContextmenu();
  nextTick(() => {
    ctxMenuFileItem = item;
    ctxMenuDropdownItems.value = getGroupCtxMenuItems(item) || [];
    ctxMenuDropdownVirtualRef.value = triggerEl;
    if (ctxMenuDropdownItems.value.length) {
      nextTick(() => {
        ctxMenuDropdownRef.value && ctxMenuDropdownRef.value.handleOpen();
      });
    }
  });
};

/** Group selected */
const handleGroupSelect = (item, isTreeClick) => {
  if (groupSelected.value != null && mobile.value) {
    splitRef.value?.toggleCollapse?.(true);
  }
  if (item && item.id != null) {
    if (!isTreeClick) {
      fileGroupRef.value?.setSelectedGroup?.(item);
    }
    if (groupSelected.value == null || groupSelected.value.id !== item.id) {
      groupSelected.value = item;
      queryData(item.id);
    }
  }
};

/** Refresh data event after edit completion */
const handleReloadData = () => {
  isEditFile.value ? queryData() : queryGroup();
};

/** Add group */
const openGroupAdd = (item) => {
  isEditFile.value = false;
  groupParentId.value = item ? item.id : void 0;
  showGroupAdd.value = true;
};

/** Rename group or file */
const openRenameItem = (item, isFile) => {
  isEditFile.value = !!isFile;
  renameData.value = item;
  showRename.value = true;
};

/** Move file to specified group */
const openMoveItem = (item, isFile) => {
  isEditFile.value = !!isFile;
  moveData.value = item;
  showMove.value = true;
};

/** Delete group or file */
const removeItem = (item, isFile) => {
  isEditFile.value = !!isFile;
  messageBoxIns
    ?.confirm?.(`Are you sure you want to delete "${item.name}"?`, 'System prompts', {
      type: 'warning',
      draggable: true
    })
    ?.then?.(() => {
      const loading = messageIns?.loading?.({
        message: 'Deleting...',
        plain: true
      });
      removeUserFile(item.id)
        .then((msg) => {
          loading?.close?.();
          messageIns?.success?.(msg);
          handleReloadData();
        })
        .catch((e) => {
          loading?.close?.();
          messageIns?.error?.(e.message);
        });
    })
    ?.catch?.(() => { });
};

/** Group right-click menu item click event */
const handleItemCommand = (command) => {
  if (ctxMenuFileItem == null) {
    return;
  }
  if (command === 'rename') {
    openRenameItem(ctxMenuFileItem);
  } else if (command === 'add') {
    openGroupAdd(ctxMenuFileItem);
  } else if (command === 'remove') {
    removeItem(ctxMenuFileItem);
  }
};

/** Close request loading icon */
const hideLoading = () => {
  loading.value = false;
};

/** Open request loading icon */
const showLoading = () => {
  loading.value = true;
};

/** Query group data */
const queryGroup = () => {
  showLoading();
  listUserFiles({ isDirectory: 1 })
    .then((list) => {
      const result = [
        { id: -1, name: 'all' },
        { id: 0, name: 'Ungrouped' },
        ...toTree({
          data: list,
          idField: 'id',
          parentIdField: 'parentId'
        })
      ];
      groupData.value = result;
      const oldSelected =
        groupSelected.value == null
          ? void 0
          : findTree(result, (d) => d.id === groupSelected.value?.id);
      groupSelected.value = null;
      nextTick(() => {
        handleGroupSelect(oldSelected || result[0]);
      });
    })
    .catch((e) => {
      groupData.value = [];
      messageIns.error(e.message);
      hideLoading();
    });
};

/** Refresh file data */
const queryData = (parentId) => {
  showLoading();
  const params = parentId == null ? void 0 : { page: 1, name: '', parentId };
  fileListRef.value && fileListRef.value.queryData(params);
};

/** Clear file selected */
const clearSelections = () => {
  fileListRef.value && fileListRef.value.clearSelections();
};

/** Popup open event */
const handleOpen = () => {
  clearSelections();
  queryGroup();
};

/** Popup close event */
const handleClose = () => {
  clearSelections();
  emit('close');
};

/** Cancel button click event */
const handleCancel = () => {
  visible.value = false;
};

/** Confirm button click event */
const handleConfirm = () => {
  showLoading();
  emit('done', fileListRef.value?.getSelections?.() || []);
};

defineExpose({ hideLoading });
</script>

<style lang="scss">
@use './style/index.scss';
</style>
