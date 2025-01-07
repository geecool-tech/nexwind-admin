import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const AUTH: AppRouteRecordRaw = {
  path: '/auth',
  name: 'auth',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.user',
    icon: 'icon-user',
    requiresAuth: true,
    order: 7,
  },
  children: [
    {
      path: 'admin-list',
      name: 'AdminList',
      component: () => import('@/views/auth/admin-list/index.vue'),
      meta: {
        locale: 'auth.admin.list.name',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'admin-log',
      name: 'AdminLog',
      component: () => import('@/views/auth/admin-log/admin-log.vue'),
      meta: {
        locale: 'auth.admin.log.name',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'menu-list',
      name: 'MenuList',
      component: () => import('@/views/auth/menu-list/index.vue'),
      meta: {
        locale: 'auth.rule.list.name',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'api-list',
      name: 'ApiList',
      component: () => import('@/views/auth/api-list/index.vue'),
      meta: {
        locale: 'auth.api.list.name',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'roles-list',
      name: 'RolesList',
      component: () => import('@/views/auth/roles-list/index.vue'),
      meta: {
        locale: 'auth.role.list.name',
        requiresAuth: true,
        roles: ['*'],
      },
    },
  ],
};

export default AUTH;
