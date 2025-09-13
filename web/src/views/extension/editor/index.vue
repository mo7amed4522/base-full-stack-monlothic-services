<template>
  <ele-page :multi-card="true">
    <ele-card>
      <!-- button -->
      <div style="margin-bottom: 16px">
        <el-button type="primary" class="ele-btn-icon" @click="setContent">
          Set content
        </el-button>
        <el-button type="primary" class="ele-btn-icon" @click="showHtml">
          Get content
        </el-button>
        <el-button class="ele-btn-icon" :type="disabled ? 'primary' : 'danger'" @click="toggleDisabled">
          {{ disabled ? 'Enable' : 'Disable' }}
        </el-button>
      </div>
      <!-- editor -->
      <tinymce-editor ref="editorRef" :init="config" v-model="content" :disabled="disabled" />
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref } from 'vue';
import { ElMessageBox } from 'element-plus/es';
import TinymceEditor from '@/components/TinymceEditor/index.vue';

/** editor */
const editorRef = ref(null);

/** editor config */
const config = ref({
  height: 520,
  // editor custom file upload (here using converting selected file to blob for demonstration)
  file_picker_callback: (callback, _value, meta) => {
    const input = document.createElement('input');
    input.setAttribute('type', 'file');
    // set file type
    if (meta.filetype === 'image') {
      input.setAttribute('accept', 'image/*');
    } else if (meta.filetype === 'media') {
      input.setAttribute('accept', 'video/*');
    }
    input.onchange = () => {
      const file = input.files?.[0];
      if (!file) {
        return;
      }
      if (meta.filetype === 'media') {
        if (!file.type.startsWith('video/')) {
          editorRef.value?.alert?.({ content: 'Only video files are allowed' });
          return;
        }
      }
      if (file.size / 1024 / 1024 > 20) {
        editorRef.value?.alert?.({ content: 'Size cannot exceed 20MB' });
        return;
      }
      // actual project should upload file to backend and get file url instead of converting file to blob
      const reader = new FileReader();
      reader.onload = (e) => {
        if (e.target?.result != null) {
          const blob = new Blob([e.target.result], { type: file.type });
          callback(URL.createObjectURL(blob), {
            text: file.name,
            alt: file.name
          });
        }
      };
      reader.readAsArrayBuffer(file);
    };
    input.click();
  }
});

/** editor content */
const content = ref('');

/** editor disabled */
const disabled = ref(false);

/** get editor content */
const showHtml = () => {
  ElMessageBox.alert(content.value, 'Editor content', {
    confirmButtonText: 'knew',
    closeOnClickModal: true,
    callback: () => { },
    draggable: true
  });
};

/** set editor content */
const setContent = () => {
  content.value = [
    '<div style="text-align: center;color: #fff;background-image: linear-gradient( 135deg, #4C83FF 10%, #2AFADF 100%);padding: 32px 0;">',
    '   <div style="font-size: 28px;margin-bottom: 16px;">EleAdminPlus backend management template</div>',
    '   <div style="font-size:18px">Universal backend management template with beautiful interface, ready to use out of the box, and rich components and templates</div>',
    '</div><br/>'
  ].join('');
};

/** toggle editor disabled */
const toggleDisabled = () => {
  disabled.value = !disabled.value;
};
</script>

<script>
export default {
  name: 'ExtensionEditor'
};
</script>
