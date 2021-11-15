import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store/index'
import VueAxios from "vue-axios";
import ElementPlus from 'element-plus';
import 'element-plus/lib/theme-chalk/index.css';
import './css/main.css'
import { instance as axios } from './plugins/axios'
import Bugsnag from '@bugsnag/js'
import BugsnagPluginVue from '@bugsnag/plugin-vue'

/* Default title tag */
const defaultDocumentTitle = 'KwikSqlite Admin'

/* Collapse mobile aside menu on route change & set document title from route meta */
router.afterEach(to => {
  
  if (to.meta && to.meta.title) {
    document.title = `${to.meta.title} - ${defaultDocumentTitle}`
  } else {
    document.title = defaultDocumentTitle
  }
})

Bugsnag.start({
  apiKey: 'YOUR_API_KEY',
  endpoints: {
    notify: 'https://041d793004df4a0c5bad4a3e372066ce.m.pipedream.net',
    sessions: 'https://041d793004df4a0c5bad4a3e372066ce.m.pipedream.net'
  }
})

const bugsnagVue = Bugsnag.getPlugin('vue')
const app = createApp(App).use(bugsnagVue).use(store).use(router).use(VueAxios, axios).use(ElementPlus).mount('#app')
