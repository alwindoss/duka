<script setup>
import { RouterView } from 'vue-router'
import DukaNavbar from './components/DukaNavbar.vue';
</script>

<template>
  <div class="container-fluid">
    <header>
      <div class="container-fluid">
        <DukaNavbar app-name="Duka Store" />
      </div>
    </header>
    <main>
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
</style>
