<template>
  <el-form ref="formRef" :model="form" :rules="rules" label-width="108px" style="padding-right: 60px"
    @submit.prevent="">
    <el-form-item label="Payment Account" prop="account">
      <el-select clearable v-model="form.account" placeholder="Please select payment account" class="ele-fluid">
        <el-option value="plus@eleadmin.com" label="plus@eleadmin.com" />
        <el-option value="pro@eleadmin.com" label="pro@eleadmin.com" />
      </el-select>
    </el-form-item>
    <el-form-item label="Payment Account" prop="receiver">
      <el-input clearable v-model="form.receiver" placeholder="Please input payment account">
        <template #prepend>
          <el-select v-model="form.pay" style="width: 100px">
            <el-option value="alipay" label="Alipay" />
            <el-option value="wxpay" label="WeChat Pay" />
          </el-select>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item label="Payee Name" prop="name">
      <el-input clearable v-model="form.name" placeholder="Please input payee name" />
    </el-form-item>
    <el-form-item label="Transfer Amount" prop="amount">
      <el-input clearable v-model="form.amount" placeholder="Please input transfer amount" :prefix-icon="prefixIcon" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" :loading="loading" @click="submit">
        Next Step
      </el-button>
    </el-form-item>
  </el-form>
</template>

<script setup>
import { ref, reactive, h } from 'vue';
import { numberReg } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';

const prefixIcon = h('span', { style: { fontStyle: 'normal' } }, '￥');

const emit = defineEmits(['done']);

/** 提交状态 */
const loading = ref(false);

/** 表单实例 */
const formRef = ref(null);

/** 表单数据 */
const [form] = useFormData({
  account: 'plus@eleadmin.com',
  receiver: 'test@example.com',
  pay: 'alipay',
  name: 'Alex',
  amount: '500'
});

/** 表单验证规则 */
const rules = reactive({
  account: [
    {
      required: true,
      message: 'Please select payment account',
      type: 'string',
      trigger: 'blur'
    }
  ],
  receiver: [
    {
      required: true,
      message: 'Please input payment account',
      type: 'string',
      trigger: 'blur'
    }
  ],
  name: [
    {
      required: true,
      message: 'Please input payee name',
      type: 'string',
      trigger: 'blur'
    }
  ],
  amount: [
    {
      required: true,
      message: 'Please input transfer amount',
      type: 'string',
      trigger: 'blur'
    },
    {
      pattern: numberReg,
      message: 'Please input legal amount number',
      type: 'string',
      trigger: 'blur'
    }
  ]
});

/** 步骤一提交 */
const submit = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    loading.value = true;
    setTimeout(() => {
      loading.value = false;
      emit('done', form);
    }, 300);
  });
};
</script>
