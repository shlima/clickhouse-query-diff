// https://nuxt.com/docs/api/configuration/nuxtn-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: false },
  css: ['../assets/scss/main.scss'],
  ssr: true, // SSR включен, но мы будем генерировать HTML во время сборки
  nitro: {
    preset: 'static'
  },
  runtimeConfig: {
    public: {
      apiPath: 'http://127.0.0.1:3333', // can be overridden by NUXT_PUBLIC_API_PATH environment variable
      version: '0.0.0'
    }
  },
})
