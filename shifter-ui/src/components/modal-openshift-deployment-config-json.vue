<template>
  <div
    class="container flex-row absolute items-center w-full h-full bg-shifter-black-soft border rounded-2xl overflow-y-auto"
    :class="showJSONModal ? 'visible' : 'invisible'"
  >
    <div class="container flex-row mx-auto bg-shifter-black- text-xl">
      <div
        class="container flex mx-auto justify-end p-4 bg-shifter-black-mute gap-4"
      >
        <a
          @click="closeModal"
          class="rounded bg-shifter-red-soft rounded border border-shifter-red-soft px-6 my-1 hover:bg-shifter-red-soft hover:animate-pulse"
          >Close</a
        >
      </div>
    </div>
    <div class="container flex mx-auto justify-end px-4 py-10">
      <div class="container flex flex-row mx-auto">
        <div
          class="container flex w-1/4 mr-4 text-lg font-bold hover:animate-pulse cursor cursor-pointer"
        >
          Deployment Config JSON
        </div>
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
