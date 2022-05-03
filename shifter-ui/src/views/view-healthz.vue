<script setup>
<<<<<<<< HEAD:shifter-ui/src/views/view-healthz.vue
// Vue Component Imports
import ShifterV1StatusHealthz from "../components/component-shifter-status-healthz.vue";
========
import ShifterV1StatusHealthz from '../../components/Shifter/v1/Status/Healthz.vue';
>>>>>>>> a550bec (Added NavBar and Routing as Well as some Simple Form Concepts):shifter-ui/src/views/Status/HealthzView.vue
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
