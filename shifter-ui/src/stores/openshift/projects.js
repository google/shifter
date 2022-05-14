import { shifterConfig } from "@/main";
import axios from "axios";
import { defineStore } from "pinia";

import { useConfigurationsClusters } from "../configurations/clusters";
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
      };
      this.fetching = true;
      try {
        const response = await axios(config);
        try {
          this.osProjects = response.data.projects.items;
        } catch (err) {
          this.osProjects = [];
          console.error("Error", err);
          return err;
        }
      } catch (err) {
        this.osProjects = [];
      }
      this.fetching = false;
    },
  },
});
