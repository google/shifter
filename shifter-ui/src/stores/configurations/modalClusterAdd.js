// Pinia Store Imports
import { defineStore } from "pinia";

// Pinia Store Definition
export const useModalClusterAdd = defineStore(
  "shifter-config-clusters-add-modal",
  {
    state: () => {
      return {
        show: false,
      };
    },

    getters: {
      showModal(state) {
        return state.show;
      },
    },

    actions: {
      async closeModal() {
        this.show = false;
      },
      async openModal() {
        this.show = true;
      },
    },
  }
);
