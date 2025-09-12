<template>
  <ele-card header="Drawer component">
    <option-item label="Direction" :responsive="false">
      <el-radio-group v-model="direction">
        <el-radio value="rtl" label="Right" />
        <el-radio value="ltr" label="Left" />
        <el-radio value="btt" label="Bottom" />
        <el-radio value="ttb" label="Top" />
      </el-radio-group>
    </option-item>
    <option-item label="Show header" :responsive="false" style="margin-top: 16px">
      <el-switch v-model="withHeader" size="small" />
    </option-item>
    <option-item label="Show footer" :responsive="false">
      <el-switch v-model="withFooter" size="small" />
    </option-item>
    <option-item label="Limit in body" :responsive="false">
      <el-switch v-model="inner" size="small" />
    </option-item>
    <div style="margin-top: 20px">
      <el-button class="ele-btn-icon" @click="openDrawer">Open drawer</el-button>
    </div>
  </ele-card>
  <ele-drawer :direction="direction" :withHeader="withHeader" :inner="inner" :size="430" title="Drawer component"
    :append-to-body="true" style="max-width: 100%" v-model="visible" :body-style="{ paddingBottom: '8px' }"
    @closed="cancelDrawer">
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
    <template v-if="withFooter" #footer>
      <el-button @click="cancelDrawer">Cancel</el-button>
      <el-button type="primary" @click="onSave">Save</el-button>
    </template>
  </ele-drawer>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';

/** Is the drawer open? */
const visible = ref(false);

/** 方向 */
const direction = ref('rtl');

/** Show header */
const withHeader = ref(true);

/** Show footer */
const withFooter = ref(true);

/** Limit in body */
const inner = ref(false);

/** Open drawer */
const openDrawer = () => {
  visible.value = true;
};

/** Form instance */
const formRef = ref(null);

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

/** Close drawer */
const cancelDrawer = () => {
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
    EleMessage.success('Save success');
  });
};
</script>
