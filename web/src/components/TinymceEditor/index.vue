<!-- Customize tinymce editor -->
<template>
  <component v-if="inlineEditor" :is="tagName" :id="elementId" />
  <textarea v-else :id="elementId"></textarea>
  <!-- Customize file picker popup -->
  <FilePicker title="File Selection" :limit="1" :accept="filePickerAccept ? filePickerAccept + '*' : void 0"
    :params="{ contentType: filePickerAccept }" :base-index="19990099" v-model="showFilePicker"
    @done="handleFilePickerDone" @close="handleFilePickerClose" />
</template>

<script setup>
import {
  ref,
  watch,
  onMounted,
  onBeforeUnmount,
  onActivated,
  onDeactivated,
  nextTick,
  useAttrs
} from 'vue';
import FilePicker from '@/components/FilePicker/index.vue';
import tinymce from 'tinymce/tinymce';
import 'tinymce/themes/silver';
import 'tinymce/icons/default';
import 'tinymce/plugins/code';
import 'tinymce/plugins/preview';
import 'tinymce/plugins/fullscreen';
import 'tinymce/plugins/paste';
import 'tinymce/plugins/searchreplace';
//import 'tinymce/plugins/save';
//import 'tinymce/plugins/autosave';
import 'tinymce/plugins/link';
import 'tinymce/plugins/autolink';
import 'tinymce/plugins/image';
import 'tinymce/plugins/media';
import 'tinymce/plugins/table';
import 'tinymce/plugins/codesample';
import 'tinymce/plugins/lists';
import 'tinymce/plugins/advlist';
import 'tinymce/plugins/hr';
import 'tinymce/plugins/charmap';
import 'tinymce/plugins/emoticons';
import 'tinymce/plugins/anchor';
import 'tinymce/plugins/directionality';
import 'tinymce/plugins/pagebreak';
import 'tinymce/plugins/quickbars';
import 'tinymce/plugins/nonbreaking';
import 'tinymce/plugins/visualblocks';
import 'tinymce/plugins/visualchars';
import 'tinymce/plugins/wordcount';
import 'tinymce/plugins/emoticons/js/emojis';
import { storeToRefs } from 'pinia';
import { useThemeStore } from '@/store/modules/theme';
import {
  DEFAULT_CONFIG,
  DARK_CONFIG,
  uuid,
  bindHandlers,
  openAlert
} from './util';

defineOptions({ name: 'TinymceEditor' });

const props = defineProps({
  /** Editor unique id */
  id: String,
  /** v-model */
  modelValue: String,
  /** Editor configuration */
  init: Object,
  /** Whether to use inline mode */
  inline: Boolean,
  /** model events */
  modelEvents: {
    type: String,
    default: 'change input undo redo'
  },
  /** Inline mode tag name */
  tagName: {
    type: String,
    default: 'div'
  },
  /** Whether to disable */
  disabled: Boolean,
  /** Whether to follow the theme of the framework */
  autoTheme: {
    type: Boolean,
    default: true
  },
  /** Whether to use dark theme when not following the theme of the framework */
  darkTheme: Boolean
});

const emit = defineEmits(['update:modelValue']);

const attrs = useAttrs();
const themeStore = useThemeStore();
const { darkMode } = storeToRefs(themeStore);

/** Editor unique id */
const elementId = props.id || uuid('tiny-vue');

/** Editor instance */
let editorIns = null;

/** Whether to use inline mode */
const inlineEditor = props.init?.inline || props.inline;

/** Update value */
const updateValue = (value) => {
  emit('update:modelValue', value);
};

/** Set content */
const setContent = (value) => {
  if (
    editorIns &&
    typeof value === 'string' &&
    value !== editorIns.getContent()
  ) {
    editorIns.setContent(value);
  }
};

/** Render editor */
const render = () => {
  const isDark = props.autoTheme ? darkMode.value : props.darkTheme;
  tinymce.init({
    ...DEFAULT_CONFIG,
    ...(isDark ? DARK_CONFIG : {}),
    ...props.init,
    selector: `#${elementId}`,
    readonly: props.disabled,
    inline: inlineEditor,
    setup: (editor) => {
      editorIns = editor;
      editor.on('init', (e) => {
        // Echo initial value
        if (props.modelValue) {
          setContent(props.modelValue);
        }
        // v-model
        editor.on(props.modelEvents, () => {
          updateValue(editor.getContent());
        });
        // valid events
        bindHandlers(e, attrs, editor);
      });
      if (typeof props.init?.setup === 'function') {
        props.init.setup(editor);
      }
    },
    // File upload calls file selector
    file_picker_callback: (callback, _value, meta) => {
      if (meta.filetype === 'image') {
        filePickerAccept.value = 'image/';
      } else if (meta.filetype === 'media') {
        filePickerAccept.value = 'video/';
      } else {
        filePickerAccept.value = void 0;
      }
      showFilePicker.value = true;
      currentEditorCallback = callback;
    }
  });
};

/** Destroy editor */
const destory = () => {
  if (tinymce != null && editorIns != null) {
    tinymce.remove(editorIns);
    editorIns = null;
  }
};

/** Show alert */
const alert = (option) => {
  openAlert(editorIns, option);
};

defineExpose({ editorIns, alert });

watch(
  () => props.modelValue,
  (val, prevVal) => {
    if (val !== prevVal) {
      setContent(val || '');
    }
  }
);

watch(
  () => props.disabled,
  (disable) => {
    if (editorIns !== null) {
      if (typeof editorIns.mode?.set === 'function') {
        editorIns.mode.set(disable ? 'readonly' : 'design');
      } else {
        editorIns.setMode(disable ? 'readonly' : 'design');
      }
    }
  }
);

watch(
  () => props.tagName,
  () => {
    destory();
    nextTick(() => {
      render();
    });
  }
);

watch(darkMode, () => {
  if (props.autoTheme) {
    destory();
    nextTick(() => {
      render();
    });
  }
});

onMounted(() => {
  render();
});

onBeforeUnmount(() => {
  destory();
});

onActivated(() => {
  render();
});

onDeactivated(() => {
  destory();
});

/** Whether to show file picker */
const showFilePicker = ref(false);

/** File picker accept */
const filePickerAccept = ref();

/** Current editor callback */
let currentEditorCallback = null;

/** File picker done */
const handleFilePickerDone = (files) => {
  showFilePicker.value = false;
  if (files.length) {
    currentEditorCallback(files[0].url, {
      text: files[0].name,
      alt: files[0].name
    });
  }
  currentEditorCallback = null;
};

/** File picker close */
const handleFilePickerClose = () => {
  currentEditorCallback = null;
};
</script>

<style>
body .tox.tox-tinymce-aux,
body.tox-fullscreen .tox.tox-tinymce-aux {
  z-index: 19990000;
}

.tox-menu.tox-collection.tox-collection--list {
  max-height: 420px !important;
}

textarea[id^='tiny-vue'] {
  width: 0;
  height: 0;
  margin: 0;
  padding: 0;
  opacity: 0;
  box-sizing: border-box;
}
</style>
