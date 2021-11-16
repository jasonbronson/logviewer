import { createRouter, createWebHashHistory } from "vue-router";
import Home from "../views/Home";
const routes = [
  {
    // Document title tag
    // We combine it with defaultDocumentTitle set in `src/main.js` on router.afterEach hook
    meta: {
      title: "Home",
    },
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    meta: {
      title: "Log",
    },
    path: "/log/:logname",
    name: "log",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "table" */ "../views/LogViewer"),
  },
  {
    meta: {
      title: "Login",
    },
    path: "/login",
    name: "login",
    component: () => import(/* webpackChunkName: "login" */ "../views/Login"),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    return savedPosition || { x: 0, y: 0 };
  },
});

export default router;
