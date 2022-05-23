// Shifter Import Config
import { shifterConfig } from "@/main";
// Axios Imports
import axios from "axios";
// Pinia Store Imports
import { defineStore } from "pinia";

// Pinia Store Definition
export const useDownloadsObjects = defineStore(
  "shifter-api-v1-downloads-objects",
  {
    state: () => {
      return {
        downloads: {},
        fetching: false,
      };
    },

    getters: {
      byId(state) {
        return (downloadId) =>
          state.downloads.find(
            (download) => download.items.downloadId === downloadId
          );
      },
      all(state) {
        return state.downloadItems;
      },
    },

    actions: {
      async get(downloadId) {
        var url = shifterConfig.API_BASE_URL + "/shifter/downloads/";
        // If Specific Download ID is Requested
        if (downloadId !== null) {
          url = url + downloadId;
        }

        const config = {
          method: "get",
          url: url,
          headers: {},
          data: {},
        };
        try {
          const response = await axios(config);
          try {
            console.log(response);
          } catch (err) {
            console.error("Error", err);
            return err;
          }
        } catch (err) {
          console.error("Error", err);
        }
      },
    },
  }
);
