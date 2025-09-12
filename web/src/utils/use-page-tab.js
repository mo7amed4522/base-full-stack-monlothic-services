import { useRouter, useRoute } from 'vue-router';
import { EleMessage } from 'ele-admin-plus/es';
import { useThemeStore } from '@/store/modules/theme';
import { HOME_PATH, LAYOUT_PATH, REDIRECT_PATH } from '@/config/setting';

/**
 * Tab operation hook
 */
export function usePageTab() {
  const HOME_ROUTE = HOME_PATH || LAYOUT_PATH;
  const route = useRoute();
  const { push, replace } = useRouter();
  const themeStore = useThemeStore();

  /**
   * Get the tab key corresponding to the current route
   */
  const getRouteTabKey = () => {
    const { path, fullPath, meta } = route;
    return meta.tabUnique === false ? fullPath : path;
  };

  /**
   * The tab key corresponding to the current route
   */
  const routeTabKey = getRouteTabKey();

  /**
   * Refresh the current route
   */
  const reloadPageTab = (option) => {
    if (route.path.includes(REDIRECT_PATH)) {
      return;
    }
    if (!option || !option.fullPath) {
      // Refresh the current route
      setPageTab({ fullPath: route.fullPath, refresh: true });
      replace({ path: REDIRECT_PATH + route.path, query: route.query });
    } else {
      // Refresh the specified tab
      setPageTab({ fullPath: option.fullPath, refresh: true });
      replace(REDIRECT_PATH + option.fullPath);
    }
  };

  /**
   * Close the current tab
   */
  const finishPageTab = () => {
    removePageTab({ key: routeTabKey, active: getRouteTabKey() });
  };

  /**
   * Close the specified tab
   */
  const removePageTab = (option) => {
    themeStore
      .tabRemove(option)
      .then((result) => {
        onRemoveDone(result);
      })
      .catch(() => {
        EleMessage.error('The current tab cannot be closed');
      });
  };

  /**
   * Close the left tab
   */
  const removeLeftPageTab = (option) => {
    themeStore
      .tabRemoveLeft(option)
      .then((result) => {
        onRemoveDone(result);
      })
      .catch(() => {
        EleMessage.error('The left tab cannot be closed');
      });
  };

  /**
   * Close the right tab
   */
  const removeRightPageTab = (option) => {
    themeStore
      .tabRemoveRight(option)
      .then((result) => {
        onRemoveDone(result);
      })
      .catch(() => {
        EleMessage.error('The right tab cannot be closed');
      });
  };

  /**
   * Close the other tab
   */
  const removeOtherPageTab = (option) => {
    themeStore
      .tabRemoveOther(option)
      .then((result) => {
        onRemoveDone(result);
      })
      .catch(() => {
        EleMessage.error('The other tab cannot be closed');
      });
  };

  /**
   * Close all tabs
   */
  const removeAllPageTab = (option) => {
    themeStore
      .tabRemoveAll(option)
      .then((result) => {
        onRemoveDone(result);
      })
      .catch(() => {
        EleMessage.error('The other tab cannot be closed');
      });
  };

  /**
   * Page tab removal method completion operation
   */
  const onRemoveDone = ({ path, home }) => {
    if (path) {
      push(path);
    } else if (home) {
      push(HOME_ROUTE);
    }
  };

  /**
   * Update page tab data
   * @param data Page tab data
   */
  const setPageTabs = (data) => {
    themeStore.setTabs(data);
  };

  /**
   * Clear page tab data after login
   */
  const cleanPageTabs = () => {
    setPageTabs([]);
  };

  /**
   * Add a tab
   * @param data Tab data
   */
  const addPageTab = (data) => {
    themeStore.tabAdd(data);
  };

  /**
   * Set a tab
   * @param data Tab data
   */
  const setPageTab = (data) => {
    themeStore.tabSetItem(data);
  };

  /**
   * Set the title of the current tab
   * @param title Title
   */
  const setPageTabTitle = (title) => {
    setPageTab({ key: routeTabKey, title });
  };

  /**
   * Go to the home route after login
   * @param from Login before address
   */
  const goHomeRoute = (from) => {
    replace(from ? decodeURIComponent(from) : HOME_ROUTE);
  };

  return {
    reloadPageTab,
    finishPageTab,
    removePageTab,
    removeLeftPageTab,
    removeRightPageTab,
    removeOtherPageTab,
    removeAllPageTab,
    setPageTabs,
    cleanPageTabs,
    addPageTab,
    setPageTab,
    setPageTabTitle,
    getRouteTabKey,
    goHomeRoute,
    routeTabKey
  };
}
