// Pinia Store Imports
import { defineStore } from "pinia";

// Pinia Store Definition
export const useJSONModal = defineStore("shifter-config-convert-json-modal", {
  state: () => {
    return {
      content: {},
      showModal: false,
    };
  },

  getters: {
    getContent(state) {
      return state.content;
    },
    showJSONModal(state) {
      return state.showModal;
    },
  },

  actions: {
    async closeModal() {
      this.showModal = false;
    },
    async openModal(content) {
      this.content = content;
      this.showModal = true;
    },

    async setContnet(content) {
      this.content = content;
    },
  },
});
