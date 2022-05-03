<script setup>
import { useShifterV1StatusSettingz } from '../../../../stores/shifter/v1/status/settingz'
import Title from "../../../../components/Forms/Elements/Titles/Title.vue"
</script>

<template>
    <div>
      <Title>
        <!--<template #icon>
          <DocumentationIcon />
        </template>-->
        <template #title>Shifter Server Status</template>
        <template #subtitle>View the status and connection information of your Shifter Serv</template>
      </Title>
      <div>
        <ui>
          <li>Shifter Server Port: {{data.runningPort}}</li>
          <li>Shifter Server Host: {{data.runningHost}}</li>
          <li>Shifter Server Storage Type: {{data.storageDescription}} ({{data.storageType}})</li>
        </ui>
      </div>
    </div>
</template>

<script>
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

<style scoped> 
</style>