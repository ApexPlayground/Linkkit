<script setup>
import { computed } from 'vue'

const props = defineProps({
    user: {
        type: Object,
        required: true
    }
})

const userInitials = computed(() => {
    if (!props.user?.name) return ''

    const names = props.user.name.trim().split(' ')
    if (names.length === 1) {
        return names[0].substring(0, 2).toUpperCase()
    }
    return (names[0][0] + names[names.length - 1][0]).toUpperCase()
})
</script>

<template>
    <div class="flex items-center gap-3 p-3 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors cursor-pointer">
        <div
            class="w-10 h-10 rounded-full bg-green-400 text-white font-semibold flex items-center justify-center shrink-0">
            {{ userInitials }}
        </div>
        <div class="flex-1 min-w-0">
            <p class="text-sm font-semibold text-gray-900 truncate">{{ user.name }}</p>
            <p class="text-xs text-gray-700 truncate">{{ user.email }}</p>
        </div>
    </div>
</template>