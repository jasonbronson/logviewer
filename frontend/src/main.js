import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store/index'
import VueAxios from "vue-axios";
import ElementPlus from 'element-plus';
import 'element-plus/lib/theme-chalk/index.css';
import './css/main.css'
import { instance as axios } from './plugins/axios'
// import Bugsnag from '@bugsnag/js'
// import BugsnagPluginVue from '@bugsnag/plugin-vue'

/* Default title tag */
const defaultDocumentTitle = 'Log Viewer'


// Bugsnag.start({
//   apiKey: 'YOUR_API_KEY',
//   endpoints: {
//     notify: '',
//     sessions: ''
//   }
// })

// const bugsnagVue = Bugsnag.getPlugin('vue')
const app = createApp(App).use(store).use(router).use(VueAxios, axios).use(ElementPlus).mount('#app')
//app.use(bugsnagVue)