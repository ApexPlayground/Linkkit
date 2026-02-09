<script setup>
const props = defineProps({
    label: {
        type: String,
        required: true
    },
    icon: {
        type: String,
        required: true
    },
    description: {
        type: String,
        default: ''
    },
    isActive: {
        type: Boolean,
        default: false
    }
})

const emit = defineEmits(['click'])
</script>

<template>
    <button @click="emit('click')" :class="[
        'w-full flex items-center gap-3 px-4 py-3 rounded-lg transition-all duration-200',
        'group',
        isActive
            ? 'bg-green-400 text-white'
            : 'text-gray-700 hover:bg-green-400/10'
    ]">
        <i :class="[
            icon,
            'text-lg transition-colors',
            isActive ? 'text-white' : 'text-gray-700 group-hover:text-green-400'
        ]"></i>
        <div class="flex-1 text-left">
            <p class="text-sm font-medium">{{ label }}</p>
            <p v-if="description" :class="[
                'text-xs',
                isActive ? 'text-white/80' : 'text-gray-700'
            ]">{{ description }}</p>
        </div>
        <i v-if="isActive" class="pi pi-chevron-right text-white text-xs"></i>
    </button>
</template>