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
export const useOSDeploymentConfigs = defineStore(
  "shifter-api-v1-openshift-deploymentConfigs",
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
          data: { ...storeConfigClusters.getCluster(clusterId) },
          timeout: 2000,
        };
        try {
          storeConfigLoading.startLoading(
            "Loading...",
            "Fetching OpenShift Deployment Configurations"
          );
          this.osDeploymentConfigs = [];
          this.osDeploymentConfigs = await axios(config)
            .then((response) => {
              // handle success
              storeConfigLoading.endLoading();
              return response.data.deploymentConfigs.items;
            })
            .catch((err) => {
              notifyAxiosError(
                err,
                "Problem Loading OpenShift Deployment Configurations",
                6000
              );
              storeConfigLoading.endLoading();
              return err;
            });
        } catch (err) {
          this.osDeploymentConfigs = [];
          notifyAxiosError(
            err,
            "Problem Loading OpenShift Deployment Configurations",
            6000
          );
          storeConfigLoading.endLoading();
          return err;
        }
      },
    },
  }
);
