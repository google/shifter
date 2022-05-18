<script setup></script>

<template>
  <div
    class="
      container
      flex-row
      absolute
      insert-0
      items-center
      w-full
      h-full
      bg-shifter-black-soft
      border
      rounded-2xl
      overflow-y-scroll
    "
    :class="showJSONModal ? 'visible' : 'invisible'"
  >
    <div class="container flex-row mx-auto m-10 bg-shifter-black- text-xl">
      <a
        @click="closeModal"
        class="
          rounded
          bg-shifter-red-soft
          rounded
          border border-shifter-red-soft
          px-6
          my-1
          hover:bg-shifter-red-soft hover:animate-pulse
        "
        >Close</a
      >
    </div>
    <div class="container flex mx-auto justify-end px-10">
      <div class="container flex-col mx-auto">
        <div class="container flex">Deployment Config</div>
        <div class="container flex">
          <pre v-html="content"></pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from "pinia";
import { useJSONModal } from "../stores/convert/jsonModal";
export default {
  data() {
    return {};
  },

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
/*pre {
  background-color: ghostwhite;
  border: 1px solid silver;
  padding: 10px 20px;
  margin: 20px;
}*/
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