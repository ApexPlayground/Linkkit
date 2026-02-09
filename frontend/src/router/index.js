import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      beforeEnter: (to, from, next) => {
        const auth = useAuthStore()
        auth.isAuthenticated ? next('/dashboard') : next()
      },
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue'),
    },
    {
      path: '/signup',
      name: 'signup',
      component: () => import('../views/SignUpView.vue'),
      beforeEnter: (to, from, next) => {
        const auth = useAuthStore()
        auth.isAuthenticated ? next('/dashboard') : next()
      },
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue'),
      beforeEnter: (to, from, next) => {
        const auth = useAuthStore()
        auth.isAuthenticated ? next('/dashboard') : next()
      },
    },
  ],
})

export default router
