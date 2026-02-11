<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useThemeStore } from '@/stores/theme'
import MobileHeader from './MobileHeader.vue'
import SidebarHeader from './SidebarHeader.vue'
import SidebarNavigation from './SidebarNavigation.vue'
import SidebarFooter from './SidebarFooter.vue'

const sidebarVisible = ref(false)
const router = useRouter()
const authStore = useAuthStore()
const themeStore = useThemeStore()

const isAuthenticated = computed(() => authStore.isAuthenticated)



const menuItems = [
    {
        label: 'Home',
        icon: 'pi pi-home',
        route: '/home',
        description: 'Shorten link or generate QR code'
    },
    {
        label: 'Links',
        icon: 'pi pi-link',
        route: '/home/links',
        description: 'Manage your links'
    },
    {
        label: 'QR Codes',
        icon: 'pi pi-qrcode',
        route: '/home/qr-codes',
        description: 'Manage QR codes'
    },
    {
        label: 'Analytics',
        icon: 'pi pi-chart-line',
        route: '/home/analytics',
        description: 'View statistics'
    },

]

const navigateTo = (route) => {
    router.push(route)
    sidebarVisible.value = false
}

const handleLogout = async () => {
    sidebarVisible.value = false
    authStore.logout()
    router.replace('/login')
}


const toggleSidebar = () => {
    sidebarVisible.value = !sidebarVisible.value
}

const closeSidebar = () => {
    sidebarVisible.value = false
}
</script>

<template>
    <div class="min-h-screen bg-gray-50">
        <!-- Mobile Header -->
        <template v-if="isAuthenticated">
            <MobileHeader :user="authStore.user" @toggle-sidebar="toggleSidebar" @navigate="navigateTo"
                @logout="handleLogout" />

            <!-- Overlay for mobile -->
            <div v-if="sidebarVisible" class="lg:hidden fixed inset-0 bg-black/50 z-40" @click="closeSidebar"></div>

            <!-- Sidebar -->
            <aside :class="[
                'fixed top-0 left-0 h-full bg-white border-r border-gray-200 z-50 transition-transform duration-300 ease-in-out',
                'w-80 flex flex-col',
                sidebarVisible ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'
            ]">
                <SidebarHeader :user="authStore.user" :show-close-button="true" @close="closeSidebar" />
                <SidebarNavigation :items="menuItems" @navigate="navigateTo" />
                <SidebarFooter v-model:is-dark="themeStore.isDark" @navigate="navigateTo" @logout="handleLogout" />
            </aside>
        </template>

        <!-- Main Content -->
        <main :class="[
            'transition-all duration-300',
            'pt-16 lg:pt-0 lg:ml-72 min-h-screen'
        ]">
            <div class="p-6">
                <router-view />
            </div>
        </main>
    </div>
</template>