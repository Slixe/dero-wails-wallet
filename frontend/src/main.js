import Vue from 'vue'
import App from './App.vue'
import router from './router'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'

Vue.config.productionTip = false
Vue.prototype.$name = "DERO"
Vue.prototype.$donations = false
Vue.prototype.$selectColor = "grey lighten-4"

Vue.use(Vuetify)
new Vue({
  render: h => h(App),
  router,
  vuetify: new Vuetify({
    icons: {
      iconfont: 'md',
    },
    theme: {
      dark: true,
    }
  }),
  created() {
    this.$router.push('/')
  }
}).$mount('#app')