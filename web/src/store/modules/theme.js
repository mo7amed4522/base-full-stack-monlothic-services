/**
 * Theme state management
 */
import { defineStore } from 'pinia';
import { changeColor } from 'ele-admin-plus/es/utils/theme-util';
import { THEME_CACHE_NAME, TAB_KEEP_ALIVE } from '@/config/setting';

/**
 * Default value
 */
const DEFAULT_STATE = Object.freeze({
  /** Tab data */
  tabs: [],
  /** Whether to collapse the sidebar */
  collapse: false,
  /** Whether the sidebar is in compact style */
  compact: false,
  /** Whether the content area is maximized */
  maximized: false,
  /** Whether the tab bar is needed */
  tabBar: true,
  /** Layout type */
  layout: 'default',
  /** Sidebar layout type */
  sidebarLayout: 'default',
  /** Header style */
  headerStyle: 'light',
  /** Sidebar style */
  sidebarStyle: 'dark',
  /** Tab style */
  tabStyle: 'simple',
  /** Whether the header is fixed */
  fixedHeader: true,
  /** Whether the sidebar is fixed */
  fixedSidebar: true,
  /** Whether the content area is fixed */
  fixedBody: true,
  /** Whether the content area is maximized */
  fluid: true,
  /** Whether the logo is in the header */
  logoInHeader: false,
  /** Whether the sidebar menu is colorful */
  colorfulIcon: false,
  /** Whether the sidebar menu is unique opened */
  uniqueOpened: true,
  /** Whether the home tab is fixed */
  fixedHome: true,
  /** Whether the tab bar is placed in the header */
  tabInHeader: false,
  /** Routing transition animation */
  transitionName: 'slide-right',
  /** Whether color weak mode is enabled */
  weakMode: false,
  /** Whether dark mode is enabled */
  darkMode: false,
  /** Theme color */
  color: null,
  /** Content area width */
  contentWidth: null,
  /** Whether rounded theme is enabled */
  roundedTheme: true,
  /** Menu trigger mode */
  menuItemTrigger: 'click',
  /** Whether responsive layout is enabled */
  responsive: true
});

/**
 * Read cache configuration
 */
function getCacheSetting() {
  try {
    const value = localStorage.getItem(THEME_CACHE_NAME);
    if (value) {
      const cache = JSON.parse(value);
      if (typeof cache === 'object') {
        return cache;
      }
    }
  } catch (e) {
    console.error(e);
  }
  return {};
}

/**
 * Cache configuration
 */
function cacheSetting(key, value) {
  const cache = getCacheSetting();
  if (cache[key] !== value) {
    cache[key] = value;
    localStorage.setItem(THEME_CACHE_NAME, JSON.stringify(cache));
  }
}

/**
 * Switch responsive layout
 */
function changeResponsive(responsive) {
  const classes = 'ele-body-limited';
  if (responsive) {
    document.body.classList.remove(classes);
  } else {
    document.body.classList.add(classes);
  }
}

/**
 * Switch rounded theme
 */
function changeRoundedTheme(roundedTheme) {
  const classes = 'rounded';
  const $html = document.querySelector('html');
  if ($html && roundedTheme) {
    $html.classList.add(classes);
  } else if ($html) {
    $html.classList.remove(classes);
  }
}

/**
 * Switch color weak mode
 */
function changeWeakMode(weakMode) {
  const classes = 'ele-admin-weak';
  if (weakMode) {
    document.body.classList.add(classes);
  } else {
    document.body.classList.remove(classes);
  }
}

/**
 * Switch theme
 */
function changeTheme(value, dark) {
  return new Promise((resolve, reject) => {
    try {
      changeColor(value, dark);
      resolve();
    } catch (e) {
      reject(e);
    }
  });
}

export const useThemeStore = defineStore('theme', {
  state: () => {
    const state = { ...DEFAULT_STATE };
    // Reading local cache
    const cache = getCacheSetting();
    Object.keys(state).forEach((key) => {
      if (typeof cache[key] !== 'undefined') {
        state[key] = cache[key];
      }
    });
    return state;
  },
  getters: {
    /** Components that need to be cached */
    keepAliveInclude() {
      if (!TAB_KEEP_ALIVE || !this.tabBar || !this.tabs) {
        return [];
      }
      const components = new Set();
      this.tabs.forEach((t) => {
        if (t.meta?.keepAlive !== false && !t.refresh && t.components) {
          t.components.forEach((c) => {
            if (typeof c === 'string' && c) {
              components.add(c);
            }
          });
        }
      });
      return Array.from(components);
    }
  },
  actions: {
    setTabs(value) {
      this.tabs = value;
      //cacheSetting('tabs', value);
    },
    setCollapse(value) {
      this.collapse = value;
    },
    setCompact(value) {
      this.compact = value;
    },
    setMaximized(value) {
      this.maximized = value;
    },
    setTabBar(value) {
      this.tabBar = value;
      cacheSetting('tabBar', value);
    },
    setLayout(value) {
      this.layout = value;
      cacheSetting('layout', value);
    },
    setSidebarLayout(value) {
      this.sidebarLayout = value;
      cacheSetting('sidebarLayout', value);
    },
    setHeaderStyle(value) {
      this.headerStyle = value;
      cacheSetting('headerStyle', value);
    },
    setSidebarStyle(value) {
      this.sidebarStyle = value;
      cacheSetting('sidebarStyle', value);
    },
    setTabStyle(value) {
      this.tabStyle = value;
      cacheSetting('tabStyle', value);
    },
    setFixedHeader(value) {
      this.fixedHeader = value;
      cacheSetting('fixedHeader', value);
    },
    setFixedSidebar(value) {
      this.fixedSidebar = value;
      cacheSetting('fixedSidebar', value);
    },
    setFixedBody(value) {
      this.fixedBody = value;
      cacheSetting('fixedBody', value);
    },
    setFluid(value) {
      this.fluid = value;
      cacheSetting('fluid', value);
    },
    setLogoInHeader(value) {
      this.logoInHeader = value;
      cacheSetting('logoInHeader', value);
    },
    setColorfulIcon(value) {
      this.colorfulIcon = value;
      cacheSetting('colorfulIcon', value);
    },
    setUniqueOpened(value) {
      this.uniqueOpened = value;
      cacheSetting('uniqueOpened', value);
    },
    setFixedHome(value) {
      this.fixedHome = value;
      cacheSetting('fixedHome', value);
    },
    setTabInHeader(value) {
      this.tabInHeader = value;
      cacheSetting('tabInHeader', value);
    },
    setTransitionName(value) {
      this.transitionName = value;
      cacheSetting('transitionName', value);
    },
    setContentWidth(value) {
      this.contentWidth = value;
    },
    setWeakMode(value) {
      changeWeakMode(value);
      this.weakMode = value;
      cacheSetting('weakMode', value);
    },
    /**
     * Switch to dark mode
     * @param value Is it dark mode?
     */
    async setDarkMode(value) {
      await changeTheme(this.color, value);
      this.darkMode = value;
      cacheSetting('darkMode', value);
    },
    /**
     * Switch theme color
     * @param value Theme color
     */
    async setColor(value) {
      await changeTheme(value, this.darkMode);
      this.color = value;
      cacheSetting('color', value);
    },
    /**
     * Reset theme
     */
    async resetSetting() {
      const excludes = ['tabs', 'collapse', 'contentWidth'];
      Object.keys(DEFAULT_STATE).forEach((key) => {
        if (!excludes.includes(key)) {
          this[key] = DEFAULT_STATE[key];
        }
      });
      localStorage.removeItem(THEME_CACHE_NAME);
      changeResponsive(this.responsive);
      changeRoundedTheme(this.roundedTheme);
      changeWeakMode(this.weakMode);
      await changeTheme(this.color, this.darkMode);
    },
    /**
     * Recover theme
     */
    recoverTheme() {
      // Close responsive layout
      if (!this.responsive) {
        changeResponsive(false);
      }
      // Enable rounded theme
      if (this.roundedTheme) {
        changeRoundedTheme(true);
      }
      // Enable weak mode
      if (this.weakMode) {
        changeWeakMode(true);
      }
      // Recover theme color
      if (this.color || this.darkMode) {
        changeTheme(this.color, this.darkMode).catch((e) => {
          console.error(e);
        });
      }
    },
    /**
     * Add tab or update tab data
     * @param data Tab data
     */
    tabAdd(data) {
      const i = this.tabs.findIndex((d) => d.key === data.key);
      if (i === -1) {
        this.setTabs([...this.tabs, data]);
      } else if (data.fullPath !== this.tabs[i].fullPath) {
        const temps = [...this.tabs];
        temps[i] = data;
        this.setTabs(temps);
      }
    },
    /**
     * Close tab
     */
    async tabRemove({ key, active }) {
      const i = this.tabs.findIndex((t) => t.key === key || t.fullPath === key);
      if (i === -1) {
        return {};
      }
      const t = this.tabs[i];
      if (!t.closable || (t.home && (this.tabs.length === 1 || this.fixedHome))) {
        return Promise.reject();
      }
      const path = this.tabs[i + (i === 0 ? 1 : -1)]?.fullPath;
      this.setTabs(this.tabs.filter((_d, j) => j !== i));
      return t.key === active ? { path, home: !path } : {};
    },
    /**
     * Close left tab
     */
    async tabRemoveLeft({ key, active }) {
      let index = -1; // Selected tab index
      for (let i = 0; i < this.tabs.length; i++) {
        if (this.tabs[i].key === active) {
          index = i;
        }
        if (this.tabs[i].key === key) {
          if (i === 0) {
            break;
          }
          const temp = this.tabs.filter((d, j) => !d.closable && j < i);
          if (temp.length === i) {
            break;
          }
          const path = index === -1 ? void 0 : this.tabs[i].fullPath;
          this.setTabs(temp.concat(this.tabs.slice(i)));
          return { path };
        }
      }
      return Promise.reject();
    },
    /**
     * Close right tab
     */
    async tabRemoveRight({ key, active }) {
      let index = -1; // Selected tab index
      for (let i = 0; i < this.tabs.length; i++) {
        if (this.tabs[i].key === active) {
          index = i;
        }
        if (this.tabs[i].key === key) {
          if (i === this.tabs.length - 1) {
            break;
          }
          const temp = this.tabs.filter((d, j) => !d.closable && j > i);
          if (temp.length === this.tabs.length - i - 1) {
            break;
          }
          const path = index === -1 ? this.tabs[i].fullPath : void 0;
          this.setTabs(this.tabs.slice(0, i + 1).concat(temp));
          return { path };
        }
      }
      return Promise.reject();
    },
    /**
     * Close other tab
     */
    async tabRemoveOther({ key, active }) {
      let path; // Close after jump address
      const temps = this.tabs.filter((d) => {
        if (d.key === key) {
          path = d.fullPath;
        }
        return !d.closable || d.key === key;
      });
      if (temps.length === this.tabs.length) {
        return Promise.reject();
      }
      this.setTabs(temps);
      return key === active ? {} : { path };
    },
    /**
     * Close all tab
     */
    async tabRemoveAll({ active }) {
      if (this.tabs.length === 1 && this.tabs[0].home) {
        return Promise.reject();
      }
      const temps = this.tabs.filter((t) => !t.closable || (t.home && this.fixedHome));
      if (temps.length === this.tabs.length) {
        return Promise.reject();
      }
      const t = active ? this.tabs.find((d) => d.key === active) : void 0;
      const jump = t != null && t.closable === true; // Whether to jump after closing
      if (!temps.length) {
        const h = this.tabs.find((d) => d.home);
        if (!h) {
          this.setTabs([]);
          return { home: true };
        }
        this.setTabs([h]);
        return { home: t?.home ? void 0 : true };
      }
      this.setTabs(temps);
      return { path: jump ? temps[0].fullPath : void 0 };
    },
    /**
     * Set tab item
     * @param data Tab data
     */
    tabSetItem(data) {
      if (!data.key && !data.fullPath) {
        if (!data.path) {
          return;
        }
        this.tabs.forEach((d) => {
          if (data.path === d.path) {
            this.tabSetItem({ ...data, key: d.key });
          }
        });
        return;
      }
      const k = data.key ? 'key' : 'fullPath';
      const i = this.tabs.findIndex((d) => data[k] === d[k]);
      if (i === -1) {
        return;
      }
      const item = { ...this.tabs[i] };
      if (data.title) {
        item.title = data.title;
      }
      if (typeof data.closable === 'boolean') {
        item.closable = data.closable;
      }
      if (typeof data.refresh === 'boolean') {
        item.refresh = data.refresh;
      }
      if (data.components) {
        item.components = data.components;
      }
      const temps = [...this.tabs];
      temps[i] = item;
      this.setTabs(temps);
    },
    /** Modify menu trigger mode */
    setMenuItemTrigger(value) {
      this.menuItemTrigger = value;
      cacheSetting('menuItemTrigger', value);
    },
    /** Switch rounded theme */
    setRoundedTheme(value) {
      changeRoundedTheme(value);
      this.roundedTheme = value;
      cacheSetting('roundedTheme', value);
    },
    /** Modify responsive switch */
    setResponsive(value) {
      changeResponsive(value);
      this.responsive = value;
      cacheSetting('responsive', value);
    }
  }
});
