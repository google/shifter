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
        <h3 class="text-lg font-bold">
          {{ namespace.metadata.name }}
        </h3>
        <button @click="toggleOpen">
          <svg style="width: 24px; height: 24px" viewBox="0 0 24 24">
            <path
              fill="currentColor"
              d="M11,6H13V14L16.5,10.5L17.92,11.92L12,17.84L6.08,11.92L7.5,10.5L11,14V6M12,22A10,10 0 0,1 2,12A10,10 0 0,1 12,2A10,10 0 0,1 22,12A10,10 0 0,1 12,22M12,20A8,8 0 0,0 20,12A8,8 0 0,0 12,4A8,8 0 0,0 4,12A8,8 0 0,0 12,20Z"
            />
          </svg>
        </button>
      </div>
      <!-- Content -->
      <div class="accordion-content px-5 pt-0 overflow-hidden max-h-0">
        <!-- Deployment Conig Maps -->
        <!--<p class="text-md font-bold pb-2">
          OpenShift Deployment Configurations
        </p>

        <ListConvertDeploymentConfigObjects
          class="pb-4"
          :namespace="namespace.metadata.name"
          :visible="isOpen"
        />-->
      </div>
    </div>
  </div>
</template>

<script>
// Pinia Store Imports
//import { useOSProjects } from "../stores/openshift/projects";
// Plugin & Package Imports
/*import { mapState } from "pinia";

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
        //header.querySelector(".fas").classList.remove("fa-plus");
        //header.querySelector(".fas").classList.add("fa-minus");
        //header.parentElement.classList.add("bg-indigo-50");
        this.isOpen = true;
      } else {
        accordionContent.style.maxHeight = `0px`;
        //header.querySelector(".fas").classList.add("fa-plus");
        //header.querySelector(".fas").classList.remove("fa-minus");
        //header.parentElement.classList.remove("bg-indigo-50");
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
};*/
</script>
