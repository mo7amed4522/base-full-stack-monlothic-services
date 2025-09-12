<template>
  <ele-card header="Drag, stretch, full screen">
    <el-form label-width="120px" style="max-width: 380px" @submit.prevent="">
      <el-form-item label="Can it be dragged out of the boundary?">
        <el-select multiple clearable v-model="moveOut" placeholder="Cannot be dragged out of the boundary"
          class="ele-fluid" :popper-options="{ strategy: 'fixed' }">
          <el-option value="top" label="top" />
          <el-option value="bottom" label="bottom" />
          <el-option value="left" label="left" />
          <el-option value="right" label="right" />
        </el-select>
      </el-form-item>
      <el-form-item label="Whether it can be stretched">
        <el-select v-model="resizable" class="ele-fluid" :popper-options="{ strategy: 'fixed' }">
          <el-option value="false" label="Unstretchable size" />
          <el-option value="true" label="Stretchable size" />
          <el-option value="horizontal" label="Only horizontally stretchable" />
          <el-option value="vertical" label="Only vertically stretchable" />
        </el-select>
      </el-form-item>
      <el-form-item label="Whether it can be maximized">
        <el-switch v-model="maxable" size="small" />
      </el-form-item>
      <el-form-item label="Close to reset position">
        <el-switch v-model="resetOnClose" size="small" />
      </el-form-item>
      <el-form-item label="Limit in body area">
        <el-switch v-model="inner" size="small" @change="onConfigChange" />
      </el-form-item>
      <el-form-item label="Default open position">
        <el-select clearable v-model="position" placeholder="Please select" class="ele-fluid" @change="onConfigChange"
          :popper-options="{ strategy: 'fixed' }">
          <el-option value="top" label="Top" />
          <el-option value="bottom" label="Bottom" />
          <el-option value="left" label="Left" />
          <el-option value="right" label="Right" />
          <el-option value="leftTop" label="Top left" />
          <el-option value="leftBottom" label="Bottom left" />
          <el-option value="rightTop" label="Top right" />
          <el-option value="rightBottom" label="Bottom right" />
          <el-option value="center" label="Center" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" class="ele-btn-icon" @click="openDialog">
          Open dialog
        </el-button>
      </el-form-item>
    </el-form>
  </ele-card>
  <ele-modal :key="modalKey" form :width="400" title="Advanced dialog" v-model="visible" :move-out="moveOut"
    :resizable="modalResizable" :maxable="maxable" :inner="inner" :reset-on-close="resetOnClose" :position="position"
    :append-to-body="true" :z-index="2001" @closed="cancelDialog">
    <el-form ref="formRef" :model="form" :rules="rules" label-width="70px" @submit.prevent="">
      <el-form-item label="Username" prop="nickname">
        <el-input clearable v-model="form.nickname" placeholder="Please input username" />
      </el-form-item>
      <el-form-item label="Sex" prop="sex">
        <el-select clearable v-model="form.sex" placeholder="Please select sex" class="ele-fluid">
          <el-option value="male">Male</el-option>
          <el-option value="female">Female</el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="Phone" prop="phone">
        <el-input clearable v-model="form.phone" placeholder="Please input phone" />
      </el-form-item>
      <el-form-item label="Email" prop="email">
        <el-input clearable v-model="form.email" placeholder="Please input email" />
      </el-form-item>
      <el-form-item label="Introduction">
        <el-input :rows="4" type="textarea" v-model="form.introduction" placeholder="Please input introduction" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="cancelDialog">Cancel</el-button>
      <el-button type="primary" @click="onSave">Save</el-button>
    </template>
  </ele-modal>
</template>

<script setup>
import { ref, reactive, computed } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';

/** Used to update pop-up windows */
const modalKey = ref(0);

/** Is the pop-up window open? */
const visible = ref(false);

/** Whether it can be dragged out of the boundary */
const moveOut = ref([]);

/** Whether it can be stretched */
const resizable = ref('false');

/** Whether to display the maximization toggle button */
const maxable = ref(true);

/** Whether to reset the position after closing */
const resetOnClose = ref(true);

/** Whether to limit in the body area */
const inner = ref(false);

/** Default position */
const position = ref();

/** Form instance */
const formRef = ref(null);

/** Pop-up stretch properties */
const modalResizable = computed(() => {
  return resizable.value === 'true'
    ? true
    : resizable.value === 'false'
      ? false
      : resizable.value;
});

/** Pop-up window property change */
const onConfigChange = () => {
  modalKey.value = modalKey.value + 1;
};

/** Form data */
const [form, resetFields] = useFormData({
  nickname: '',
  sex: void 0,
  phone: '',
  email: '',
  introduction: ''
});

/** Form validation rules */
const rules = reactive({
  nickname: [
    {
      required: true,
      message: 'Please input username',
      type: 'string',
      trigger: 'blur'
    }
  ],
  sex: [
    {
      required: true,
      message: 'Please select sex',
      type: 'string',
      trigger: 'blur'
    }
  ],
  phone: [
    {
      required: true,
      message: 'Please input phone',
      type: 'string',
      trigger: 'blur'
    }
  ],
  email: [
    {
      required: true,
      message: 'Please input email',
      type: 'string',
      trigger: 'blur'
    }
  ]
});

/** Open dialog */
const openDialog = () => {
  if (!visible.value) {
    visible.value = true;
  }
};

/** Close dialog */
const cancelDialog = () => {
  resetFields();
  formRef.value?.clearValidate?.();
  visible.value = false;
};

/** Save edit */
const onSave = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    EleMessage.success('Save successful');
  });
};
</script>
