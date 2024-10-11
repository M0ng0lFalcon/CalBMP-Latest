import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/about",
    name: "About",
    component: () => import("../views/About.vue"),
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../views/Login.vue"),
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("../views/Register.vue"),
  },
  {
    path: "/userInputStep1",
    name: "UserInputStep1",
    component: () => import("../views/UserInputStep1.vue"),
  },
  {
    path: "/userInputStep2",
    name: "UserInputStep2",
    component: () => import("../views/UserInputStep2.vue"),
  },
  {
    path: "/visualization",
    name: "Visualization",
    component: () => import("../views/Visualization.vue"),
  },
  {
    path: "/history",
    name: "History",
    component: () => import("../views/History.vue"),
  },
  {
    path: "/manual",
    name: "Manual",
    component: () => import("../views/Manual.vue"),
  },
  {
    path: "/change_password",
    name: "ChangePassword",
    component: () => import("../views/ChangePassword"),
  },
  {
    path: "/bmp",
    name: "bmp",
    component: () => import("../views/Bmp"),
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
