import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css'

import Vue from 'vue'
import Material from 'vue-material'

import App from './app.component'
import router from './router'

Vue.config.productionTip = false

Vue.use(Material)

const application = new Vue({
  router,
  render (h) {
    return h(App)
  }
})

application.$mount('#app')
