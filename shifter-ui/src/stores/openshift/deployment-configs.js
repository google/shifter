import { shifterConfig } from "@/main";
import axios from "axios";
import { defineStore } from "pinia";

import { useConfigurationsClusters } from "../configurations/clusters";
const configurationsClusters = useConfigurationsClusters();

export const useOSDeploymentConfigs = defineStore(
  "openshift-deploymentConfigs",
  {
    state: () => {
      return {
        osDeploymentConfigs: [],
        fetching: false,
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
        };
        this.fetching = true;
        try {
          const response = await axios(config);
          try {
            this.osDeploymentConfigs = response.data.deploymentConfigs.items;
          } catch (err) {
            this.osDeploymentConfigs = [];
            console.error("Error", err);
            return err;
          }
        } catch (err) {
          this.osDeploymentConfigs = [];
        }
        this.fetching = false;
      },
    },
  }
);
