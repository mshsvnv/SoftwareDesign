import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
// import Login from '../views/Login.vue'
// import Cart from '../views/Cart.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/cart',
      name: 'cart',
      component: () => import('../views/Cart.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login.vue')
    },
    {
      path: '/registration',
      name: 'registration',
      component: () => import('../views/Registration.vue')
    }
  ]
})

export default router
