// Shifter Import Config
import { shifterConfig } from "@/main";
// Notifications Imports
import { notifyAxiosError, shifterConversionSuccess } from "@/notifications";
// Axios Imports
import axios from "axios";
// Pinia Store Imports
import { defineStore } from "pinia";
// External Pinia Store Imports
import { useConfigurationsClusters } from "../configurations/clusters";
import { useConfigurationsLoading } from "../configurations/loading";
import { useOSProjects } from "../openshift/projects";
// Instansitate Pinia Store Objects
const storeConfigClusters = useConfigurationsClusters();
const storeConfigLoading = useConfigurationsLoading();

const storeOSProjects = useOSProjects();

// Pinia Store Definition
export const useConvertObjects = defineStore("shifter-api-v1-convert-objects", {
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
      //if (this.cluster.id === undefined) {
      // TODO <<-- Subscribe to this action in project refresh
      // Refresh OpenShift Projects
      storeOSProjects.fetch(clusterId);
      // When the Cluster is Changed. Reset the State and Update Cluster
      this.$reset();
      // Set Cluster from Cluster ID
      this.cluster = {
        ...storeConfigClusters.getCluster(clusterId),
      };
      //}
    },
    async add(deploymentconfig) {
      try {
        const newConversionItem = {
          namespace: storeOSProjects.getByName(
            deploymentconfig.metadata.namespace
          ),
          deploymentConfig: deploymentconfig,
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
        timeout: 10000,
      };
      try {
        storeConfigLoading.startLoading(
          "Shifting...",
          "Standby while we convert the workloads."
        );
        return await axios(config)
          .then((response) => {
            // handle success
            shifterConversionSuccess("Converted OpenShift objects.");
            // Clear Conversion State on Successful Conversion
            this.$reset();
            storeConfigLoading.endLoading();
            return response;
          })
          .catch((err) => {
            notifyAxiosError(err, "Error Converting Workloads", 6000);
            storeConfigLoading.endLoading();
            return err;
          });
      } catch (err) {
        notifyAxiosError(err, "Error Converting Workloads", 6000);
        storeConfigLoading.endLoading();
        return err;
      }
    },
  },
});
