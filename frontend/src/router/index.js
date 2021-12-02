import { createRouter, createWebHashHistory } from "vue-router";
import Home from "../views/Home";
import Login from "../views/Login";

import store from "../store/index.js";
import { localstorage } from "../services/storage/localStorageService";
const routes = [
  {
    // Document title tag
    // We combine it with defaultDocumentTitle set in `src/main.js` on router.afterEach hook
    meta: {
      title: "Home",
      requiresAuth: true,
    },
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    meta: {
      title: "Login",
    },
    path: "/login",
    name: "Login",
    component: Login,
  },
  {
    meta: {
      title: "Log",
      requiresAuth: true,
    },
    path: "/log/:logname",
    name: "log",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "table" */ "../views/LogViewer"),
  },
  // {
  //   meta: {
  //     title: "Login",
  //   },
  //   path: "/login/",
  //   name: "login",
  //   component: () => import(/* webpackChunkName: "login" */ "../views/Login"),
  // },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    return savedPosition || { x: 0, y: 0 };
  },
});

router.beforeEach(function(to, from, next) {
  var aValue = localstorage.getToken();
  store.commit("setAuthenticated", aValue);
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    if (!store.state.auth.isAuthenticated) {
      next({
        path: "/login",
      });
    } else {
      next();
    }
  } else {
    next(); // make sure to always call next()!
  }
});

export default router;
