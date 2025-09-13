<template>
  <EleModal :width="460" title="Move to" :zIndex="baseIndex" :appendToBody="false" v-bind="modalProps || {}"
    v-model="visible" @open="handleOpen">
    <div class="file-picker-move-wrapper">
      <ElTree :data="groupData.filter((d) => d.id !== -1)" nodeKey="id" :props="{ label: 'name' }"
        :expandOnClickNode="false" :defaultExpandAll="true" class="file-picker-move-tree"
        @node-click="handleGroupSelect">
        <template #default="scope">
          <img src="/ele-file-list/ic_file_folder.png" class="file-picker-tree-icon" />
          <span :class="[
            'el-tree-node__label',
            { 'is-active': scope.data.id === selectedId }
          ]">
            {{ scope.data.name }}
          </span>
          <ElRadio :value="true" :modelValue="scope.data.id === selectedId" />
        </template>
      </ElTree>
    </div>
    <template #footer>
      <ElButton @click="handleCancel">Cancel</ElButton>
      <ElButton type="primary" :loading="loading" @click="save">
        Save
      </ElButton>
    </template>
  </EleModal>
</template>

<script setup>
import { ref } from 'vue';
import { updateUserFile } from '@/api/system/user-file';

const props = defineProps({
  /** File data */
  data: Object,
  /** Group data */
  groupData: Array,
  /** Modal parameters */
  modalProps: Object,
  /** Base index */
  baseIndex: Number,
  /** Message instance */
  messageIns: [Object, Function]
});

const emit = defineEmits(['done']);

/** Modal visibility */
const visible = defineModel({ type: Boolean });

/** Loading state */
const loading = ref(false);

/** Selected group id */
const selectedId = ref();

/** Group select event */
const handleGroupSelect = (item) => {
  selectedId.value = item.id;
};

/** Modal close event */
const handleCancel = () => {
  visible.value = false;
};

/** Save event */
const save = () => {
  const parentId = selectedId.value;
  if (parentId == null) {
    props.messageIns?.error?.('Please select a group');
    return;
  }
  loading.value = true;
  updateUserFile({ id: props.data?.id, parentId })
    .then((msg) => {
      loading.value = false;
      props.messageIns?.success?.(msg);
      handleCancel();
      emit('done');
    })
    .catch((e) => {
      loading.value = false;
      props.messageIns?.error?.(e.message);
    });
};

/** Modal open event */
const handleOpen = () => {
  if (props.data) {
    selectedId.value = props.data.parentId;
  }
};
</script>
