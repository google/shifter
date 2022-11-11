<!--
 Copyright 2022 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

<template>
  <div class="container flex mx-auto m-6 items-center">
    <div
      class="container flex-row mx-auto bg-shifter-black-mute justify-center rounded-2xl py-6"
    >
      <!-- DOWNLOAD SELECTION -->
      <div class="container flex-row mx-auto justify-center py-12">
        <div class="container flex-row justify-center items-center">
          <div class="flex justify-center bold text-4xl m-2 text-shifter-white-soft">
            Download Conversion Files
          </div>
          <div class="flex justify-center text-baseline m-2">
            {{ displayName }}
          </div>
        </div>
        <div
          v-show="singleItem"
          class="container flex mx-auto justify-center my-4"
        >
          <div class="container flex mx-auto justify-center my-4">
            <a
              class="uppercase rounded px-6 py-2 bg-shifter-red-soft hover:bg-shifter-red animate-pulse"
              :onclick="download"
              >Download Files</a
            >
          </div>
        </div>
        <div
          v-show="!singleItem"
          class="container flex mx-auto justify-center my-4"
        >
          <p>List All Downloads</p>
          {{ downloads }}
        </div>
      </div>
      <!-- END DOWNLOAD SELECTION -->
    </div>
  </div>
</template>

<script>
// Pinia Store Imports
import { useDownloadsObjects } from "../stores/downloads/downloads";
// Plugin & Package Imports
import { mapState, mapActions } from "pinia";

export default {
  data() {
    return {
      defaultFileName: "Shifter Conversion Package",
      downloadData: null,
    };
  },
  computed: {
    ...mapState(useDownloadsObjects, {
      downloads: "all",
    }),

    singleItem() {
      return this.$route.params.downloadId !== undefined;
    },
    downloadId() {
      if (this.singleItem) {
        return this.$route.params.downloadId;
      } else return null;
    },
    displayName() {
      if (
        this.downloadData !== undefined &&
        this.downloadData !== null &&
        this.downloadData.suid !== undefined &&
        this.downloadData.suid !== null &&
        this.downloadData.suid.displayName !== undefined &&
        this.downloadData.suid.displayName !== null
      ) {
        return this.downloadData.suid.displayName;
      } else return this.defaultFileName;
    },
  },
  methods: {
    ...mapActions(useDownloadsObjects, { getDownloads: "get" }),
    ...mapActions(useDownloadsObjects, { getFile: "getFile" }),

    download() {
      this.getFile(this.downloadId, this.displayName);
    },
  },

  created() {
    this.getDownloads(this.downloadId).then((response) => {
      this.downloadData = response.data;
    });
  },
};
</script>
