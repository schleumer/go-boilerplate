import * as types from '../mutation-types'

const state = {
  booted: false
}

const getters = {}

const actions = {
  'system:boot' ({dispatch, commit}) {
    dispatch('auth:preload').then(() => {
      commit(types.SYSTEM_BOOTED)
    })
  }
}

const mutations = {
  [types.SYSTEM_BOOTED] (state) {
    state.booted = true
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
