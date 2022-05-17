import { createApp } from "vue";

// Import Pinia for State Management
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";
// Import Tailwind CSS Configuration
import "./index.css";

// Import Config JSON File
import Config from "./env-config.json";

// Create Vue Application
const app = createApp(App);

// Create Pinia State Store
const store = createPinia();

// Setup Shifter Environment Variable Configuration:
app.config.globalProperties.$shifterConfig = {
  // API Base URL Format "{domain.com}:{port}/api/vX"
  API_BASE_URL: Config.SHIFTER_SERVER_ENDPOINT,
};

// Add Pinia Store to Vue Application
app.use(store);
// Add Vue Router to Vue Application
app.use(router);

app.mount("#app");

// Export Shifter Environment Variable Configuration
export const shifterConfig = app.config.globalProperties.$shifterConfig;
