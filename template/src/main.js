import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import "nly-adminlte-vue/dist/adminlte/css/adminlte.css";
import "nly-adminlte-vue/dist/adminlte/fontawesome-free/css/all.css";
import "nly-adminlte-vue/dist/adminlte/icon/iconfont.css";
import "nly-adminlte-vue/dist/nly-adminlte-vue.css";
import "./assets/open-iconic/font/css/open-iconic-bootstrap.css";
import { NlyAdminlteVue, NlyAdminlteVueIcons } from "nly-adminlte-vue";
Vue.use(NlyAdminlteVue);
Vue.use(NlyAdminlteVueIcons);

import Toast from "./utils/toast";
import Api from "./utils/http";
Vue.prototype.$toast = Toast;
Vue.prototype.$api = Api;

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
