<script setup>
import Title from "../components/form-title.vue";
import TitleSection from "../components/form-title-section.vue";
import FormClusterManagement from "../components/form-cluster-management.vue";
</script>

<template>
  <div>
    <Title>
      <!--<template #icon>
            <DocumentationIcon />
          </template>-->
      <template #title>Manage OpenShift Cluster Connections</template>
      <template #subtitle
        >Manage your OpenShift cluster connection details</template
      >
    </Title>
    <TitleSection>
      <!--<template #icon>
            <DocumentationIcon />
          </template>-->
      <template #title>Cluster Connections</template>
      <template #subtitle>Create, Edit or Delete Cluster Connections</template>
    </TitleSection>
    <FormClusterManagement class="bg-dummy-green" />
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
  methods: {},
};
</script>

<style scoped></style>
