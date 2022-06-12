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
  </div>
</template>
