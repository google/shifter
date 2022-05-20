<template>
  <div class="container flex mx-auto m-6 items-center">
    <div
      class="container flex-row mx-auto bg-shifter-black-mute justify-center rounded-2xl py-6"
    >
      <!-- DOWNLOAD SELECTION -->
      <div class="container flex-row mx-auto justify-center py-12">
        <div class="container flex-row justify-center items-center">
          <div class="flex justify-center bold text-4xl m-2">
            Download Conversion Files
          </div>
          <div class="flex justify-center text-baseline m-2">
            Select OpenShift cluster from which you would like to convert
            workloads
          </div>
        </div>
        <div
          v-show="singleItem"
          class="container flex mx-auto justify-center my-4"
        >
          <p>Single Item: {{ singleItem }}</p>
          <p>Download ID: {{ downloadId }}</p>
          {{ downloads }}
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
    return {};
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
  },
  methods: {
    ...mapActions(useDownloadsObjects, { getDownloads: "get" }),
  },

  created() {
    this.getDownloads(this.downloadId);
  },
};
</script>
