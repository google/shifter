<script setup>
import { useConfigurationsClusters } from '../../stores/configurations/clusters'
</script>

<template>
<div class="container flex mx-auto m-6 items-center">
    <div class="container flex-row mx-auto bg-shifter-black-mute justify-center rounded-2xl py-6">
        <div class="container flex mx-auto justify-center py-6 gap-8 uppercase pb-12">
            <div v-for="step in convertSteps" :key=step.id class="container flex justify-center items-center">
                <div class="flex rounded-full shadow-2xl bg-shifter-black px-4 py-2 text-shifter-red-muted text-bold">{{step.id}}</div> 
                <div class="flex ml-4">{{step.title}}</div> 
            </div>
        </div>
        <div class="container flex-row mx-auto justify-center py-12">
            <div class="container flex-row justify-center items-center">
                <div class="flex justify-center bold text-4xl m-2">Cluster Selection</div>
                <div class="flex justify-center text-baseline m-2">Select OpenShift cluster from which you would like to convert workloads</div>
            </div>
            <div class="container flex mx-auto justify-center my-4">
                <select class="flex justify-center w-1/4 p-2 m-2 bg-shifter-black rounded" id="cluster">
                    <option v-for="cluster in configurationClusters" :key=cluster.id :value="cluster.shifter">{{cluster.shifter.clusterConfig.connectionName}}</option>
                </select>
            </div>
        </div>
         <div class="container flex mx-auto justify-end px-10">
                <a class="uppercase">Next</a>
        </div>
    </div>
    
</div>
</template>

<script>
import { mapState, mapActions } from 'pinia'

export default {
  data() {
    return {
      convert: {
          shifter :{

          }
      },
      convertSteps: [
        {
          id: 1,
          title: "Cluster Selection",
        },
        {
          id: 2,
          title: "Object Selection",
        },
        {
          id: 3,
          title: "Review",
        },
        {
          id: 4,
          title: "Shift Workflows",
        },
         {
          id: 5,
          title: "Summary",
        },
      ],
    }
  },
  computed: {
    ...mapState(useConfigurationsClusters, {configurationClusters: 'getActiveClusters'})



  },

  methods: {
    //...mapActions(useConfigurationsClusters, ['fetchClusters']),
  },

  created() {
    //this.fetchClusters();
  }
};
</script>
