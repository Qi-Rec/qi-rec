import Vue from 'vue';
import App from './App.vue';
import axios from 'axios';

Vue.config.productionTip = false;

// Устанавливаем базовый URL для axios
axios.defaults.baseURL = 'http://0.0.0.0:3000';

// Добавляем axios в прототип Vue, чтобы можно было использовать его в компонентах через this.$axios
Vue.prototype.$axios = axios;

new Vue({
  render: h => h(App),
}).$mount('#app');