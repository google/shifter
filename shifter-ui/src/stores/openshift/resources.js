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

// Shifter Import Config
import { shifterConfig } from "@/main";
// Notifications Imports
import { notifyAxiosError } from "@/notifications";
// Axios Imports
import axios from "axios";
// Pinia Store Imports
import { defineStore } from "pinia";
// External Pinia Store Imports
import { useConfigurationsClusters } from "../configurations/clusters";
import { useConfigurationsLoading } from "../configurations/loading";
// Instansitate Pinia Store Objects
const storeConfigClusters = useConfigurationsClusters();
const storeConfigLoading = useConfigurationsLoading();

// Pinia Store Definition
export const useOSResources = defineStore(
  "shifter-api-v1-openshift-resources",
  {
    state: () => {
      return {
        osResources: [],
      };
    },

    getters: {
      all(state) {
        return state.osResources;
      },
      getByKind(state) {
        return (kind) =>
          state.osResources.find((resource) => reosurce.metadata.kind === kind);
      },
      getByName(state) {
        return (name) =>
          state.osResources.find((resource) => resource.metadata.name === name);
      },
      getByUid(state) {
        return (uid) =>
          state.osResources.find((resource) => resource.metadata.uid === uid);
      },
    },

    actions: {
      async fetch(clusterId) {
        // API Endpoint Configuration
        const config = {
          method: "post",
          url: shifterConfig.API_BASE_URL + "/openshift/resources/",
          headers: {},
          data: { ...storeConfigClusters.getCluster(clusterId) },
          timeout: 2000,
        };
        try {
          storeConfigLoading.startLoading(
            "Loading...",
            "Fetching OpenShift Resources"
          );
          this.osResources = [];
          this.osResources = await axios(config)
            .then((response) => {
              // handle success
              storeConfigLoading.endLoading();
              return response.data.resources.items;
            })
            .catch((err) => {
              notifyAxiosError(
                err,
                "Problem Loading OpenShift Resources",
                6000
              );
              storeConfigLoading.endLoading();
              return err;
            });
        } catch (err) {
          this.osResources = [];
          notifyAxiosError(err, "Problem Loading OpenShift Resources", 6000);
          storeConfigLoading.endLoading();
          return err;
        }
      },
    },
  }
);
