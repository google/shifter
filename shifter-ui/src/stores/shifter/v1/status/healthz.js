import { shifterConfig } from "../../../../main"
import axios from "axios"
import { defineStore } from 'pinia'

// API Endpoint Configuration
const config = {
    method: 'get',
    url: (shifterConfig.API_BASE_URL+'/status/healthz'),
    headers: {
    
    },
    data : null
};

export const useShifterV1StatusHealthz = defineStore('shifter-v1-status-healthz', {
  state: () => {
    return {
      data: {
        message: "Running Connection Tests."
      },
      fetching: false
    }
  },

  getters: {
    results(state) {
      return state.data;
    },

    isFetching(state) {
      return state.fetching;
    }
  },

  actions: {
    async fetchHealthz() {
      this.fetching = true;
      try{
        const response = await axios(config);
      try {
        const result = await response.data
        this.data = result;
      } catch (err) {
        this.data = [];
        console.error('Error loading Shifter Server, Status, Healthz API:', err);
        return err;
      }
      }
      catch (err) {
        this.data = {
          message: "Shifter Server Unreachable, Timeout"
        }
      }
      this.fetching = false;
    }
  }
})