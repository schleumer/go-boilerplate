import axios from 'axios'
import store from './store'

export const client = axios.create({
  baseURL: process.env.BACKEND_URL
})

client.interceptors.request.use(function (config) {
  if (store.state.auth.token) {
    config.headers.Authorization = `Bearer ${store.state.auth.token}`
  }

  return config
}, (error) => Promise.reject(error))
