<script setup>
// Vue Component Imports
import OpenshiftDeploymentConfigList from "./openshift-deploymentconfig-list.vue";
</script>
<template>
  <div
    class="container mb-6 border rounded-xl bg-shifter-black-soft overflow-hidden"
  >
    <div class="flex flex-row items-center my-4">
      <!-- Title -->
      <div class="container flex">
        <p class="text-xl font-bold mx-6">{{ namespace.metadata.name }}</p>
      </div>
      <!-- Actions -->
      <div class="container flex flex-row-reverse gap-3 mx-6">
        <!-- Action: Expand Section -->
        <div class="flex">
          <!-- Material Design - SVG - plus-circle-outline -->
          <svg
            style="width: 24px; height: 24px"
            viewBox="0 0 24 24"
            @click="openSection"
            v-show="!isOpen"
          >
            <path
              fill="currentColor"
              d="M12,20C7.59,20 4,16.41 4,12C4,7.59 7.59,4 12,4C16.41,4 20,7.59 20,12C20,16.41 16.41,20 12,20M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M13,7H11V11H7V13H11V17H13V13H17V11H13V7Z"
            />
          </svg>
          <!-- Material Design - SVG - minus-circle -->
          <svg
            style="width: 24px; height: 24px"
            viewBox="0 0 24 24"
            @click="closeSection"
            v-show="isOpen"
          >
            <path
              fill="currentColor"
              d="M17,13H7V11H17M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z"
            />
          </svg>
        </div>
        <!-- End Action -->
        <!-- Action: Select All -->
        <div class="flex">
          <!-- Material Design - SVG - checkbox-multiple-marked-circle-outline -->
          <svg
            style="width: 24px; height: 24px"
            viewBox="0 0 24 24"
            @click="selectAll"
          >
            <path
              fill="currentColor"
              d="M14,2A8,8 0 0,0 6,10A8,8 0 0,0 14,18A8,8 0 0,0 22,10H20C20,13.32 17.32,16 14,16A6,6 0 0,1 8,10A6,6 0 0,1 14,4C14.43,4 14.86,4.05 15.27,4.14L16.88,2.54C15.96,2.18 15,2 14,2M20.59,3.58L14,10.17L11.62,7.79L10.21,9.21L14,13L22,5M4.93,5.82C3.08,7.34 2,9.61 2,12A8,8 0 0,0 10,20C10.64,20 11.27,19.92 11.88,19.77C10.12,19.38 8.5,18.5 7.17,17.29C5.22,16.25 4,14.21 4,12C4,11.7 4.03,11.41 4.07,11.11C4.03,10.74 4,10.37 4,10C4,8.56 4.32,7.13 4.93,5.82Z"
            />
          </svg>
        </div>
        <!-- End Action -->
      </div>
    </div>
    <!-- Start Deployment Configs -->
    <div
      class="container flex flex-col bg-shifter-black-soft overflow-hidden"
      v-show="isOpen"
    >
      <OpenshiftDeploymentConfigList
        :namespace="namespace.metadata.name"
      ></OpenshiftDeploymentConfigList>
    </div>
    <!-- End Deployment Configs -->
  </div>
</template>

<script>
// Pinia Store Imports
import { useOSProjects } from "../stores/openshift/projects";
// Plugin & Package Imports
import { mapState } from "pinia";

export default {
  props: {
    uid: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      isOpen: false,
    };
  },

  methods: {
    openSection() {
      this.isOpen = true;
    },
    closeSection() {
      this.isOpen = false;
    },
    selectAll() {
      alert("Selecting All");
    },
  },

  computed: {
    ...mapState(useOSProjects, { getByUid: "getByUid" }),

    namespace() {
      return this.getByUid(this.uid);
    },
  },
};
</script>
