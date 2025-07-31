<template>
  <NuxtLayout>
    <NuxtPage/>
  </NuxtLayout>
</template>

<script setup lang="ts">
import {AppConfig, useAppConfig} from "~/lib/app_config"
const config: AppConfig = useAppConfig()

useHead({
  title: `v${config.version}`,
})
</script>