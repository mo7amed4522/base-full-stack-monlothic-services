<!-- markdown parsing -->
<template>
  <!-- eslint-disable vue/no-v-html -->
  <div ref="rootRef" v-html="content" class="markdown-body" @click="handleClick">
  </div>
</template>

<script setup>
import { ref, watch, onMounted, onBeforeUnmount, nextTick } from 'vue';
import { getProcessor } from 'bytemd';

const props = defineProps({
  value: String,
  config: Object
});

/** Root Node */
const rootRef = ref(null);

/** Parsed Content */
const content = ref(null);

/** Bound Plugins */
const cbs = ref([]);

/** Bind Plugin */
const on = () => {
  if (props.config?.plugins && rootRef.value && content.value) {
    cbs.value = props.config.plugins.map(({ viewerEffect }) => {
      return (
        viewerEffect &&
        viewerEffect({
          markdownBody: rootRef.value,
          file: content.value
        })
      );
    });
  }
};

/** Unbind Plugin */
const off = () => {
  if (cbs.value) {
    cbs.value.forEach((cb) => cb && cb());
  }
};

/** Handle Click Event for Anchor Links */
const handleClick = (e) => {
  const $ = e.target;
  if ($.tagName !== 'A') {
    return;
  }
  const href = $.getAttribute('href');
  if (!href || !href.startsWith('#')) {
    return;
  }
  const $root = rootRef.value;
  if ($root) {
    const dest = $root.querySelector('#user-content-' + href.slice(1));
    if (dest) {
      dest.scrollIntoView();
    }
  }
};

// Parse Markdown
watch(
  [() => props.value, () => props.config],
  () => {
    try {
      content.value = getProcessor(
        Object.assign({}, props.config)
      ).processSync(props.value);
    } catch (e) {
      console.error(e);
    }
    off();
    nextTick(() => {
      on();
    });
  },
  {
    immediate: true,
    deep: true
  }
);

onMounted(() => {
  on();
});

onBeforeUnmount(() => {
  off();
});
</script>
