// import { createStore } from "vuex";
import { localstorage } from "../services/storage/localStorageService";
import jwtDecode from "jwt-decode";
import api from '@/api'

const state = {
    userName: null,
    isAuthenticated: false,
}

const getters = {
}
const mutations = {
    /* User */
    user(state, payload) {
      if (payload.name) {
        state.userName = payload.name;
      }
    },
    setAuthenticated(state, value) {
      try {
        localstorage.setToken(value);
        let token = jwtDecode(value);
  
        if (token) {
          state.isAuthenticated = true;
        } else {
          state.isAuthenticated = false;
        }
      } catch (e) {
        console.log("err", e);
        state.isAuthenticated = false;
      }
    },
}
const actions = {
  login({ commit }, payload) {
    return api.auth
      .login(payload)
      .then((res) => {
        return res
      })
      .catch((err) => {
        console.log('login error: ', err)
      })
  },
}
export default {
    actions,
    getters,
    mutations,
    state
}
