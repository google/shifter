import { defineStore } from "pinia";

export const useConvertObjects = defineStore("configurations-convert-object", {
  state: () => {
    return {
      items: [],
      fetching: false,
    };
  },

  getters: {
    selected(state) {
      return state.items;
    },
    contains(state) {
      return (item) => {
        if (state.items.length >= 1) {
          var idx = state.items.findIndex((object) => {
            return object.metadata.uid === item.metadata.uid;
          });

          if (idx >= 0) {
            return true;
          }
        }
        return false;
      };
    },
  },

  actions: {
    async addItem(item) {
      try {
        this.items.push(item);
      } catch (error) {
        console.log(error);
        return error;
      }
      return;
    },

    async removeItem(item) {
      if (this.items.length >= 1) {
        var idx = this.items.findIndex((object) => {
          return object.metadata.uid === item.metadata.uid;
        });
        if (idx >= 0) {
          this.items.splice(idx, 1);
        }
      }
    },
  },
});
