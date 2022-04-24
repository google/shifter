<script setup>

import HelloWorld from "@/components/HelloWorld.vue";
import HomeDetailsView from "./HomeDetailsView.vue";
</script>

<template>
  <header>
    <!--<img
      alt="Shifter"
      class="logo"
      src="@/assets/logo.svg"
      width="125"
      height="125"
    />-->

    <div class="wrapper">
      <HelloWorld msg="Shifter" />
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
        <a target="_blank" href="https://github.com/google/shifter">Source</a>
      </nav>
    </div>
  </header>
  <HomeDetailsView />
</template>


<script>
import { shifterConfig } from "@/main"
export default {
  
  created() {
    console.log("---- Hello Home View -----")
    console.log(shifterConfig)
    console.log("---- ================== -----")
  }
};
</script>