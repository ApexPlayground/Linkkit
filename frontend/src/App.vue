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
  <div class="min-h-screen bg-white text-black transition-colors duration-300">
    <Nav @toggle-dark-mode="toggleDarkMode" :is-dark="isDark"></Nav>
    <RouterView />
  </div>
</template>