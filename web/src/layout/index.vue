<template>
  <ele-pro-layout :menus="menus" :tabs="tabs" :collapse="collapse" :compact="compact" :maximized="maximized"
    :tab-bar="tabBar ? (tabInHeader ? 'header' : true) : false"
    :breadcrumb="layout === 'default' && (!tabBar || !tabInHeader)" :layout="layout" :sidebar-layout="sidebarLayout"
    :header-style="headerStyle" :sidebar-style="sidebarStyle" :tab-style="tabStyle" :fixed-header="fixedHeader"
    :fixed-sidebar="fixedSidebar" :fixed-body="fixedBody" :fluid="fluid" :logo-in-header="logoInHeader"
    :colorful-icon="colorfulIcon" :unique-opened="uniqueOpened" :fixed-home="fixedHome" :home-path="HOME_PATH"
    :redirectPath="REDIRECT_PATH" :locale="locale" :i18n="i18n" :tab-sortable="!mobileDevice" :tab-context-menu="{
      iconProps: { size: 15 },
      popperOptions: { strategy: 'fixed' }
    }" :tab-context-menus="tabContext" :nav-trigger="layout === 'top' ? void 0 : menuItemTrigger"
    :box-trigger="menuItemTrigger" :keep-alive="TAB_KEEP_ALIVE" :transition-name="transitionName"
    :ellipsis-props="{ hideTimeout: 800 }" :responsive="responsive" :sidebar-custom-style="{ width: '215px' }"
    @update:collapse="updateCollapse" @update:maximized="updateMaximized" @tabAdd="addPageTab" @tabClick="onTabClick"
    @tabRemove="removePageTab" @tabContextMenu="onTabContextMenu" @tabSortChange="setPageTabs"
    @bodySizeChange="onBodySizeChange">
    <router-layout />
    <!-- logo -->
    <template #logoTitle>
      <h1>{{ PROJECT_NAME }}</h1>
    </template>
    <!-- Button on the left side of the top bar -->
    <template #left="{ sidebar }">
      <!-- Collapse Sidebar -->
      <layout-tool v-if="sidebar" @click="updateCollapse(!collapse)">
        <el-icon style="transform: scale(1.14)">
          <MenuUnfoldOutlined v-if="collapse" />
          <MenuFoldOutlined v-else />
        </el-icon>
      </layout-tool>
      <!-- Refresh -->
      <layout-tool v-if="layout !== 'top' && layout !== 'mix' && !(tabBar && tabInHeader)" class="hidden-sm-and-down"
        @click="reloadPageTab()">
        <el-icon style="transform: scale(1.09)">
          <ReloadOutlined />
        </el-icon>
      </layout-tool>
    </template>
    <!-- Button on the right side of the top bar -->
    <template #right>
      <!-- Fullscreen toggle -->
      <layout-tool class="hidden-sm-and-down" @click="toggleFullscreen">
        <el-icon style="transform: scale(1.18)">
          <CompressOutlined v-if="isFullscreen" style="stroke-width: 4" />
          <ExpandOutlined v-else style="stroke-width: 4" />
        </el-icon>
      </layout-tool>
      <!-- Language switch -->
      <layout-tool :class="{ 'hidden-sm-and-down': tabBar && tabInHeader }">
        <i18n-icon :icon-style="{ transform: 'scale(1.15)' }" />
      </layout-tool>
      <!-- Notification -->
      <layout-tool :class="{ 'hidden-sm-and-down': tabBar && tabInHeader }">
        <header-notice />
      </layout-tool>
      <!-- User info -->
      <layout-tool>
        <header-user />
      </layout-tool>
      <!-- Theme setting -->
      <layout-tool @click="openSetting">
        <el-icon>
          <MoreOutlined />
        </el-icon>
      </layout-tool>
    </template>
    <!-- Page tab right dropdown menu -->
    <template v-if="tabBar && !tabInHeader" #tabExtra="{ active }">
      <tab-dropdown :items="tabExtra" :dropdown-props="{
        iconProps: { size: 15 },
        popperOptions: { strategy: 'fixed' }
      }" @menuClick="(key) => onTabDropdownMenu(key, active)" />
    </template>
    <!-- Collapse both sidebars -->
    <template #boxBottom>
      <div :style="{ flexShrink: 0, padding: roundedTheme ? '4px 8px' : 0 }">
        <layout-tool style="height: 32px" @click="updateCompact(!compact)">
          <el-icon style="transform: scale(1.05)">
            <MenuUnfoldOutlined v-if="compact" />
            <MenuFoldOutlined v-else />
          </el-icon>
        </layout-tool>
      </div>
    </template>
    <!-- Global Footer -->
    <template #footer>
      <!-- <page-footer /> -->
    </template>
    <!-- Menu icon -->
    <template #icon="{ icon, item }">
      <el-icon v-if="icon" v-bind="item.meta?.props?.iconProps || {}">
        <component :is="icon" :style="item.meta?.props?.iconStyle" />
      </el-icon>
    </template>
    <!-- Page tab title -->
    <template #tabTitle="{ label, item }">
      <el-icon v-if="item.meta?.icon" class="ele-tab-icon" v-bind="item.meta?.props?.iconProps || {}">
        <component :is="item.meta.icon" :style="item.meta?.props?.iconStyle" />
      </el-icon>
      <span :style="item.meta?.icon ? { paddingLeft: '4px' } : {}">
        {{ label }}
      </span>
    </template>
  </ele-pro-layout>

  <!-- Theme setting drawer -->
  <setting-drawer v-model="settingVisible" />
</template>

<script setup>
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';
import { useI18n } from 'vue-i18n';
import { LayoutTool, TabDropdown, requestFullscreen, exitFullscreen, checkFullscreen, EleMessage } from 'ele-admin-plus/es';
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  ReloadOutlined,
  ExpandOutlined,
  CompressOutlined,
  MoreOutlined,
  CloseOutlined,
  ArrowLeftOutlined,
  ArrowRightOutlined,
  MinusCircleOutlined,
  CloseCircleOutlined
} from '@/components/icons';
import { PROJECT_NAME, HOME_PATH, REDIRECT_PATH, TAB_KEEP_ALIVE } from '@/config/setting';
import { useUserStore } from '@/store/modules/user';
import { useThemeStore } from '@/store/modules/theme';
import { useMobileDevice } from '@/utils/use-mobile';
import { usePageTab } from '@/utils/use-page-tab';
import RouterLayout from '@/components/RouterLayout/index.vue';
import HeaderUser from './components/header-user.vue';
import HeaderNotice from './components/header-notice.vue';
import I18nIcon from './components/i18n-icon.vue';
import SettingDrawer from './components/setting-drawer.vue';

const { push } = useRouter();
const { t, locale } = useI18n();
const { addPageTab, removePageTab, removeAllPageTab, removeLeftPageTab, removeRightPageTab, removeOtherPageTab, reloadPageTab, setPageTabs } = usePageTab();
const { mobileDevice } = useMobileDevice();
const userStore = useUserStore();
const themeStore = useThemeStore();

/** Menu data */
const { menus } = storeToRefs(userStore);

/** mock data */

// const items = shallowRef([
//   {
//     path: '/dashboard/workplace',
//     meta: {
//       title: 'front page'
//     }
//   },
//   {
//     meta: {
//       path: '/member',
//       title: 'member management'
//     },
//     children: [
//       {
//         path: '/member/index',
//         meta: {
//           title: 'member management'
//         }
//       },
//       {
//         path: '/risk/index',
//         meta: {
//           title: 'risk user'
//         }
//       },
//       {
//         path: '/member/grade',
//         meta: {
//           title: 'member grade'
//         }
//       },
//       {
//         path: '/member/level',
//         meta: {
//           title: 'member level'
//         }
//       },
//       {
//         path: '/member/demoAccount',
//         meta: {
//           title: 'demo account'
//         }
//       }
//     ]
//   }
// ]);

/** Layout style */
const {
  tabs,
  collapse,
  compact,
  maximized,
  tabBar,
  layout,
  sidebarLayout,
  headerStyle,
  sidebarStyle,
  tabStyle,
  fixedHeader,
  fixedSidebar,
  fixedBody,
  fluid,
  logoInHeader,
  colorfulIcon,
  transitionName,
  uniqueOpened,
  fixedHome,
  tabInHeader,
  roundedTheme,
  menuItemTrigger,
  responsive
} = storeToRefs(themeStore);

/** Whether to full screen */
const isFullscreen = ref(false);

/** Whether to show the theme setting drawer */
const settingVisible = ref(false);

/** Page tab right-click menu */
const tabContext = computed(() => {
  return [
    {
      title: t('layout.tabs.reload'),
      command: 'reload',
      icon: ReloadOutlined,
      iconStyle: { transform: 'scale(0.98)' }
    },
    {
      title: t('layout.tabs.close'),
      command: 'close',
      icon: CloseOutlined
    },
    {
      title: t('layout.tabs.closeLeft'),
      command: 'left',
      icon: ArrowLeftOutlined,
      divided: true
    },
    {
      title: t('layout.tabs.closeRight'),
      command: 'right',
      icon: ArrowRightOutlined
    },
    {
      title: t('layout.tabs.closeOther'),
      command: 'other',
      icon: MinusCircleOutlined,
      divided: true
    },
    {
      title: t('layout.tabs.closeAll'),
      command: 'all',
      icon: CloseCircleOutlined
    }
  ];
});

/** Page tab extra menu */
const tabExtra = computed(() => {
  const isMax = maximized.value;
  return [
    {
      title: t(`layout.tabs.${isMax ? 'fullscreenExit' : 'fullscreen'}`),
      command: 'fullscreen',
      icon: isMax ? CompressOutlined : ExpandOutlined
    },
    ...tabContext.value
  ];
});

/** Sidebar collapse toggle */
const updateCollapse = (value) => {
  themeStore.setCollapse(value);
};

/** Sidebar compact toggle */
const updateCompact = (value) => {
  themeStore.setCompact(value);
};

/** Content fullscreen toggle */
const updateMaximized = (value) => {
  themeStore.setMaximized(value);
};

/** Page tab click event */
const onTabClick = (option) => {
  const { key, active, item } = option;
  const path = item?.fullPath || key;
  if (key !== active && path) {
    push(path);
  }
};

/** Content area size change event */
const onBodySizeChange = ({ width }) => {
  themeStore.setContentWidth(width ?? null);
  isFullscreen.value = checkFullscreen();
};

/** Fullscreen toggle */
const toggleFullscreen = () => {
  if (isFullscreen.value) {
    exitFullscreen();
    isFullscreen.value = false;
    return;
  }
  try {
    requestFullscreen();
    isFullscreen.value = true;
  } catch (e) {
    console.error(e);
    EleMessage.error('Your browser does not support fullscreen mode');
  }
};

/** Page tab right-click menu click event */
const onTabContextMenu = (option) => {
  const { command, key, item, active } = option;
  if (command === 'reload') {
    reloadPageTab({ fullPath: item?.fullPath || key });
  } else if (command === 'close') {
    removePageTab({ key, active });
  } else if (command === 'left') {
    removeLeftPageTab({ key, active });
  } else if (command === 'right') {
    removeRightPageTab({ key, active });
  } else if (command === 'other') {
    removeOtherPageTab({ key, active });
  } else if (command === 'all') {
    removeAllPageTab({ key, active });
  }
};

/** Page tab right-click menu click event */
const onTabDropdownMenu = (command, active) => {
  if (command === 'reload') {
    reloadPageTab();
  } else if (command === 'fullscreen') {
    updateMaximized(!maximized.value);
  } else {
    onTabContextMenu({ command, key: active, active });
  }
};

/** Page tab title internationalization */
const i18n = ({ menu, locale }) => {

  // If lang exists, JSON.parse processes the data
  if (menu?.meta?.lang && typeof menu?.meta?.lang === 'string') {
    menu.meta.lang = JSON.parse(menu?.meta?.lang)
  }

  if (locale && menu?.meta?.lang && menu.meta.lang[locale]) {
    return menu.meta.lang[locale];
  }
  return menu?.component ? void 0 : menu?.meta?.title;
};

/** Open theme setting drawer */
const openSetting = () => {
  settingVisible.value = true;
};
</script>

<script>
import * as MenuIcons from './menu-icons';

export default {
  name: 'Layout',
  components: MenuIcons
};
</script>
<style scoped>
.el-menu {
  .el-sub-menu {
    height: 40px;
  }
}
</style>
