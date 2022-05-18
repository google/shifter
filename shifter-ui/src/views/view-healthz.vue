<script setup>
// Vue Component Imports
import ShifterV1StatusHealthz from "../components/component-shifter-status-healthz.vue";
</script>

<template>
  <div class="healthz">
    <ShifterV1StatusHealthz />
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .healthz {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
