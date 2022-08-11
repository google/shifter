<!--
 Copyright 2022 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

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
