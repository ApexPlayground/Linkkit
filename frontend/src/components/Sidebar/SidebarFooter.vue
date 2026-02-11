<script setup>
import { useRoute } from 'vue-router'
import ToggleSwitch from 'primevue/toggleswitch'

const props = defineProps({
    isDark: {
        type: Boolean,
        required: true
    }
})

const emit = defineEmits(['navigate', 'logout', 'update:isDark'])
const route = useRoute()

const isActive = (routePath) => {
    return route.path === routePath || route.path.startsWith(routePath + '/')
}

const handleThemeChange = (value) => {
    emit('update:isDark', value)
}
</script>

<template>
    <div class="p-4 border-t border-gray-200 space-y-2">
        <!-- Dark Mode Toggle -->
        <div class="flex items-center justify-between px-4 py-2.5 hover:bg-green-400/10 rounded-lg transition-colors">
            <div class="flex items-center gap-3">
                <i :class="isDark ? 'pi pi-moon' : 'pi pi-sun'" class="text-green-400"></i>
                <span class="text-sm font-medium text-gray-700">
                    {{ isDark ? 'Dark' : 'Light' }} Mode
                </span>
            </div>
            <ToggleSwitch :model-value="isDark" @update:model-value="handleThemeChange" />
        </div>

        <!-- Settings Button -->
        <button @click="emit('navigate', '/dashboard/settings')" :class="[
            'w-full flex items-center gap-3 px-4 py-2.5 rounded-lg transition-all duration-200 text-left',
            isActive('/dashboard/settings')
                ? 'bg-green-400 text-white'
                : 'text-gray-700 hover:bg-green-400/10 hover:text-green-400'
        ]">
            <i :class="[
                'pi pi-cog text-base',
                isActive('/dashboard/settings') ? 'text-white' : 'text-green-400'
            ]"></i>
            <span class="text-sm font-medium">Settings</span>
        </button>

        <!-- Logout Button -->
        <button @click="emit('logout')"
            class="w-full flex items-center gap-3 px-4 py-2.5 rounded-lg transition-all duration-200 text-left text-red-600 hover:bg-red-500/20">
            <i class="pi pi-sign-out text-base"></i>
            <span class="text-sm font-medium">Logout</span>
        </button>
    </div>
</template>