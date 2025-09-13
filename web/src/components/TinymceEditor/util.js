import { uploadFile } from '@/api/system/file';
const BASE_URL = import.meta.env.BASE_URL;
//const BASE_URL = location.protocol + '//' + location.host + import.meta.env.BASE_URL;

/** Default plugins */
const PLUGINS = [
  'code',
  'preview',
  'fullscreen',
  'paste',
  'searchreplace',
  //'save',
  //'autosave',
  'link',
  'autolink',
  'image',
  'media',
  'table',
  'codesample',
  'lists',
  'advlist',
  'hr',
  'charmap',
  'emoticons',
  'anchor',
  'directionality',
  'pagebreak',
  'quickbars',
  'nonbreaking',
  'visualblocks',
  'visualchars',
  'wordcount'
].join(' ');

/** Default toolbar layout */
const TOOLBAR = [
  'fullscreen',
  'preview',
  'code',
  '|',
  'undo',
  'redo',
  '|',
  'forecolor',
  'backcolor',
  '|',
  'bold',
  'italic',
  'underline',
  'strikethrough',
  '|',
  'alignleft',
  'aligncenter',
  'alignright',
  'alignjustify',
  '|',
  'outdent',
  'indent',
  '|',
  'numlist',
  'bullist',
  '|',
  'formatselect',
  'fontselect',
  'fontsizeselect',
  '|',
  'link',
  'image',
  'media',
  'emoticons',
  'charmap',
  'anchor',
  'pagebreak',
  'codesample',
  '|',
  'ltr',
  'rtl'
].join(' ');

/** Default configuration */
export const DEFAULT_CONFIG = {
  height: 300,
  branding: false,
  skin_url: BASE_URL + 'tinymce/skins/ui/oxide',
  content_css: BASE_URL + 'tinymce/skins/content/default/content.min.css',
  language_url: BASE_URL + 'tinymce/langs/zh_CN.js',
  language: 'zh_CN',
  plugins: PLUGINS,
  toolbar: TOOLBAR,
  draggable_modal: true,
  toolbar_mode: 'sliding',
  quickbars_insert_toolbar: '',
  convert_urls: false,
  images_upload_handler: (blobInfo, success, error) => {
    const file = blobInfo.blob();
    /* if (file.size / 1024 / 1024 > 100) {
      error('File size cannot exceed 100MB');
      return;
    } */
    uploadFile(file, void 0, file.name)
      .then((res) => {
        success(res.url);
      })
      .catch((e) => {
        console.error(e);
        error(e.message);
      });
  },
  file_picker_types: 'media image file',
  file_picker_callback: () => {}
};

/** Dark theme configuration */
export const DARK_CONFIG = {
  skin_url: BASE_URL + 'tinymce/skins/ui/oxide-dark',
  content_css: BASE_URL + 'tinymce/skins/content/dark/content.min.css'
};

/** Supported events */
export const VALID_EVENTS = [
  'onActivate',
  'onAddUndo',
  'onBeforeAddUndo',
  'onBeforeExecCommand',
  'onBeforeGetContent',
  'onBeforeRenderUI',
  'onBeforeSetContent',
  'onBeforePaste',
  'onBlur',
  'onChange',
  'onClearUndos',
  'onClick',
  'onContextMenu',
  'onCopy',
  'onCut',
  'onDblclick',
  'onDeactivate',
  'onDirty',
  'onDrag',
  'onDragDrop',
  'onDragEnd',
  'onDragGesture',
  'onDragOver',
  'onDrop',
  'onExecCommand',
  'onFocus',
  'onFocusIn',
  'onFocusOut',
  'onGetContent',
  'onHide',
  'onInit',
  'onKeyDown',
  'onKeyPress',
  'onKeyUp',
  'onLoadContent',
  'onMouseDown',
  'onMouseEnter',
  'onMouseLeave',
  'onMouseMove',
  'onMouseOut',
  'onMouseOver',
  'onMouseUp',
  'onNodeChange',
  'onObjectResizeStart',
  'onObjectResized',
  'onObjectSelected',
  'onPaste',
  'onPostProcess',
  'onPostRender',
  'onPreProcess',
  'onProgressState',
  'onRedo',
  'onRemove',
  'onReset',
  'onSaveContent',
  'onSelectionChange',
  'onSetAttrib',
  'onSetContent',
  'onShow',
  'onSubmit',
  'onUndo',
  'onVisualAid'
];

let unique = 0;

/**
 * Generate editor id
 */
export function uuid(prefix) {
  const time = Date.now();
  const random = Math.floor(Math.random() * 1000000000);
  unique++;
  return prefix + '_' + random + unique + String(time);
}

/**
 * Bind events
 */
export function bindHandlers(initEvent, listeners, editor) {
  const validEvents = VALID_EVENTS.map((event) => event.toLowerCase());
  Object.keys(listeners)
    .filter((key) => validEvents.includes(key.toLowerCase()))
    .forEach((key) => {
      const handler = listeners[key];
      if (typeof handler === 'function') {
        if (key === 'onInit') {
          handler(initEvent, editor);
        } else {
          editor.on(key.substring(2), (e) => handler(e, editor));
        }
      }
    });
}

/**
 * Open alert dialog
 */
export function openAlert(editor, option = {}) {
  editor?.windowManager?.open?.({
    title: option.title ?? 'Alert',
    body: {
      type: 'panel',
      items: [
        {
          type: 'htmlpanel',
          html: `<p>${option.content ?? ''}</p>`
        }
      ]
    },
    buttons: [
      {
        type: 'cancel',
        name: 'closeButton',
        text: 'Sure',
        primary: true
      }
    ]
  });
}
