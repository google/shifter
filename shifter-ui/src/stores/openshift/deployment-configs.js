import { shifterConfig } from "@/main";
import axios from "axios";
import { defineStore } from "pinia";

import { useConfigurationsLoading } from "../configurations/loading";
import { useConfigurationsClusters } from "../configurations/clusters";
const configurationsLoading = useConfigurationsLoading();
const configurationsClusters = useConfigurationsClusters();

export const useOSDeploymentConfigs = defineStore(
  "openshift-deploymentConfigs",
  {
    state: () => {
      return {
        osDeploymentConfigs: [],
      };
    },

    getters: {
      all(state) {
        return state.osDeploymentConfigs;
      },
      getByName(state) {
        return (name) =>
          state.osDeploymentConfigs.find(
            (deploymentconfig) => deploymentconfig.metadata.name === name
          );
      },
      getByUid(state) {
        return (uid) =>
          state.osDeploymentConfigs.find(
            (deploymentconfig) => deploymentconfig.metadata.uid === uid
          );
      },
    },

    actions: {
      async fetch(clusterId) {
        // API Endpoint Configuration
        const config = {
          method: "post",
          url: shifterConfig.API_BASE_URL + "/openshift/deploymentconfigs/",
          headers: {},
          data: { ...configurationsClusters.getCluster(clusterId) },
          timeout: 1000,
        };
        try {
          configurationsLoading.startLoading(
            "Loading...",
            "Fetching OpenShift Deployment Configurations"
          );
          this.osDeploymentConfigs = [];
          this.osDeploymentConfigs = await axios(config)
            .then((response) => {
              // handle success
              console.log(response);
              configurationsLoading.endLoading();
              return response.data.deploymentConfigs.items;
            })
            .catch((err) => {
              console.error("Error", err);
              configurationsLoading.endLoading(err);
              return err;
            });
        } catch (err) {
          this.osDeploymentConfigs = [];
          console.error("Error", err);
          configurationsLoading.endLoading(err);
        }
      },
    },
  }
);
