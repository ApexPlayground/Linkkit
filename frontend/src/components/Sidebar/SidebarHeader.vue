<script setup>
import Button from 'primevue/button'
import SidebarUserCard from './SidebarUserCard.vue'
import logo from '/lk_logo.png'

const props = defineProps({
  user: {
    type: Object,
    required: true,
  },
  showCloseButton: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits(['close'])
</script>

<template>
  <div class="p-6 border-b border-gray-200">
    <div class="flex items-center justify-between mb-6">
      <h1 class="flex gap-3 items-center text-4xl font-bold text-gray-900">
        <img :src="logo" alt="LinkKit logo" class="h-9 md:h-12 rounded-lg bg-amber-700" />Linkkit
      </h1>
      <Button v-if="showCloseButton" icon="pi pi-times" text rounded severity="secondary" class="lg:hidden! w-8! h-8!"
        @click="emit('close')" />
    </div>

    <SidebarUserCard :user="user" />
  </div>
</template>
