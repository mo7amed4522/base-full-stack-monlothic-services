<!-- 机构编辑弹窗 -->
<template>
  <ele-modal form :width="550" :model-value="modelValue" :title="isUpdate ? 'Add a merchant' : 'Modify the merchant'"
    @update:modelValue="updateModelValue">
    <el-form ref="formRef" :model="form" :rules="rules" label-width="120px" @submit.prevent="">
      <el-form-item label="Username" required>
        <el-input v-model="form.content" placeholder="Please enter the template content" clearable />
      </el-form-item>
      <el-form-item label="Login password" required>
        <el-input v-model="form.content" placeholder="Please enter the template content" clearable />
      </el-form-item>
      <el-form-item label="Payment password" required>
        <el-input v-model="form.content" placeholder="Please enter the template content" clearable />
      </el-form-item>
      <el-form-item label="Is IP limited" required>
        <el-radio-group v-model="form.suoha">
          <el-radio :value="1">Open</el-radio>
          <el-radio :value="0">Close</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="Is payment enabled" required>
        <el-radio-group v-model="form.suoha">
          <el-radio :value="1">Open</el-radio>
          <el-radio :value="0">Close</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="Status">
        <el-switch v-model="form.enable" :active-value="1" :inactive-value="0" />
      </el-form-item>
      <el-form-item label="Room description" required>
        <el-input v-model="form.content" type="textarea" :rows="6" placeholder="Please enter the template content"
          clearable />
      </el-form-item>

      <el-form-item label="Game category" required>
        <el-select v-model="form.game_type" :empty-values="[null, undefined]" @change="getGameList">
          <el-option v-for="item in props.gameTypeList" :key="item.type_id" :value="item.type_id"
            :label="item.type_name" />
        </el-select>
      </el-form-item>

      <el-form-item label="" required>
        <el-select v-model="form.game_id" :empty-values="[null, undefined]">
          <el-option v-for="item in gameList" :key="item.gameid" :value="item.gameid" :label="item.title" />
        </el-select>
      </el-form-item>

      <el-form-item label="" required v-if="form.type == 0">
        <el-input v-model="form.content" type="textarea" :rows="6" placeholder="Please enter the template content"
          clearable />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="updateModelValue(false)">Cancel</el-button>
      <el-button type="primary" :loading="loading" @click="save">Save</el-button>
    </template>
  </ele-modal>

  <!--  <ele-modal form :width="620" :model-value="modelValue" :title="isUpdate ? '新增' : '修改'" @update:modelValue="updateModelValue">-->
  <!--    <el-form ref="formRef" :model="form" :rules="rules" label-width="80px" @submit.prevent="">-->
  <!--      <el-row :gutter="16">-->
  <!--        <el-col :sm="18" :xs="24">-->
  <!--          <el-form-item label="路径" prop="path">-->
  <!--            <el-input clearable :maxlength="20" v-model="form.path" placeholder="请输入机构名称" />-->
  <!--          </el-form-item>-->
  <!--          <el-form-item label="api分组">-->
  <!--            <el-select v-model="form.apiGroup">-->
  <!--              <el-option v-for="item in apiGroup" :key="item.key" :value="item.value" :label="item.value"></el-option>-->
  <!--            </el-select>-->
  <!--          </el-form-item>-->
  <!--          <el-form-item label="请求类型">-->
  <!--            <el-select v-model="form.method">-->
  <!--              <el-option value="POST" label="创建(POST)" />-->
  <!--              <el-option value="GET" label="查看(GET)" />-->
  <!--              <el-option value="PUT" label="更新(PUT)" />-->
  <!--              <el-option value="DELETE" label="删除(DELETE)" />-->
  <!--            </el-select>-->
  <!--          </el-form-item>-->
  <!--          <el-form-item label="api简介" prop="description">-->
  <!--            <el-input clearable :maxlength="100" v-model="form.description" placeholder="请输入机构全称" />-->
  <!--          </el-form-item>-->
  <!--        </el-col>-->
  <!--      </el-row>-->
  <!--    </el-form>-->
  <!--    <template #footer>-->
  <!--      <el-button @click="updateModelValue(false)">取消</el-button>-->
  <!--      <el-button type="primary" :loading="loading" @click="save">保存</el-button>-->
  <!--    </template>-->
  <!--  </ele-modal>-->
</template>

<script setup>
import { ref, reactive, watch } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';
import { addOrganization, updateOrganization } from '@/api/system/organization';


const emit = defineEmits(['done', 'update:modelValue']);

const props = defineProps({
  /** 弹窗是否打开 */
  modelValue: Boolean,
  /** 修改回显的数据 */
  data: Object,
  apiGroup: {
    type: Array,
    required: true
  }
});

/** 是否是修改 */
const isUpdate = ref(false);

/** 提交状态 */
const loading = ref(false);

/** 表单实例 */
const formRef = ref(null);

/** 表单数据 */
const [form, resetFields, assignFields] = useFormData({
  path: '',
  method: '',
  description: '',
  apiGroup: '',
  sortNumber: void 0
});

/** 表单验证规则 */
const rules = reactive({
  organizationName: [
    {
      required: true,
      message: 'Please enter the organization name',
      type: 'string',
      trigger: 'blur'
    }
  ],
  organizationFullName: [
    {
      required: true,
      message: 'Please enter the organization full name',
      type: 'string',
      trigger: 'blur'
    }
  ],
  organizationType: [
    {
      required: true,
      message: 'Please select the organization type',
      type: 'string',
      trigger: 'blur'
    }
  ],
  sortNumber: [
    {
      required: true,
      message: 'Please enter the sort number',
      type: 'number',
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
    const saveOrUpdate = isUpdate.value ? updateOrganization : addOrganization;
    saveOrUpdate({ ...form, parentId: form.parentId || 0 })
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

watch(
  () => props.modelValue,
  (modelValue) => {
    if (modelValue) {
      if (props.data) {
        assignFields({
          ...props.data,
          parentId: props.data.parentId || void 0
        });
        isUpdate.value = true;
      } else {
        form.parentId = props.organizationId;
        isUpdate.value = false;
      }
    } else {
      resetFields();
      formRef.value?.clearValidate?.();
    }
  }
);
</script>
