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
  <div
    class="container flex-row absolute items-center w-full h-full bg-shifter-black-soft border rounded-2xl overflow-y-auto"
    :class="showJSONModal ? 'visible' : 'invisible'"
  >
    <div class="container flex-row mx-auto backdrop-blur-sm text-xl">
      <div
        class="container flex mx-auto justify-end p-4 bg-shifter-black-mute gap-4"
      >
      <div class="container flex font-bold hover:animate-pulse cursor cursor-pointer text-shifter-white-soft justify-left">Resource Source</div>
        <a @click="closeModal" class="rounded bg-shifter-red-soft rounded border border-shifter-red-soft px-6 my-1 hover:bg-shifter-red-soft over:animate-pulse">Close</a>
      </div>
    </div>
    <div class="container flex mx-auto justify-end px-4 py-10">
      <div class="container flex flex-row mx-auto">
        <div class="container flex w-3/4 bg-shifter-black-soft text-sm">
          <pre v-html="content"></pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// Pinia Store Imports
import { useJSONModal } from "../stores/convert/modalJSON";
// Plugin & Package Imports
import { mapState, mapActions } from "pinia";

export default {
  methods: {
    ...mapActions(useJSONModal, { closeModal: "closeModal" }),
    ...mapActions(useJSONModal, { openModal: "openModal" }),

    replacer(match, pIndent, pKey, pVal, pEnd) {
      // JSON Key
      var key = "<span class=text-shifter-red-soft>";
      // JSON Value
      var val = "<span class=text-shifter-white-soft>";
      // JSON String
      var str = "<span class=text-shifter-white-soft>";
      var r = pIndent || "";
      if (pKey) r = r + key + pKey.replace(/[": ]/g, "") + "</span>: ";
      if (pVal) r = r + (pVal[0] == '"' ? str : val) + pVal + "</span>";
      return r + (pEnd || "");
    },
    prettyPrint(obj) {
      var jsonLine = /^( *)("[\w]+": )?("[^"]*"|[\w.+-]*)?([,[{])?$/gm;
      return JSON.stringify(obj, null, 3)
        .replace(/&/g, "&amp;")
        .replace(/\\"/g, "&quot;")
        .replace(/</g, "&lt;")
        .replace(/>/g, "&gt;")
        .replace(jsonLine, this.replacer);
    },
  },
  computed: {
    ...mapState(useJSONModal, {
      showJSONModal: "showJSONModal",
    }),
    ...mapState(useJSONModal, {
      getContent: "getContent",
    }),

    content() {
      return this.prettyPrint(this.getContent);
    },
  },

  created() {},
};
</script>

<style scoped>
.scrollbar::-webkit-scrollbar-track {
  border-radius: 100vh;
  background: #f7f4ed;
}
.json-key {
  color: brown;
}
.json-value {
  color: navy;
}
.json-string {
  color: olive;
}
</style>
