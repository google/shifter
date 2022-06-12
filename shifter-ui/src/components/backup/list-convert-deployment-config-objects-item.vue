<template>
  <div class="container flex-row mx-auto justify-center bg-shifter-black">
    <!--{{ namespace }}-->

    <!-- What is term -->
    <div>
      <!-- header -->
      <div
        class="accordion-header flex flex-row mx-auto space-x-5 px-5 items-center h-16"
      >
        <div>
          <h3 class="text-md font-semi-bold">
            {{ deploymentconfig.metadata.name }}
          </h3>
        </div>
        <div class="flex justify-center">
          <a
            @click="openModal(deploymentconfig)"
            class="rounded bg-shifter-red-soft px-6 my-1 hover:bg-shifter-red hover:animate-pulse"
            >View</a
          >
        </div>
        <div class="flex justify-center">
          <a
            v-if="isSelected(deploymentconfig)"
            @click="deploymentConfigRemove(deploymentconfig)"
            class="rounded border border-shifter-red-soft bg-shifter-red-soft px-6 my-1"
            >Remove</a
          >
          <a
            v-else
            @click="deploymentConfigAdd(deploymentconfig)"
            class="rounded border border-shifter-red-soft px-6 my-1 hover:bg-shifter-red-soft hover:animate-pulse"
            >Select</a
          >
        </div>

        <!--<h3 @click="toggleOpen">{{ namespace.metadata.name }}</h3>-->
      </div>
      <!-- Content -->
      <div class="accordion-content px-5 pt-0 overflow-hidden max-h-0">
        <p class="leading-6 font-light pl-9 text-justify">
          <!-- {{ namespace.metadata }}-->
        </p>
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
import { useJSONModal } from "../stores/convert/jsonModal";
import { useConvertObjects } from "../stores/convert/convert";
// Plugin & Package Imports
import { mapActions, mapState } from "pinia";
export default {
  props: {
    deploymentconfig: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {};
  },

  methods: {
    ...mapActions(useConvertObjects, { deploymentConfigRemove: "remove" }),
    ...mapActions(useConvertObjects, { deploymentConfigAdd: "add" }),
    ...mapActions(useJSONModal, { openModal: "openModal" }),
  },
  computed: {
    ...mapState(useConvertObjects, {
      isSelected: "isSelected",
    }),
  },
};
</script>

getByUid
