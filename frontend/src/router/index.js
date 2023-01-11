import { createRouter, createWebHistory } from 'vue-router'


const routes = [
  {
    path: '/users',
    name: 'Users',
    component: () => import("@/views/HomeView")
  },
  {
    path: '',
    // component: () => import("@/views/HomeView")
    redirect: '/users'
  },
  {
    path: '/',
    name: 'HomeView',
    // component: () => import("@/views/HomeView"),
    redirect: '/users'
  },
  {
    path: '/user/:id',
    name: 'User',
    component: () => import("@/views/UserView"),
  },
  {
    path: '/:catchAll(.*)',
    redirect: '/users'
  }
]

const router = createRouter({
  history: createWebHistory(),
  // base: process.env.BASE_URL,
  routes
})

export default router;
