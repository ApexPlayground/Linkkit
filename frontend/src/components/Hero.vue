<script setup>
import { ref } from 'vue'
import Tabs from 'primevue/tabs'
import TabList from 'primevue/tablist'
import Tab from 'primevue/tab'
import TabPanels from 'primevue/tabpanels'
import TabPanel from 'primevue/tabpanel'

import InputText from 'primevue/inputtext'
import Button from 'primevue/button'

const items = ref([
    { value: 'link', label: 'Link', icon: 'pi pi-link' },
    { value: 'qr', label: 'QR Code', icon: 'pi pi-qrcode' },
])

const value = ref('')
</script>
<template>
    <div class="mx-auto w-[95%] sm:w-[90%] md:w-[70%]  mt-16 sm:mt-24">
        <!-- Header -->
        <div class="text-center mb-12 md:mb-16">
            <h1 class="text-3xl sm:text-4xl md:text-6xl font-bold mb-4 md:mb-6">
                URL Management Made Simple
            </h1>
            <p class="text-base sm:text-lg md:text-xl max-w-2xl mx-auto">
                Shorten, store, update, and track your links, generate QR codes, and manage everything from
                one simple, powerful URL toolkit designed with you in mind.
            </p>
        </div>

        <!-- Tabs Card -->
        <div class="card p-4 sm:p-6 md:p-8 border border-gray-200 rounded-2xl shadow-lg ">
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
                        <div class="flex flex-col gap-4 md:gap-6">
                            <!-- Title -->
                            <h2 class="text-2xl sm:text-3xl md:text-4xl font-semibold sm:text-left mb-3 sm:mb-7">
                                Shorten a long link
                            </h2>

                            <div class="space-y-2">
                                <p class="font-semibold">Paste your long link here</p>
                                <div class="flex flex-col sm:flex-row gap-3 sm:gap-2">
                                    <InputText type="text" placeholder="Enter your long URL"
                                        class="flex-1 text-sm sm:text-base" fluid />
                                    <Button label="Shorten" class="w-full sm:w-auto sm:min-w-35 py-3 sm:py-2" />
                                </div>
                            </div>
                        </div>
                    </TabPanel>

                    <!-- QR Tab -->
                    <TabPanel value="qr">
                        <div class="flex flex-col gap-4 md:gap-6">
                            <!-- Title -->
                            <h2 class="text-2xl sm:text-3xl md:text-4xl font-semibold sm:text-left mb-3 sm:mb-7">
                                Generate QR Code
                            </h2>

                            <div class="space-y-2">
                                <p class="font-semibold">Enter QR Code destination</p>
                                <div class="flex flex-col sm:flex-row gap-3 sm:gap-2">
                                    <InputText type="text" placeholder="Enter URL for QR code"
                                        class="flex-1 text-sm sm:text-base" fluid />
                                    <Button label="Generate" class="w-full sm:w-auto sm:min-w-35 py-3 sm:py-2" />
                                </div>
                            </div>
                        </div>
                    </TabPanel>
                </TabPanels>
            </Tabs>
        </div>
        <div class="text-center px-4 py-12">
            <h2 class="text-3xl font-bold mb-3">
                Sign up for free and simplify your analysis
            </h2>

            <div class="flex flex-col sm:flex-row justify-center gap-4 sm:gap-6 mb-8">
                <span class="flex items-center justify-center sm:justify-start gap-2">
                    <i class="pi pi-check text-green-500"></i>
                    <span class="text-sm sm:text-base">Free URL Shortening</span>
                </span>
                <span class="flex items-center justify-center sm:justify-start gap-2">
                    <i class="pi pi-check text-green-500"></i>
                    <span class="text-sm sm:text-base">Free URL Analysis</span>
                </span>
                <span class="flex items-center justify-center sm:justify-start gap-2">
                    <i class="pi pi-check text-green-500"></i>
                    <span class="text-sm sm:text-base">Free QR Code Generation</span>
                </span>
            </div>


        </div>
    </div>

</template>
