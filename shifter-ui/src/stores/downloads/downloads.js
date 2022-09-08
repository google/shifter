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
// Notifications Imports
import { notifyAxiosError } from "@/notifications";
// Axios Imports
import axios from "axios";
// Pinia Store Imports
import { defineStore } from "pinia";
// External Pinia Store Imports
import { useConfigurationsLoading } from "../configurations/loading";
// Instansitate Pinia Store Objects
const storeConfigLoading = useConfigurationsLoading();

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
      async get(downloadId = null) {
        this.$reset();
        var url = shifterConfig.API_BASE_URL + "/shifter/downloads/";
        // If Specific Download ID is Requested
        if (downloadId !== null) {
          url = url + downloadId;
        }

        const config = {
          method: "get",
          responseType: "json",
          url: url,
          headers: {},
          data: {},
        };
        try {
          storeConfigLoading.startLoading(
            "Searching...",
            "Standby while we locate your converted files."
          );
          return await axios(config)
            .then(function (response) {
              // handle success
              storeConfigLoading.endLoading();
              return response;
            })
            .catch((err) => {
              notifyAxiosError(err, "Error Locating Downloads", 6000);
              storeConfigLoading.endLoading();
              return err;
            });
        } catch (err) {
          notifyAxiosError(err, "Error Locating Downloads", 6000);
          storeConfigLoading.endLoading();
          return err;
        }
      },

      async getFile(
        downloadId = null,
        fileName = "Shifter Conversion Package"
      ) {
        var url =
          shifterConfig.API_BASE_URL +
          "/shifter/downloads/" +
          downloadId +
          "/file";

        const config = {
          method: "get",
          responseType: "blob",
          url: url,
          headers: {},
          data: {},
        };
        try {
          /*storeConfigLoading.startLoading(
            "Searching...",
            "Standby while we do your converted files."
          );*/
          return await axios(config)
            .then(function (response) {
              // handle success
              const url = window.URL.createObjectURL(new Blob([response.data]));
              const link = document.createElement("a");
              link.href = url;
              link.setAttribute("download", fileName + ".zip"); //or any other extension
              document.body.appendChild(link);
              link.click();

              return response;
            })
            .catch((err) => {
              notifyAxiosError(err, "Error Download Requested File", 6000);
              //storeConfigLoading.endLoading();
              return err;
            });
        } catch (err) {
          notifyAxiosError(err, "Error Download Requested File", 6000);
          //storeConfigLoading.endLoading();
          return err;
        }
      },
    },
  }
);
