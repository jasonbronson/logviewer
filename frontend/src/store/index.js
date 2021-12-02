// import { createStore } from "vuex";
import Vue from 'vue'
import Vuex from 'vuex'
import log from "./log.js"
import createPersistedState from 'vuex-persistedstate'
import auth from './auth.js'

// Vue.use(Vuex)
let _store

const createStore = () => {
  return (
    store ||
    new Vuex.Store({
      plugins: [
        createPersistedState({
          path: ['user'],
          reducer(val) {
            if (val === null) {
              return {}
            }
            return val
          }
        })
      ],
      modules: {
        log: log,
        auth: auth,
      }
    })
  )
}

const store = createStore()

export default store
