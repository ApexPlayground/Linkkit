<script setup>
import { ref, onMounted } from 'vue'
import Nav from './components/HeaderNav.vue'
import { RouterView } from 'vue-router'

const isDark = ref(false)

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




  <div class="absolute inset-0 z-0" :style="{
    backgroundImage: 'linear-gradient(to right, #e2e8f0 1px, transparent 1px), linear-gradient(to bottom, #e2e8f0 1px, transparent 1px)',
    backgroundSize: '20px 30px',
    WebkitMaskImage: 'radial-gradient(ellipse 70% 60% at 50% 0%, #000 60%, transparent 100%)',
    maskImage: 'radial-gradient(ellipse 70% 60% at 50% 0%, #000 60%, transparent 100%)'
  }" />
  <div class="min-h-screen  bg-white text-black transition-colors duration-300">
    <Nav @toggle-dark-mode="toggleDarkMode" :is-dark="isDark"></Nav>
    <div class="mx-auto py-24">
      <RouterView />
    </div>
  </div>
</template>
