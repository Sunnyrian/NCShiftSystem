import { createRouter, createWebHashHistory, createWebHistory, RouteRecordRaw} from 'vue-router'

import Login from '../pages/Login.vue'
import Register from '../pages/Register.vue'
import Home from '../pages/Index.vue';

const routes: Array<RouteRecordRaw> = [
    { path: '/', redirect: '/Login'},
    { path: '/Login', component: Login},
    { path: '/Register', component: Register},
    { path: '/Home', component: Home},
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes,
})

export default router