<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from 'primevue/usetoast'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import Message from 'primevue/message'
import QRCode from 'qrcode'
import api from '@/services/api'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()
const toast = useToast()

// Add this helper function
const showToast = (severity, summary, detail) => {
    if (toast && typeof toast.add === 'function') {
        toast.add({ severity, summary, detail, life: 3000 })
    } else {
        console.warn('Toast not available:', { severity, summary, detail })
    }
}

const error = ref('')
const qrUrl = ref('')
const qrData = ref(null)
const qrCodeDataUrl = ref('')
const showSuccessDialog = ref(false)
const validationError = ref('')
const isGenerating = ref(false)

const isValidUrl = (url) => {
    try {
        const urlObj = new URL(url)
        if (!['http:', 'https:'].includes(urlObj.protocol)) {
            return false
        }
        if (!urlObj.hostname || urlObj.hostname.length < 3) {
            return false
        }
        return true
    } catch {
        return false
    }
}

const handleQRSubmit = async () => {
    if (!authStore.isAuthenticated) {
        router.push('/login')
        return
    }

    validationError.value = ''
    error.value = ''

    if (!qrUrl.value.trim()) {
        validationError.value = 'Please enter a URL'
        return
    }

    const validUrl = qrUrl.value.trim()

    if (!isValidUrl(validUrl)) {
        validationError.value = 'Please enter a valid URL (e.g., https://example.com)'
        return
    }

    isGenerating.value = true

    try {
        const response = await api.post('v1/qr/generate', {
            original_url: validUrl,
        })

        qrData.value = response.data

        const qrScanUrl = response.data.qr_url
        const dataUrl = await QRCode.toDataURL(qrScanUrl, {
            width: 400,
            margin: 2,
            color: {
                dark: '#000000',
                light: '#FFFFFF',
            },
            errorCorrectionLevel: 'M',
        })

        qrCodeDataUrl.value = dataUrl
        showSuccessDialog.value = true
        qrUrl.value = ''

        showToast('success', 'Success!', 'QR code generated and saved')
    } catch (err) {
        console.error('QR Generation Error:', err)

        if (err.response?.status === 401) {
            error.value = 'Session expired. Please login again.'
            setTimeout(() => router.push('/login'), 2000)
        } else if (err.response?.data?.error) {
            error.value = err.response.data.error
        } else {
            error.value = 'Failed to generate QR code. Please try again.'
        }
    } finally {
        isGenerating.value = false
    }
}

const downloadQRCode = () => {
    const link = document.createElement('a')
    link.href = qrCodeDataUrl.value
    link.download = `qrcode-${qrData.value?.id || 'download'}.png`
    link.click()

    showToast('success', 'Downloaded', 'QR code downloaded')
}

const copyQRToClipboard = async () => {
    try {
        const response = await fetch(qrCodeDataUrl.value)
        const blob = await response.blob()

        await navigator.clipboard.write([
            new ClipboardItem({
                [blob.type]: blob,
            }),
        ])

        showToast('success', 'Copied!', 'QR code copied to clipboard')
    } catch (err) {
        console.error('Failed to copy QR code:', err)
        showToast('error', 'Copy Failed', 'Could not copy QR code')
    }
}

const copyQRLink = async () => {
    if (!qrData.value?.qr_url) return

    try {
        await navigator.clipboard.writeText(qrData.value.qr_url)
        showToast('success', 'Copied!', 'QR link copied to clipboard')
    } catch (err) {
        showToast('error', 'Copy Failed', 'Could not copy link')
    }
}
</script>

<template>
    <div class="flex flex-col gap-4 md:gap-6">
        <h2 class="text-2xl text-gray-700 sm:text-3xl md:text-4xl font-semibold sm:text-left mb-3 sm:mb-7">
            Generate QR Code
        </h2>

        <div class="space-y-2">
            <p class="font-semibold text-gray-700">Enter QR Code destination</p>

            <Message v-if="validationError" severity="error" :closable="false">
                {{ validationError }}
            </Message>

            <Message v-if="error" severity="error" :closable="false">
                {{ error }}
            </Message>

            <form @submit.prevent="handleQRSubmit">
                <div class="flex flex-col sm:flex-row gap-3 sm:gap-2">
                    <InputText v-model="qrUrl" type="text" placeholder="Enter URL for QR code"
                        class="flex-1 text-sm sm:text-base" fluid />
                    <Button type="submit" label="Generate" :loading="isGenerating"
                        class="w-full sm:w-auto sm:min-w-35 py-3 sm:py-2" />
                </div>
            </form>
        </div>
    </div>

    <Dialog v-model:visible="showSuccessDialog" modal :dismissableMask="true"
        :style="{ width: '90vw', maxWidth: '550px' }" :pt="{
            root: { class: 'rounded-xl shadow-2xl' },
            header: { class: 'border-b border-gray-200 pb-4' },
            content: { class: 'pt-6' },
        }">
        <template #header>
            <div class="flex items-center gap-3">
                <div class="bg-green-100 text-green-600 rounded-full p-2">
                    <i class="pi pi-check text-xl"></i>
                </div>
                <div>
                    <h3 class="text-xl font-semibold text-gray-800">QR Code Generated!</h3>
                    <p class="text-sm text-gray-500 mt-1">Your QR code is ready to use</p>
                </div>
            </div>
        </template>

        <div v-if="qrData" class="flex flex-col gap-5">
            <!-- QR Code Display -->
            <div
                class="flex justify-center bg-gradient-to-br from-gray-50 to-gray-100 p-6 rounded-xl border border-gray-200 shadow-sm">
                <img :src="qrCodeDataUrl" alt="Generated QR Code" class="max-w-full h-auto rounded-lg shadow-md" />
            </div>

            <!-- QR Link Section -->
            <div
                class="border border-blue-200 bg-blue-50/50 p-4 rounded-xl space-y-2 hover:shadow-sm transition-shadow">
                <div class="flex justify-between items-start">
                    <div class="flex items-center gap-2">
                        <i class="pi pi-link text-blue-600"></i>
                        <span class="text-sm font-semibold text-gray-700">Scan URL</span>
                    </div>
                    <Button icon="pi pi-copy" @click="copyQRLink" size="small" severity="secondary" text rounded
                        v-tooltip.top="'Copy link'" />
                </div>
                <p class="text-xs text-gray-600 break-all font-mono bg-white p-2 rounded border border-blue-100">
                    {{ qrData.qr_url }}
                </p>
            </div>

            <!-- Destination Section -->
            <div class="border border-green-200 bg-green-50/50 p-4 rounded-xl space-y-2">
                <div class="flex items-center gap-2">
                    <i class="pi pi-external-link text-green-600"></i>
                    <span class="text-sm font-semibold text-gray-700">Destination</span>
                </div>
                <p class="text-xs text-gray-600 break-all font-mono bg-white p-2 rounded border border-green-100">
                    {{ qrData.original_url }}
                </p>
            </div>

            <!-- Action Buttons -->
            <div class="flex flex-col sm:flex-row gap-2 justify-end pt-2 border-t border-gray-200">
                <Button label="Close" @click="showSuccessDialog = false" severity="secondary" text
                    class="order-3 sm:order-1" />
                <Button label="Copy Image" @click="copyQRToClipboard" icon="pi pi-copy" severity="secondary" outlined
                    class="order-2" />
                <Button label="Download" @click="downloadQRCode" icon="pi pi-download" class="order-1 sm:order-3" />
            </div>
        </div>
    </Dialog>
</template>
