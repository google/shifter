<script setup>
// Vue Component Imports
import ListConvertDeploymentConfigObjectsItem from "./list-convert-deployment-config-objects-item.vue";
</script>
<template>
  <div class="container flex-row mx-auto justify-center">
    <ListConvertDeploymentConfigObjectsItem
      v-for="deploymentConfig in osDeploymentConfigs"
      :key="deploymentConfig.metadata.uid"
      :deploymentconfig="deploymentConfig"
    ></ListConvertDeploymentConfigObjectsItem>
  </div>
</template>

<script>
// Shifter Import Config
import { shifterConfig } from "@/main";
// Notifications Imports
import { notifyAxiosError } from "@/notifications";
// Axios Imports
import axios from "axios";
// External Pinia Store Imports
import { useConvertObjects } from "../stores/convert/convert";
import { useConfigurationsClusters } from "../stores/configurations/clusters";
import { useConfigurationsLoading } from "../stores/configurations/loading";
// Instansitate Pinia Store Objects
const storeConvertObjects = useConvertObjects();
const storeConfigClusters = useConfigurationsClusters();
const storeConfigLoading = useConfigurationsLoading();
// Plugin & Package Imports
import { mapState } from "pinia";
export default {
  props: {
    visible: {
      type: Boolean,
      required: false,
    },
    namespace: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      osDeploymentConfigs: [],
    };
  },
  watch: {
    // whenever question changes, this function will run
    visible(newState, oldState) {
      if (oldState === false && newState === true) {
        this.fetch();
      }
    },
  },

  methods: {
    async fetch() {
      // API Endpoint Configuration
      const config = {
        method: "post",
        url:
          shifterConfig.API_BASE_URL +
          "/openshift/deploymentconfigs/" +
          this.namespace,
        headers: {},
        data: {
          ...storeConfigClusters.getCluster(
            storeConvertObjects.selectedCluster.id
          ),
        },
        timeout: 2000,
      };
      try {
        /*storeConfigLoading.startLoading(
          "Loading...",
          "Fetching OpenShift Deployment Configurations"
        );*/
        this.osDeploymentConfigs = [];
        this.osDeploymentConfigs = await axios(config)
          .then((response) => {
            // handle success
            storeConfigLoading.endLoading();
            return response.data.deploymentConfigs.items;
          })
          .catch((err) => {
            console.log(err);
            notifyAxiosError(
              err,
              "Problem Loading OpenShift Deployment Configurations",
              6000
            );
            storeConfigLoading.endLoading();
            return err;
          });
      } catch (err) {
        console.log(err);
        this.osDeploymentConfigs = [];
        notifyAxiosError(
          err,
          "Problem Loading OpenShift Deployment Configurations",
          6000
        );
        storeConfigLoading.endLoading();
        return err;
      }
    },
  },

  computed: {
    ...mapState(useConfigurationsClusters, {
      configurationClusters: "getActiveClusters",
    }),
    ...mapState(useConfigurationsClusters, {
      configurationClusters: "getActiveClusters",
    }),
  },
};
</script>
