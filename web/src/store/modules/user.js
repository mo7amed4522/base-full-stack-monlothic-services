/**
 * Login user status management
 */
import { defineStore } from 'pinia';
import { toTree, mapTree, isExternalLink } from 'ele-admin-plus/es';
import { getUserInfo } from '@/api/layout';
/** Directly specify menu data */
const USER_MENUS = null;

export const useUserStore = defineStore('user', {
  state: () => ({
    /** Current login user information */
    info: null,
    /** Current login user menus */
    menus: null,
    /** Current login user authorities */
    authorities: [
      // Development permissions - add all required permissions for testing
      'odds:read',
      'odds:write',
      'tips:read',
      'tips:write',
      'hot-matches:read',
      'hot-matches:write'
    ],
    /** Current login user roles */
    roles: ['admin'],
    /** Dictionary data cache */
    dicts: {}
  }),
  actions: {
    /**
     * Request login user personal information/authorities/roles/menus
     */
    async fetchUserInfo() {
      const result = await getUserInfo().catch((e) => console.error(e));
      if (!result) {
        return {};
      }
      // User Information
      // this.setInfo(result);
      // 用户权限
      // if (result.authorities) {
      //   this.authorities =
      //     result.authorities.map((d) => d.authority).filter((a) => !!a) ?? [];
      // }
      // 用户角色
      // this.roles = result.roles?.map?.((d) => d.roleCode) ?? [];
      // 用户菜单, 过滤掉按钮类型并转为children形式
      // const { menus, homePath } = formatMenus(
      //   USER_MENUS ??
      //     toTree({
      //       data: result.authorities?.filter?.((d) => d.menu_type !== 1),
      //       idField: 'id',
      //       parentIdField: 'parent_id'
      //     })
      // );
      const { menus, homePath } = formatMenus(USER_MENUS ?? result.menus);
      this.setMenus(menus);
      return { menus, homePath };
    },
    /**
     * Update user information
     */
    setInfo(value) {
      this.info = value;
    },
    /**
     * Update menu data
     */
    setMenus(menus) {
      this.menus = menus;
    },
    /**
     * Update menu badge
     * @param path Menu path
     * @param value Badge value
     * @param type Badge type
     */
    setMenuBadge(path, value, type) {
      this.menus = mapTree(this.menus, (m) => {
        if (path === m.path) {
          const meta = m.meta || {};
          return {
            ...m,
            meta: {
              ...meta,
              props: {
                ...meta.props,
                badge: value == null ? void 0 : { value, type }
              }
            }
          };
        }
        return m;
      });
    },
    /**
     * Update dictionary data
     */
    setDicts(value, code) {
      if (code == null) {
        this.dicts = value;
        return;
      }
      this.dicts[code] = value;
    }
  }
});

/**
 * Menu data processing format required by EleProLayout
 * @param data Menu data
 * @param childField Field name of child level
 */
function formatMenus(data, childField = 'children') {
  let homePath;
  let homeTitle;
  const menus = mapTree(
    data,
    (item) => {
      const meta = typeof item.meta === 'string' ? JSON.parse(item.meta || '{}') : item.meta;
      const menu = {
        path: item.path,
        component: item.component,
        meta: {
          title: item.name,
          icon: item.icon,
          hide: item.hidden,
          ...meta
        }
      };
      const children = item[childField] ? item[childField].filter((d) => !(d.meta?.hide ?? d.hidden)) : void 0;
      if (!children?.length) {
        if (!homePath && menu.path && !isExternalLink(menu.path)) {
          homePath = menu.path;
          homeTitle = menu.meta?.title;
        }
      } else {
        const childPath = children[0].path;
        if (childPath) {
          if (!menu.redirect) {
            menu.redirect = childPath;
          }
          if (!menu.path) {
            menu.path = childPath.substring(0, childPath.lastIndexOf('/'));
          }
        }
      }
      if (!menu.path) {
        console.error('Menu path cannot be empty and must be unique:', item);
        return;
      }
      return menu;
    },
    childField
  );
  return { menus, homePath, homeTitle };
}
