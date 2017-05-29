export const types = {
  APPEND: 'chaos:APPEND'
}

const messageBuilder = (field, message, form = 'global', type = 'error') => {
  return {
    form: form || 'global',
    type: type || 'error',
    field,
    message
  }
}

const state = {
  messages: []
}

const getters = {}

const actions = {
  'chaos:verify' ({commit}, {form, data}) {
    if (data.errors && Array.isArray(data.errors) && data.errors.length > 0) {
      commit(types.APPEND, data.errors.map(e => messageBuilder(e.field, e.message, form, 'error')))
    }

    if (data.messages && Array.isArray(data.messages) && data.messages.length > 0) {
      commit(types.APPEND, data.messages.map(e => messageBuilder(e.field, e.message, form, 'message')))
    }
  },
  'chaos:fire' ({commit}, {form, type, field, message}) {
    commit(types.APPEND, messageBuilder(field, message, form, type))
  }
}

const mutations = {
  [types.APPEND] (state, message) {
    state.messages = state.messages.concat(message)
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
