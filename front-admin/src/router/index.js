import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '仪表盘', icon: 'dashboard' }
    }]
  },

  {
    path: '/write',
    component: Layout,
    redirect: '/write/article',
    name: 'write',
    meta: { title: '写作管理', icon: 'form' },
    children: [
      {
        path: 'article',
        name: 'article',
        component: () => import('@/views/article/index'),
        meta: { title: '文章管理' }
      },
      {
        path: 'category',
        name: 'category',
        component: () => import('@/views/category/index'),
        meta: {title: '分类管理'}
      },
      // {
      //   path: 'comments',
      //   name: 'comments',
      //   component: () => import('@/views/table/index'),
      //   meta: {title: '评论管理'}
      // },
      {
        path: 'recycle',
        name: 'recycle',
        component: () => import('@/views/article/recycle'),
        meta: { title: '回收站'}
      },
    ]
  },
  {
    path: '/system',
    component: Layout,
    redirect: '/system',
    name: 'system',
    meta: { title: '系统设置', icon: 'system' },
    children: [
      {
        path: 'nav',
        name: 'nav',
        component: () => import('@/views/nav/index'),
        meta: { title: '导航管理' }
      },
      {
        path: 'flink',
        name: 'flink',
        component: () => import('@/views/flink/index'),
        meta: { title: '友情链接' }
      },
      {
        path: 'seo',
        name: 'seo',
        component: () => import('@/views/seo/index'),
        meta: { title: '站点设置' }
      },
    ]
  },
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
