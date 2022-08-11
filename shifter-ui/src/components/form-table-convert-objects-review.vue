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
  <div class="container flex-row mx-auto justify-center">
    <table class="container table-auto">
      <thead class="uppercase text-shifter-red-soft bg-shifter-black text-lg">
        <tr>
          <th>Namespace/Project</th>
		  <th>Resource Kind</th>
          <th>Resource Name</th>
          <!--<th>View</th>-->
          <th>Select</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="r in all" :key="r.resource.UID">
          <td>
            {{ r.resource.Namespace }}
          </td>
		  <td>
			{{ r.resource.Kind }}
          </td>
          <td>
            {{ r.resource.Name }}
          </td>
          <!--<td>
            <div class="flex justify-center">
              <a
                class="rounded bg-shifter-red-soft px-6 my-1 hover:bg-shifter-red hover:animate-pulse"
                >View</a
              >
            </div>
          </td>-->
          <td>
            <div class="flex justify-center">
              <a
                @click="resRemove(r)"
                class="rounded border border-shifter-red-soft bg-shifter-red-soft px-6 my-1"
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
import { useConvertObjects } from "../stores/convert/convertv2";
// Plugin & Package Imports
import { mapState, mapActions } from "pinia";

export default {
  methods: {
    ...mapActions(useConvertObjects, { resRemove: "remove" }),
    ...mapActions(useConvertObjects, { resAdd: "add" }),
  },
  computed: {
    ...mapState(useConvertObjects, {
      isSelected: "isSelected",
    }),
    ...mapState(useConvertObjects, { all: "all" }),
  },
};
</script>
