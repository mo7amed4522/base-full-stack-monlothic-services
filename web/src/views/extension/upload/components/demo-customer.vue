<template>
  <ele-card header="Custom Upload">
    <div style="margin-bottom: 24px">
      Custom upload component simplifies the process of handling upload events and converting data formats. It can be
      used just like a normal input field.
    </div>
    <el-form ref="formRef" :model="form" :rules="rules" label-width="100px" style="max-width: 460px" @submit.prevent="">
      <el-form-item label="Real Name" prop="realName">
        <el-input clearable :maxlength="20" v-model="form.realName" placeholder="Please input real name" />
      </el-form-item>
      <el-form-item label="ID Card Front" prop="idCardImg1">
        <image-upload :limit="1" v-model="form.idCardImg1" :item-style="{ width: '260px', height: '160px', margin: 0 }"
          :button-style="{ width: '260px', height: '160px', margin: 0 }" />
      </el-form-item>
      <el-form-item label="ID Card Back" prop="idCardImg2">
        <image-upload :limit="1" v-model="form.idCardImg2" :item-style="{ width: '260px', height: '160px', margin: 0 }"
          :button-style="{ width: '260px', height: '160px', margin: 0 }" />
      </el-form-item>
      <el-form-item label="Complaint Content" prop="content">
        <el-input :rows="4" type="textarea" v-model="form.content" placeholder="Please input complaint content" />
      </el-form-item>
      <el-form-item label="Scene Photos" prop="images">
        <image-upload :limit="3" v-model="form.images" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="save">Submit</el-button>
        <el-button type="primary" @click="setData">Set Data</el-button>
      </el-form-item>
    </el-form>
  </ele-card>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';
import ImageUpload from '@/components/ImageUpload/index.vue';

/** 表单实例 */
const formRef = ref(null);

/** 表单数据 */
const [form, _resetFields, assignFields] = useFormData({
  id: void 0,
  realName: '',
  /** 身份证正面 */
  idCardImg1: '',
  /** 身份证反面 */
  idCardImg2: '',
  content: '',
  /** 现场照片 */
  images: ''
});

/** 表单验证规则 */
const rules = reactive({
  realName: [
    {
      required: true,
      message: 'Please input real name',
      type: 'string',
      trigger: 'blur'
    }
  ],
  idCardImg1: [
    {
      required: true,
      message: 'Please upload ID card front',
      type: 'string',
      trigger: 'change'
    },
    {
      type: 'string',
      trigger: 'change',
      validator: (_rule, value, callback) => {
        if (value && value.startsWith('blob:')) {
          return callback(new Error('Image is not uploaded yet'));
        }
        callback();
      }
    }
  ],
  idCardImg2: [
    {
      required: true,
      message: 'Please upload ID card back',
      type: 'string',
      trigger: 'change'
    },
    {
      type: 'string',
      trigger: 'change',
      validator: (_rule, value, callback) => {
        if (value && value.startsWith('blob:')) {
          return callback(new Error('Image is not uploaded yet'));
        }
        callback();
      }
    }
  ],
  content: [
    {
      required: true,
      message: 'Please input complaint content',
      type: 'string',
      trigger: 'blur'
    }
  ],
  images: [
    {
      required: true,
      message: 'Please upload scene photos',
      type: 'string',
      trigger: 'change'
    },
    {
      type: 'string',
      trigger: 'change',
      validator: (_rule, value, callback) => {
        if (
          value &&
          JSON.parse(value).some((d) => !d || d.startsWith('blob:'))
        ) {
          return callback(new Error('Image is not uploaded yet'));
        }
        callback();
      }
    }
  ]
});

/** 保存编辑 */
const save = () => {
  formRef.value?.validate?.((valid) => {
    if (!valid) {
      return;
    }
    console.log(JSON.parse(JSON.stringify(form)));
    EleMessage.success('Data has been printed in the console');
  });
};

/** 回显数据 */
const setData = () => {
  const data = {
    id: 1,
    realName: 'Zhang San',
    idCardImg1:
      'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg',
    idCardImg2:
      'https://cdn.eleadmin.com/20200610/WLXm7gp1EbLDtvVQgkeQeyq5OtDm00Jd.jpg',
    content: '2栋楼下的路灯已经坏了几天了都没人来维修一下',
    images:
      '["https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg","https://cdn.eleadmin.com/20200610/WLXm7gp1EbLDtvVQgkeQeyq5OtDm00Jd.jpg"]'
  };
  assignFields(data);
  formRef.value?.clearValidate?.();
};
</script>
