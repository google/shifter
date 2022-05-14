<script setup>
//import { useOSProjects } from "../stores/openshift/projects";
//import { useOSDeploymentConfigs } from "../stores/openshift/deployment-configs";

import { useConfigurationsClusters } from "../stores/configurations/clusters";
import { useConvertObjects } from "../stores/convert/convert";
import FormTableConvertObjects from "../components/form-table-convert-objects.vue";
import FormTableConvertObjectsReview from "../components/form-table-convert-objects-review.vue";
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

      <!-- STEP 3 OBJECT Review -->
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
        <a
          v-show="currentStep === maxSteps"
          class="uppercase rounded px-6 py-2 bg-shifter-black hover:bg-shifter-red hover:animate-pulse"
          :onclick="convertStep"
          >Convert</a
        >
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from "pinia";
//import { shifterConfig } from "@/main";
//import axios from "axios";

export default {
  setup() {
    //const convertObjects = useConvertObjects();
    //const { items } = storeToRefs(convertObjects);
    //return {
    //  items,
    //};
  },
  data() {
    return {
      clusterId: "",
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
          enabled: true,
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
    ...mapState(useConfigurationsClusters, {
      getSelectedCluster: "getCluster",
    }),
    ...mapState(useConvertObjects, { dcSelected: "selected" }),

    activeSteps() {
      return this.convertSteps.filter((step) => step.enabled);
    },
    maxSteps() {
      return this.convertSteps.filter((step) => step.enabled).length;
    },

    /*selectedCluster() {
      return this.getSelectedCluster(1);
    },*/

    /*selectedDeploymentConfigs() {
      return this.dcSelected;
    },*/
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

    //async convertStep() {
    /*const config = {
        method: "post",
        url: shifterConfig.API_BASE_URL + "/shifter/convert/",
        headers: {},
        data: {
          shifter: JSON.parse(
            JSON.stringify({ ...this.getSelectedCluster(0).shifter })
          ),
          items: JSON.parse(JSON.stringify([...this.dcSelected])),
        },
      };
      this.fetching = true;
      try {
        const response = await axios(config);
        try {
          console.log(response);
          //this.data = response.data.deploymentConfigs.items;
          this.fetching = false;
        } catch (err) {
          this.data = [];
          console.error(
            "Error Fetching OpenShift Object from Shifter Server.",
            err
          );
          this.fetching = false;
          return err;
        }
      } catch (err) {
        this.data = {
          message: "Shifter Server Unreachable, Timeout",
        };
        this.fetching = false;
      }*/
    // },
  },

  created() {
    if (this.configurationClusters.length >= 1) {
      this.clusterId = this.configurationClusters[0].id;
      this.changeCluster();
    }
  },
};
</script>
