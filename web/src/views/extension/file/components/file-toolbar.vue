<template>
  <ele-toolbar theme="plain" class="demo-file-toolbar">
    <el-upload action="" :show-upload-list="false" :before-upload="onUpload"
      style="display: inline-block; vertical-align: middle">
      <el-button type="primary" class="ele-btn-icon" :icon="UploadOutlined">
        Upload
      </el-button>
    </el-upload>
    <el-button class="ele-btn-icon" :icon="FolderAddOutlined" style="margin-left: 12px" @click="openFolderAdd">
      New folder
    </el-button>
    <el-button plain type="danger" :icon="DeleteOutlined" :disabled="!selections.length"
      class="ele-btn-icon hidden-xs-only" @click="removeBatch">
      Delete
    </el-button>
    <template #tools>
      <!-- Search box -->
      <div style="max-width: 240px" class="hidden-sm-and-down">
        <el-input v-model="search" placeholder="Search your files" :prefix-icon="SearchOutlined" />
      </div>
      <!-- Display mode switch -->
      <el-icon class="demo-file-tool" @click="toggleShowType">
        <TableOutlined v-if="grid" />
        <AppstoreOutlined v-else />
      </el-icon>
    </template>
  </ele-toolbar>
  <!-- New folder popup -->
  <folder-edit v-model="folderAddVisible" :parent-id="parentId" @done="onDone" />
</template>

<script setup>
import { ref } from 'vue';
import { ElMessageBox } from 'element-plus/es';
import { EleMessage } from 'ele-admin-plus/es';
import {
  UploadOutlined,
  FolderAddOutlined,
  DeleteOutlined,
  SearchOutlined,
  TableOutlined,
  AppstoreOutlined
} from '@/components/icons';
import { uploadFile } from '@/api/system/file';
import { addUserFile, removeUserFiles } from '@/api/system/user-file';
import FolderEdit from './folder-edit.vue';

const props = defineProps({
  /** Whether to display in grid mode */
  grid: Boolean,
  /** Selected data */
  selections: Array,
  /** Parent id */
  parentId: Number
});

const emit = defineEmits(['update:grid', 'done']);

/** Search keyword */
const search = ref('');

/** Whether the new folder popup is open */
const folderAddVisible = ref(false);

/** Upload */
const onUpload = (file) => {
  if (file.size / 1024 / 1024 > 100) {
    EleMessage.error('Size cannot exceed 100MB');
    return false;
  }
  const loading = EleMessage.loading({
    message: 'Uploading...',
    plain: true
  });
  uploadFile(file)
    .then((data) => {
      addUserFile({
        name: data.name,
        isDirectory: 0,
        parentId: props.parentId,
        path: data.path,
        length: data.length,
        contentType: data.contentType
      })
        .then(() => {
          loading.close();
          EleMessage.success('Upload successful');
          onDone();
        })
        .catch((e) => {
          loading.close();
          EleMessage.error(e.message);
        });
    })
    .catch((e) => {
      loading.close();
      EleMessage.error(e.message);
    });
  return false;
};

/** Open new folder popup */
const openFolderAdd = () => {
  folderAddVisible.value = true;
};

/** Batch delete */
const removeBatch = () => {
  const rows = props.selections;
  ElMessageBox.confirm(
    'Are you sure you want to delete “' + rows.map((d) => d.name).join(', ') + '”?',
    'System prompt',
    { type: 'warning', draggable: true }
  )
    .then(() => {
      const loading = EleMessage.loading({
        message: 'Requesting...',
        plain: true
      });
      removeUserFiles(rows.map((d) => d.id))
        .then((msg) => {
          loading.close();
          EleMessage.success(msg);
          onDone();
        })
        .catch((e) => {
          loading.close();
          EleMessage.error(e.message);
        });
    })
    .catch(() => { });
};

/** Refresh list data */
const onDone = () => {
  emit('done');
};

/** Display mode switch */
const toggleShowType = () => {
  emit('update:grid', !props.grid);
};
</script>

<style lang="scss" scoped>
.demo-file-toolbar :deep(.el-upload-list) {
  display: none;
}

/* Icon button */
.demo-file-tool {
  font-size: 18px;
  margin-left: 16px;
  cursor: pointer;
}
</style>
