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

// Shifter Import Config
import { shifterConfig } from "@/main";
// Axios Imports
import axios from "axios";
// Pinia Store Imports
import { defineStore } from "pinia";

// API Endpoint Configuration
const config = {
  method: "get",
  url: shifterConfig.API_BASE_URL + "/status/healthz",
  headers: {},
  data: null,
};

// Pinia Store Definition
export const useShifterV1StatusHealthz = defineStore(
  "shifter-api-v1-status-healthz",
  {
    state: () => {
      return {
        data: {
          message: "Running Connection Tests.",
        },
        fetching: false,
      };
    },

    getters: {
      results(state) {
        return state.data;
      },

      isFetching(state) {
        return state.fetching;
      },
    },

    actions: {
      async fetchHealthz() {
        this.fetching = true;
        try {
          const response = await axios(config);
          try {
            const result = await response.data;
            this.data = result;
          } catch (err) {
            this.data = [];
            console.error(
              "Error loading Shifter Server, Status, Healthz API:",
              err
            );
            return err;
          }
        } catch (err) {
          this.data = {
            message: "Shifter Server Unreachable, Timeout",
          };
        }
        this.fetching = false;
      },
    },
  }
);
