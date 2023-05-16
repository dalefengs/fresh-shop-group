import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [{
  path: '/',
  redirect: '/wechat'
},
{
  path: '/init',
  name: 'Init',
  component: () => import('@/view/init/index.vue')
},
{
  path: '/wechat',
  name: 'Login',
  component: () => import('@/view/wechat/index.vue')
}
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
