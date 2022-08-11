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

<script setup>
// Vue Component Imports
</script>
<template>
  <div class="container">
    <div class="flex flex-row items-center my-1">
      <!-- Title -->
      <div class="container flex">
        <p class="text-sm italic mx-6">
	  {{ resource.Kind }}
          {{ resource.Name }}
        </p>
      </div>
      <!-- Actions -->
      <div class="container flex flex-row-reverse gap-3 mx-6">
        <!-- Action: Show Resource -->
        <div class="flex">
          <!-- Material Design - SVG - check-circle -->
          <svg
            style="width: 24px; height: 24px"
            viewBox="0 0 24 24"
            @click="openModal(resource)"
          >
            <path
              fill="currentColor"
              d="M12 4.5C7 4.5 2.7 7.6 1 12C2.7 16.4 7 19.5 12 19.5H13.1C13 19.2 13 18.9 13 18.5C13 17.9 13.1 17.4 13.2 16.8C12.8 16.9 12.4 17 12 17C9.2 17 7 14.8 7 12S9.2 7 12 7 17 9.2 17 12C17 12.3 17 12.6 16.9 12.9C17.6 12.7 18.3 12.5 19 12.5C20.2 12.5 21.3 12.8 22.3 13.5C22.6 13 22.8 12.5 23 12C21.3 7.6 17 4.5 12 4.5M12 9C10.3 9 9 10.3 9 12S10.3 15 12 15 15 13.7 15 12 13.7 9 12 9M19 21V19H15V17H19V15L22 18L19 21"
            />
          </svg>
        </div>
        <!-- End Action -->
        <!-- Action: Select/Unselect Resource -->
        <div class="flex">
          <!-- Material Design - SVG - check-circle -->
          <svg
            style="width: 24px; height: 24px"
            viewBox="0 0 24 24"
            v-if="isSelected(resource)"
            @click="resourceRemove(resource)"
          >
            <path
              fill="currentColor"
              d="M12 2C6.5 2 2 6.5 2 12S6.5 22 12 22 22 17.5 22 12 17.5 2 12 2M10 17L5 12L6.41 10.59L10 14.17L17.59 6.58L19 8L10 17Z"
            />
          </svg>
          <!-- Material Design - SVG - check-circle-outline -->
          <svg
            style="width: 24px; height: 24px"
            viewBox="0 0 24 24"
            v-else
            @click="resourceAdd(resource)"
          >
            <path
              fill="currentColor"
              d="M12 2C6.5 2 2 6.5 2 12S6.5 22 12 22 22 17.5 22 12 17.5 2 12 2M12 20C7.59 20 4 16.41 4 12S7.59 4 12 4 20 7.59 20 12 16.41 20 12 20M16.59 7.58L10 14.17L7.41 11.59L6 13L10 17L18 9L16.59 7.58Z"
            />
          </svg>
        </div>
        <!-- End Action -->
      </div>
    </div>
  </div>
  <div class="container bg-shifter-black-soft"></div>
</template>

<script>
// Pinia Store Imports
import { useJSONModal } from "../stores/convert/modalJSON";
//import { useConvertObjects } from "../stores/convert/convert";
import { useConvertObjects } from "../stores/convert/convertv2";
// Plugin & Package Imports
import { mapActions, mapState } from "pinia";
export default {
  props: {
    resource: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {};
  },

  methods: {
    ...mapActions(useConvertObjects, { resourceRemove: "remove" }),
    ...mapActions(useConvertObjects, { resourceAdd: "add" }),
    ...mapActions(useJSONModal, { openModal: "openModal" }),
  },
  computed: {
    ...mapState(useConvertObjects, {
      isSelected: "isSelected",
    }),
  },
};
</script>
