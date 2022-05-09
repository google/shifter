import { defineStore } from 'pinia'

export const useConfigurationsClusters = defineStore('configurations-clusters', {
  state: () => {
    return {
      clusters: [
        {   
            id: 0,
            enabled: true,
            shifter: {
                clusterConfig:{
                    connectionName: "OKD Shifter DEV Server v3.11",
                    baseUrl: "https://console.okd.shifter.cloud:8443",
                    bearerToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InJvYm90LXRva2VuLThod3dtIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6InJvYm90Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiNGIyZjdmZDItYmZmOS0xMWVjLWI5Y2MtNDIwMTBhMDAwMDAzIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6cm9ib3QifQ.BshucYnBWMlBPtcGWBk8EbCGWGTjXYWoExrspYKvCnLrP8Iqs4p-SA-3RCvZ8fe0JyxwE-1GIDIgaxH9270303BGXPcdEJFUNW06j_ZPv5EZYn-C8UQxCmRpfZBe9QdtOEre6YSmcL-GrfE4fdaHAkYN1Yv5Fr6k_JPt3XGRHD2Hy0By14eFG7oW3b2rj15_lYWW-vfQaYSbcG8HViPF62lqQL0KSZwySOC82jpMdKUCdtuuBC01Gzd9NfTGTeBo3VSQ6mTTOchL1T9GX6oD-wu4J9NOmbPYZPln0Q-_8h6H488Mc4xTfzLdkb136-GJAwT7GgksYm3UMx_sXkj36A"
                }
            }
        },
        {
            id: 1,
            enabled: false,
            shifter: {
                clusterConfig:{
                    connectionName: "OKD Shifter TST Server v3.11",
                    baseUrl: "https://console.okd.shifter.cloud:8443",
                    bearerToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InJvYm90LXRva2VuLThod3dtIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6InJvYm90Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiNGIyZjdmZDItYmZmOS0xMWVjLWI5Y2MtNDIwMTBhMDAwMDAzIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6cm9ib3QifQ.BshucYnBWMlBPtcGWBk8EbCGWGTjXYWoExrspYKvCnLrP8Iqs4p-SA-3RCvZ8fe0JyxwE-1GIDIgaxH9270303BGXPcdEJFUNW06j_ZPv5EZYn-C8UQxCmRpfZBe9QdtOEre6YSmcL-GrfE4fdaHAkYN1Yv5Fr6k_JPt3XGRHD2Hy0By14eFG7oW3b2rj15_lYWW-vfQaYSbcG8HViPF62lqQL0KSZwySOC82jpMdKUCdtuuBC01Gzd9NfTGTeBo3VSQ6mTTOchL1T9GX6oD-wu4J9NOmbPYZPln0Q-_8h6H488Mc4xTfzLdkb136-GJAwT7GgksYm3UMx_sXkj36A"
                }
            }
        },
        {
            id: 2,
            enabled: true,
            shifter: {
                clusterConfig:{
                    connectionName: "OKD Shifter PRD Server v3.11",
                    baseUrl: "https://console.okd.shifter.cloud:8443",
                    bearerToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InJvYm90LXRva2VuLThod3dtIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6InJvYm90Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiNGIyZjdmZDItYmZmOS0xMWVjLWI5Y2MtNDIwMTBhMDAwMDAzIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6cm9ib3QifQ.BshucYnBWMlBPtcGWBk8EbCGWGTjXYWoExrspYKvCnLrP8Iqs4p-SA-3RCvZ8fe0JyxwE-1GIDIgaxH9270303BGXPcdEJFUNW06j_ZPv5EZYn-C8UQxCmRpfZBe9QdtOEre6YSmcL-GrfE4fdaHAkYN1Yv5Fr6k_JPt3XGRHD2Hy0By14eFG7oW3b2rj15_lYWW-vfQaYSbcG8HViPF62lqQL0KSZwySOC82jpMdKUCdtuuBC01Gzd9NfTGTeBo3VSQ6mTTOchL1T9GX6oD-wu4J9NOmbPYZPln0Q-_8h6H488Mc4xTfzLdkb136-GJAwT7GgksYm3UMx_sXkj36A"
                }
            }
        },
      ],
      fetching: false
    }
  },

  getters: {
    getActiveClusters(state) {
        const activeClusters = state.clusters.filter((cluster) => cluster.enabled)
        return activeClusters
      },
    getAllClusters(state) {
      return state.clusters;
    },
  },

  actions: {
  }
})