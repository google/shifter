<template>
  <div
    class="container mb-6 border rounded-xl bg-shifter-black-soft overflow-hidden"
  >
    <div class="flex flex-row items-center my-4">
      <!-- Title -->
      <div class="container flex">
        <p class="text-xl font-bold mx-6">
          {{ clusterconfig.shifter.clusterConfig.connectionName }}
        </p>
      </div>
      <!-- Actions -->
      <div class="container flex flex-row-reverse gap-3 mx-6">
        <!-- Action: Delete Cluster Configuration -->
        <div class="flex">
          <!-- Material Design - SVG - trash-can -->
          <svg
            style="width: 24px; height: 24px"
            viewBox="0 0 24 24"
            @click="deleteConfig(clusterconfig)"
          >
            <path
              fill="currentColor"
              d="M9,3V4H4V6H5V19A2,2 0 0,0 7,21H17A2,2 0 0,0 19,19V6H20V4H15V3H9M9,8H11V17H9V8M13,8H15V17H13V8Z"
            />
          </svg>
        </div>
        <!-- End Action -->
        <!-- Action: Edit Cluster Configuration -->
        <div class="flex">
          <!-- Material Design - SVG - pencil-circle -->
          <svg
            style="width: 24px; height: 24px"
            viewBox="0 0 24 24"
            @click="toggleEditConfig()"
          >
            <path
              fill="currentColor"
              d="M12,2C6.47,2 2,6.47 2,12C2,17.53 6.47,22 12,22C17.53,22 22,17.53 22,12C22,6.47 17.53,2 12,2M15.1,7.07C15.24,7.07 15.38,7.12 15.5,7.23L16.77,8.5C17,8.72 17,9.07 16.77,9.28L15.77,10.28L13.72,8.23L14.72,7.23C14.82,7.12 14.96,7.07 15.1,7.07M13.13,8.81L15.19,10.87L9.13,16.93H7.07V14.87L13.13,8.81Z"
            />
          </svg>
        </div>
        <!-- End Action -->
        <!-- Action: Show Cluster Configuration Section -->
        <div class="flex">
          <!-- Material Design - SVG - pencil-circle -->
          <svg
            style="width: 24px; height: 24px"
            viewBox="0 0 24 24"
            @click="toggleConfig()"
          >
            <path
              fill="currentColor"
              d="M12 4.5C7 4.5 2.7 7.6 1 12C2.7 16.4 7 19.5 12 19.5H13.1C13 19.2 13 18.9 13 18.5C13 17.9 13.1 17.4 13.2 16.8C12.8 16.9 12.4 17 12 17C9.2 17 7 14.8 7 12S9.2 7 12 7 17 9.2 17 12C17 12.3 17 12.6 16.9 12.9C17.6 12.7 18.3 12.5 19 12.5C20.2 12.5 21.3 12.8 22.3 13.5C22.6 13 22.8 12.5 23 12C21.3 7.6 17 4.5 12 4.5M12 9C10.3 9 9 10.3 9 12S10.3 15 12 15 15 13.7 15 12 13.7 9 12 9M19 21V19H15V17H19V15L22 18L19 21"
            />
          </svg>
        </div>
        <!-- End Action -->
      </div>
    </div>
    <!-- Show Openshift Cluster Config -->
    <div
      class="flex flex-col bg-shifter-black-soft overflow-hidden px-6 pb-2"
      v-show="visibleConfig"
    >
      <div class="flex flex-col overflow-hidden">
        <div class="flex flex-row my-2">
          <div class="w-2/6 font-bold">Cluster Connection Name</div>
          <div class="w-3/6 overflow-x-auto">
            <input
              type="text"
              class="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-shifter-black-soft bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
              id="exampleFormControlInput1"
              :placeholder="clusterconfig.shifter.clusterConfig.connectionName"
              disabled
            />
          </div>
        </div>
        <div class="flex flex-row my-2">
          <div class="w-2/6 font-bold">Base URL</div>
          <div class="w-3/6 overflow-x-auto">
            <input
              type="text"
              class="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-shifter-black-soft bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
              id="exampleFormControlInput1"
              :placeholder="clusterconfig.shifter.clusterConfig.baseUrl"
              disabled
            />
          </div>
        </div>
        <div class="flex flex-row my-2 grow-0">
          <div class="w-2/6 font-bold">Openshift User Bearer Token</div>
          <div class="w-3/6 overflow-x-auto">
            <textarea
              class="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-shifter-black-soft bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
              id="exampleFormControlTextarea1"
              rows="3"
              :placeholder="clusterconfig.shifter.clusterConfig.bearerToken"
              disabled
            ></textarea>
          </div>
        </div>
      </div>
    </div>
    <!-- End  Show Openshift Cluster Config -->
    <!-- Edit Openshift Cluster Config -->
    <div
      class="flex flex-col bg-shifter-black-soft overflow-hidden px-6 pb-2"
      v-show="visibleEditConfig"
    >
      <div class="flex flex-col overflow-hidden">
        <div class="flex flex-row my-2">
          <div class="w-2/6 font-bold">Cluster Connection Name</div>
          <div class="w-3/6 overflow-x-auto">
            <input
              type="text"
              class="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-shifter-black-soft bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
              id="exampleFormControlInput1"
              :placeholder="clusterconfig.shifter.clusterConfig.connectionName"
            />
          </div>
        </div>
        <div class="flex flex-row my-2">
          <div class="w-2/6 font-bold">Base URL</div>
          <div class="w-3/6 overflow-x-auto">
            <input
              type="text"
              class="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-shifter-black-soft bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
              id="exampleFormControlInput1"
              :placeholder="clusterconfig.shifter.clusterConfig.baseUrl"
            />
          </div>
        </div>
        <div class="flex flex-row my-2 grow-0">
          <div class="w-2/6 font-bold">Openshift User Bearer Token</div>
          <div class="w-3/6 overflow-x-auto">
            <textarea
              class="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-shifter-black-soft bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
              id="exampleFormControlTextarea1"
              rows="3"
              :placeholder="clusterconfig.shifter.clusterConfig.bearerToken"
            ></textarea>
          </div>
        </div>
      </div>
      <div>
        <div class="container flex mx-auto justify-end px-10 gap-4">
          <a
            class="uppercase rounded px-6 py-2 bg-shifter-black hover:bg-shifter-red hover:animate-pulse"
            @click="commitEdit"
            >Save</a
          >
          <a
            class="uppercase rounded px-6 py-2 bg-shifter-black hover:bg-shifter-red hover:animate-pulse"
            @click="cancelEdit"
            >Cancel</a
          >
        </div>
      </div>
    </div>
    <!-- End Edit Openshift Cluster Config -->
  </div>
</template>

<script>
// Pinia Store Imports
import { useModalClusterDelete } from "../stores/configurations/modalClusterDelete";
//import { useConfigurationsClusters } from "../stores/configurations/clusters";
// Plugin & Package Imports
import { mapActions } from "pinia";
export default {
  props: {
    clusterconfig: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      visibleEditConfig: false,
      visibleConfig: false,
    };
  },

  methods: {
    ...mapActions(useModalClusterDelete, { openModal: "openModal" }),
    deleteConfig(clusterConfig) {
      this.openModal(clusterConfig);
    },

    toggleEditConfig() {
      if (this.visibleEditConfig === true) {
        // If Visible Hide
        this.visibleEditConfig = false;
        return;
      }
      this.visibleConfig = false; // Can't show both at the same time
      // Make Visible
      this.visibleEditConfig = true;
    },
    toggleConfig() {
      if (this.visibleConfig === true) {
        // If Visible Hide
        this.visibleConfig = false;
        return;
      }
      this.visibleEditConfig = false; // Can't show both at the same time
      // Make Visible
      this.visibleConfig = true;
    },

    cancelEdit() {
      this.toggleEditConfig();
    },
    commitEdit() {
      alert(this.clusterconfig);
    },
  },
  computed: {},
};
</script>
