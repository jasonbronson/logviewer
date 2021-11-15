import { createStore } from "vuex";

const store = createStore({
  state: {
    /* User */
    userName: null,
    logs: [],
    selectedLog: "",
  },
  getters: {
    getLogs(state) {
      return state.logs;
    },
  },
  mutations: {
    setSelectedLog(state, value) {
      state.selectedLog = value;
    },
    setLogs(state, value) {
      state.logs = value;
    },
    /* User */
    user(state, payload) {
      if (payload.name) {
        state.userName = payload.name;
      }
    },
  },
  actions: {
    // asideMobileToggle ({ commit, state }, payload = null) {
    //   const isShow = payload !== null ? payload : !state.isAsideMobileExpanded
    //   document.getElementById('app').classList[isShow ? 'add' : 'remove']('ml-60')
    //   document.documentElement.classList[isShow ? 'add' : 'remove']('m-clipped')
    //   commit('basic', {
    //     key: 'isAsideMobileExpanded',
    //     value: isShow
    //   })
    // },
    // formScreenToggle ({ commit, state }, payload) {
    //   commit('basic', {
    //     key: 'isFormScreen',
    //     value: payload
    //   })
    //   document.documentElement.classList[payload ? 'add' : 'remove']('form-screen')
    // }
  },
  modules: {},
});

export default store;
