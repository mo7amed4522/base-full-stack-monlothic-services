<template>
  <ele-page>
    <ele-card header="Vertical menu">
      <div style="display: flex; flex-wrap: wrap">
        <option-item label="Dark theme" :responsive="false" style="margin: 0 12px 12px 0">
          <el-switch v-model="dark" size="small" />
        </option-item>
        <option-item label="Colorful icon" :responsive="false" style="margin: 0 12px 12px 0">
          <el-switch v-model="colorful" size="small" />
        </option-item>
        <option-item label="Collapse menu" :responsive="false" style="margin: 0 12px 12px 0">
          <el-switch v-model="collapse" size="small" />
        </option-item>
      </div>
      <div style="display: flex; flex-wrap: wrap">
        <div style="margin: 0 12px 12px 0">
          <el-button type="primary" class="ele-btn-icon" @click="changeData()">
            Group menu items
          </el-button>
        </div>
        <div style="margin: 0 12px 12px 0">
          <el-button type="primary" class="ele-btn-icon" @click="changeColor()">
            Change menu icon color
          </el-button>
        </div>
      </div>
      <div style="max-width: 240px">
        <ele-menus mode="vertical" :items="menuItems" :collapse="collapse" :collapse-transition="true"
          :default-active="menuActive" :theme="dark ? 'dark' : 'light'" :colorful="colorful"
          :text-ellipsis-tooltip="true" />
      </div>
    </ele-card>
    <ele-card header="Horizontal menu">
      <div style="display: flex; flex-wrap: wrap">
        <option-item label="Dark theme" :responsive="false" style="margin: 0 12px 12px 0">
          <el-switch v-model="navDark" size="small" />
        </option-item>
        <option-item label="Colorful icon" :responsive="false" style="margin: 0 12px 12px 0">
          <el-switch v-model="navColorful" size="small" />
        </option-item>
      </div>
      <div style="display: flex; flex-wrap: wrap">
        <div style="margin: 0 12px 12px 0">
          <el-button type="primary" class="ele-btn-icon" @click="changeColor('nav')">
            Change menu icon color
          </el-button>
        </div>
      </div>
      <div style="max-width: 800px">
        <ele-menus mode="horizontal" :items="navItems" :default-active="navActive" :theme="navDark ? 'dark' : 'light'"
          :colorful="navColorful" :text-ellipsis-tooltip="true" />
      </div>
    </ele-card>
    <ele-card header="Compact menu">
      <div style="display: flex; flex-wrap: wrap">
        <option-item label="Dark theme" :responsive="false" style="margin: 0 12px 12px 0">
          <el-switch v-model="boxDark" size="small" />
        </option-item>
        <option-item label="Colorful icon" :responsive="false" style="margin: 0 12px 12px 0">
          <el-switch v-model="boxColorful" size="small" />
        </option-item>
        <option-item label="Hide title" :responsive="false" style="margin: 0 12px 12px 0">
          <el-switch v-model="boxCollapse" size="small" />
        </option-item>
      </div>
      <div style="display: flex; flex-wrap: wrap">
        <div style="margin: 0 12px 12px 0">
          <el-button type="primary" class="ele-btn-icon" @click="changeData('box')">
            Group menu items
          </el-button>
        </div>
        <div style="margin: 0 12px 12px 0">
          <el-button type="primary" class="ele-btn-icon" @click="changeColor('box')">
            Change menu icon color
          </el-button>
        </div>
      </div>
      <div :style="{
        width: boxCollapse ? '66px' : '88px',
        transition: 'width 0.2s'
      }">
        <ele-menus mode="compact" :items="boxItems" :collapse="boxCollapse" :default-active="boxActive"
          :theme="boxDark ? 'dark' : 'light'" :colorful="boxColorful" :text-ellipsis-tooltip="true" />
      </div>
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref, shallowRef } from 'vue';
import {
  SettingOutlined,
  CalendarOutlined,
  LogOutlined,
  LinkOutlined,
  UserOutlined
} from '@/components/icons';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';

/** Whether the vertical menu is dark theme */
const dark = ref(false);

/** Whether the vertical menu is colorful icon */
const colorful = ref(false);

/** Whether the vertical menu is collapsed */
const collapse = ref(false);

/** Whether the horizontal menu is dark theme */
const navDark = ref(false);

/** Whether the horizontal menu is colorful icon */
const navColorful = ref(false);

/** Whether the compact menu is dark theme */
const boxDark = ref(false);

/** Whether the compact menu is colorful icon */
const boxColorful = ref(false);

/** Whether the compact menu is collapsed */
const boxCollapse = ref(false);

/** Vertical menu selected */
const menuActive = ref('1-1');

/** Horizontal menu selected */
const navActive = ref('1-1');

/** Compact menu selected */
const boxActive = ref('1-1');

/** Get menu data */
const getMenuData = () => {
  return [
    {
      index: '1',
      title: 'System management',
      icon: SettingOutlined,
      children: [
        {
          index: '1-1',
          title: 'User management'
        },
        {
          index: '1-2',
          title: 'Role management',
          badge: { value: 1 }
        },
        {
          index: '1-3',
          title: 'Permission management'
        }
      ]
    },
    {
      index: '2',
      title: 'Log management',
      icon: CalendarOutlined,
      badge: { value: 'New', type: 'warning' },
      children: [
        {
          index: '2-1',
          title: 'Login log',
          icon: LinkOutlined
        },
        {
          index: '2-2',
          title: 'Operation log if the title is very long',
          icon: LinkOutlined,
          badge: { isDot: true, type: 'success' }
        }
      ]
    },
    {
      index: '3',
      title: 'Order management',
      icon: LogOutlined,
      popperClass: 'demo-menu-popper',
      children: [
        {
          index: '3-1',
          title: 'Order query',
          icon: LinkOutlined
        },
        {
          index: '3-2',
          title: 'Refund record',
          icon: LinkOutlined
        },
        {
          index: '3-3',
          title: 'Order report',
          icon: LinkOutlined,
          children: [
            {
              index: '3-3-1',
              title: 'Daily report',
              icon: LinkOutlined
            },
            {
              index: '3-3-2',
              title: 'Weekly report if the title is very long',
              icon: LinkOutlined,
              disabled: true
            },
            {
              index: '3-3-3',
              title: 'Monthly report',
              icon: LinkOutlined
            },
            {
              index: '3-3-4',
              title: 'Quarterly report',
              group: true,
              children: [
                {
                  index: '3-3-4-1',
                  title: 'First quarter',
                  icon: LinkOutlined
                },
                {
                  index: '3-3-4-2',
                  title: 'Second quarter',
                  icon: LinkOutlined
                }
              ]
            }
          ]
        },
        {
          index: '3-4',
          title: 'Bill statistics',
          group: true,
          children: [
            {
              index: '3-4-1',
              title: 'Sales statistics',
              icon: LinkOutlined
            },
            {
              index: '3-4-2',
              title: 'Refund statistics',
              icon: LinkOutlined
            }
          ]
        }
      ]
    },
    {
      index: '4',
      title: 'Personal center',
      icon: UserOutlined,
      badge: { value: 100 }
    }
  ];
};

/** Vertical menu data */
const menuItems = shallowRef(getMenuData());

/** Horizontal menu data */
const navItems = shallowRef(getMenuData());

/** Compact menu data */
const boxItems = shallowRef(getMenuData());

/** First level menu changed to grouped form */
const changeData = (type) => {
  const result = getMenuData().map((d) => {
    const t = { ...d };
    if (t.children) {
      t.group = true;
      t.icon = void 0;
      if (t.title) {
        t.title = t.title.replace('management', '');
      }
      t.children = t.children.map((c) => {
        return { ...c, icon: c.icon ?? LinkOutlined };
      });
    }
    return t;
  });
  if (type === 'box') {
    boxItems.value = result;
    return;
  }
  menuItems.value = result;
};

/** Custom icon background color */
const changeColor = (type) => {
  const colors = {
    1: 'linear-gradient(-225deg, #231557 0%, #44107A 29%, #FF1361 67%, #FFF800 100%)',
    2: 'linear-gradient( 135deg, #FFF3B0 10%, #CA26FF 100%)',
    3: 'linear-gradient(-225deg, #69EACB 0%, #EACCF8 48%, #6654F1 100%)',
    4: '#faad14',
    '1-1': '#f5222d',
    '1-2': '#fa541c',
    '1-3': '#fa8c16',
    '2-1': '#2f54eb',
    '2-2': '#722ed1',
    '3-1': '#eb2f96',
    '3-2':
      'linear-gradient(135deg, rgb(67, 203, 255) 10%, rgb(151, 8, 204) 100%)',
    '3-3':
      'linear-gradient(135deg, rgb(160, 254, 101) 10%, rgb(250, 1, 109) 100%)',
    '3-4': 'linear-gradient(to right, #f83600 0%, #f9d423 100%)'
  };
  const data =
    type === 'box'
      ? boxItems.value
      : type === 'nav'
        ? navItems.value
        : menuItems.value;
  const result = data.map((d) => {
    const t = { ...d };
    if (!t.group) {
      if (t.index && colors[t.index]) {
        if (colors[t.index].startsWith('linear-gradient(')) {
          t.iconProps = { style: { backgroundImage: colors[t.index] } };
        } else {
          t.iconProps = { style: { backgroundColor: colors[t.index] } };
        }
      }
    } else if (t.children) {
      t.children = t.children.map((c) => {
        const m = { ...c };
        if (m.index && colors[m.index]) {
          if (colors[m.index].startsWith('linear-gradient(')) {
            m.iconProps = { style: { backgroundImage: colors[m.index] } };
          } else {
            m.iconProps = { style: { backgroundColor: colors[m.index] } };
          }
        }
        return m;
      });
    }
    return t;
  });
  if (type === 'box') {
    boxItems.value = result;
    return;
  }
  if (type === 'nav') {
    navItems.value = result;
    return;
  }
  menuItems.value = result;
};
</script>

<script>
export default {
  name: 'ExtensionMenu'
};
</script>

<style lang="scss">
.demo-menu-popper.ele-menu>.el-menu--popup-container>.el-menu--popup {
  overflow: visible;

  .el-sub-menu>.el-popper {
    position: absolute !important;
  }
}
</style>
