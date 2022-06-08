<script setup>
// Vue Component Imports
import ListConvertDeploymentConfigObjects from "./list-convert-deployment-config-objects.vue";
</script>
<template>
  <div
    class="container flex-row mx-auto justify-center bg-shifter-black-soft rounded"
  >
    <!--{{ namespace }}-->

    <!-- What is term -->
    <div class="transition hover:bg-indigo-50">
      <!-- header -->
      <div
        @click="toggleOpen"
        class="accordion-header cursor-pointer transition flex space-x-5 items-center h-10"
      >
        <i class="fas fa-plus"></i>
        <h3 class="text-lg font-bold" @click="toggleOpen">
          {{ namespace.metadata.name }}
        </h3>
      </div>
      <!-- Content -->
      <div class="accordion-content px-5 pt-0 overflow-hidden max-h-0">
        <!-- Deployment Conig Maps -->
        <p class="text-md font-bold pb-2">
          OpenShift Deployment Configurations
        </p>

        <ListConvertDeploymentConfigObjects
          class="pb-4"
          :namespace="namespace.metadata.name"
          :visible="isOpen"
        />

        <!-- Config Maps -->
        <!--<p class="text-md font-bold pb-2">OpenShift Config Maps</p>

        <ListConvertDeploymentConfigObjects
          class="pb-4"
          :namespace="namespace.metadata.name"
          :visible="isOpen"
        />-->

        <!--<button
          class="rounded-full bg-indigo-600 text-white font-medium font-lg px-6 py-2 my-5 ml-9"
        >
          Learn more
        </button>-->
      </div>
    </div>
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
    toggleOpen(event) {
      const header = event.target;
      const accordionContent =
        header.parentElement.querySelector(".accordion-content");
      let accordionMaxHeight = accordionContent.style.maxHeight;

      // Condition handling
      if (accordionMaxHeight == "0px" || accordionMaxHeight.length == 0) {
        accordionContent.style.maxHeight = `${
          accordionContent.scrollHeight + 32
        }px`;
        header.querySelector(".fas").classList.remove("fa-plus");
        header.querySelector(".fas").classList.add("fa-minus");
        header.parentElement.classList.add("bg-indigo-50");
        this.isOpen = true;
      } else {
        accordionContent.style.maxHeight = `0px`;
        header.querySelector(".fas").classList.add("fa-plus");
        header.querySelector(".fas").classList.remove("fa-minus");
        header.parentElement.classList.remove("bg-indigo-50");
        this.isOpen = false;
      }
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

getByUid
