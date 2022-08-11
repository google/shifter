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
import ClusterConfiguration from "../components/configutation-cluster-connections.vue";
import ShifterServerSettingz from "../components/configuration-shifter-server-settingz.vue";
import ModalConfigurationOpenshiftClusterAdd from "./modal-configuration-openshift-cluster-add.vue";
import ModalConfigurationOpenshiftClusterDelete from "./modal-configuration-openshift-cluster-delete.vue";
</script>

<template>
  <div class="grid grid-cols-1 sm:grid-cols-4 place-items-center gap-6 w-full">
    <div
      class="grid place-items-end w-full gap-6 border-r-2 border-shifter-black-soft pr-4"
    >
      <ul class="space-y-4">
        <li
          v-for="(tab, index) in tabs"
          :key="index"
          class="text-lg font-bold hover:animate-pulse cursor cursor-pointer"
          :class="{ 'text-shifter-red-muted': show === index }"
          @click="clickItem(index)"
        >
          {{ tab.title }}
        </li>
      </ul>
    </div>
    <div class="grid col-span-3 place-items-start w-full px-4">
      <transition-group name="fade-up" target="div">
        <div
          class="w-full"
          v-for="(tab, index) in tabs"
          :key="index"
          v-show="show == index"
        >
          <component :is="tab.component"></component>
        </div>
      </transition-group>
    </div>
    <!-- Openshift Cluster Config Add Modal -->
    <ModalConfigurationOpenshiftClusterAdd />
    <!-- Openshift Cluster Config Delete Modal -->
    <ModalConfigurationOpenshiftClusterDelete />
  </div>
</template>

<script>
export default {
  components: {
    ClusterConfiguration,
    ShifterServerSettingz,
  },
  data() {
    return {
      show: 0,
      tabs: [
        {
          title: "Manage Cluster Connections",
          component: "ClusterConfiguration",
        },
        {
          title: "Shifter Server Status",
          component: "ShifterServerSettingz",
        },
      ],
    };
  },
  methods: {
    clickItem(index) {
      this.show = index;
    },
  },
};
</script>

<style scoped>
.section {
  padding: 2em 0;
}
.fade-up-enter-active,
.fade-up-leave-active {
  transition: all 0.4s ease-in-out;
}
.fade-up-enter,
.fade-up-leave-to {
  height: 0;
  transform: translateY(50px);
  opacity: 0;
}
</style>
