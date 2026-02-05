<script setup>

</script>

<template>
    <!-- Use Cases Hero -->
    <div class="mx-auto w-[95%] sm:w-[90%] md:w-[50%] mt-16 sm:mt-24">
        <div v-animateonscroll="{ enterClass: 'animate-enter fade-in-10 spin-in-[-45deg] slide-in-from-t-16 animate-duration-1000' }"
            class="text-center mb-12">
            <h2 class="text-3xl sm:text-4xl md:text-5xl font-bold mb-4">
                Perfect For Every Use Case
            </h2>
            <p class="text-xl">
                From social media to marketing campaigns, we've got you covered
            </p>
        </div>

        <div class="grid md:grid-cols-2 gap-8 mx-auto max-w-7xl ">
            <!-- Use Case 1 -->
            <div v-animateonscroll="{ enterClass: 'animate-enter fade-in-10 slide-in-from-b-20 animate-duration-1000' }"
                class="flex gap-4">
                <div class="shrink-0">
                    <div
                        class="w-12 h-12 bg-blue-100 dark:bg-blue-900/30 rounded-full flex items-center justify-center">
                        <i class="pi pi-users text-xl text-blue-500"></i>
                    </div>
                </div>
                <div>
                    <h3 class="text-xl font-bold mb-2">Social Media Managers</h3>
                    <p class="text-lg">
                        Share clean, trackable links across Instagram, Twitter, and LinkedIn
                    </p>
                </div>
            </div>

            <!-- Use Case 2 -->
            <div v-animateonscroll="{ enterClass: 'animate-enter fade-in-10 slide-in-from-b-20 animate-duration-1000' }"
                class="flex gap-4">
                <div class="shrink-0">
                    <div
                        class="w-12 h-12 bg-purple-100 dark:bg-purple-900/30 rounded-full flex items-center justify-center">
                        <i class="pi pi-megaphone text-xl text-purple-500"></i>
                    </div>
                </div>
                <div>
                    <h3 class="text-xl font-bold mb-2">Marketing Teams</h3>
                    <p class="text-lg">
                        Create branded short links for campaigns and track performance
                    </p>
                </div>
            </div>

            <!-- Use Case 3 -->
            <div v-animateonscroll="{ enterClass: 'animate-enter fade-in-10 slide-in-from-b-20 animate-duration-1000' }"
                class="flex gap-4">
                <div class="shrink-0">
                    <div
                        class="w-12 h-12 bg-green-100 dark:bg-green-900/30 rounded-full flex items-center justify-center">
                        <i class="pi pi-shopping-cart text-xl text-green-500"></i>
                    </div>
                </div>
                <div>
                    <h3 class="text-xl font-bold mb-2">E-commerce Stores</h3>
                    <p class="text-lg">
                        Generate QR codes for products and promotional materials
                    </p>
                </div>
            </div>

            <!-- Use Case 4 -->
            <div v-animateonscroll="{ enterClass: 'animate-enter fade-in-10 slide-in-from-b-20 animate-duration-1000' }"
                class="flex gap-4">
                <div class="shrink-0">
                    <div
                        class="w-12 h-12 bg-orange-100 dark:bg-orange-900/30 rounded-full flex items-center justify-center">
                        <i class="pi pi-briefcase text-xl text-orange-500"></i>
                    </div>
                </div>
                <div>
                    <h3 class="text-xl font-bold mb-2">Content Creators</h3>
                    <p class="text-lg">
                        Simplify link sharing in your videos, podcasts, and articles
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>