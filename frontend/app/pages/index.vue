<template>
  <div class="query-layout">
    <div class="query-a">
      <textarea
          v-model="leftSql"
          placeholder="Запрос с PROD"
          class="query-textarea"
      />
    </div>
    <div class="query-b">
      <textarea
          v-model="rightSql"
          placeholder="Запрос из MR"
          class="query-textarea"
      />
    </div>

    <div class="query-diff">
      <div class="query-diff-html" v-html="diff"/>
    </div>
  </div>
</template>

<script setup lang="ts">
import {watch, ref} from 'vue'
import {AppConfig, useAppConfig} from "~/lib/app_config"
const config: AppConfig = useAppConfig()

const leftSql = ref('')
const rightSql = ref('')
const diff = ref('')

watch([leftSql, rightSql], ([newA, newB], [prevA, prevB]) => {
    load();
});

async function load() {
  await $fetch(`${config.apiPath}/api/select`, {
    method: 'POST',
    body: {
      left: leftSql.value,
      right: rightSql.value,
    }
  }).then(res => diff.value = res)
}
</script>
