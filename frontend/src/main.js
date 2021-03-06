import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'


axios.defaults.withCredentials = true
axios.defaults.baseURL = 'http://127.0.0.1:8090'

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
