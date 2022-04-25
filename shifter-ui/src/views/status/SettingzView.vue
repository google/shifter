<script setup>
import ShifterV1StatusSettingz from '../../components/shifter/v1/status/settingz.vue';
</script>

<template>
  <div class="settingz">
    <ShifterV1StatusSettingz/>
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


<script>
import { shifterConfig } from "@/main"
export default {
  
  created() {
  }
};
</script>