/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Pinia Store Imports
import { defineStore } from "pinia";
// Pinia Store Definition
export const useModalClusterDelete = defineStore(
  "shifter-config-clusters-delete-modal",
  {
    state: () => {
      return {
        show: false,
        clusterId: null,
      };
    },

    getters: {
      showModal(state) {
        return state.show;
      },
      getClusterId(state) {
        return state.clusterId;
      },
    },

    actions: {
      async closeModal() {
        this.clusterId = null;
        this.show = false;
      },
      async openModal(clusterConfig) {
        this.clusterId = clusterConfig.id;
        this.show = true;
      },
    },
  }
);
