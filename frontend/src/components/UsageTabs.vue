<script setup>
import { ref } from 'vue'
import Tabs from 'primevue/tabs'
import TabList from 'primevue/tablist'
import Tab from 'primevue/tab'
import TabPanels from 'primevue/tabpanels'
import TabPanel from 'primevue/tabpanel'

import LinkShortener from './LinkShortner.vue'
import QRGenerator from './QRGenerator.vue'

const items = ref([
  { value: 'link', label: 'Link', icon: 'pi pi-link' },
  { value: 'qr', label: 'QR Code', icon: 'pi pi-qrcode' },
])
</script>

<template>
  <!-- Tabs Card -->
  <div class="card p-4 sm:p-6 md:p-8 border bg-white/80 backdrop-blur-md border-gray-200 rounded-2xl shadow-lg">
    <Tabs value="link">
      <!-- Tab Headers -->
      <TabList class="flex flex-wrap justify-center gap-2 mb-4 sm:mb-6 px-2">
        <Tab v-for="tab in items" :key="tab.value" :value="tab.value"
          class="flex items-center justify-center gap-2 px-4 py-2.5 min-w-30 flex-1 sm:flex-initial">
          <i :class="tab.icon" class="text-sm sm:text-base" />
          <span class="text-sm sm:text-base whitespace-nowrap">{{ tab.label }}</span>
        </Tab>
      </TabList>

      <!-- Tab Panels -->
      <TabPanels class="px-2 sm:px-4">
        <!-- Link Tab -->
        <TabPanel value="link">
          <LinkShortener />
        </TabPanel>

        <!-- QR Tab -->
        <TabPanel value="qr">
          <QRGenerator />
        </TabPanel>
      </TabPanels>
    </Tabs>
  </div>
</template>