import Vue from 'vue'
import App from './App.vue'

// GoCaptcha
import "go-captcha-vue/dist/style.css"
import GoCaptcha from "go-captcha-vue"

// Cache Testing
// import "../cache/dist/style.css"
// import GoCaptcha from "../cache"

// Element
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:9001/';

Vue.config.productionTip = false

// use element ui
Vue.use(ElementUI)
// use go captcha
Vue.use(GoCaptcha)

new Vue({
  render: h => h(App),
}).$mount('#app')