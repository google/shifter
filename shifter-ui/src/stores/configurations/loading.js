/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Pinia Store Imports
import { defineStore } from "pinia";
// Store Contants
const defaultTitle = "Loading...";
const defaultSubTitle = "Shifter is working on your request.";

// Pinia Store Definition
export const useConfigurationsLoading = defineStore("shifter-config-loading", {
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
