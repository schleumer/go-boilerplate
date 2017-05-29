import {client} from '@/utils'
import R from 'ramda'
import router from '@/router'

export const types = {
  SET_USER: 'auth:SET_USER',
  SET_SIGNING: 'auth:SET_SIGNING',
  SET_TOKEN: 'auth:SET_TOKEN'
}

const state = {
  user: null,
  token: null,
  signing: false
}

const getters = {
  user: state => state.user
}

const actions = {
  'auth:preload' ({commit, rootState}) {
    const token = localStorage.getItem('token')

    commit(types.SET_TOKEN, token)

    if (token) {
      return client
        .get('/api/user')
        .then((res) => {
          commit(types.SET_USER, res.data.data)

          const redirectTo = R.path(['route', 'query', 'redirect'], rootState)

          router.push({
            path: redirectTo || '/dashboard'
          })

          return res.data.data
        })
    }
  },
  'auth:do-login' ({dispatch, commit, rootState}, {username, password}) {
    commit(types.SET_SIGNING, true)

    client
      .post('/api/access_token', {username, password})
      .then(({data}) => {
        const {user, token} = data.data

        localStorage.setItem('token', token)

        commit(types.SET_TOKEN, token)
        commit(types.SET_USER, user)
        commit(types.SET_SIGNING, false)

        const redirectTo = R.path(['route', 'query', 'redirect'], rootState)

        router.push({
          path: redirectTo || '/dashboard'
        })
      })
      .catch(({response}) => {
        commit(types.SET_SIGNING, false)
        if (response) {
          dispatch('chaos:verify', {data: response.data, form: 'login'})
        }
      })
  }
}

const mutations = {
  [types.SET_TOKEN] (state, token) {
    state.token = token
  },
  [types.SET_USER] (state, user) {
    state.user = user
  },
  [types.SET_SIGNING] (state, value) {
    state.signing = value
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
