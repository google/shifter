<script setup></script>

<template>
  <div class="container flex-row mx-auto justify-center">
    <table class="container table-auto">
      <thead class="uppercase text-shifter-red-soft bg-shifter-black text-lg">
        <!--  class="justify-center items-center uppercase text-shifter-red-soft bg-shifter-black text-lg" -->
        <tr>
          <th>Namespace/Project</th>
          <th>Deployment Name</th>
          <th>View</th>
          <th>Select</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="dc in all" :key="dc.metadata.uid">
          <td>
            {{ dc.metadata.namespace }}
          </td>
          <td>
            {{ dc.metadata.name }}
          </td>
          <td>
            <div class="flex justify-center">
              <a
                @click="openModal(dc)"
                class="
                  rounded
                  bg-shifter-red-soft
                  px-6
                  my-1
                  hover:bg-shifter-red hover:animate-pulse
                "
                >View</a
              >
            </div>
          </td>
          <td>
            <div class="flex justify-center">
              <a
                v-if="isSelected(dc)"
                @click="dcRemove(dc)"
                class="rounded bg-shifter-red-soft px-6 my-1"
                >Remove</a
              >
              <a
                v-else
                @click="dcAdd(dc)"
                class="
                  rounded
                  border border-shifter-red-soft
                  px-6
                  my-1
                  hover:bg-shifter-red-soft hover:animate-pulse
                "
                >Select</a
              >
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
//import { shifterConfig } from "@/main";
import { mapActions, mapState } from "pinia";
//import { useOSProjects } from "../stores/openshift/projects";
import { useOSDeploymentConfigs } from "../stores/openshift/deployment-configs";
import { useConvertObjects } from "../stores/convert/convert";
import { useJSONModal } from "../stores/convert/jsonModal";
//import axios from "axios";

// API Endpoint Configuration
export default {
  setup() {
    //const oSProjects = useOSProjects();
    //const oSDeploymentConfigs = useOSDeploymentConfigs();
    return {
      // oSProjects,
      // oSDeploymentConfigs,
    };
  },

  data() {
    return {
      //osDeploymentConfigs: null,
      //osProjects: null,
      //fetching: false,
    };
  },
  methods: {
    ...mapActions(useConvertObjects, { dcRemove: "remove" }),
    ...mapActions(useConvertObjects, { dcAdd: "add" }),
    ...mapActions(useJSONModal, { openModal: "openModal" }),
    /* async fetchDeploymentConfigs() {
      const config = {
        method: "post",
        url: shifterConfig.API_BASE_URL + "/openshift/deploymentconfigs/",
        headers: {},
        data: { ...this.selectedCluster },
      };

      this.fetching = true;
      try {
        const response = await axios(config);
        try {
          this.osDeploymentConfigs = response.data.deploymentConfigs.items;
        } catch (err) {
          this.osDeploymentConfigs = [];
          console.error(
            "Error Fetching OpenShift Object from Shifter Server.",
            err
          );
          return err;
        }
      } catch (err) {
        this.osDeploymentConfigs = {
          message: "Shifter Server Unreachable, Timeout",
        };
      }
      this.fetching = false;
    },

    async fetchProjects() {
      const config = {
        method: "post",
        url: shifterConfig.API_BASE_URL + "/openshift/projects/",
        headers: {},
        data: { ...this.selectedCluster },
      };

      this.fetching = true;
      try {
        const response = await axios(config);
        try {
          this.osProjects = response.data.projects.items;
        } catch (err) {
          this.osProjects = [];
          console.error(
            "Error Fetching OpenShift Object from Shifter Server.",
            err
          );
          return err;
        }
      } catch (err) {
        this.osProjects = {
          message: "Shifter Server Unreachable, Timeout",
        };
      }
      this.fetching = false;
    },*/
  },

  computed: {
    /*...mapState(useConfigurationsClusters, {
      getSelectedCluster: "getCluster",
    }),*/
    ...mapState(useConvertObjects, {
      isSelected: "isSelected",
    }),

    ...mapState(useOSDeploymentConfigs, { all: "all" }),

    // ...mapState(useConvertObjects, { dcSelected: "selected" }),

    /*selectedCluster() {
      return this.getSelectedCluster(1);
    },*/
    loadedDeploymentConfigs() {
      return this.osDeploymentConfigs;
    },
  },

  created() {
    //this.fetchDeploymentConfigs();
    //this.fetchProjects();
    //oSProjects;
    //this.oSDeploymentConfigs.fetch();
  },
};
</script>
