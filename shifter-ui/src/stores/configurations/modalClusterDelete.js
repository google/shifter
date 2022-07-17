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
