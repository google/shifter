<script setup>
import { useConfigurationsClusters } from "../stores/configurations/clusters";
import { useConvertObjects } from "../stores/convert/convert-objects";
</script>

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
        <tr v-for="dc in selectedDeploymentConfigs" :key="dc.metadata.uid">
          <td>
            {{ dc.metadata.namespace }}
          </td>
          <td>
            {{ dc.metadata.name }}
          </td>
          <td>
            <div class="flex justify-center">
              <a
                class="rounded bg-shifter-red-soft px-6 my-1 hover:bg-shifter-red hover:animate-pulse"
                >View</a
              >
            </div>
          </td>
          <td>
            <div class="flex justify-center">
              <a
                v-if="dcIsSelected(dc)"
                @click="dcRemove(dc)"
                class="rounded bg-shifter-red-soft px-6 my-1"
                >Remove</a
              >
              <a
                v-else
                @click="dcAdd(dc)"
                class="rounded border border-shifter-red-soft px-6 my-1 hover:bg-shifter-red-soft hover:animate-pulse"
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
import { shifterConfig } from "@/main";
import { mapState, mapActions } from "pinia";
import axios from "axios";

// API Endpoint Configuration
export default {
  data() {
    return {
      data: null,
      fetching: false,
    };
  },
  methods: {
    ...mapActions(useConvertObjects, { dcRemove: "removeItem" }),
    ...mapActions(useConvertObjects, { dcAdd: "addItem" }),

    async fetchObjects() {
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
          this.data = response.data.deploymentConfigs.items;
        } catch (err) {
          this.data = [];
          console.error(
            "Error Fetching OpenShift Object from Shifter Server.",
            err
          );
          return err;
        }
      } catch (err) {
        this.data = {
          message: "Shifter Server Unreachable, Timeout",
        };
      }
      this.fetching = false;
    },
  },

  computed: {
    ...mapState(useConfigurationsClusters, {
      getSelectedCluster: "getCluster",
    }),
    ...mapState(useConvertObjects, { dcIsSelected: "contains" }),
    ...mapState(useConvertObjects, { dcSelected: "selected" }),

    selectedDeploymentConfigs() {
      return this.dcSelected;
    },

    selectedCluster() {
      return this.getSelectedCluster(1);
    },
    deploymentConfigs() {
      return this.data;
    },
  },

  mounted() {
    this.fetchObjects();
  },
};
</script>
