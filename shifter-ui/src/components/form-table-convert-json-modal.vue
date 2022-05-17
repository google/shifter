<script setup></script>

<template>
  <div
    class="
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
    "
    :class="showJSONModal ? 'visible' : 'invisible'"
  >
    <div class="container flex mx-auto justify-end px-10 gap-4">
      <a @click="closeModal">close</a>
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
  },
  created() {},
};
</script>
