<script setup>
import Button from 'primevue/button'
import SidebarUserCard from './SidebarUserCard.vue'

const props = defineProps({
    user: {
        type: Object,
        required: true
    },
    showCloseButton: {
        type: Boolean,
        default: false
    }
})

const emit = defineEmits(['close'])
</script>

<template>
    <div class="p-6 border-b border-gray-200">
        <div class="flex items-center justify-between mb-6">
            <h1 class="text-2xl font-bold text-gray-900">LinkKit</h1>
            <Button v-if="showCloseButton" icon="pi pi-times" text rounded severity="secondary"
                class="lg:hidden! w-8! h-8!" @click="emit('close')" />
        </div>

        <SidebarUserCard :user="user" />
    </div>
</template>