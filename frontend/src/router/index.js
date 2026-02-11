import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/LandingPage.vue'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'landing',
      component: HomeView,
      beforeEnter: (to, from, next) => {
        const auth = useAuthStore()
        auth.isAuthenticated ? next('/home') : next()
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
        auth.isAuthenticated ? next('/home') : next()
      },
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue'),
      beforeEnter: (to, from, next) => {
        const auth = useAuthStore()
        auth.isAuthenticated ? next('/home') : next()
      },
    },
    {
      path: '/home',
      name: 'home',
      component: () => import('../views/HomeView.vue'),
      beforeEnter: (to, from, next) => {
        const auth = useAuthStore()
        auth.isAuthenticated ? next() : next('/login')
      },
      children: [
        {
          path: '',
          name: 'home-content',
          component: () => import('../views/HomeContentView.vue'),
        },
        {
          path: '/home/links',
          name: 'links',
          component: () => import('../views/LinksView.vue'),
        },
        {
          path: '/home/qr-codes',
          name: 'qr-code',
          component: () => import('../views/QR-CodeView.vue'),
        },
        {
          path: '/home/analytics',
          name: 'analytics',
          component: () => import('../views/AnalyticsView.vue'),
        },
      ],
    },
  ],
})

export default router
