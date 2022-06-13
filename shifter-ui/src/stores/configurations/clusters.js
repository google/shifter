// @vueuse/core Imports
import { useStorage } from "@vueuse/core";
// Pinia Store Imports
import { defineStore } from "pinia";
// Pinia Store Definition
export const useConfigurationsClusters = defineStore(
  "shifter-config-clusters",
  {
    state: () => {
      return {
        clusters: useStorage("shifter-config-clusters", []),
        fetching: false,
      };
    },

    getters: {
      getNextId(state) {
        if (state.clusters.length >= 1) {
          return this.clusters[this.clusters.length - 1].id + 1;
        }
        return 0;
      },
      getActiveClusters(state) {
        return state.clusters.filter((cluster) => cluster.enabled);
      },
      getAllClusters(state) {
        return state.clusters.sort((a, b) =>
          a.shifter.clusterConfig.connectionName >
          b.shifter.clusterConfig.connectionName
            ? 1
            : b.shifter.clusterConfig.connectionName >
              a.shifter.clusterConfig.connectionName
            ? -1
            : 0
        );
      },
      getCluster(state) {
        return (clusterId) =>
          state.clusters.find((cluster) => cluster.id === clusterId);
      },
    },
    actions: {
      async deleteCluster(clusterId) {
        if (this.clusters.length >= 1) {
          var idx = this.clusters.findIndex((object) => {
            if (clusterId !== undefined) {
              return object.id === clusterId;
            }
          });
          if (idx >= 0) {
            this.clusters.splice(idx, 1);
            return;
          } else {
            throw TypeError("Failed to remove Cluster Configuration");
          }
        }
      },
      async addCluster(cluster) {
        this.clusters.push({
          id: this.getNextId,
          enabled: cluster.enabled,
          shifter: {
            clusterConfig: {
              connectionName: cluster.clusterConfig.connectionName,
              username: "",
              password: "",
              baseUrl: cluster.clusterConfig.baseUrl,
              bearerToken: cluster.clusterConfig.bearerToken,
            },
          },
        });
        // if Cluster ID already Exists then Clean up
        if (cluster.id !== null) {
          this.deleteCluster(cluster.id);
        }
      },
    },
  }
);
