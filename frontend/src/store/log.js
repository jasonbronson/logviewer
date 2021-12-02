// import { createStore } from "vuex";
import { localstorage } from "../services/storage/localStorageService";
import jwtDecode from "jwt-decode";
import api from '@/api'

const state = {
    logs: [],
    selectedLog: "",
}

const getters = {
    getLogs(state) {
      return state.logs;
    },
}
const mutations = {
    setSelectedLog(state, value) {
      state.selectedLog = value;
    },
    setLogs(state, value) {
      state.logs = value;
    },
}
const actions = {
  getLogs({ commit }, payload) {
    return api.logs
      .getLogs(payload.logName, payload.pageLimit, payload.pageOffset)
      .then((res) => {
        return res.data
      })
      .catch((err) => {
        console.log('getLogs error: ', err)
      })
  },
  getSearchContent({ commit }, payload) {
    return api.logs
      .getSearchContent(payload.filename, payload.keysearch, payload.pageLimit, payload.pageOffset)
      .then((res) => {
        return res.data
      })
      .catch((err) => {
        console.log('getSearchContent error: ', err)
      })
  },
}
export default {
    actions,
    getters,
    mutations,
    state
}
