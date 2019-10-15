import Vue from 'vue'
import App from './App.vue'
import router from './router.js'
import store from './store/store.js'

Vue.config.productionTip = false

import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import "./vee-validate";

Vue.use(BootstrapVue)

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')

router.beforeEach((to, from, next) => {
  if(to.matched.some(url => url.meta.isPublic) || store.state.loggedIn) {
      next()
  } else {
      next('login')
  }
});