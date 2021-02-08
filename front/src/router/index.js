import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import HelloWorld from '@/components/HelloWorld'
import Signup from '@/components/Signup'
import Signin from '@/components/Signin'
import Balance from '@/components/Balance'
import Ticker from '@/components/Ticker'


Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '*',
    redirect: 'signin'
  },
  {
    path: '/',
    name: 'HelloWorld',
    component: HelloWorld
  },
  {
    path: '/signup',
    name: 'Signup',
    component: Signup
  },
  {
    path: '/signin',
    name: 'Signin',
    component: Signin
  },
  {
    path: '/balance',
    name: 'Balance',
    component: Balance
  },
  {
    path: '/ticker',
    name: 'Ticker',
    component: Ticker
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router





