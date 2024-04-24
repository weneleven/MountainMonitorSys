import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: () => import('@/views/Login.vue')
    },
    {
      path: '/manager',
      name: 'manager',
      component: () => import('@/views/manager/Manager.vue')
    },
    {
      path: '/course',
      name: 'course',
      component: () => import('@/views/manager/Course.vue')
    },
    {
      path: '/person',
      name: 'home',
      component: () => import('@/views/manager/Home.vue')
    },
    {
      path: '/sensor',
      name: 'sensor',
      component: () => import('@/views/manager/sensor.vue')
    },
    {
      path: '/dataview',
      name: 'dataview',
      component: () => import('@/views/manager/dataview.vue')
    },
    {
      path: '/analysis',
      name: 'analysis',
      component: () => import('@/views/manager/analysis.vue')
    },
    {
      path: '/upload',
      name: 'upload',
      component: () => import('@/views/manager/upload.vue')
    },
    {
      path: '/alert',
      name: 'alert',
      component: () => import('@/views/manager/alert.vue')
    },
    {
      path: '/test',
      name: 'test',
      component: () => import('@/views/test.vue')
    },
    {
      path: '/analysis1',
      name: 'analysis1',
      component: () => import('@/views/manager/analysis1.vue')
    },
    {
      path: '/warnview',
      name: 'warnview',
      component: () => import('@/views/manager/warn/warnview.vue')
    },
    {
      path: '/warnanalysis',
      name: 'warnanalysis',
      component: () => import('@/views/manager/warn/warnanalysis.vue')
    },
    {
      path: '/infowindow',
      name: 'infowindow',
      component: () => import('@/views/manager/map/infowindow.vue')
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('@/views/manager/dashboard/index.vue')
    },
    {
      path: '/displaceline',
      name: 'diplceline',
      component: () => import('@/views/manager/displacement/LineChart.vue')
    }
  ]
})

export default router
