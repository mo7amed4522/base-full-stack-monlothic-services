<!-- File display list -->
<template>
  <div class="demo-file-list">
    <ele-file-list :data="data" :grid="grid" :sort="sort" :order="order" :sortable="true" :box-choose="true"
      selection-type="checkbox" :selections="selections" :context-menus="contextMenus" :context-menu-props="{
        menuStyle: { minWidth: '120px' },
        iconProps: { size: 15 },
        popperOptions: { strategy: 'fixed' }
      }" :icons="localIcons" :small-icons="localSmallIcons"
      :style="{ minHeight: '400px', minWidth: grid ? 'auto' : '456px' }" @item-click="onItemClick"
      @sort-change="onSortChange" @update:selections="updateSelections" @item-context-menu="onCtxMenuClick" />
    <div v-if="!data.length" class="demo-file-list-empty">
      <el-empty description="No data" :image-size="80" />
    </div>
  </div>
  <!-- Image preview -->
  <div style="display: none">
    <el-image ref="previewRef" v-if="previewOption.visible" :src="previewOption.currentSrc"
      :preview-src-list="previewImages" :initial-index="previewOption.current" :hide-on-click-modal="true"
      :preview-teleported="true" />
  </div>
  <!-- File rename popup -->
  <name-edit v-model="nameEditVisible" :data="nameEditData" @done="onDone" />
</template>

<script setup>
import { ref, reactive, nextTick } from 'vue';
import { ElMessageBox } from 'element-plus/es';
import { EleMessage } from 'ele-admin-plus/es';
import {
  localIcons,
  localSmallIcons
} from 'ele-admin-plus/es/ele-file-list/icons';
import {
  DownloadOutlined,
  EditOutlined,
  DragOutlined,
  DeleteOutlined
} from '@/components/icons';
import { removeUserFile } from '@/api/system/user-file';
import NameEdit from './name-edit.vue';

const props = defineProps({
  /** Parent directory id */
  parentId: Number,
  /** File list data */
  data: {
    type: Array,
    required: true
  },
  /** Sort field */
  sort: String,
  /** Sort order */
  order: String,
  /** Selected data */
  selections: {
    type: Array,
    required: true
  },
  /** Whether to display in grid mode */
  grid: {
    type: Boolean,
    required: true
  }
});

const emit = defineEmits([
  'sort-change',
  'update:selections',
  'go-directory',
  'done'
]);

/** Image preview component */
const previewRef = ref(null);

/** Image preview configuration */
const previewOption = reactive({
  currentSrc: '',
  current: 0,
  visible: false
});

/** Image preview list */
const previewImages = ref([]);

/** File rename popup is open */
const nameEditVisible = ref(false);

/** File rename data */
const nameEditData = ref();

/** File list sort order change */
const onSortChange = (option) => {
  emit('sort-change', option);
};

/** Update selected data */
const updateSelections = (value) => {
  emit('update:selections', value);
};

/** item click event */
const onItemClick = (item) => {
  console.log('onItemClick:', JSON.parse(JSON.stringify(item)));
  if (item.isDirectory) {
    // Enter folder
    emit('go-directory', item);
  } else if (isImageFile(item)) {
    // Preview image file
    previewItemImage(item);
  } else {
    // Select or deselect file
    updateSelections(
      props.selections.includes(item)
        ? props.selections.filter((d) => d !== item)
        : [...props.selections, item]
    );
  }
};

/** Right-click menu click event */
const onCtxMenuClick = (option) => {
  const { key, item } = option;
  if (key === 'open') {
    // Open file
    if (item.isDirectory || isImageFile(item)) {
      onItemClick(item);
    } else {
      window.open(item.url);
    }
  } else if (key === 'download') {
    // Download file
    if (typeof item.downloadUrl === 'string') {
      window.open(item.downloadUrl);
    }
  } else if (key === 'edit') {
    // Rename file
    nameEditData.value = item;
    nameEditVisible.value = true;
  } else if (key === 'remove') {
    // Remove file
    removeItem(item);
  } else if (key === 'move') {
    // Move
    EleMessage.info({ message: `Click move to: ${item.name}`, plain: true });
  }
};

/** Remove file */
const removeItem = (item) => {
  ElMessageBox.confirm(`Are you sure you want to delete “${item.name}”?`, 'System prompt', {
    type: 'warning',
    draggable: true
  })
    .then(() => {
      const loading = EleMessage.loading({
        message: 'Requesting...',
        plain: true
      });
      removeUserFile(item.id)
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

/** Check if it is an image file */
const isImageFile = (item) => {
  return (
    typeof item.contentType === 'string' &&
    item.contentType.startsWith('image/') &&
    item.url
  );
};

/** Preview image file */
const previewItemImage = (item) => {
  const temp = props.data.filter((d) => isImageFile(d));
  const index = temp.indexOf(item);
  if (index !== -1) {
    previewImages.value = temp.map((d) => d.url);
    previewOption.currentSrc = item.url;
    previewOption.current = index;
    previewOption.visible = true;
    nextTick(() => {
      previewRef.value?.$el?.querySelector?.('img')?.click?.();
    });
  }
};

/** Right-click menu */
const contextMenus = (item) => {
  if (!item) {
    return [];
  }
  const menus = [
    {
      title: 'Rename',
      command: 'edit',
      icon: EditOutlined,
      divided: item.isDirectory
    },
    {
      title: 'Move to',
      command: 'move',
      icon: DragOutlined
    },
    {
      title: 'Delete',
      command: 'remove',
      icon: DeleteOutlined,
      divided: true,
      danger: true
    }
  ];
  if (!item.isDirectory) {
    menus.unshift({
      title: 'Download',
      command: 'download',
      icon: DownloadOutlined,
      divided: true
    });
  }
  menus.unshift({ title: 'Open', command: 'open' });
  return menus;
};
</script>

<style lang="scss" scoped>
.demo-file-list {
  position: relative;
  overflow-x: auto;

  .demo-file-list-empty {
    position: absolute;
    top: 100px;
    left: 50%;
    transform: translateX(-50%);
  }
}
</style>
