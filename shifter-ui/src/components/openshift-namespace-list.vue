<script setup>
// Vue Component Imports
import OpenshiftNamespaceListItem from "./openshift-namespace-list-item.vue";
</script>
<template>
  <div class="container">
    <div class="flex items-center mx-6 mb-2">
      <div class="text-xl font-bold">OpenShift Namespaces</div>
      <div class="text-sm pl-2">({{ itemCount }})</div>
    </div>
    <div class="flex flex-col">
      <OpenshiftNamespaceListItem
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
