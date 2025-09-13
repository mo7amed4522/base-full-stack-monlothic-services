<!-- I18n language switch -->
<template>
  <ele-dropdown :items="[
    { title: 'Simplified Chinese', command: 'zh_CN' },
    // { title: 'Traditional Chinese', command: 'zh_TW' },
    { title: 'English', command: 'en' }
  ]" :model-value="locale" :placement="placement" :menu-style="{ minWidth: '108px' }" :popper-options="{
      modifiers: [{ name: 'offset', options: { offset: [0, 5] } }]
    }" style="line-height: inherit" @command="changeLanguage">
    <div :style="{
      height: '100%',
      display: 'flex',
      alignItems: 'center',
      outline: 'none'
    }">
      <slot>
        <el-icon>
          <GlobalOutlined :style="iconStyle" />
        </el-icon>
      </slot>
    </div>
  </ele-dropdown>
</template>

<script setup>
import { useI18n } from 'vue-i18n';
import { GlobalOutlined } from '@/components/icons';
import { setCacheLang } from '@/i18n/use-locale';

defineProps({
  /** placement */
  placement: {
    type: String,
    default: 'bottom'
  },
  /** Custom icon style */
  iconStyle: Object
});

const { locale } = useI18n();

/** Change language */
const changeLanguage = (command) => {
  locale.value = command;
  setCacheLang(command);
};
</script>
