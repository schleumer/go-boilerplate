import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import Login from '@/components/Login'
import Dashboard from '@/components/Dashboard'
import DashboardIndex from '@/components/DashboardIndex'
import store from '../store'

Vue.use(Router)

function requireAuth (to, from, next) {
  if (store.state.auth.user === null) {
    next({
      path: '/login',
      query: {redirect: to.fullPath}
    })
  } else {
    next()
  }
}

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Hello',
      component: Hello
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/dashboard',
      component: Dashboard,
      children: [
        {
          path: '/',
          name: 'Dashboard',
          component: DashboardIndex
        }
      ],
      beforeEnter: requireAuth
    }
  ]
})
