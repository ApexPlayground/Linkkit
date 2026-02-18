<script setup>
import { onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useThemeStore } from '@/stores/theme'
import Nav from './components/HeaderNav.vue'
import FooterView from './views/FooterView.vue'
import { RouterView } from 'vue-router'
import Toast from 'primevue/toast'

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
  backgroundSize: '20px 20px, 100% 100%',
}))

// Check if current route is dashboard
const isHomeRoute = computed(() => route.path.startsWith('/home'))

onMounted(() => {
  themeStore.initializeTheme()
})
</script>

<template>
  <div class="min-h-screen bg-white text-black transition-colors duration-300 relative">
    <div v-if="!isHomeRoute" class="absolute inset-0 z-0" :style="gridStyle" />
    <div :class="isHomeRoute ? '' : 'relative z-10'">
      <Nav v-if="!isHomeRoute" />
      <div :class="isHomeRoute ? '' : 'mx-auto py-24'">
        <Transition name="page" mode="out-in">
          <RouterView />
        </Transition>
      </div>

      <FooterView v-if="!isHomeRoute" />
    </div>


    <Toast />
  </div>
</template>

<style>
.page-enter-active,
.page-leave-active {
  transition:
    opacity 0.2s ease,
    transform 0.2s ease;
}

.page-enter-from,
.page-leave-to {
  opacity: 0;
  transform: translateY(10px);
}

.page-enter-to,
.page-leave-from {
  opacity: 1;
  transform: translateY(0);
}
</style>
