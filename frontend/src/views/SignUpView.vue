<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router'
import InputText from 'primevue/inputtext';
import Password from 'primevue/password';
import Button from 'primevue/button';
import Message from 'primevue/message';
import InlineMessage from 'primevue/inlinemessage';

import api from '@/services/api';

const router = useRouter()

const formData = ref({
    name: '',
    email: '',
    password: ''
});

const loading = ref(false)
const error = ref(null)

const handleSubmit = async () => {
    loading.value = true
    error.value = null

    try {
        const { data } = await api.post('v1/users/signup', formData.value)

        localStorage.setItem('token', data.token)
        router.push('/dashboard')
    } catch (err) {
        error.value =
            err.response?.data?.message || 'Something went wrong'
    } finally {
        loading.value = false
    }
}

const handleGoogleSignup = () => {
    console.log('Google signup clicked');
};
</script>

<template>
    <div class="flex items-center justify-center p-4 lg:p-8 mt-16 sm:mt-24">
        <div class="bg-white rounded-3xl shadow-2xl w-full max-w-6xl overflow-hidden grid lg:grid-cols-2">

            <!-- Right Side - Form -->
            <div class="p-12 lg:p-16 flex flex-col justify-center">
                <div class="max-w-md w-full mx-auto">
                    <h2 class="text-4xl lg:text-5xl font-bold mb-3">Create Account</h2>
                    <p class="mb-8 text-lg text-gray-500">Sign up to get started</p>

                    <!-- Error Message -->
                    <Message v-if="error" severity="error" :closable="true" @close="error = null" class="mb-6">
                        {{ error }}
                    </Message>

                    <form @submit.prevent="handleSubmit" class="space-y-6">
                        <!-- Name Input -->
                        <div>
                            <label for="name" class="block text-base font-medium mb-2">
                                <i class="pi pi-user mr-2"></i>Name
                            </label>
                            <InputText id="name" v-model="formData.name" placeholder="Enter your name"
                                class="w-full text-lg" size="large" :disabled="loading" required />
                        </div>

                        <!-- Email Input -->
                        <div>
                            <label for="email" class="block text-base font-medium mb-2">
                                <i class="pi pi-envelope mr-2"></i>Email
                            </label>
                            <InputText id="email" v-model="formData.email" type="email" placeholder="Enter your email"
                                class="w-full text-lg" size="large" :disabled="loading" required />
                        </div>

                        <!-- Password Input -->
                        <div>
                            <label for="password" class="block text-base font-medium mb-2">
                                <i class="pi pi-lock mr-2"></i>Password
                            </label>
                            <Password id="password" v-model="formData.password" placeholder="Enter your password"
                                toggleMask fluid class="w-full" inputClass="text-lg" :disabled="loading" required>
                                <template #header>
                                    <h6>Pick a password</h6>
                                </template>
                                <template #footer>
                                    <p class="mt-2">Suggestions:</p>
                                    <ul class="pl-2 ml-2 mt-0" style="line-height: 1.5">
                                        <li>At least one lowercase</li>
                                        <li>At least one uppercase</li>
                                        <li>At least one numeric</li>
                                        <li>Minimum 8 characters</li>
                                    </ul>
                                </template>
                            </Password>
                        </div>

                        <!-- Sign Up Button -->
                        <Button type="submit" label="Sign Up" class="w-full text-lg py-4 mt-2" severity="primary"
                            size="large" :loading="loading" :disabled="loading">
                            <template #icon>
                                <i class="pi pi-user-plus"></i>
                            </template>
                        </Button>

                        <!-- Divider -->
                        <div class="relative my-8">
                            <div class="absolute inset-0 flex items-center">
                                <div class="w-full border-t border-gray-300"></div>
                            </div>
                            <div class="relative flex justify-center text-base">
                                <span class="px-4 bg-white text-gray-500">Or continue with</span>
                            </div>
                        </div>

                        <!-- Google Sign Up Button -->
                        <Button type="button" @click="handleGoogleSignup" severity="secondary" outlined
                            class="w-full text-lg py-4" size="large" :disabled="loading">
                            <svg class="w-6 h-6 mr-3" viewBox="0 0 24 24">
                                <path fill="#4285F4"
                                    d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" />
                                <path fill="#34A853"
                                    d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" />
                                <path fill="#FBBC05"
                                    d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" />
                                <path fill="#EA4335"
                                    d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" />
                            </svg>
                            Sign up with Google
                        </Button>
                    </form>

                    <p class="text-center text-base mt-8">
                        Already have an account?
                        <RouterLink to="/login" class="text-green-400 hover:text-green-400/90 font-medium ml-1">Sign in
                        </RouterLink>
                    </p>
                </div>
            </div>

            <!-- Left Side - Illustration -->
            <div
                class="hidden lg:flex bg-gradient-to-br from-green-400 to-emerald-600 p-12 flex-col justify-center items-center relative">
                <div class="text-white z-10 text-center">
                    <h1 class="text-5xl font-bold mb-6">Hey There!</h1>
                    <p class="text-xl text-white mb-8">Join our community and unlock amazing features</p>
                    <div class="w-full max-w-md">
                        <img src="/reg.svg" alt="Sign up illustration" class="w-full h-auto" />
                    </div>
                </div>

                <!-- Decorative elements -->
                <div class="absolute top-10 left-10 w-20 h-20 border-4 border-white/30 rounded-full"></div>
                <div class="absolute bottom-10 right-10 w-32 h-32 border-4 border-white/30 rounded-full"></div>
            </div>
        </div>
    </div>
</template>