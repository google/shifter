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
