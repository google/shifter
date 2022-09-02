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
// Vue Component Imports
import OpenshiftNamespaceListItem from "./openshift-namespace-list-item.vue";
</script>
<template>
  <div class="container">
    <div class="flex items-center mx-6 mb-2">
      <div class="text-xl font-bold text-shifter-white-soft">Cluster Namespaces</div>
      <div class="text-sm pl-2 text-shifter-white-soft">({{ itemCount }})</div>
    </div>
    <div class="flex flex-col">
      <OpenshiftNamespaceListItem
        :v-if="itemCount > 0"
        v-for="namespace in all"
        :key="namespace.metadata.uid"
        :uid="namespace.metadata.uid"
      />
    </div>
  </div>
</template>

<script>
// Pinia Store Imports
import { useOSProjects } from "../stores/openshift/projects";
// Plugin & Package Imports
import { mapState } from "pinia";
export default {
  data() {
    return {};
  },

  methods: {},

  computed: {
    ...mapState(useOSProjects, { all: "all" }),
    itemCount() {
      if (this.all !== undefined && this.all.length >= 0) {
        return this.all.length;
      }
      return 0;
    },
  },
};
</script>
