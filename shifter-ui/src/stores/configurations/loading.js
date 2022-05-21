import { defineStore } from "pinia";

const defaultTitle = "Loading...";
const defaultSubTitle = "Shifter is working on your request.";

export const useConfigurationsLoading = defineStore("configurations-loading", {
  state: () => {
    return {
      loading: false,
      title: defaultTitle,
      subtitle: defaultSubTitle,
    };
  },

  getters: {
    isLoading(state) {
      return state.loading;
    },
    getTitle(state) {
      return state.title;
    },
    getSubTitle(state) {
      return state.subtitle;
    },
  },

  actions: {
    startLoading(title = defaultTitle, subtitle = defaultSubTitle) {
      // Set Loading Splash - Title
      if (title === null || title == undefined) {
        // Use Default Title
        this.title = defaultTitle;
      } else {
        this.title = title;
      }
      // Set Loading Splash - Subtitle
      if (subtitle === null || subtitle == undefined) {
        // Use Default SubTitle
        this.subtitle = defaultSubTitle;
      } else {
        this.subtitle = subtitle;
      }
      // Commence Loading
      this.loading = true;
    },
    endLoading() {
      setTimeout(() => {
        // End Loading Splash
        this.loading = false;
      }, 300);
    },
  },
});
