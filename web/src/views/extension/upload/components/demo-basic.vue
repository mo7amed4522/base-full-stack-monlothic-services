<template>
  <ele-card header="Basic Example">
    <ele-upload-list :limit="8" :drag="true" v-model="images" :readonly="readonly" @upload="onUpload"
      @retry="(item) => onUpload(item, true)" @remove="onRemove" />
    <div style="display: flex; align-items: center">
      <el-button type="primary" @click="getData">Get Data</el-button>
      <div style="line-height: 22px; margin-left: 22px">Readonly:</div>
      <div style="flex: 1; padding-left: 12px">
        <el-radio-group v-model="readonly">
          <el-radio :value="false" label="No" />
          <el-radio :value="true" label="Yes" />
        </el-radio-group>
      </div>
    </div>
  </ele-card>
</template>

<script setup>
import { ref } from 'vue';
import { EleMessage } from 'ele-admin-plus/es';

/** 数据 */
const images = ref([
  {
    key: 1,
    name: 'c184eef391ae48dba87e3057e70238fb.jpg',
    url: 'https://cdn.eleadmin.com/20200609/c184eef391ae48dba87e3057e70238fb.jpg',
    status: 'done'
  },
  {
    key: 2,
    name: 'WLXm7gp1EbLDtvVQgkeQeyq5OtDm00Jd.jpg',
    url: 'https://cdn.eleadmin.com/20200610/WLXm7gp1EbLDtvVQgkeQeyq5OtDm00Jd.jpg',
    status: 'done'
  },
  {
    key: 3,
    name: 'f6bc05af944a4f738b54128717952107.jpg',
    url: 'https://cdn.eleadmin.com/20200609/f6bc05af944a4f738b54128717952107.jpg',
    status: 'done'
  }
]);

/** 只读 */
const readonly = ref(false);

/** 获取当前数据 */
const getData = () => {
  console.log(JSON.parse(JSON.stringify(images.value)));
  EleMessage.success('Data has been printed in the console');
};

/** 上传事件 */
const onUpload = (d, retry) => {
  if (!d.file) {
    return;
  }
  if (!d.file.type.startsWith('image')) {
    EleMessage.error('Only image files are allowed');
    return;
  }
  if (d.file.size / 1024 / 1024 > 2) {
    EleMessage.error('Size cannot exceed 2MB');
    return;
  }
  if (!retry) {
    images.value.push({ ...d });
  }
  const item = images.value.find((t) => t.key === d.key);
  console.log(JSON.parse(JSON.stringify(item)));
  if (!item) {
    return;
  }
  // 模拟上传
  item.status = 'uploading';
  if (item.progress == null) {
    item.progress = 20;
  }
  const timer = setInterval(() => {
    if (Math.round(Math.random() * 10) < 3) {
      item.status = 'exception';
      EleMessage.error('Upload failed, server busy');
      clearInterval(timer);
      return;
    }
    if (item.progress == null) {
      item.progress = 20;
    } else {
      item.progress += 20;
    }
    if (item.progress === 100) {
      item.status = 'done';
      clearInterval(timer);
    }
  }, 1000);
};

/** 删除事件 */
const onRemove = (item) => {
  images.value.splice(images.value.indexOf(item), 1);
};
</script>
