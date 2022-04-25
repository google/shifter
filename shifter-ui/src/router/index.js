import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/Home.vue";
//import HomeView from "../views/HomeView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/about",
      name: "about",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/AboutView.vue"),
    },
    {
      path: "/status/healthz",
      name: "healthz",
      component: () => import("../views/status/HealthzView.vue"),
    },
    {
      path: "/status/settingz",
      name: "settingz",
      component: () => import("../views/status/SettingzView.vue"),
    },
  ],
});

export default router;
