/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { createApp } from "vue";

// Import Pinia for State Management
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";
// Import Tailwind CSS Configuration
import "./index.css";

// Import Config JSON File
import Config from "./env-config.json";

// Import Toast Notifications
import Toast from "vue-toastification";
// Import Toast Notifications CSS
import "vue-toastification/dist/index.css";

// Create Vue Application
const app = createApp(App);

// Create Pinia State Store
const store = createPinia();

// Setup Shifter Environment Variable Configuration:
app.config.globalProperties.$shifterConfig = {
  // API Base URL Format "{domain.com}:{port}/api/vX"
  API_BASE_URL: Config.SHIFTER_SERVER_ENDPOINT,
};

// Add Toast Notifications to Vue Application
app.use(Toast, {
  // https://github.com/Maronato/vue-toastification#installation
  // You can set your default Toast Notifications options here
  transition: "Vue-Toastification__bounce",
  maxToasts: 20,
  newestOnTop: true,
});

// Add Pinia Store to Vue Application
app.use(store);
// Add Vue Router to Vue Application
app.use(router);

app.mount("#app");

// Export Shifter Environment Variable Configuration
export const shifterConfig = app.config.globalProperties.$shifterConfig;
