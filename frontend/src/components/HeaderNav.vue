<script setup>
import { ref, computed } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import ToggleSwitch from 'primevue/toggleswitch'

const props = defineProps({
  isDark: Boolean
})

const isOpen = ref(false)
const showProfileMenu = ref(false)
const emit = defineEmits(['toggleDarkMode'])

const authStore = useAuthStore()
const router = useRouter()

// Get user initials from name
const userInitials = computed(() => {
  if (!authStore.user?.name) return ''

  const names = authStore.user.name.trim().split(' ')
  if (names.length === 1) {

    return names[0].substring(0, 2).toUpperCase()
  }
  // Multiple names like "first last" = "FL"
  return (names[0][0] + names[names.length - 1][0]).toUpperCase()
})

const handleLogout = () => {
  authStore.logout()
  showProfileMenu.value = false
  isOpen.value = false
  router.push('/login')
}
</script>

<template>
  <nav class="fixed top-7 left-1/2 -translate-x-1/2 w-[90%] md:w-[70%]
         backdrop-blur-sm bg-white/50
         border border-gray-200
         rounded-2xl px-6 py-6 z-50 shadow-lg
         will-change-[backdrop-filter] transform-gpu">
    <div class="flex items-center justify-between">
      <RouterLink to="/" class="text-2xl md:text-3xl font-bold">
        LinkKit
      </RouterLink>

      <!-- Desktop Navigation -->
      <div class="hidden md:flex text-xl space-x-2 font-medium">
        <RouterLink to="/about" class="hover:bg-gray-200 rounded-lg px-3 py-1.5 transition-colors duration-300">
          About
        </RouterLink>
        <RouterLink to="/tools" class="hover:bg-gray-200 rounded-lg px-3 py-1.5 transition-colors duration-300">
          Tools
        </RouterLink>
        <RouterLink v-if="authStore.isAuthenticated" to="/dashboard"
          class="hover:bg-gray-200 rounded-lg px-3 py-1.5 transition-colors duration-300">
          Dashboard
        </RouterLink>
      </div>

      <!-- Desktop Actions -->
      <div class="hidden md:flex items-center space-x-4 text-xl font-medium">
        <ToggleSwitch v-model="props.isDark" @update:model-value="$emit('toggleDarkMode')" />

        <!-- Show login/signup if not authenticated -->
        <template v-if="!authStore.isAuthenticated">
          <RouterLink to="/login" class="hover:bg-gray-200 rounded-lg px-3 py-1.5 transition-colors duration-300">
            Login
          </RouterLink>
          <RouterLink to="/signup"
            class="bg-green-400 text-white hover:opacity-80 rounded-lg px-3 py-1.5 transition-colors duration-300">
            Sign up Free
          </RouterLink>
        </template>

        <!-- Show profile if authenticated -->
        <div v-else class="relative">
          <button @click="showProfileMenu = !showProfileMenu" class="w-10 h-10 rounded-full bg-green-400 text-white font-semibold 
                   hover:opacity-80 transition-opacity duration-300 flex items-center justify-center">
            {{ userInitials }}
          </button>

          <!-- Profile Dropdown -->
          <transition enter-active-class="transition duration-200 ease-out" enter-from-class="opacity-0 scale-95"
            enter-to-class="opacity-100 scale-100" leave-active-class="transition duration-150 ease-in"
            leave-from-class="opacity-100 scale-100" leave-to-class="opacity-0 scale-95">
            <div v-if="showProfileMenu" class="absolute right-0 mt-2 w-48 bg-white border border-gray-200 
                     rounded-xl shadow-lg py-2 overflow-hidden">
              <div class="px-4 py-2 border-b border-gray-200">
                <p class="font-semibold text-sm">{{ authStore.user?.name }}</p>
                <p class="text-xs text-gray-500">{{ authStore.user?.email }}</p>
              </div>
              <RouterLink @click="showProfileMenu = false" to="/dashboard"
                class="block px-4 py-2 text-sm hover:bg-gray-100 transition-colors">
                Dashboard
              </RouterLink>
              <RouterLink @click="showProfileMenu = false" to="/profile"
                class="block px-4 py-2 text-sm hover:bg-gray-100 transition-colors">
                Profile Settings
              </RouterLink>
              <button @click="handleLogout"
                class="w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-gray-100 transition-colors">
                Logout
              </button>
            </div>
          </transition>
        </div>
      </div>

      <!-- Mobile: Dark Mode Toggle + Hamburger Menu -->
      <div class="md:hidden flex items-center gap-3">
        <div class="flex items-center">
          <ToggleSwitch v-model="props.isDark" @update:model-value="$emit('toggleDarkMode')" class="scale-90" />
        </div>

        <!-- Profile circle for mobile if authenticated -->
        <button v-if="authStore.isAuthenticated" class="w-8 h-8 rounded-full bg-green-400 text-white text-sm font-semibold 
                 flex items-center justify-center">
          {{ userInitials }}
        </button>

        <button class="text-2xl flex items-center" @click="isOpen = !isOpen">
          <i v-if="!isOpen" class="pi pi-bars"></i>
          <i v-else class="pi pi-times"></i>
        </button>
      </div>
    </div>

    <!-- Mobile Menu -->
    <transition enter-active-class="transition duration-200 ease-out" enter-from-class="opacity-0 -translate-y-2"
      enter-to-class="opacity-100 translate-y-0" leave-active-class="transition duration-150 ease-in"
      leave-from-class="opacity-100 translate-y-0" leave-to-class="opacity-0 -translate-y-2">
      <div v-if="isOpen"
        class="md:hidden mt-4 bg-gray-100 rounded-xl p-4 flex flex-col space-y-3 border border-gray-200">
        <RouterLink @click="isOpen = false" to="/about" class="py-2 transition-colors">About</RouterLink>
        <RouterLink @click="isOpen = false" to="/tools" class="py-2 transition-colors">Tools</RouterLink>

        <template v-if="authStore.isAuthenticated">
          <RouterLink @click="isOpen = false" to="/dashboard" class="py-2 transition-colors">Dashboard</RouterLink>
          <RouterLink @click="isOpen = false" to="/profile" class="py-2 transition-colors">Profile Settings</RouterLink>
          <button @click="handleLogout" class="py-2 text-left text-red-600 transition-colors">
            Logout
          </button>
        </template>

        <template v-else>
          <RouterLink @click="isOpen = false" to="/login" class="py-2 transition-colors">Login</RouterLink>
          <RouterLink @click="isOpen = false" to="/signup"
            class="bg-green-400 rounded-lg px-6 py-3 text-center text-white transition-colors duration-300">
            Sign up Free
          </RouterLink>
        </template>
      </div>
    </transition>
  </nav>
</template>