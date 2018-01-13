import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css'

import Vue from 'vue'
import Material from 'vue-material'
import Router from 'vue-router'

import App from './app.component'
import router from './router'
import { RouterService } from './router/router.service'

Vue.use(Router)
Vue.use(Material)

Vue.config.productionTip = false

const application = new Vue({
  router,
  render (h) {
    return h(App)
  }
})

let { hash } = location
if (hash && hash.startsWith('#')) {
  hash = hash.substr(1)
}

RouterService.push({
  path: hash || '/'
})

// render the application on this component
application.$mount('#app')
