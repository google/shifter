// Pinia Store Imports
import { defineStore } from "pinia";

// Pinia Store Definition
export const useConfigurationsClusters = defineStore(
  "shifter-config-clusters",
  {
    state: () => {
      return {
        clusters: [],
        /*clusters: [
          {
            id: 0,
            enabled: true,
            shifter: {
              clusterConfig: {
                connectionName: "OKD Shifter DEV Server v3.11",
                username: "",
                password: "",
                baseUrl: "https://console.okd.shifter.cloud:8443",
                bearerToken:
                  "eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InJvYm90LXRva2VuLThod3dtIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6InJvYm90Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiNGIyZjdmZDItYmZmOS0xMWVjLWI5Y2MtNDIwMTBhMDAwMDAzIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6cm9ib3QifQ.BshucYnBWMlBPtcGWBk8EbCGWGTjXYWoExrspYKvCnLrP8Iqs4p-SA-3RCvZ8fe0JyxwE-1GIDIgaxH9270303BGXPcdEJFUNW06j_ZPv5EZYn-C8UQxCmRpfZBe9QdtOEre6YSmcL-GrfE4fdaHAkYN1Yv5Fr6k_JPt3XGRHD2Hy0By14eFG7oW3b2rj15_lYWW-vfQaYSbcG8HViPF62lqQL0KSZwySOC82jpMdKUCdtuuBC01Gzd9NfTGTeBo3VSQ6mTTOchL1T9GX6oD-wu4J9NOmbPYZPln0Q-_8h6H488Mc4xTfzLdkb136-GJAwT7GgksYm3UMx_sXkj36A",
              },
            },
          },
          {
            id: 1,
            enabled: false,
            shifter: {
              clusterConfig: {
                connectionName: "OKD Shifter TST Server v3.11",
                username: "",
                password: "",
                baseUrl: "https://console.okd.shifter.cloud:8443",
                bearerToken:
                  "eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InJvYm90LXRva2VuLThod3dtIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6InJvYm90Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiNGIyZjdmZDItYmZmOS0xMWVjLWI5Y2MtNDIwMTBhMDAwMDAzIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6cm9ib3QifQ.BshucYnBWMlBPtcGWBk8EbCGWGTjXYWoExrspYKvCnLrP8Iqs4p-SA-3RCvZ8fe0JyxwE-1GIDIgaxH9270303BGXPcdEJFUNW06j_ZPv5EZYn-C8UQxCmRpfZBe9QdtOEre6YSmcL-GrfE4fdaHAkYN1Yv5Fr6k_JPt3XGRHD2Hy0By14eFG7oW3b2rj15_lYWW-vfQaYSbcG8HViPF62lqQL0KSZwySOC82jpMdKUCdtuuBC01Gzd9NfTGTeBo3VSQ6mTTOchL1T9GX6oD-wu4J9NOmbPYZPln0Q-_8h6H488Mc4xTfzLdkb136-GJAwT7GgksYm3UMx_sXkj36A",
              },
            },
          },
          {
            id: 2,
            enabled: true,
            shifter: {
              clusterConfig: {
                connectionName: "OKD Shifter PRD Server v3.11",
                username: "",
                password: "",
                baseUrl: "https://console.okd.shifter.cloud:8443",
                bearerToken:
                  "eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InJvYm90LXRva2VuLThod3dtIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6InJvYm90Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiNGIyZjdmZDItYmZmOS0xMWVjLWI5Y2MtNDIwMTBhMDAwMDAzIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6cm9ib3QifQ.BshucYnBWMlBPtcGWBk8EbCGWGTjXYWoExrspYKvCnLrP8Iqs4p-SA-3RCvZ8fe0JyxwE-1GIDIgaxH9270303BGXPcdEJFUNW06j_ZPv5EZYn-C8UQxCmRpfZBe9QdtOEre6YSmcL-GrfE4fdaHAkYN1Yv5Fr6k_JPt3XGRHD2Hy0By14eFG7oW3b2rj15_lYWW-vfQaYSbcG8HViPF62lqQL0KSZwySOC82jpMdKUCdtuuBC01Gzd9NfTGTeBo3VSQ6mTTOchL1T9GX6oD-wu4J9NOmbPYZPln0Q-_8h6H488Mc4xTfzLdkb136-GJAwT7GgksYm3UMx_sXkj36A",
              },
            },
          },
          {
            id: 3,
            enabled: true,
            shifter: {
              clusterConfig: {
                connectionName:
                  "OKD Shifter PRD Server v3.11 Username & Password",
                username: "shifter",
                password: "P4@FpHbMn&YqFEfn",
                baseUrl: "https://console.okd.shifter.cloud:8443",
                bearerToken: "",
              },
            },
          },
        ],*/
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
