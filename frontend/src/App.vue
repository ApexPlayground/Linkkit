<script setup>
import { onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useThemeStore } from '@/stores/theme'
import Nav from './components/HeaderNav.vue'
import FooterView from './views/FooterView.vue'
import { RouterView } from 'vue-router'

const route = useRoute()
const themeStore = useThemeStore()

const gridStyle = computed(() => ({
  background: themeStore.isDark ? '#1a1a1a' : '#ffffff',
  backgroundImage: themeStore.isDark
    ? `
      radial-gradient(circle at 1px 1px, rgba(255, 255, 255, 0.15) 1px, transparent 0),
      radial-gradient(circle at 50% 50%, rgba(52,211,153,0.25) 0%, rgba(52,211,153,0.1) 40%, transparent 80%)
    `
    : `
      radial-gradient(circle at 1px 1px, rgba(0, 0, 0, 0.35) 1px, transparent 0),
      radial-gradient(circle at 50% 50%, rgba(16,185,129,0.25) 0%, rgba(16,185,129,0.1) 40%, transparent 80%)
    `,
  backgroundSize: '20px 20px, 100% 100%'
}))

// Check if current route is dashboard
const isDashboardRoute = computed(() => {
  return route.path.startsWith('/dashboard')
})

onMounted(() => {
  themeStore.initializeTheme()
})
</script>

<template>
  <div class="min-h-screen bg-white text-black transition-colors duration-300 relative">
    <div v-if="!isDashboardRoute" class="absolute inset-0 z-0" :style="gridStyle" />
    <div :class="isDashboardRoute ? '' : 'relative z-10'">
      <Nav v-if="!isDashboardRoute" />
      <div :class="isDashboardRoute ? '' : 'mx-auto py-24'">
        <RouterView />
      </div>
      <FooterView v-if="!isDashboardRoute" />
    </div>
  </div>
</template>