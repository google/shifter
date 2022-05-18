<script setup>
// Vue Component Imports
import Title from "../components/form-title.vue";
import TitleSection from "../components/form-title-section.vue";
import FormClusterManagement from "../components/form-cluster-management.vue";
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
    <FormClusterManagement class="bg-dummy-green" />
  </div>
</template>
