<script setup>
// Vue Component Imports
import ListConvertObjectsItem from "./list-convert-namespace-objects-item.vue";
</script>
<template>
  <div class="container flex-row mx-auto justify-center mx-4">
    <p class="text-lg font-bold">OpenShift Namespaces ({{ itemCount }})</p>
    <ListConvertObjectsItem
      class="flex-row mx-auto justify-left border border-shifter-black m-3"
      v-for="namespace in all"
      :key="namespace.metadata.uid"
      :uid="namespace.metadata.uid"
    ></ListConvertObjectsItem>
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
