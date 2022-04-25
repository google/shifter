<script setup>
import ShifterV1StatusHealthz from '../../components/shifter/v1/status/healthz.vue';
</script>

<template>
  <div class="healthz">
    <ShifterV1StatusHealthz/>
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


<script>
import { shifterConfig } from "@/main"
export default {
  
  created() {
  }
};
</script>