<script setup>
import { ref, onMounted, computed } from 'vue'
import Nav from './components/HeaderNav.vue'
import { RouterView } from 'vue-router'

const isDark = ref(false)

const gridStyle = computed(() => ({
  background: isDark.value ? '#1a1a1a' : '#ffffff',
  backgroundImage: isDark.value
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

onMounted(() => {
  const savedMode = localStorage.getItem('darkMode')
  isDark.value = savedMode ? savedMode === 'true' : window.matchMedia('(prefers-color-scheme: dark)').matches
  if (isDark.value) {
    document.documentElement.classList.add('dark')
  }
})

const toggleDarkMode = () => {
  isDark.value = !isDark.value
  if (isDark.value) {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }
  localStorage.setItem('darkMode', isDark.value.toString())
}
</script>

<template>
  <div class="min-h-screen bg-white text-black transition-colors duration-300 relative">
    <div class="absolute inset-0 z-0" :style="gridStyle" />
    <div class="relative z-10">
      <Nav @toggle-dark-mode="toggleDarkMode" :is-dark="isDark"></Nav>
      <div class="mx-auto py-24">
        <RouterView />
      </div>
    </div>
  </div>
</template>