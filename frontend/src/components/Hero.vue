<script setup>
import UsageTabs from './UsageTabs.vue'
</script>
<template>
    <div class="mx-auto w-[95%] sm:w-[90%] md:w-[50%] mt-16 sm:mt-24">
        <!-- Header -->
        <div v-animateonscroll="{
            enterClass: 'animate-enter fade-in-10 slide-in-from-t-20 animate-duration-1000',
        }">
            <div class="text-center mb-12 md:mb-16">
                <h1 class="text-3xl sm:text-4xl md:text-6xl font-bold mb-4 md:mb-6">
                    URL Management Made Simple
                </h1>
                <p class="text-base sm:text-lg md:text-xl mx-auto">
                    Shorten, store, update, and track your links, generate QR codes, and manage everything
                    from one simple, powerful URL toolkit designed with you in mind .
                </p>
            </div>
        </div>

        <div v-animateonscroll="{
            enterClass: 'animate-enter fade-in-10 slide-in-from-b-20 animate-duration-1000',
        }">
            <UsageTabs />

            <div class="text-center px-4 pt-12">
                <h2 class="text-3xl font-bold mb-3">Sign up for free and simplify your analysis</h2>

                <div class="flex flex-col sm:flex-row justify-center gap-4 sm:gap-6">
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
    </div>
</template>
