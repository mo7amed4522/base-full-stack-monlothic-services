<template>
  <ele-card header="Layout basic usage" :body-style="{ padding: '12px' }">
    <frame-card :style="mobile ? { maxWidth: '420px' } : {}">
      <div ref="wrapperRef" :style="{
        height: '520px',
        overflow: 'auto',
        minWidth: mobile ? void 0 : '800px'
      }" class="demo-layout-wrapper">
        <ele-admin-layout ref="layoutRef" :height="520" :headerMenus="headerMenus" :headerActive="headerActive"
          :sidebarMenus="sidebarMenus" :sidebarActive="sidebarActive" :sideboxMenus="sideboxMenus"
          :sideboxActive="sideboxActive" :tabs="tabs" :tabActive="tabActive" :levels="levels" :layout="layout"
          :sidebarLayout="sidebarLayout" :tabBar="tabBar" :breadcrumb="mobile ? false : breadcrumb"
          :headerStyle="headerStyle" :sidebarStyle="sidebarStyle" :tabStyle="tabStyle" :fixedHeader="fixedHeader"
          :fixedSidebar="fixedSidebar" :fixedBody="fixedBody" :logoInHeader="logoInHeader" :maximized="maximized"
          v-model:collapse="collapse" :compact="compact" :mobile="mobile" :fixedHome="true"
          :menu-scroll-to-active="false" :style="{
            '--ele-header-height': headerHeight + 'px',
            '--ele-sidebar-width': sidebarWidth + 'px',
            '--ele-sidebox-width': sideboxWidth + 'px',
            '--ele-sidebar-mix-width': sidebarMixWidth + 'px'
          }">
          <template #logo>
            <img src="@/assets/logo.svg" />
          </template>
          <template #logoTitle>
            <h1>Full Stack Admin</h1>
          </template>
          <template v-if="layout !== 'top'" #left>
            <layout-tool @click="collapse = !collapse">
              <el-icon style="transform: scale(1.14)">
                <MenuUnfoldOutlined v-if="collapse" />
                <MenuFoldOutlined v-else />
              </el-icon>
            </layout-tool>
            <layout-tool>
              <el-icon style="transform: scale(1.09)">
                <ReloadOutlined />
              </el-icon>
            </layout-tool>
          </template>
          <template #right>
            <layout-tool>
              <el-icon style="transform: scale(1.18)">
                <ExpandOutlined style="stroke-width: 4" />
              </el-icon>
            </layout-tool>
            <layout-tool>
              <el-icon>
                <MoreOutlined />
              </el-icon>
            </layout-tool>
          </template>
          <template #boxBottom>
            <div :style="{ flexShrink: 0, padding: '4px 8px' }">
              <layout-tool style="height: 32px" @click="compact = !compact">
                <el-icon style="transform: scale(1.05)">
                  <MenuUnfoldOutlined v-if="compact" />
                  <MenuFoldOutlined v-else />
                </el-icon>
              </layout-tool>
            </div>
          </template>
          <div :style="{
            flex: 'auto',
            padding: tabStyle === 'button' && tabBar === true ? '0 12px' : '12px 12px 0 12px',
            boxSizing: 'border-box'
          }">
            <ele-card :body-style="{ padding: '16px 8px 16px 14px' }" style="margin: 0">
              <option-item label="Layout">
                <el-radio-group v-model="layout">
                  <el-radio value="default" label="Sidebar navigation" />
                  <el-radio value="top" label="Top navigation" />
                  <el-radio value="mix" label="Mix navigation" />
                  <el-radio value="side" label="Only sidebar navigation" />
                </el-radio-group>
              </option-item>
              <option-item label="Sidebar layout">
                <el-radio-group v-model="sidebarLayout">
                  <el-radio value="default" label="Single sidebar" />
                  <el-radio value="mix" label="Double sidebar" />
                  <el-radio value="box" label="Only first level sidebar" />
                </el-radio-group>
              </option-item>
              <option-item label="Header style">
                <el-radio-group :disabled="layout === 'side'" v-model="headerStyle">
                  <el-radio value="light" label="Light style" />
                  <el-radio value="dark" label="Dark style" />
                  <el-radio value="primary" label="Primary style" />
                </el-radio-group>
              </option-item>
              <option-item label="Sidebar style">
                <el-radio-group :disabled="layout === 'top'" v-model="sidebarStyle">
                  <el-radio value="light" label="Light style" />
                  <el-radio value="dark" label="Dark style" />
                </el-radio-group>
              </option-item>
              <option-item label="Tab bar">
                <el-radio-group v-model="tabBar">
                  <el-radio :value="true" label="Show" />
                  <el-radio value="header" label="Show in header" />
                  <el-radio :value="false" label="Hide" />
                </el-radio-group>
              </option-item>
              <option-item label="Tab style">
                <el-radio-group :disabled="!tabBar" v-model="tabStyle">
                  <el-radio value="simple" label="Simple style" />
                  <el-radio value="indicator" label="Indicator style" />
                  <el-radio value="tag" label="Tag style" />
                  <el-radio value="button" label="Button style" />
                </el-radio-group>
              </option-item>
              <option-item label="Breadcrumb">
                <el-radio-group v-model="breadcrumb">
                  <el-radio :value="true" label="Show" />
                  <el-radio :value="false" label="Hide" />
                </el-radio-group>
              </option-item>
              <option-item label="Content area scroll">
                <el-radio-group v-model="fixedBody">
                  <el-radio :value="true" label="Show" />
                  <el-radio :value="false" label="Hide" />
                </el-radio-group>
              </option-item>
              <option-item label="Fixed header">
                <el-radio-group :disabled="fixedBody" v-model="fixedHeader">
                  <el-radio :value="true" label="Show" />
                  <el-radio :value="false" label="Hide" />
                </el-radio-group>
              </option-item>
              <option-item label="Fixed sidebar">
                <el-radio-group :disabled="fixedBody" v-model="fixedSidebar">
                  <el-radio :value="true" label="Show" />
                  <el-radio :value="false" label="Hide" />
                </el-radio-group>
              </option-item>
              <option-item label="Logo position">
                <el-radio-group :disabled="layout === 'top' || layout === 'side' || mobile" v-model="logoInHeader">
                  <el-radio :value="true" label="Header" />
                  <el-radio :value="false" label="Sidebar" />
                </el-radio-group>
              </option-item>
              <option-item label="Maximize content area">
                <el-radio-group v-model="maximized">
                  <el-radio :value="false" label="Hide" />
                  <el-radio :value="true" label="Show" />
                  <el-radio value="expanded" label="Without tab bar" />
                </el-radio-group>
              </option-item>
              <option-item label="Collapse sidebar">
                <el-radio-group :disabled="(layout === 'top' || sidebarLayout === 'box') && !mobile" v-model="collapse">
                  <el-radio :value="false" label="Hide" />
                  <el-radio :value="true" label="Show" />
                </el-radio-group>
              </option-item>
              <option-item label="Collapse sidebar first level">
                <el-radio-group :disabled="layout === 'top' || sidebarLayout === 'default'" v-model="compact">
                  <el-radio :value="false" label="Hide" />
                  <el-radio :value="true" label="Show" />
                </el-radio-group>
              </option-item>
              <option-item label="Mobile style">
                <el-radio-group v-model="mobile">
                  <el-radio :value="true" label="Show" />
                  <el-radio :value="false" label="Hide" />
                </el-radio-group>
              </option-item>
              <option-item label="Header height">
                <div style="max-width: 160px">
                  <el-slider v-model="headerHeight" :min="38" :max="68" />
                </div>
              </option-item>
              <option-item v-if="sidebarLayout === 'mix'" label="Sidebar width">
                <div style="max-width: 160px">
                  <el-slider v-model="sidebarMixWidth" :min="100" :max="260" />
                </div>
              </option-item>
              <option-item v-else label="Sidebar width">
                <div style="max-width: 160px">
                  <el-slider v-model="sidebarWidth" :min="120" :max="280" />
                </div>
              </option-item>
              <option-item v-if="sidebarLayout === 'mix'" label="Sidebar first level width">
                <div style="max-width: 160px">
                  <el-slider v-model="sideboxWidth" :min="60" :max="120" />
                </div>
              </option-item>
            </ele-card>
          </div>
          <div style="padding: 16px 0; text-align: center; opacity: 0.6">Footer&emsp;Footer&emsp;Footer</div>
        </ele-admin-layout>
      </div>
    </frame-card>
  </ele-card>
</template>

<script setup>
import { CalendarOutlined, ExpandOutlined, IdcardOutlined, LogOutlined, MenuFoldOutlined, MenuUnfoldOutlined, MoreOutlined, ReloadOutlined, SettingOutlined, UserOutlined } from '@/components/icons';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';
import { LayoutTool } from 'ele-admin-plus/es';
import { ref, shallowRef } from 'vue';
import FrameCard from './frame-card.vue';

const wrapperRef = ref(null);

const layoutRef = ref(null);

const headerMenus = shallowRef([
  {
    index: '/personnel',
    title: 'Personnel System',
    icon: UserOutlined
  },
  {
    index: '/finance',
    title: 'Finance System',
    icon: SettingOutlined
  }
]);

const headerActive = ref('/personnel');

const sidebarMenus = shallowRef([
  {
    index: '/system',
    title: 'System Management',
    icon: SettingOutlined,
    children: [
      {
        index: '/system/user',
        title: 'User Management',
        icon: UserOutlined
      },
      {
        index: '/system/role',
        title: 'Role Management',
        icon: IdcardOutlined
      }
    ]
  },
  {
    index: '/log',
    title: 'Log Management',
    icon: LogOutlined,
    children: [
      {
        index: '/log/login',
        title: 'Login Log',
        icon: CalendarOutlined
      }
    ]
  },
  {
    index: '/profile',
    title: 'Profile',
    icon: UserOutlined
  }
]);

const sidebarActive = ref('/system/user');

const sideboxMenus = shallowRef([
  {
    index: '/system',
    title: 'System',
    group: true,
    children: [
      {
        index: '/system/user',
        title: 'User Management',
        icon: UserOutlined
      },
      {
        index: '/system/role',
        title: 'Role Management',
        icon: IdcardOutlined
      }
    ]
  },
  {
    index: '/log',
    title: 'Other',
    group: true,
    children: [
      {
        index: '/log/login',
        title: 'Login Log',
        icon: CalendarOutlined
      },
      {
        index: '/profile',
        title: 'Profile',
        icon: UserOutlined
      }
    ]
  }
]);

const sideboxActive = ref('/system/user');

const tabs = ref([
  { name: '/system/user', label: 'User Management', closable: true },
  { name: '/system/role', label: 'Role Management', closable: true }
]);

const tabActive = ref('/system/user');

const levels = ref([{ title: 'Home' }, { title: 'System Management' }, { title: 'User Management' }]);

const layout = ref('default');

const sidebarLayout = ref('mix');

const tabBar = ref(true);

const breadcrumb = ref(true);

const headerStyle = ref('light');

const sidebarStyle = ref('light');

const tabStyle = ref('button');

const fixedHeader = ref(true);

const fixedSidebar = ref(true);

const fixedBody = ref(true);

const logoInHeader = ref(false);

const maximized = ref(false);

const collapse = ref(false);

const compact = ref(false);

const mobile = ref(false);

const headerHeight = ref(54);

const sidebarWidth = ref(220);

const sideboxWidth = ref(80);

const sidebarMixWidth = ref(160);
</script>

<style lang="scss" scoped>
.demo-layout-wrapper :deep(.option-item) {
  align-items: flex-start;

  .option-item-label {
    width: 106px;
    height: 32px;
    line-height: 32px;
  }

  &+.option-item {
    margin-top: 16px;
  }
}
</style>
