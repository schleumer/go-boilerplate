// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'

import App from './App'
import router from './router'
import store from './store'
import {sync} from 'vuex-router-sync'

Vue.config.productionTip = false

const Helpers = function () {
}
Helpers.install = function (Vue, options) {
  Vue.prototype.$nextParent = function (type) {
    const parent = this.$parent
    if (parent && parent.$options.name === type) {
      return parent
    } else if (parent) {
      return parent.$nextParent(type)
    }

    return null
  }
}

Vue.use(Helpers)

sync(store, router, {moduleName: 'route'})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: {App}
})
