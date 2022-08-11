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

<template>
  <div
    class="z-50 fixed top-30 bg-shifter-red-soft flex items-center justify-center h-full w-screen"
    :class="showModal ? 'visible' : 'invisible'"
  >
    <div
      class="container flex absolute items-center w-1/3 h-1/4 bg-shifter-black-soft border rounded-2xl overflow-y-auto"
    >
      <div class="container flex flex-col mx-auto w-full gap-5">
        <div class="container flex flex-row mx-auto w-full justify-center">
          <div class="flex flex-row text-xl text-center font-bold">
            Delete Selected Openshift Cluster Configuration
          </div>
        </div>
        <div
          class="container flex flex-row mx-auto w-full text-xl gap-5 justify-center"
        >
          <div class="flex flex-col w-1/4">
            <a
              @click="deleteClusterConfig"
              class="rounded bg-shifter-red-soft rounded border border-shifter-red-soft py-2 text-center font-bold text-sm"
              >Yes</a
            >
          </div>
          <div class="flex flex-col w-1/4">
            <a
              @click="closeModal"
              class="rounded bg-shifter-red-soft rounded border border-shifter-red-soft py-2 text-center font-bold text-sm"
              >Cancel</a
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// Notifications Imports
import {
  shifterConfigurationUpdateSuccess,
  shifterConfigurationUpdateError,
} from "@/notifications";
// Pinia Store Imports
import { useModalClusterDelete } from "../stores/configurations/modalClusterDelete";
import { useConfigurationsClusters } from "../stores/configurations/clusters";
// Plugin & Package Imports
import { mapState, mapActions } from "pinia";

export default {
  methods: {
    ...mapActions(useConfigurationsClusters, {
      deleteCluster: "deleteCluster",
    }),
    ...mapActions(useModalClusterDelete, { closeModal: "closeModal" }),
    deleteClusterConfig() {
      this.deleteCluster(this.getClusterId)
        .then(() => {
          shifterConfigurationUpdateSuccess(
            "Openshift Cluster Configuration Deleted"
          );
          this.closeModal();
        })
        .catch(() => {
          shifterConfigurationUpdateError(
            "Unable to Delete Openshift Cluster Configuration"
          );
          this.closeModal();
        });
    },
  },
  computed: {
    ...mapState(useModalClusterDelete, {
      showModal: "showModal",
    }),
    ...mapState(useModalClusterDelete, {
      getClusterId: "getClusterId",
    }),
  },
};
</script>

<style scoped></style>
