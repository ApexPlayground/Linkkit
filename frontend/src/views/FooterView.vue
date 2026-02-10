<script setup>
import { ref } from 'vue'
import { RouterLink } from 'vue-router'
import Button from 'primevue/button'
import Divider from 'primevue/divider'

const currentYear = ref(new Date().getFullYear())

const productLinks = [
    { label: 'Features', path: '/features', icon: 'pi pi-star' },
    { label: 'Tools', path: '/tools', icon: 'pi pi-wrench' }
]

const companyLinks = [
    { label: 'About', path: '/about', icon: 'pi pi-info-circle' }
]

const legalLinks = [
    { label: 'Privacy Policy', path: '/privacy' },
    { label: 'Terms of Service', path: '/terms' }
]

const socialLinks = [
    {
        name: 'GitHub',
        icon: 'pi pi-github',
        url: 'https://github.com/ApexPlayground',
        severity: 'secondary'
    },
    {
        name: 'LinkedIn',
        icon: 'pi pi-linkedin',
        url: 'https://www.linkedin.com/in/divine-eboigbe-a63483205/',
        severity: 'info'
    }
]
</script>

<template>
    <footer class="bg-white border-t border-gray-200 mt-auto">
        <div class="max-w-7xl mx-auto px-6 py-12 lg:px-8">
            <!-- Main Footer Content -->
            <div class="grid grid-cols-1 md:grid-cols-3 gap-8 lg:gap-12">

                <!-- Brand Section -->
                <div class="md:col-span-2">
                    <RouterLink to="/" class="inline-block">
                        <h2 class="text-3xl font-bold text-gray-900 mb-4">LinkKit</h2>
                    </RouterLink>
                    <p class="text-gray-600 mb-6 text-sm leading-relaxed max-w-md">
                        Shorten, customize, and track your links with powerful analytics.
                        Build better connections with your audience.
                    </p>

                    <!-- Social Links -->
                    <div class="flex gap-2">
                        <Button v-for="social in socialLinks" :key="social.name" :icon="social.icon"
                            :severity="social.severity" text rounded :aria-label="social.name"
                            @click="window.open(social.url, '_blank')" class="!w-10 !h-10" />
                    </div>
                </div>

                <!-- Links Section -->
                <div class="grid grid-cols-2 gap-8">
                    <!-- Product Links -->
                    <div>
                        <h3 class="text-sm font-semibold text-gray-900 uppercase tracking-wider mb-4">
                            Product
                        </h3>
                        <ul class="space-y-3">
                            <li v-for="link in productLinks" :key="link.path">
                                <RouterLink :to="link.path"
                                    class="flex items-center gap-2 text-gray-600 hover:text-green-400 transition-colors duration-200 text-sm">
                                    <i :class="link.icon" class="text-xs"></i>
                                    {{ link.label }}
                                </RouterLink>
                            </li>
                        </ul>
                    </div>

                    <!-- Company Links -->
                    <div>
                        <h3 class="text-sm font-semibold text-gray-900 uppercase tracking-wider mb-4">
                            ME
                        </h3>
                        <ul class="space-y-3">
                            <li v-for="link in companyLinks" :key="link.path">
                                <RouterLink :to="link.path"
                                    class="flex items-center gap-2 text-gray-600 hover:text-green-400 transition-colors duration-200 text-sm">
                                    <i :class="link.icon" class="text-xs"></i>
                                    {{ link.label }}
                                </RouterLink>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>

            <Divider />

            <!-- Bottom Footer -->
            <div class="flex flex-col md:flex-row justify-between items-center gap-4">
                <!-- Copyright -->
                <p class="text-gray-500 text-sm">
                    © {{ currentYear }} <a class=" text-green-400 hover:text-green-400/80"
                        href="https://divineeboigbe.netlify.app/">Apexplayground</a>. All rights
                    reserved.
                </p>

                <!-- Legal Links -->
                <div class="flex flex-wrap justify-center gap-6">
                    <RouterLink v-for="link in legalLinks" :key="link.path" :to="link.path"
                        class="text-gray-500 hover:text-green-400 text-sm transition-colors duration-200">
                        {{ link.label }}
                    </RouterLink>
                </div>
            </div>
        </div>
    </footer>
</template>