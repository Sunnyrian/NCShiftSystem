import { createRouter, createWebHashHistory, createWebHistory, RouteRecordRaw} from 'vue-router'

import Login from '../pages/Login.vue'
import Register from '../pages/Register.vue'

const routes: Array<RouteRecordRaw> = [
    { path: '/', redirect: '/Login'},
    { path: '/Login', component: Login},
    { path: '/Register', component: Register}
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes,
})

export default router