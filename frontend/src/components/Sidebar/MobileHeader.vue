<script setup>
import { ref, computed } from 'vue'
import Button from 'primevue/button'

const props = defineProps({
    user: {
        type: Object,
        required: true
    }
})

const emit = defineEmits(['toggleSidebar', 'navigate', 'logout'])

const showProfileMenu = ref(false)

const userInitials = computed(() => {
    if (!props.user?.name) return ''

    const names = props.user.name.trim().split(' ')
    if (names.length === 1) {
        return names[0].substring(0, 2).toUpperCase()
    }
    return (names[0][0] + names[names.length - 1][0]).toUpperCase()
})

const handleNavigate = (route) => {
    emit('navigate', route)
    showProfileMenu.value = false
}

const handleLogout = () => {
    emit('logout')
    showProfileMenu.value = false
}
</script>

<template>
    <div
        class="lg:hidden fixed top-0 left-0 right-0 bg-white border-b border-gray-200 px-4 py-3 flex items-center justify-between z-40">
        <div class="flex items-center gap-3">
            <Button icon="pi pi-bars" text rounded severity="secondary" @click="emit('toggleSidebar')"
                class="w-10! h-10!" />
            <h1 class="text-xl font-bold">LinkKit</h1>
        </div>

        <!-- Mobile User Avatar with Dropdown -->
        <div class="relative">
            <button @click="showProfileMenu = !showProfileMenu"
                class="w-9 h-9 rounded-full bg-green-400 text-white text-sm font-semibold flex items-center justify-center hover:opacity-80 transition-opacity">
                {{ userInitials }}
            </button>

            <!-- Mobile Profile Dropdown -->
            <transition enter-active-class="transition duration-200 ease-out" enter-from-class="opacity-0 scale-95"
                enter-to-class="opacity-100 scale-100" leave-active-class="transition duration-150 ease-in"
                leave-from-class="opacity-100 scale-100" leave-to-class="opacity-0 scale-95">
                <div v-if="showProfileMenu"
                    class="absolute right-0 mt-2 w-48 bg-white border border-gray-200 rounded-xl shadow-lg py-2 overflow-hidden">
                    <div class="px-4 py-2 border-b border-gray-200">
                        <p class="font-semibold text-sm truncate">{{ user.name }}</p>
                        <p class="text-xs text-gray-500 truncate">{{ user.email }}</p>
                    </div>
                    <button @click="handleNavigate('/dashboard/settings')"
                        class="w-full text-left px-4 py-2 text-sm hover:bg-gray-100 transition-colors">
                        <i class="pi pi-cog mr-2"></i>Settings
                    </button>
                    <button @click="handleLogout"
                        class="w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-gray-100 transition-colors">
                        <i class="pi pi-sign-out mr-2"></i>Logout
                    </button>
                </div>
            </transition>
        </div>
    </div>
</template>