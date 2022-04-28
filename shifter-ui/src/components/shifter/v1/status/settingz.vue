<script setup>
import {useShifterV1StatusSettingz} from '../../../../stores/shifter/v1/status/settingz'
import DetailsListItem from '../../../DetailsListItem.vue';
import DocumentationIcon from "../../../icons/IconDocumentation.vue";
</script>

<template>
<main>
  <DetailsListItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>Shifter Server API Endpoint</template>

   {{shifterConfig.API_BASE_URL}}
  </DetailsListItem>

  <DetailsListItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>Shifter Server Port</template>

   {{data.runningPort}}
  </DetailsListItem>

  <DetailsListItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>Shifter Server Host</template>

   {{data.runningHost}}
  </DetailsListItem>

  <DetailsListItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>Shifter Server Storage Type</template>

   {{data.storageDescription}} ({{data.storageType}})
   
  </DetailsListItem>

  <DetailsListItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>Shifter Server Storage Output Path</template>

   {{data.storageOutputPath}}
   
  </DetailsListItem>

    <DetailsListItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>Shifter Server Storage Source Path</template>

   {{data.storageSourcePath}}
   
  </DetailsListItem>
  </main>
</template>

<script>
import { shifterConfig } from "@/main"
import { mapState, mapActions } from 'pinia'

export default {
  computed: {
    ...mapState(useShifterV1StatusSettingz, {data: 'results'})
  },

  methods: {
    ...mapActions(useShifterV1StatusSettingz, ['fetchSettingz']),
  },

  created() {
    // when the template is created, we call this action
    this.fetchSettingz();
  }
};
</script>