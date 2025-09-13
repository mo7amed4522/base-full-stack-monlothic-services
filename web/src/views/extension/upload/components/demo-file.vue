<template>
  <ele-card header="附件上传">
    <div style="margin-bottom: 24px">
      It can be used not only for image upload scenarios, but also for file upload scenarios
    </div>
    <el-form ref="formRef" :model="form" :rules="rules" label-width="100px" style="max-width: 460px" @submit.prevent="">
      <el-form-item label="Complaint content" prop="content">
        <el-input :rows="4" type="textarea" v-model="form.content" placeholder="Please input complaint content" />
      </el-form-item>
      <el-form-item label="Attachment" prop="files">
        <file-upload :limit="8" v-model="form.files" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="save">Submit</el-button>
        <el-button type="primary" @click="setData">Display data</el-button>
      </el-form-item>
    </el-form>
  </ele-card>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';
import { useFormData } from '@/utils/use-form-data';
import FileUpload from '@/components/FileUpload/index.vue';

/** 表单实例 */
const formRef = ref(null);

/** 表单数据 */
const [form, _resetFields, assignFields] = useFormData({
  id: void 0,
  content: '',
  /** 附件 */
  files: ''
});

/** 表单验证规则 */
const rules = reactive({
  content: [
    {
      required: true,
      message: 'Please input complaint content',
      type: 'string',
      trigger: 'blur'
    }
  ],
  files: [
    {
      required: true,
      message: 'Please upload scene attachment',
      type: 'string',
      trigger: 'change'
    },
    {
      type: 'string',
      trigger: 'change',
      validator: (_rule, value, callback) => {
        if (
          value &&
          JSON.parse(value).some((d) => !d.url || d.url.startsWith('blob:'))
        ) {
          return callback(new Error('Scene attachment has not been uploaded yet'));
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
    content: 'The street light at the bottom of 2nd floor has been broken for a few days now and no one has come to repair it',
    files:
      '[{"url":"https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg","name":"c184eef391ae48dba87e3057e70238fb.jpg"},{"url":"https://cdn.eleadmin.com/20200610/WLXm7gp1EbLDtvVQgkeQeyq5OtDm00Jd.jpg","name":"WLXm7gp1EbLDtvVQgkeQeyq5OtDm00Jd.jpg"},{"url":"https://cdn.eleadmin.com/20200610/20200708224450.pdf","name":"20200708224450.pdf"}]'
  };
  assignFields(data);
  formRef.value?.clearValidate?.();
};
</script>
