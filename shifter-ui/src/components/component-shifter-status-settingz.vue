<template>
  <main>
    <ui>
      <li>Shifter Server Port: {{ data.runningPort }}</li>
      <li>Shifter Server Host: {{ data.runningHost }}</li>
      <li>
        Shifter Server Storage Type: {{ data.storageDescription }} ({{
          data.storageType
        }})
      </li>
    </ui>
  </main>
</template>

<script>
// Pinia Store Imports
import { useShifterV1StatusSettingz } from "../stores/shifter/v1/status/settingz";
// Plugin & Package Imports
import { mapState, mapActions } from "pinia";

export default {
  computed: {
    ...mapState(useShifterV1StatusSettingz, { data: "results" }),
  },
  methods: {
    ...mapActions(useShifterV1StatusSettingz, ["fetchSettingz"]),
  },
  created() {
    // when the template is created, we call this action
    this.fetchSettingz();
  },
};
</script>
