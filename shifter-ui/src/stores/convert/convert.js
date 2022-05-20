import { defineStore } from "pinia";

import { useConfigurationsLoading } from "../configurations/loading";
import { useOSProjects } from "../openshift/projects";
import { useOSDeploymentConfigs } from "../openshift/deployment-configs";

import { useConfigurationsClusters } from "../configurations/clusters";
const configurationsClusters = useConfigurationsClusters();

import { shifterConfig } from "@/main";
import axios from "axios";

const configurationsLoading = useConfigurationsLoading();
const oSProjects = useOSProjects();
const oSDeploymentConfigs = useOSDeploymentConfigs();

export const useConvertObjects = defineStore("convert-objects", {
  state: () => {
    return {
      cluster: {},
      conversionItems: [],
      fetching: false,
    };
  },

  getters: {
    selectedCluster(state) {
      return state.cluster;
    },
    all(state) {
      return state.conversionItems;
    },
    isSelected(state) {
      return (item) => {
        if (state.conversionItems.length >= 1) {
          var idx = state.conversionItems.findIndex((object) => {
            if (item.metadata !== undefined) {
              return object.deploymentConfig.metadata.uid === item.metadata.uid;
            }
          });

          if (idx >= 0) {
            return true;
          }
        }
        return false;
      };
    },
  },

  actions: {
    async setCluster(clusterId) {
      if (this.cluster.id !== clusterId || this.cluster.id === undefined) {
        // TODO <<-- Subscribe to this action in project refresh
        // Refresh OpenShift Projects
        oSProjects.fetch(clusterId);
        // Refresh OpenShift Deployment Configs
        oSDeploymentConfigs.fetch(clusterId);
        // When the Cluster is Changed. Reset the State and Update Cluster
        this.$reset();
        // Set Cluster from Cluster ID
        this.cluster = {
          ...configurationsClusters.getCluster(clusterId),
        };
      }
    },
    async add(deploymentconfig) {
      try {
        const newConversionItem = {
          namespace: oSProjects.getByName(deploymentconfig.metadata.namespace),
          deploymentConfig: oSDeploymentConfigs.getByUid(
            deploymentconfig.metadata.uid
          ),
        };
        this.conversionItems.push(newConversionItem);
      } catch (error) {
        console.log(error);
        return error;
      }
      return;
    },

    async remove(item) {
      if (this.conversionItems.length >= 1) {
        var idx = this.conversionItems.findIndex((object) => {
          if (item.metadata !== undefined) {
            return object.deploymentConfig.metadata.uid === item.metadata.uid;
          }
          if (item.metadata === undefined) {
            return (
              object.deploymentConfig.metadata.uid ===
              item.deploymentConfig.metadata.uid
            );
          }
        });
        if (idx >= 0) {
          this.conversionItems.splice(idx, 1);
        }
      }
    },

    async convert() {
      const config = {
        method: "post",
        url: shifterConfig.API_BASE_URL + "/shifter/convert/",
        headers: {},
        data: {
          shifter: JSON.parse(JSON.stringify({ ...this.selectedCluster }))
            .shifter,
          items: JSON.parse(JSON.stringify([...this.all])),
        },
        timeout: 1000,
      };
      try {
        configurationsLoading.startLoading(
          "Shifting...",
          "Standby while we convert the workloads."
        );
        await axios(config)
          .then(function (response) {
            // handle success
            console.log(response);
          })
          .catch((err) => {
            console.error("Error", err);
            configurationsLoading.endLoading();
            return err;
          });
      } catch (err) {
        console.error("Error", err);
        configurationsLoading.endLoading();
      }
    },
  },
});
