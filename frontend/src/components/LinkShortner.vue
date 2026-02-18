<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from 'primevue/usetoast'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import Message from 'primevue/message'

import api from '@/services/api'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()
const toast = useToast()

const error = ref('')
const longUrl = ref('')
const shortUrl = ref('')
const showSuccessDialog = ref(false)
const validationError = ref('')

const showToast = (severity, summary, detail) => {
    if (toast && typeof toast.add === 'function') {
        toast.add({ severity, summary, detail, life: 3000 })
    } else {
        console.warn('Toast not available:', { severity, summary, detail })
    }
}

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


        if (data.message === "Link already exists") {
            validationError.value = 'You have already shortened this link.'
        } else {
            validationError.value = ''
        }

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
        showToast('success', 'Copied!', 'Link copied to clipboard')
    } catch (err) {
        console.error('Failed to copy:', err)
        showToast('error', 'Copy Failed', 'Could not copy Link')
    }
}

const openLink = (url) => {
    window.open(url)
}
</script>

<template>
    <div class="flex flex-col gap-4 md:gap-6">
        <!-- Title -->
        <h2 class="text-2xl sm:text-3xl md:text-4xl font-semibold sm:text-left mb-3 sm:mb-7 text-gray-700">
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
                    <InputText v-model="longUrl" type="text" placeholder="Enter your long URL"
                        class="flex-1 text-sm sm:text-base" fluid />
                    <Button type="submit" label="Shorten" class="w-full sm:w-auto sm:min-w-35 py-3 sm:py-2" />
                </div>
            </form>
        </div>
    </div>

    <!-- Success Dialog -->
    <Dialog v-model:visible="showSuccessDialog" modal header="Link Shortened Successfully!"
        :style="{ width: '90vw', maxWidth: '500px' }">
        <div class="flex flex-col gap-4">
            <p class="text-gray-600">Your shortened link is ready:</p>
            <p class="text-red-600">
                {{ validationError || 'Your shortened link is ready:' }}
            </p>


            <div class="flex gap-2">
                <InputText :value="shortUrl" readonly class="flex-1" fluid />
                <Button icon="pi pi-copy" @click="copyToClipboard" severity="secondary"
                    v-tooltip.top="'Copy to clipboard'" />
            </div>


            <div class="flex gap-2 justify-end mt-2">
                <Button label="Close" @click="showSuccessDialog = false" severity="secondary" />
                <Button label="Open Link" @click="() => openLink(shortUrl)" icon="pi pi-external-link" />
            </div>
        </div>
    </Dialog>
</template>