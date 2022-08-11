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
