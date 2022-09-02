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
  <div class="z-50 fixed top-30 backdrop-blur-sm flex items-center justify-center h-full w-screen" :class="showModal ? 'visible' : 'invisible'" >
    <div class="container flex absolute items-center w-1/2 h-1/2 bg-shifter-black-soft border rounded-2xl overflow-y-auto">
      <div class="container flex flex-col mx-auto w-full gap-5">
        <div class="container flex flex-row mx-auto w-full justify-center">
          <div class="flex flex-row text-xl text-center font-bold text-shifter-white-soft" >
            Create Cluster Configuration
          </div>
        </div>
        <div class="container flex flex-col mx-auto w-full text-xl gap-5 justify-center">
          <div class="flex flex-col overflow-hidden text-base px-6">
            <div class="flex flex-row my-2">
              <div class="text-shifter-white-soft w-2/6 font-bold">Connection Name</div>
              <div class="w-4/6 overflow-x-auto">
                <input
                  type="text" 
                  class="form-control block w-full px-3 py-1.5 text-base font-normal text-shifter-black bg-shifter-grey-background bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-slate-500 focus:border-blue-600 focus:outline-none"
                  placeholder="New Cluster Connection Name"
                  v-model.trim.lazy="shifter.clusterConfig.connectionName"
                />
              </div>
            </div>
            <div class="flex flex-row my-2">
              <div class="w-2/6 text-shifter-white-soft font-bold">Endpoint URL</div>
              <div class="w-4/6 overflow-x-auto">
                <input
                  type="text"
                  class="form-control block w-full px-3 py-1.5 text-base font-normal text-shifter-black bg-shifter-grey-background bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                  placeholder="example: https://0.0.0.0:443"
                  v-model.trim.lazy="shifter.clusterConfig.baseUrl"
                />
              </div>
            </div>
            <div class="flex flex-row my-2 grow-0">
              <div class="w-2/6 text-shifter-white-soft font-bold">Authentication Token</div>
              <div class="w-4/6 overflow-x-auto">
                <textarea
                  class="form-control block w-full px-3 py-1.5 text-base font-normal text-shifter-black bg-shifter-grey-background bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                  rows="3"
                  placeholder="User token"
                  v-model.trim.lazy="shifter.clusterConfig.bearerToken"
                ></textarea>
              </div>
            </div>
          </div>
          <div class="container flex flex-row mx-auto w-full text-xl gap-5 justify-end">
            <div class="flex flex-col w-1/4">
              <a
                @click="closeModal"
                class="rounded px-6 bg-shifter-red-soft rounded border border-shifter-red-soft py-2 text-center font-bold text-sm"
                >Cancel</a
              >
            </div>
            <div class="flex flex-col w-1/4">
              <a 
                @click="addClusterConfig"
                v-show="formValid"
                class="disabled rounded bg-shifter-red-soft rounded border border-shifter-red-soft py-2 text-center font-bold text-sm"
                >Create Configuration</a
              >
              <a
                v-show="!formValid"
                class="opacity-25 rounded bg-shifter-black-soft rounded border border-shifter-red-mute py-2 text-center font-bold text-sm"
                >Create Configuration</a
              >
            </div>
            
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
import { useModalClusterAdd } from "../stores/configurations/modalClusterAdd";
import { useConfigurationsClusters } from "../stores/configurations/clusters";
// Plugin & Package Imports
import { mapState, mapActions } from "pinia";

export default {
  data() {
    return {
      shifter: {
        id: null,
        enabled: true,
        clusterConfig: {
          connectionName: null,
          baseUrl: null,
          bearerToken: null,
        },
      },
    };
  },
  methods: {
    ...mapActions(useConfigurationsClusters, {
      addCluster: "addCluster",
    }),
    ...mapActions(useModalClusterAdd, { closeModal: "closeModal" }),
    addClusterConfig() {
      this.addCluster(this.shifter)
        .then(() => {
          shifterConfigurationUpdateSuccess(
            "Openshift Cluster Configuration created"
          );
          this.closeModal();
        })
        .catch(() => {
          shifterConfigurationUpdateError(
            "Unable to create new Openshift Cluster Configuration"
          );
          this.closeModal();
        });
    },
  },
  computed: {
    ...mapState(useModalClusterAdd, {
      showModal: "showModal",
    }),

    validateBaseUrl() {
      if (
        this.shifter.clusterConfig.baseUrl !== null &&
        this.shifter.clusterConfig.baseUrl !== undefined &&
        this.shifter.clusterConfig.baseUrl.length >= 1
      ) {
        return true;
      }
      return false;
    },

    validateConnectionName() {
      if (
        this.shifter.clusterConfig.connectionName !== null &&
        this.shifter.clusterConfig.connectionName !== undefined &&
        this.shifter.clusterConfig.connectionName.length >= 1
      ) {
        return true;
      }
      return false;
    },

    validateBearerToken() {
      if (
        this.shifter.clusterConfig.bearerToken !== null &&
        this.shifter.clusterConfig.bearerToken !== undefined &&
        this.shifter.clusterConfig.bearerToken.length >= 1
      ) {
        return true;
      }
      return false;
    },

    formValid() {
      if (
        this.validateBaseUrl &&
        this.validateBearerToken &&
        this.validateConnectionName
      ) {
        return true;
      }
      return false;
    },
  },
};
</script>

<style scoped></style>
