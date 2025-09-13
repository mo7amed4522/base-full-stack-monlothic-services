import RedirectLayout from '@/components/RedirectLayout/index.vue';
import { HOME_PATH, LAYOUT_PATH, REDIRECT_PATH } from '@/config/setting';
import Layout from '@/layout/index.vue';
import { eachTree, menuToRoutes } from 'ele-admin-plus/es';
const modules = import.meta.glob('/src/views/**/index.vue');

/**
 * Page tab static routes
 */
export const routes = [
  {
    path: '/login',
    component: () => import('@/views/login/index.vue'),
    meta: { title: 'Log in' }
  },
  // 404
  {
    path: '/:path(.*)*',
    component: () => import('@/views/exception/404/index.vue')
  }

  // {
  //   path: '/finance',
  //   component: Layout,
  //   children: [
  //     {
  //       path: '/deposit/index',
  //       component: () => import('@/views/finance/deposit/index.vue'),
  //       meta: { title: 'Deposit audit' }
  //     },
  //     {
  //       path: '/withdrawal/index',
  //       component: () => import('@/views/finance/withdrawal/index.vue'),
  //       meta: { title: 'Withdrawal management' }
  //     }
  //   ]
  // },
  //
  // {
  //   path: '/member',
  //   component: Layout,
  //   children: [
  //     {
  //       path: '/member/index',
  //       component: () => import('@/views/member/index/index.vue'),
  //       meta: { title: 'Member list' }
  //     },
  //     {
  //       path: '/risk/index',
  //       component: () => import('@/views/member/risk/index.vue'),
  //       meta: { title: 'Risk users' }
  //     },
  //     {
  //       path: '/member/grade',
  //       component: () => import('@/views/member/grade/index.vue'),
  //       meta: { title: 'Member grade' }
  //     },
  //     {
  //       path: '/member/level',
  //       component: () => import('@/views/member/level/index.vue'),
  //       meta: { title: 'Member level' }
  //     },
  //     {
  //       path: '/member/demoAccount',
  //       component: () => import('@/views/member/demoAccount/index.vue'),
  //       meta: { title: 'Demo account' }
  //     }
  //   ]
  // }
];

/**
 * Generate dynamic routes based on menus
 * @param menus Menu data
 * @param homePath Home path
 */
export function getMenuRoutes(menus, homePath) {
  const childRoutes = [
    // Routing for refresh
    {
      path: REDIRECT_PATH + '/:path(.*)',
      component: RedirectLayout,
      meta: { hideFooter: true }
    }
  ];
  const layoutRoutes = [
    {
      path: LAYOUT_PATH,
      component: Layout,
      redirect: HOME_PATH ?? homePath,
      children: childRoutes
    }
  ];
  // Flatten the routing processing
  eachTree(menuToRoutes(menus, getComponent, routes), (route) => {
    const temp = Object.assign({}, route, { children: void 0 });
    if (route.meta?.layout === false) {
      layoutRoutes.push(temp); // Routing that does not need an outer layout
    } else {
      childRoutes.push(temp); // Routing that needs an outer layout
    }
  });
  return layoutRoutes;
}

/**
 * Parse route components
 * @param component Component name
 */
function getComponent(component) {
  if (component) {
    const module = modules[`/src/views${component}.vue`];
    if (!module) {
      return modules[`/src/views${component}/index.vue`];
    }
    return module;
  }
}
