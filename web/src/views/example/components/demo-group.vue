<template>
  <ele-card header="Group menu">
    <div style="display: flex; flex-wrap: wrap">
      <el-button type="primary" @click="toGroup1" style="margin: 0 12px 12px 0">
        Group level 1
      </el-button>
      <el-button type="primary" @click="toGroup2" style="margin: 0 12px 12px 0">
        Group level 2
      </el-button>
    </div>
    <ele-text type="placeholder">
      <span>Group level 2 menu can view list page/card list effect, </span>
      <el-link type="primary" :underline="false" @click="resetMenu">
        Reset menu data
      </el-link>
    </ele-text>
  </ele-card>
</template>

<script setup>
import { storeToRefs } from 'pinia';
import { useUserStore } from '@/store/modules/user';

const userStore = useUserStore();
const { menus } = storeToRefs(userStore);
const tempMenus = JSON.parse(JSON.stringify(menus.value));

/** Group level 1 */
const toGroup1 = () => {
  userStore.setMenus(
    tempMenus.map((m) => {
      const hasChild = !!m.children?.length;
      const meta = m.meta || {};
      return {
        ...m,
        meta: {
          ...meta,
          //icon: hasChild ? null : meta.icon,
          props: {
            ...meta.props,
            group: hasChild
          }
        }
      };
    })
  );
};

/** Group level 2 */
const toGroup2 = () => {
  userStore.setMenus(
    tempMenus.map((m) => {
      return {
        ...m,
        children: m.children
          ? m.children.map((c) => {
            const hasChild = c.children
              ? !!c.children.filter((t) => !t.meta?.hide).length
              : false;
            const meta = c.meta || {};
            return {
              ...c,
              meta: {
                ...meta,
                //icon: hasChild ? null : meta.icon,
                props: {
                  ...meta.props,
                  group: hasChild
                }
              }
            };
          })
          : void 0
      };
    })
  );
};

/** Reset menu data */
const resetMenu = () => {
  userStore.setMenus(JSON.parse(JSON.stringify(tempMenus)));
};
</script>
