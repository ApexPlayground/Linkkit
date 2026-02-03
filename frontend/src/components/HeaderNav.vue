<script setup>
import { ref } from 'vue'
import { RouterLink } from 'vue-router'
import ToggleSwitch from 'primevue/toggleswitch'

import Slider from 'primevue/slider';


const props = defineProps({
  isDark: Boolean
})

const isOpen = ref(false)
const emit = defineEmits(['toggleDarkMode'])
</script>

<template>
  <nav class="fixed top-7 left-1/2 -translate-x-1/2 w-[90%] md:w-[70%]
                bg-white/80 backdrop-blur-sm
                border border-gray-200
                rounded-2xl px-6 py-6 z-50 shadow-lg transition-all duration-300">
    <div class="flex items-center justify-between">
      <RouterLink to="/" class="text-2xl md:text-3xl font-bold">
        LinkKit
      </RouterLink>

      <div class="hidden md:flex text-xl space-x-2 font-medium">
        <RouterLink to="/about" class="hover:bg-gray-100 rounded-lg px-3 py-1.5 transition-all duration-500">
          About
        </RouterLink>
        <RouterLink to="/tools" class="hover:bg-gray-100 rounded-lg px-3 py-1.5 transition-all duration-500">
          Tools
        </RouterLink>
        <RouterLink to="/dashboard" class="hover:bg-gray-100 rounded-lg px-3 py-1.5 transition-all duration-500">
          Dashboard
        </RouterLink>






      </div>

      <div class="hidden md:flex items-center space-x-4 text-xl font-medium">
        <ToggleSwitch v-model="props.isDark" @update:model-value="$emit('toggleDarkMode')" />
        <RouterLink to="/login" class="hover:bg-gray-100 rounded-lg px-3 py-1.5 transition-all duration-500">
          Login
        </RouterLink>
        <RouterLink to="/signup"
          class="bg-black text-white hover:opacity-80 rounded-lg px-3 py-1.5 transition-all duration-500">
          Sign up Free
        </RouterLink>
      </div>

      <button class="md:hidden text-2xl" @click="isOpen = !isOpen">
        <i v-if="!isOpen" class="pi pi-bars"></i>
        <i v-else class="pi pi-times"></i>
      </button>
    </div>

    <transition enter-active-class="transition duration-200 ease-out" enter-from-class="opacity-0 -translate-y-2"
      enter-to-class="opacity-100 translate-y-0" leave-active-class="transition duration-150 ease-in"
      leave-from-class="opacity-100 translate-y-0" leave-to-class="opacity-0 -translate-y-2">
      <div v-if="isOpen"
        class="md:hidden mt-4 bg-gray-100 rounded-xl p-4 flex flex-col space-y-3 border border-gray-200">
        <RouterLink @click="isOpen = false" to="/" class="py-2 transition-colors">Home</RouterLink>
        <RouterLink @click="isOpen = false" to="/about" class="py-2 transition-colors">About</RouterLink>
        <RouterLink @click="isOpen = false" to="/tools" class="py-2 transition-colors">Tools</RouterLink>
        <RouterLink @click="isOpen = false" to="/dashboard" class="py-2 transition-colors">Dashboard</RouterLink>
        <RouterLink @click="isOpen = false" to="/login" class="py-2 transition-colors">Login</RouterLink>
        <ToggleSwitch v-model="props.isDark" @update:model-value="$emit('toggleDarkMode')" class="py-2" />
        <RouterLink @click="isOpen = false" to="/signup"
          class="bg-orange-400 hover:bg-orange-500 rounded-lg px-6 py-3 text-center text-white transition-all duration-300">
          Sign up Free
        </RouterLink>
      </div>
    </transition>
  </nav>
</template>