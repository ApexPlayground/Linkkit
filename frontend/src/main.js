import { createApp } from 'vue'
import PrimeVue from 'primevue/config'
import Aura from '@primeuix/themes/aura'
import { createPinia } from 'pinia'
import 'primeicons/primeicons.css'
import App from './App.vue'
import router from './router'
import './assets/main.css'

const app = createApp(App)
app.use(PrimeVue, {
  theme: {
    preset: Aura,
    options: {
      cssLayer: {
        name: 'primevue',
        order: 'theme, base, primevue',
      },
      darkModeSelector: '.dark',
    },
  },
})
app.use(createPinia())
app.use(router)
app.mount('#app')
