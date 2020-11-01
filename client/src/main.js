import Vue from 'vue'
import App from './App.vue'
import PaperDashboard from "./plugins/paperDashboard";
import router from "./router/index";
import axios from 'axios';
import store from './store/index.js';
import VueMaterial from 'vue-material';
import 'vue-material/dist/vue-material.min.css';
import 'vue-material/dist/theme/default.css';
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
Vue.use(Antd);

Vue.use(PaperDashboard);
Vue.use(VueMaterial); //dùng full component

Vue.config.productionTip = false

axios.defaults.baseURL = "http://localhost:5000/"


new Vue({
  router: router,
  store: store,
  render: h => h(App),
}).$mount('#app')
