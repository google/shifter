<script setup>
import HomeTitle from "../components/Home/HomeTitle.vue";
import HomeDetails from "../components/Home/HomeDetails.vue"
</script>

<template>
  <header>
    <img
      alt="Shifter"
      class="logo"
      src="../assets/logo-shifter.svg"
      width="125"
      height="125"
    />

    <div class="wrapper">
      <HomeTitle msg="Shifter" />
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
        <a target="_blank" href="https://github.com/google/shifter">Source</a>
      </nav>
    </div>
  </header>
  <HomeDetails />
</template>


<script>
import { shifterConfig } from "../main"
export default {
  
  created() {
  }
};
</script>