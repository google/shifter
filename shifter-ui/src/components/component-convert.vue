<script setup>
import { useConfigurationsClusters } from "../stores/configurations/clusters";
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
        >
          <div
            class="flex rounded-full shadow-2xl bg-shifter-black px-4 py-2 text-shifter-red-muted text-bold"
          >
            {{ step.id }}
          </div>
          <div class="flex ml-4">{{ step.title }}</div>
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
          >
            <option
              v-for="cluster in configurationClusters"
              :key="cluster.id"
              :value="cluster.shifter"
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
        <div class="container flex mx-auto justify-center my-4"></div>
      </div>
      <!-- END STEP 2 OBJECT SELECTION -->

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
  </div>
</template>

<script>
import { mapState } from "pinia";

export default {
  data() {
    return {
      currentStep: 1,
      convert: {
        shifter: {},
      },
      convertSteps: [
        {
          id: 1,
          title: "Cluster Selection",
          enabled: true,
        },
        {
          id: 2,
          title: "Object Selection",
          enabled: true,
        },
        {
          id: 3,
          title: "Review",
          enabled: false,
        },
        {
          id: 4,
          title: "Shift Workflows",
          enabled: false,
        },
        {
          id: 5,
          title: "Summary",
          enabled: false,
        },
      ],
    };
  },
  computed: {
    ...mapState(useConfigurationsClusters, {
      configurationClusters: "getActiveClusters",
    }),

    activeSteps() {
      return this.convertSteps.filter((step) => step.enabled);
    },
    maxSteps() {
      return this.convertSteps.filter((step) => step.enabled).length;
    },
  },

  methods: {
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

    //...mapActions(useConfigurationsClusters, ['fetchClusters']),
  },

  created() {
    //this.fetchClusters();
  },
};
</script>
