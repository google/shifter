<template>
  <div class="container flex-row mx-auto justify-center">
    <table class="container table-auto">
      <thead class="uppercase text-shifter-red-soft bg-shifter-black text-lg">
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
                class="rounded bg-shifter-red-soft px-6 my-1 hover:bg-shifter-red hover:animate-pulse"
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
// Pinia Store Imports
import { useOSDeploymentConfigs } from "../stores/openshift/deployment-configs";
import { useConvertObjects } from "../stores/convert/convert";
import { useJSONModal } from "../stores/convert/jsonModal";
// Plugin & Package Imports
import { mapActions, mapState } from "pinia";

export default {
  methods: {
    ...mapActions(useConvertObjects, { dcRemove: "remove" }),
    ...mapActions(useConvertObjects, { dcAdd: "add" }),
    ...mapActions(useJSONModal, { openModal: "openModal" }),
  },

  computed: {
    ...mapState(useConvertObjects, {
      isSelected: "isSelected",
    }),
    ...mapState(useOSDeploymentConfigs, { all: "all" }),
    loadedDeploymentConfigs() {
      return this.osDeploymentConfigs;
    },
  },
};
</script>
