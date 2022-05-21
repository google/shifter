import { shifterConfig } from "@/main";
import { notifyAxiosError } from "@/notifications";

import axios from "axios";
import { defineStore } from "pinia";

import { useConfigurationsLoading } from "../configurations/loading";
import { useConfigurationsClusters } from "../configurations/clusters";
const configurationsLoading = useConfigurationsLoading();
const configurationsClusters = useConfigurationsClusters();

export const useOSProjects = defineStore("openshift-projects", {
  state: () => {
    return {
      osProjects: [],
      fetching: false,
    };
  },

  getters: {
    getByName(state) {
      return (name) =>
        state.osProjects.find((project) => project.metadata.name === name);
    },
    getByUid(state) {
      return (uid) =>
        state.osProjects.find((project) => project.metadata.uid === uid);
    },
  },

  actions: {
    async fetch(clusterId) {
      // API Endpoint Configuration
      const config = {
        method: "post",
        url: shifterConfig.API_BASE_URL + "/openshift/projects/",
        headers: {},
        data: { ...configurationsClusters.getCluster(clusterId) },
        timeout: 2000,
      };
      try {
        configurationsLoading.startLoading(
          "Loading...",
          "Fetching OpenShift Namespaces & Projects"
        );
        this.osProjects = [];
        this.osProjects = await axios(config)
          .then((response) => {
            // handle success
            console.log(response);
            configurationsLoading.endLoading();
            return response.data.projects.items;
          })
          .catch((err) => {
            notifyAxiosError(
              err,
              "Problem Loading OpenShift Projects & Namespaces",
              6000
            );
            configurationsLoading.endLoading();
            return err;
          });
      } catch (err) {
        this.osProjects = [];
        notifyAxiosError(
          err,
          "Problem Loading OpenShift Projects & Namespaces",
          6000
        );
      }
    },
  },
});
