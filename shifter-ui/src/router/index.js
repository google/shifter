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

import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: () => import("../views/view-home.vue"),
    },
    {
      path: "/about",
      name: "about",
      // route level code-splitting
      // this generates a separate chunk (AboutView.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/view-about.vue"),
    },
    {
      path: "/convert",
      name: "convert",
      // route level code-splitting
      // this generates a separate chunk (Convert.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/view-convert.vue"),
    },
    {
      path: "/downloads",
      name: "downloads",
      // route level code-splitting
      // this generates a separate chunk (Download.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/view-downloads.vue"),
    },
    {
      path: "/downloads/:downloadId",
      name: "download",
      // route level code-splitting
      // this generates a separate chunk (Download.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/view-downloads.vue"),
    },
    {
      path: "/configure",
      name: "configure",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/view-configure.vue"),
    },
    {
      path: "/status/healthz",
      name: "healthz",
      component: () => import("../views/view-healthz.vue"),
    },
    {
      path: "/status/settingz",
      name: "settingz",
      component: () => import("../views/view-settingz.vue"),
    },
    {
      path: "/:pathMatch(.*)*",
      name: "404",
      component: () => import("../views/view-404.vue"),
    },
  ],
});

export default router;
