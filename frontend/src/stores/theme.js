import { defineStore } from 'pinia'

export const useThemeStore = defineStore('theme', {
  state: () => ({
    isDark: false,
  }),
  actions: {
    initializeTheme() {
      const savedMode = localStorage.getItem('darkMode')
      this.isDark = savedMode
        ? savedMode === 'true'
        : window.matchMedia('(prefers-color-scheme: dark)').matches

      this.applyTheme()

      // Use Pinia's $subscribe to watch for any state changes
      this.$subscribe(() => {
        this.applyTheme()
      })
    },
    applyTheme() {
      if (this.isDark) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
      localStorage.setItem('darkMode', this.isDark.toString())
    },
  },
})
