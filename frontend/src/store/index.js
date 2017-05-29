import Vue from 'vue'
import Vuex from 'vuex'
import * as actions from './actions'
import * as getters from './getters'
import mutations from './mutations'
import auth from './modules/auth'
import system from './modules/system'
import chaos from './modules/chaos'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
  state: {
    count: 1
  },
  actions,
  getters,
  mutations,
  modules: {
    auth,
    system,
    chaos
  },
  strict: debug,
  plugins: debug ? [] : []
})
