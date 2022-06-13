<script setup>
// Vue Component Imports
import Title from "../components/form-title.vue";
import TitleSection from "../components/form-title-section.vue";
import ConfigurationOpenshiftClusterList from "../components/configuration-openshift-cluster-list.vue";
</script>

<template>
  <div>
    <Title>
      <template #title>Manage OpenShift Cluster Connections</template>
      <template #subtitle
        >Manage your OpenShift cluster connection details</template
      >
    </Title>
    <TitleSection>
      <template #title>Cluster Connections</template>
      <template #subtitle>Create, Edit or Delete Cluster Connections</template>
    </TitleSection>
    <ConfigurationOpenshiftClusterList />
    <a
      @click="openModal"
      class="rounded bg-shifter-red-soft rounded border border-shifter-red-soft px-2 py-2 text-center font-bold text-sm"
      >Add New Cluster</a
    >
  </div>
</template>

<script>
// Pinia Store Imports
import { useModalClusterAdd } from "../stores/configurations/modalClusterAdd";
// Plugin & Package Imports
import { mapActions } from "pinia";
export default {
  data() {
    return {};
  },
  methods: {
    ...mapActions(useModalClusterAdd, { openModal: "openModal" }),
    AddConfig() {
      this.openModal();
    },
  },
};
</script>
