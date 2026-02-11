<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import Tabs from 'primevue/tabs'
import TabList from 'primevue/tablist'
import Tab from 'primevue/tab'
import TabPanels from 'primevue/tabpanels'
import TabPanel from 'primevue/tabpanel'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import Message from 'primevue/message'

import api from '@/services/api'
import { useAuthStore } from '@/stores/auth'

const items = ref([
  { value: 'link', label: 'Link', icon: 'pi pi-link' },
  { value: 'qr', label: 'QR Code', icon: 'pi pi-qrcode' },
])

const router = useRouter()
const authStore = useAuthStore()

const error = ref('')
const longUrl = ref('')
const shortUrl = ref('')
const showSuccessDialog = ref(false)
const validationError = ref('')

// URL validation function
const isValidUrl = (url) => {
  try {
    const urlObj = new URL(url)
    // Only allow http and https protocols
    if (!['http:', 'https:'].includes(urlObj.protocol)) {
      return false
    }
    // Basic domain validation
    if (!urlObj.hostname || urlObj.hostname.length < 3) {
      return false
    }
    return true
  } catch {
    return false
  }
}

const handleLinkSubmit = async () => {
  if (!authStore.isAuthenticated) {
    router.push('/login')
    return
  }

  validationError.value = ''
  error.value = ''

  if (!longUrl.value.trim()) {
    validationError.value = 'Please enter a URL'
    return
  }

  const validUrl = longUrl.value.trim()

  if (!isValidUrl(validUrl)) {
    validationError.value = 'Please enter a valid URL (e.g., https://example.com)'
    return
  }

  console.log('Submitting:', validUrl)

  try {
    const { data } = await api.post('v1/shorten', {
      long_url: validUrl,
    })
    console.log('Success:', data)

    // Store the shortened URL and show dialog
    shortUrl.value = data.short_url
    showSuccessDialog.value = true

    // Clear the input
    longUrl.value = ''
  } catch (err) {
    if (err.response && err.response.data) {
      error.value =
        err.response.data.message || err.response.data.error || 'Failed to create short link'
    } else {
      error.value = err.message || 'Something went wrong.'
    }
    console.error('Error:', error.value)
  }
}

const copyToClipboard = async () => {
  try {
    await navigator.clipboard.writeText(shortUrl.value)
    console.log('Copied to clipboard!')
  } catch (err) {
    console.error('Failed to copy:', err)
  }
}

const openLink = (url) => {
  window.open(url)
}
</script>

<template>
  <!-- Tabs Card -->
  <div
    class="card p-4 sm:p-6 md:p-8 border bg-white/80 backdrop-blur-md border-gray-200 rounded-2xl shadow-lg"
  >
    <Tabs value="link">
      <!-- Tab Headers -->
      <TabList class="flex flex-wrap justify-center gap-2 mb-4 sm:mb-6 px-2">
        <Tab
          v-for="tab in items"
          :key="tab.value"
          :value="tab.value"
          class="flex items-center justify-center gap-2 px-4 py-2.5 min-w-30 flex-1 sm:flex-initial"
        >
          <i :class="tab.icon" class="text-sm sm:text-base" />
          <span class="text-sm sm:text-base whitespace-nowrap">{{ tab.label }}</span>
        </Tab>
      </TabList>

      <!-- Tab Panels -->
      <TabPanels class="px-2 sm:px-4">
        <!-- Link Tab -->
        <TabPanel value="link">
          <div class="flex flex-col gap-4 md:gap-6">
            <!-- Title -->
            <h2
              class="text-2xl sm:text-3xl md:text-4xl font-semibold sm:text-left mb-3 sm:mb-7 text-gray-700"
            >
              Shorten a long link
            </h2>

            <div class="space-y-2">
              <p class="font-semibold text-gray-700">Paste your long link here</p>

              <!-- Validation Error Message -->
              <Message v-if="validationError" severity="error" :closable="false">
                {{ validationError }}
              </Message>

              <!-- API Error Message -->
              <Message v-if="error" severity="error" :closable="false">
                {{ error }}
              </Message>

              <form @submit.prevent="handleLinkSubmit">
                <div class="flex flex-col sm:flex-row gap-3 sm:gap-2">
                  <InputText
                    v-model="longUrl"
                    type="text"
                    placeholder="Enter your long URL"
                    class="flex-1 text-sm sm:text-base"
                    fluid
                  />
                  <Button
                    type="submit"
                    label="Shorten"
                    class="w-full sm:w-auto sm:min-w-35 py-3 sm:py-2"
                  />
                </div>
              </form>
            </div>
          </div>
        </TabPanel>

        <!-- QR Tab -->
        <TabPanel value="qr">
          <div class="flex flex-col gap-4 md:gap-6">
            <!-- Title -->
            <h2
              class="text-2xl text-gray-700 sm:text-3xl md:text-4xl font-semibold sm:text-left mb-3 sm:mb-7"
            >
              Generate QR Code
            </h2>

            <div class="space-y-2">
              <p class="font-semibold text-gray-700">Enter QR Code destination</p>
              <div class="flex flex-col sm:flex-row gap-3 sm:gap-2">
                <InputText
                  type="text"
                  placeholder="Enter URL for QR code"
                  class="flex-1 text-sm sm:text-base"
                  fluid
                />
                <Button label="Generate" class="w-full sm:w-auto sm:min-w-35 py-3 sm:py-2" />
              </div>
            </div>
          </div>
        </TabPanel>
      </TabPanels>
    </Tabs>
  </div>

  <!-- Success Dialog -->
  <Dialog
    v-model:visible="showSuccessDialog"
    modal
    header="Link Shortened Successfully!"
    :style="{ width: '90vw', maxWidth: '500px' }"
  >
    <div class="flex flex-col gap-4">
      <p class="text-gray-600">Your shortened link is ready:</p>

      <div class="flex gap-2">
        <InputText :value="shortUrl" readonly class="flex-1" fluid />
        <Button
          icon="pi pi-copy"
          @click="copyToClipboard"
          severity="secondary"
          v-tooltip.top="'Copy to clipboard'"
        />
      </div>

      <div class="flex gap-2 justify-end mt-2">
        <Button label="Close" @click="showSuccessDialog = false" severity="secondary" />
        <Button label="Open Link" @click="() => openLink(shortUrl)" icon="pi pi-external-link" />
      </div>
    </div>
  </Dialog>
</template>
