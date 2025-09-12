/** API base url */
export const API_BASE_URL = import.meta.env.VITE_API_URL;

/** Project name */
export const PROJECT_NAME = import.meta.env.VITE_APP_NAME;

/** Unprotected routes */
export const WHITE_LIST = ['/login', '/forget'];

/** Home path, if empty, take the first menu address */
export const HOME_PATH = '/dashboard/analysis';

/** Layout route address */
export const LAYOUT_PATH = '/';

/** Refresh route address */
export const REDIRECT_PATH = '/redirect';

/** Whether to cache components after opening the tab bar */
//export const TAB_KEEP_ALIVE = !import.meta.env.DEV;
export const TAB_KEEP_ALIVE = true;

/** Token cache name */
export const TOKEN_CACHE_NAME = 'token';

/** Theme cache name */
export const THEME_CACHE_NAME = 'theme';

/** I18n cache name */
export const I18N_CACHE_NAME = 'i18n-lang';

/** map key, please apply for it on the map official website */
export const MAP_KEY = '006d995d433058322319fa797f2876f5';

/** EleAdminPlus authorization code */
export const LICENSE_CODE = import.meta.env.VITE_LICENSE;
