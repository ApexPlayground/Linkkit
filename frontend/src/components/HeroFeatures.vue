<script setup></script>

<template>
    <!-- Features Grid Hero -->
    <div class="">
        <div class="mx-auto w-[95%] sm:w-[90%] md:w-[70%] pt-16 sm:pt-24">
            <!-- Header with animation -->
            <div v-animateonscroll="{
                enterClass: 'animate-enter fade-in-10 slide-in-from-b-20 animate-duration-1000',
            }" class="text-center mb-12 animation-duration-1000">
                <h2 class="text-3xl sm:text-4xl md:text-5xl font-bold mb-4">
                    Everything You Need in One Place
                </h2>
                <p class="text-xl">Powerful tools to manage, track, and optimize all your links</p>
            </div>

            <div class="grid sm:grid-cols-2 lg:grid-cols-3 gap-6">
                <!-- Card 1 -->
                <div v-animateonscroll="{
                    enterClass: 'animate-enter fade-in-10 slide-in-from-t-20 animate-duration-1000',
                }"
                    class="group relative p-6 rounded-xl bg-white shadow-lg animation-duration-1000 animation-ease-in-out overflow-hidden">
                    <div
                        class="w-12 h-12 bg-green-100 dark:bg-green-900/30 rounded-lg flex items-center justify-center mb-4">
                        <i class="pi pi-bolt text-2xl text-green-400"></i>
                    </div>
                    <h3 class="text-xl font-bold mb-2">Lightning Fast</h3>
                    <p class="text-lg">Shorten URLs in milliseconds. No waiting, no hassle.</p>
                </div>

                <!-- Card 2 -->
                <div v-animateonscroll="{
                    enterClass: 'animate-enter fade-in-10 slide-in-from-b-20 animate-duration-1000',
                }"
                    class="group relative p-6 rounded-xl bg-white shadow-lg animation-duration-1000 animation-ease-in-out overflow-hidden">
                    <div
                        class="w-12 h-12 bg-green-100 dark:bg-green-900/30 rounded-lg flex items-center justify-center mb-4">
                        <i class="pi pi-pencil text-2xl text-green-400"></i>
                    </div>
                    <h3 class="text-xl font-bold mb-2">Custom Short Links</h3>
                    <p class="text-lg">Create memorable branded links that match your style.</p>
                </div>

                <!-- Card 3 -->
                <div v-animateonscroll="{
                    enterClass: 'animate-enter fade-in-10 slide-in-from-t-20 animate-duration-1000',
                }"
                    class="group relative p-6 rounded-xl bg-white shadow-lg animation-duration-1000 animation-ease-in-out overflow-hidden">
                    <div
                        class="w-12 h-12 bg-green-100 dark:bg-green-900/30 rounded-lg flex items-center justify-center mb-4">
                        <i class="pi pi-qrcode text-2xl text-green-400"></i>
                    </div>
                    <h3 class="text-xl font-bold mb-2">QR Code Generator</h3>
                    <p class="text-lg">Generate scannable QR codes for any link instantly.</p>
                </div>

                <!-- Card 4 -->
                <div v-animateonscroll="{
                    enterClass: 'animate-enter fade-in-10 zoom-in-50 animate-duration-1000',
                }"
                    class="group relative p-6 rounded-xl bg-white shadow-lg animation-duration-1000 animation-ease-in-out overflow-hidden">
                    <div
                        class="w-12 h-12 bg-green-100 dark:bg-green-900/30 rounded-lg flex items-center justify-center mb-4">
                        <i class="pi pi-shield text-2xl text-green-400"></i>
                    </div>
                    <h3 class="text-xl font-bold mb-2">Secure & Private</h3>
                    <p class="text-lg">Your links are protected with enterprise-grade security.</p>
                </div>

                <!-- Card 5 -->
                <div v-animateonscroll="{
                    enterClass: 'animate-enter fade-in-10 zoom-in-75 animate-duration-1000',
                }"
                    class="group relative p-6 rounded-xl bg-white shadow-lg animation-duration-1000 animation-ease-in-out overflow-hidden">
                    <div
                        class="w-12 h-12 bg-green-100 dark:bg-green-900/30 rounded-lg flex items-center justify-center mb-4">
                        <i class="pi pi-sync text-2xl text-green-400"></i>
                    </div>
                    <h3 class="text-xl font-bold mb-2">Update Anytime</h3>
                    <p class="text-lg">Change destination URLs without breaking your short links.</p>
                </div>

                <!-- Card 6 -->
                <div v-animateonscroll="{
                    enterClass: 'animate-enter fade-in-10 zoom-in-50 animate-duration-1000',
                }"
                    class="group relative p-6 rounded-xl bg-white shadow-lg animation-duration-1000 animation-ease-in-out overflow-hidden">
                    <div
                        class="w-12 h-12 bg-green-100 dark:bg-green-900/30 rounded-lg flex items-center justify-center mb-4">
                        <i class="pi pi-chart-line text-2xl text-green-400"></i>
                    </div>
                    <h3 class="text-xl font-bold mb-2">Track Performance</h3>
                    <p class="text-lg">Monitor clicks and engagement across all your links.</p>
                </div>
            </div>
        </div>
    </div>
</template>
