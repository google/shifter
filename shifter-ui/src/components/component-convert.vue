<script setup>
// Vue Component Imports
import FormTableConvertObjects from "../components/form-table-convert-objects.vue";
import FormTableConvertObjectsReview from "../components/form-table-convert-objects-review.vue";
import FormTableJSONModal from "../components/form-table-convert-json-modal.vue";
</script>
<template>
  <div class="container flex mx-auto m-6 items-center">
    <div
      class="container flex-row mx-auto bg-shifter-black-mute justify-center rounded-2xl py-6"
    >
      <div
        class="container flex mx-auto justify-center py-6 gap-8 uppercase pb-12"
      >
        <div
          v-for="step in activeSteps"
          :key="step.id"
          class="container flex justify-center items-center"
          :class="
            step.id === currentStep ? 'animate-pulse text-bold' : 'animate-none'
          "
        >
          <div
            class="flex rounded-full shadow-2xl bg-shifter-black px-4 py-2 text-shifter-red-muted text-bold no-underline"
            :class="step.id === currentStep ? 'no-underline' : 'no-underline'"
          >
            {{ step.id }}
          </div>
          <div
            class="flex ml-4"
            :class="
              step.id === currentStep
                ? 'animate-pulse underline decoration-4 underline-offset-4 text-bold'
                : 'animate-none'
            "
          >
            {{ step.title }}
          </div>
        </div>
      </div>
      <!-- STEP 1 CLUSTER SELECTION -->
      <div
        v-show="currentStep === 1"
        class="container flex-row mx-auto justify-center py-12"
      >
        <div class="container flex-row justify-center items-center">
          <div class="flex justify-center bold text-4xl m-2">
            Cluster Selection
          </div>
          <div class="flex justify-center text-baseline m-2">
            Select OpenShift cluster from which you would like to convert
            workloads
          </div>
        </div>
        <div class="container flex mx-auto justify-center my-4">
          <select
            class="flex justify-center w-1/4 p-2 m-2 bg-shifter-black rounded"
            id="cluster"
            @change="changeCluster($event)"
            v-model="clusterId"
          >
            <option
              v-for="cluster in configurationClusters"
              :key="cluster.id"
              :value="cluster.id"
            >
              {{ cluster.shifter.clusterConfig.connectionName }}
            </option>
          </select>
        </div>
      </div>
      <!-- END STEP 1 CLUSTER SELECTION -->

      <!-- STEP 2 OBJECT SELECTION -->
      <div
        v-show="currentStep == 2"
        class="container flex-row mx-auto justify-center py-12"
      >
        <div class="container flex-row justify-center items-center">
          <div class="flex justify-center bold text-4xl m-2">
            Object Selection
          </div>
          <div class="flex justify-center text-baseline m-2">
            Select deployment configurations to convert for migration.
          </div>
        </div>
        <div class="container flex mx-auto justify-center my-4">
          <FormTableConvertObjects />
        </div>
      </div>
      <!-- END STEP 2 OBJECT SELECTION -->

      <!-- STEP 3 OBJECT REVIEW -->
      <div
        v-show="currentStep == 3"
        class="container flex-row mx-auto justify-center py-12"
      >
        <div class="container flex-row justify-center items-center">
          <div class="flex justify-center bold text-4xl m-2">Object Review</div>
          <div class="flex justify-center text-baseline m-2">
            Review configurations selected for conversion.
          </div>
        </div>
        <div class="container flex mx-auto justify-center my-4">
          <FormTableConvertObjectsReview />
        </div>
      </div>
      <!-- END STEP 3 OBJECT REVIEW -->

      <!-- STEP 4 REVIEW -->
      <div
        v-show="currentStep == 4"
        class="container flex-row mx-auto justify-center py-12"
      >
        <div class="container flex-row justify-center items-center">
          <div class="flex justify-center bold text-4xl m-2">Convert</div>
          <div class="flex justify-center text-baseline m-2">
            Convert your selected workloads for GKE & Anthos.
          </div>
        </div>
        <div class="container flex mx-auto justify-center my-4">
          <a
            class="uppercase rounded px-6 py-2 bg-shifter-red-soft hover:bg-shifter-red animate-pulse"
            :onclick="convertStep"
            >Convert</a
          >
        </div>
      </div>
      <!-- END STEP 4 REVIEW -->

      <div class="container flex mx-auto justify-end px-10 gap-4">
        <a
          v-show="currentStep > 1"
          class="uppercase rounded px-6 py-2 bg-shifter-black hover:bg-shifter-red hover:animate-pulse"
          :onclick="previousStep"
          >Previous</a
        >
        <a
          v-show="currentStep < maxSteps"
          class="uppercase rounded px-6 py-2 bg-shifter-black hover:bg-shifter-red hover:animate-pulse"
          :onclick="nextStep"
          >Next</a
        >
      </div>
    </div>
    <!-- Deployment Config JSON Modal -->
    <FormTableJSONModal />
  </div>
</template>

<script>
// Pinia Store Imports
import { useConfigurationsClusters } from "../stores/configurations/clusters";
import { useConvertObjects } from "../stores/convert/convert";
import { useJSONModal } from "../stores/convert/jsonModal";
// Plugin & Package Imports
import { mapState, mapActions } from "pinia";

export default {
  data() {
    return {
      clusterId: "",
      currentStep: 1,
      convertSteps: [
        {
          id: 1,
          title: "OpenShift Cluster",
          enabled: true,
        },
        {
          id: 2,
          title: "Selections",
          enabled: true,
        },
        {
          id: 3,
          title: "Options",
          enabled: true,
        },
        {
          id: 4,
          title: "Convert",
          enabled: true,
        },
        {
          id: 5,
          title: "Results",
          enabled: false,
        },
      ],
    };
  },
  computed: {
    ...mapState(useConfigurationsClusters, {
      configurationClusters: "getActiveClusters",
    }),
    ...mapState(useConfigurationsClusters, {
      getSelectedCluster: "getCluster",
    }),
    ...mapState(useJSONModal, {
      showJSONModal: "showJSONModal",
    }),
    ...mapState(useConvertObjects, { dcSelected: "selected" }),

    activeSteps() {
      return this.convertSteps.filter((step) => step.enabled);
    },
    maxSteps() {
      return this.convertSteps.filter((step) => step.enabled).length;
    },
  },

  methods: {
    ...mapActions(useConvertObjects, { convertStep: "convert" }),
    ...mapActions(useConvertObjects, { setCluster: "setCluster" }),

    nextStep() {
      if (this.currentStep < this.maxSteps) {
        this.currentStep++;
      }
    },
    previousStep() {
      if (this.currentStep > 1) {
        this.currentStep--;
      }
    },

    changeCluster() {
      this.setCluster(this.clusterId);
    },
  },

  created() {
    if (this.configurationClusters.length >= 1) {
      this.clusterId = this.configurationClusters[0].id;
      this.changeCluster();
    }
  },
};
</script>
