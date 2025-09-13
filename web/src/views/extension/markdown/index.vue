<template>
  <ele-page>
    <ele-card header="Markdown editor">
      <byte-md-editor v-model="content" height="520px" :config="config" />
    </ele-card>
    <ele-card header="Markdown parser">
      <byte-md-viewer :value="content" :config="{ plugins: [gfm(), highlight()] }"
        style="max-width: 880px; margin: 0 auto; padding: 24px" />
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref, reactive } from 'vue';
import ByteMdEditor from '@/components/ByteMdEditor/index.vue';
import ByteMdViewer from '@/components/ByteMdViewer/index.vue';
/** Markdown editor language file */
import zh_Hans from 'bytemd/locales/zh_Hans.json';
/** Link/Strikethrough/Checkbox/Table and other plug-ins */
import gfm from '@bytemd/plugin-gfm';
/** Plug-in Chinese language file */
import zh_HansGfm from '@bytemd/plugin-gfm/locales/zh_Hans.json';
/** Code highlighting plug-in */
import highlight from '@bytemd/plugin-highlight';
/** Code highlighting plug-in style */
import 'highlight.js/styles/github-dark.css';
//import 'highlight.js/styles/github.css';
/** Preview interface style, using github markdown theme */
//import 'github-markdown-css/github-markdown-dark.css';
import 'github-markdown-css/github-markdown-light.css';
import { getDemoMD } from '@/api/example';

/** Editor content */
const content = ref('');

/** Configuration */
const config = reactive({
  locale: zh_Hans,
  plugins: [gfm({ locale: zh_HansGfm }), highlight()],
  editorConfig: { lineNumbers: true },
  uploadImages: async (files) => {
    const result = [];
    files.forEach((file) => {
      result.push({
        url: 'https://cdn.eleadmin.com/20200610/4Z0QR2L0J1XStxBh99jVJ8qLfsGsOgjU.jpg',
        alt: file.name,
        title: ''
      });
    });
    return result;
  }
});

//
getDemoMD()
  .then((res) => {
    content.value = res;
  })
  .catch((e) => {
    console.error(e);
  });
</script>

<script>
export default {
  name: 'ExtensionMarkdown'
};
</script>

<style lang="scss">
.markdown-body .highlight pre,
.markdown-body pre {
  color: #e6edf3;
  background-color: #161b22;
}
</style>
