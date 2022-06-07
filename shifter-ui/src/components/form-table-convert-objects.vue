<template>
  <div class="container flex-row mx-auto justify-center">
    <div class="container flex-row mx-auto bg-shifter-black">
      <div>
        <a
          v-show="ifPreviousItems"
          @click="previousItems()"
          class="rounded border border-shifter-red-soft px-6 my-1 hover:bg-shifter-red-soft hover:animate-pulse"
          >Previous</a
        >{{ itemsFrom }}-{{ itemsTo }} of {{ itemsTotal }}
        <a
          v-show="ifNextItems"
          @click="nextItems()"
          class="rounded border border-shifter-red-soft px-6 my-1 hover:bg-shifter-red-soft hover:animate-pulse"
          >Next</a
        >
      </div>
    </div>
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
        <tr v-for="dc in items" :key="dc.metadata.uid">
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
                class="rounded border border-shifter-red-soft bg-shifter-red-soft px-6 my-1"
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
  data() {
    return {
      pagination: {
        from: 1,
        to: 0,
        max: 8,
      },
    };
  },

  methods: {
    ...mapActions(useConvertObjects, { dcRemove: "remove" }),
    ...mapActions(useConvertObjects, { dcAdd: "add" }),
    ...mapActions(useJSONModal, { openModal: "openModal" }),

    nextItems() {
      this.pagination.from = this.pagination.from + this.itemsMax;
    },
    previousItems() {
      this.pagination.from = this.pagination.from - this.itemsMax;
    },
  },

  computed: {
    ...mapState(useConvertObjects, {
      isSelected: "isSelected",
    }),
    ...mapState(useOSDeploymentConfigs, { all: "all" }),
    loadedDeploymentConfigs() {
      return this.osDeploymentConfigs;
    },

    items() {
      return this.all.slice(this.itemsFrom - 1, this.itemsTo);
    },

    itemsMax() {
      return this.pagination.max;
    },

    itemsFrom() {
      return this.pagination.from;
    },

    itemsTo() {
      if (this.all.length <= this.itemsFrom + this.itemsMax) {
        return this.all.length;
      }
      return this.itemsFrom + this.itemsMax - 1;
    },

    itemsTotal() {
      return this.all.length;
    },

    ifPreviousItems() {
      if (this.itemsFrom >= this.itemsMax && this.itemsFrom >= 0) {
        return true;
      }
      return false;
    },
    ifNextItems() {
      if (this.itemsTotal > this.itemsTo) {
        return true;
      }
      return false;
    },
  },
};
</script>
