//import { getCurrentInstance } from 'vue'
import axios from "axios";
import { settings } from "@/api/settings";
//import { localstorage } from '../services/storage/localStorageService'
import store from "@/store/index";

const instance = axios.create({
  baseURL: settings.baseURL,
});

instance.interceptors.request.use(
  (config) => {
    // const token = localstorage.getToken()
    
    // let authToken = `Bearer ${token || process.env.DEVICE_TOKEN}`
    let authToken;
    if (
      config.url.includes("/sign-in") // dot not include the token when sign-in
    ) {
      authToken = undefined;
    }
    if (authToken) {
      config.headers["Authorization"] = authToken;
    }
    return config;
  },
  (error) => {
    Promise.reject(error);
  }
);
instance.interceptors.response.use(
  (response) => {
    // const token = response.headers['authtoken']
    // if (token) {
    // //   store.commit('localAuthenticated', {
    // //     token
    // //   })
    // }
    return response;
  },
  function(error) {
    if (error.response && error.response.status === 401) {
      //localstorage.clearAuth()
    }
    return Promise.reject(error);
  }
);

// const app = getCurrentInstance()
// console.log(app)
//Vue.prototype.$axios = instance
//Vue.prototype.$localstorage = localstorage
//app.appContext.config.globalProperties.$axios = instance
//this.$localstorage = localstorage

export { instance };
