<script setup></script>

<template>
  <div
    class="
      container
      flex
      absolute
      insert-0
      flex
      items-center
      w-full
      h-full
      justify-center
      bg-shifter-black
      border
      rounded-2xl
      truncate
    "
    :class="showJSONModal ? 'visible' : 'invisible'"
  >
    <div class="container flex-row mx-auto justify-end px-10 gap-4">
      <a @click="closeModal">close</a>
    </div>
    <div class="container flex-row mx-auto justify-end px-10 gap-4">
      {{ getContent }}
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
  },
  computed: {
    ...mapState(useJSONModal, {
      showJSONModal: "showJSONModal",
    }),
    ...mapState(useJSONModal, {
      getContent: "getContent",
    }),
  },
  created() {},
};
</script>
