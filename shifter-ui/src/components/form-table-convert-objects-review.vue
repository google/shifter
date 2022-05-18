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
        <tr v-for="dc in all" :key="dc.deploymentConfig.metadata.uid">
          <td>
            {{ dc.deploymentConfig.metadata.namespace }}
          </td>
          <td>
            {{ dc.deploymentConfig.metadata.name }}
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
                @click="dcRemove(dc)"
                class="rounded bg-shifter-red-soft px-6 my-1"
                >Remove</a
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
import { useConvertObjects } from "../stores/convert/convert";
// Plugin & Package Imports
import { mapState, mapActions } from "pinia";

export default {
  methods: {
    ...mapActions(useConvertObjects, { dcRemove: "remove" }),
    ...mapActions(useConvertObjects, { dcAdd: "add" }),
  },
  computed: {
    ...mapState(useConvertObjects, {
      isSelected: "isSelected",
    }),
    ...mapState(useConvertObjects, { all: "all" }),
  },
};
</script>
