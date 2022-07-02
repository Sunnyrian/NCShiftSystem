import { createRouter, createWebHashHistory, createWebHistory, RouteRecordRaw} from 'vue-router'

import Login from '../pages/Login.vue'
import Register from '../pages/Register.vue'
import Home from '../pages/Index.vue'
import Admin from '../pages/Admin.vue';
import NAList from '../components/NAList.vue';
import NotFound from '../pages/NotFound.vue'
import VirtualTable from '../components/VirtualTable.vue'

import { tr } from 'element-plus/lib/locale'
import axios from 'axios'
import cookie from '../api/cookie.js'


const routes: Array<RouteRecordRaw> = [
    {   
        path: '/',
        name: '/',
        redirect: '/Login',
    },
    { 
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: NotFound,
    },
    {
        path: '/Login',
        name: 'Login',
        component: Login,
    },
    { 
        path: '/Register',
        name: 'Register',
        component: Register,
    },
    { 
        path: '/Home',
        name: 'Home',
        component: Home,
    },
    {
        path: '/Admin',
        name: 'Admin',
        component: Admin,
        children: [
            {
                path: 'NAList',
                name: 'NAList',
                component: NAList,
            },
            {
                path: 'VirtualTable',
                name: 'VirtualTable',
                component: VirtualTable,
            },
        ]
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes,
})

let loginStatus:boolean

router.beforeEach((to, from) => {

    let token = cookie.get("token")
    console.log("beforeEach~")
    if ( to.name != 'Login' && to.name != 'Register'){
        if (token === null) {
            return { name: 'Login'}
        } else {
            checkLogin().then( () => {
            changeRouteToLogin(loginStatus)
            })
        }
    } else {
        checkLogin().then( () => {
            changeRouteToHome(loginStatus)
        })
    }
    
})

function changeRouteToLogin (loginStatus: boolean) {
    // 如果没登录
    if(!loginStatus) {
        router.push('/Login')
    }
}

function changeRouteToHome (loginStatus: boolean) {
    // 如果登录了
    if(loginStatus) {
        router.push('/Home')
    }
}

async function checkLogin() {

    var config = {
    method: 'get',
    url: 'portalApi/checkLogin',
    };

    await axios(config)
    .then(function (response:any) {
        loginStatus = response.data.success
    })
    .catch(function (error:any) {
    console.log(error);
    });
}



export default router