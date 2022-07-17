<script setup>
import NavBar from "./components/component-navbar.vue";
import Loading from "./components/component-loading.vue";
</script>

<template>
  <Loading class="" v-show="isLoading" />
  <div class="flex flex-col h-screen">
    <div class="hidden sm:block">
      <NavBar />
    </div>
    <div class="flex flex-grow"><RouterView /></div>
  </div>
</template>

<script>
// Pinia Store Imports
import { useConfigurationsLoading } from "../src/stores/configurations/loading";
// Plugin & Package Imports
import { mapState } from "pinia";

export default {
  computed: {
    ...mapState(useConfigurationsLoading, {
      isLoading: "isLoading",
    }),
  },
};
</script>

<style>
@import "@/assets/base.css";
</style>
