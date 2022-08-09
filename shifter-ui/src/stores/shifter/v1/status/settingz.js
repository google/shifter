// Shifter Import Config
import { shifterConfig } from "@/main";
// Axios Imports
import axios from "axios";
// Pinia Store Imports
import { defineStore } from "pinia";

// API Endpoint Configuration
const config = {
  method: "get",
  url: shifterConfig.API_BASE_URL + "/status/settingz",
  headers: {},
  data: null,
};

// Pinia Store Definition
export const useShifterV1StatusSettingz = defineStore(
  "shifter-api-v1-status-settingz",
  {
    state: () => {
      return {
        data: {
          message: "Collecting Server Settings.",
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
      async fetchSettingz() {
        this.fetching = true;
        try {
          const response = await axios(config);
          try {
            const result = await response.data;
            this.data = result;
          } catch (err) {
            this.data = [];
            console.error(
              "Error loading Shifter Server Settings, Status, Settingz API:",
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
