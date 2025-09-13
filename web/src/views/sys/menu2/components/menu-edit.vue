<!-- 编辑弹窗 -->
<template>
  <ele-modal form :width="740" :model-value="modelValue" :title="isUpdate ? 'Edit Menu' : 'New Menu'"
    @update:modelValue="updateModelValue">
    <el-form ref="formRef" :model="form" :rules="rules" label-width="92px" @submit.prevent="">
      <el-row :gutter="16">
        <el-col :sm="12" :xs="24">
          <el-form-item label="Parent Menu" prop="parentId">
            <menu-select v-model="form.parentId" />
          </el-form-item>
          <el-form-item label="Menu Name" prop="title">
            <el-input clearable :maxlength="20" v-model="form.title" placeholder="Enter menu name" />
          </el-form-item>
        </el-col>
        <el-col :sm="12" :xs="24">
          <el-form-item label="Menu Type" prop="menuType">
            <el-radio-group v-model="form.menuType" @change="onMenuTypeChange">
              <el-radio :value="0" label="Directory" />
              <el-radio :value="1" label="Menu" />
              <el-radio :value="2" label="Button" />
            </el-radio-group>
          </el-form-item>
          <el-form-item label="Open Type">
            <el-radio-group v-model="form.openType" :disabled="form.menuType === 0 || form.menuType === 2"
              @change="onOpenTypeChange">
              <el-radio :value="0" label="Component" />
              <el-radio :value="1" label="Internal Link" />
              <el-radio :value="2" label="External Link" />
            </el-radio-group>
          </el-form-item>
        </el-col>
      </el-row>
      <el-divider style="margin: 8px 0 22px 0; opacity: 0.6" />
      <el-row :gutter="16">
        <el-col :sm="12" :xs="24">
          <el-form-item label="Menu Icon" prop="icon">
            <icon-select v-model="form.icon" :disabled="form.menuType === 2" />
          </el-form-item>
          <el-form-item prop="path">
            <template #label>
              <ele-tooltip v-if="form.openType === 2"
                content="External link address must start with `http://`, `https://`, or `//`">
                <el-icon :size="15" style="align-self: center; margin-right: 4px; cursor: help">
                  <QuestionCircleOutlined style="opacity: 0.6" />
                </el-icon>
              </ele-tooltip>
              <span>{{ form.openType === 2 ? 'External Link Address' : 'Route Address' }}</span>
            </template>
            <el-input clearable :maxlength="100" v-model="form.path" :disabled="form.menuType === 2" :placeholder="form.openType === 2 ? 'Enter external link address' : 'Enter route address'
              " />
          </el-form-item>
          <el-form-item prop="component">
            <template #label>
              <ele-tooltip v-if="form.openType === 1"
                content="Internal link address must start with `http://`, `https://`, or `//`">
                <el-icon :size="15" style="align-self: center; margin-right: 4px; cursor: help">
                  <QuestionCircleOutlined style="opacity: 0.6" />
                </el-icon>
              </ele-tooltip>
              <span>{{ form.openType === 1 ? 'Internal Link Address' : 'Component Path' }}</span>
            </template>
            <el-input clearable :maxlength="100" v-model="form.component" :disabled="form.menuType === 0 ||
              form.menuType === 2 ||
              form.openType === 2
              " :placeholder="form.openType === 1 ? 'Enter internal link address' : 'Enter component path'
                " />
          </el-form-item>
        </el-col>
        <el-col :sm="12" :xs="24">
          <el-form-item label="Authority Identifier" prop="authority">
            <el-input clearable v-model="form.authority" placeholder="Enter authority identifier" :disabled="form.menuType === 0 ||
              (form.menuType === 1 && form.openType === 2)
              " />
          </el-form-item>
          <el-form-item label="Sort Number" prop="sortNumber">
            <el-input-number :min="0" :max="99999" v-model="form.sortNumber" placeholder="Enter sort number"
              controls-position="right" class="ele-fluid" />
          </el-form-item>
          <el-form-item label="Is Visible">
            <el-switch inline-prompt active-text="Yes" inactive-text="No" :model-value="form.hidden === false"
              @change="updateHideValue" />
            <ele-tooltip
              content="Select not visible to register the route but not show in the sidebar, such as adding a page should select not visible"
              :popper-style="{ maxWidth: '240px' }">
              <el-icon :size="15" style="margin-left: 16px; cursor: help">
                <QuestionCircleOutlined style="opacity: 0.6" />
              </el-icon>
            </ele-tooltip>
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item label="Route Meta" prop="meta">
        <el-input :rows="4" type="textarea" :maxlength="200" v-model="form.meta"
          placeholder="Enter JSON format route meta" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="updateModelValue(false)">Cancel</el-button>
      <el-button type="primary" :loading="loading" @click="save">
        Save
      </el-button>
    </template>
  </ele-modal>
</template>

<script setup>
import { ref, reactive, watch } from 'vue';
import { EleMessage, isExternalLink } from 'ele-admin-plus/es';
import { QuestionCircleOutlined } from '@/components/icons';
import { useFormData } from '@/utils/use-form-data';
import MenuSelect from './menu-select.vue';
import IconSelect from './icon-select.vue';
import { addMenu, updateMenu } from '@/api/system/menu';

const emit = defineEmits(['done', 'update:modelValue']);

const props = defineProps({
  /** 弹窗是否打开 */
  modelValue: Boolean,
  /** 修改回显的数据 */
  data: Object,
  /** 添加时上级菜单id */
  parentId: Number
});

/** 是否是修改 */
const isUpdate = ref(false);

/** 提交状态 */
const loading = ref(false);

/** 表单实例 */
const formRef = ref(null);

/** 表单数据 */
const [form, resetFields, assignFields] = useFormData({
  menuId: void 0,
  parentId: void 0,
  title: '',
  menuType: 0,
  openType: 0,
  icon: '',
  path: '',
  component: '',
  authority: '',
  sortNumber: void 0,
  hidden: 0,
  meta: ''
});

/** 表单验证规则 */
const rules = reactive({
  title: [
    {
      required: true,
      type: 'string',
      message: 'Enter menu name',
      trigger: 'blur'
    }
  ],
  sortNumber: [
    {
      required: true,
      type: 'number',
      message: 'Enter sort number',
      trigger: 'blur'
    }
  ],
  meta: [
    {
      type: 'string',
      validator: (_rule, value, callback) => {
        if (value) {
          const msg = 'Enter correct JSON format';
          try {
            const obj = JSON.parse(value);
            if (obj == null || typeof obj !== 'object') {
              return callback(msg);
            }
          } catch (_e) {
            return callback(msg);
          }
        }
        callback();
      },
      trigger: 'blur'
    }
  ]
});

/** 保存编辑 */
const save = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    loading.value = true;
    const saveOrUpdate = isUpdate.value ? updateMenu : addMenu;
    saveOrUpdate({
      ...form,
      // menuType 对应的值与后端不一致在前端处理
      menuType: form.menuType === 2 ? 1 : 0,
      parentId: form.parentId || 0
    })
      .then((msg) => {
        loading.value = false;
        EleMessage.success(msg);
        updateModelValue(false);
        emit('done');
      })
      .catch((e) => {
        loading.value = false;
        EleMessage.error(e.message);
      });
  });
};

/** 更新modelValue */
const updateModelValue = (value) => {
  emit('update:modelValue', value);
};

/** menuType选择改变 */
const onMenuTypeChange = () => {
  if (form.menuType === 0) {
    form.authority = '';
    form.openType = 0;
    form.component = '';
  } else if (form.menuType === 1) {
    if (form.openType === 2) {
      form.authority = '';
    }
  } else {
    form.openType = 0;
    form.icon = '';
    form.path = '';
    form.component = '';
    form.hidden = 0;
  }
};

/** openType选择改变 */
const onOpenTypeChange = () => {
  if (form.openType === 2) {
    form.component = '';
    form.authority = '';
  }
};

/** hide选择改变 */
const updateHideValue = (value) => {
  form.hidden = !value;
};

/** 判断是否是目录 */
const isDirectory = (d) => {
  return !!d.children?.length && !d.component;
};

watch(
  () => props.modelValue,
  (modelValue) => {
    if (modelValue) {
      if (props.data) {
        const isExternal = isExternalLink(props.data.path);
        const isInner = isExternalLink(props.data.component);
        // menuType对应的值与后端不一致在前端处理
        const menuType =
          props.data.menuType === 1 ? 2 : isDirectory(props.data) ? 0 : 1;
        assignFields({
          ...props.data,
          menuType,
          openType: isExternal ? 2 : isInner ? 1 : 0,
          parentId: props.data.parentId === 0 ? void 0 : props.data.parentId
        });
        isUpdate.value = true;
      } else {
        form.parentId = props.parentId;
        isUpdate.value = false;
      }
    } else {
      resetFields();
      formRef.value?.clearValidate?.();
    }
  }
);
</script>
