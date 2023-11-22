import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'mylogin',
      component: () => import('@/views/mylogin.vue')
    }
  ]
})

export default router
