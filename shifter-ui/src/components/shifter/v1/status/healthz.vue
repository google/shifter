<script setup>
import {useShifterV1StatusHealthz} from '../../../../stores/shifter/v1/status/healthz'
import DetailsListItem from '../../../DetailsListItem.vue';
import DocumentationIcon from "../../../icons/IconDocumentation.vue";
</script>

<template>
  <DetailsListItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>Shifter Server Status</template>

   {{data.message}}
  </DetailsListItem>
</template>

<script>
import { mapState, mapActions } from 'pinia'

export default {
  computed: {
    ...mapState(useShifterV1StatusHealthz, {data: 'results'})
  },

  methods: {
    ...mapActions(useShifterV1StatusHealthz, ['fetchHealthz']),
  },

  created() {
    // when the template is created, we call this action
    this.fetchHealthz();
  }
};
</script>