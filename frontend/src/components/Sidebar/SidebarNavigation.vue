<script setup>
import { useRoute } from 'vue-router'
import SidebarNavItem from './SidebarNavItem.vue'

const props = defineProps({
    items: {
        type: Array,
        required: true
    }
})

const emit = defineEmits(['navigate'])
const route = useRoute()

const isActive = (itemRoute) => {
    if (itemRoute === '/dashboard') {
        return route.path === '/dashboard'
    }
    return route.path === itemRoute
}


const handleClick = (itemRoute) => {
    emit('navigate', itemRoute)
}
</script>

<template>
    <nav class="flex-1 overflow-y-auto p-4">
        <div class="space-y-5">
            <SidebarNavItem v-for="item in items" :key="item.route" :label="item.label" :icon="item.icon"
                :description="item.description" :is-active="isActive(item.route)" @click="handleClick(item.route)" />
        </div>
    </nav>
</template>

<style scoped>
nav::-webkit-scrollbar {
    width: 6px;
}

nav::-webkit-scrollbar-track {
    background: transparent;
}

nav::-webkit-scrollbar-thumb {
    background: #d1d5db;
    border-radius: 3px;
}

nav::-webkit-scrollbar-thumb:hover {
    background: #9ca3af;
}
</style>