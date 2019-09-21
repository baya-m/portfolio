import Vue from 'vue'
import App from './App.vue'
import router from './router.js'

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')

router.beforeEach((to, from, next) => {
  if(to.matched.some(url => url.meta.isPublic)) {
      next()
  } else {
      next('login')
  }
});