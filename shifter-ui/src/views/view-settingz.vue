<script setup>
// Vue Component Imports
import ShifterV1StatusSettingz from "../components/component-shifter-status-settingz.vue";
</script>

<template>
  <div class="settingz">
    <ShifterV1StatusSettingz />
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .settingz {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
