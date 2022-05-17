import { defineStore } from "pinia";

export const useJSONModal = defineStore("configurations-convert-json-modal", {
  state: () => {
    return {
      content: {},
      showModal: true,
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
      this.content = content
      this.showModal = true;
    },

    async setContnet(content) {
      this.content = content
    },
  },
});
