<script setup>
import {useShifterV1StatusHealthz} from '../../../../stores/shifter/v1/status/healthz'
import WelcomeItem from '../../../WelcomeItem.vue';
import DocumentationIcon from "../../../icons/IconDocumentation.vue";
</script>

<template>
  <WelcomeItem>
    <template #icon>
      <!--<DocumentationIcon />-->
    </template>
    <template #heading>Shifter Server Status</template>

   {{data.message}}
  </WelcomeItem>
</template>

<script>
import { mapState, mapActions } from 'pinia'

export default {
  computed: {
    ...mapState(useShifterV1StatusHealthz, {data: 'results'})
  },

  methods: {
    ...mapActions(useShifterV1StatusHealthz, ['fetchHealthz']),

    addToCart() {
      // we'll populate this later
    }
  },

  created() {
    // when the template is created, we call this action
    this.fetchHealthz();
  }
};
</script>