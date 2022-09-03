<!--
 Copyright 2022 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

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
