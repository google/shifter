<template>
  <ui>
    <li>Shifter Server Status: {{ data.message }}</li>
  </ui>
</template>

<script>
// Pinia Store Imports
import { useShifterV1StatusHealthz } from "../stores/shifter/v1/status/healthz";
// Plugin & Package Imports
import { mapState, mapActions } from "pinia";

export default {
  computed: {
    ...mapState(useShifterV1StatusHealthz, { data: "results" }),
  },

  methods: {
    ...mapActions(useShifterV1StatusHealthz, ["fetchHealthz"]),
  },

  created() {
    // when the template is created, we call this action
    this.fetchHealthz();
  },
};
</script>
