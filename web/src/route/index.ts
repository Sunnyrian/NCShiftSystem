import { createRouter, createWebHashHistory, createWebHistory, RouteRecordRaw} from 'vue-router'

import Login from '../pages/Login.vue'
import Register from '../pages/Register.vue'
import Home from '../pages/Index.vue';
import { tr } from 'element-plus/lib/locale';
import axios from 'axios'

const routes: Array<RouteRecordRaw> = [
    {   
        path: '/',
        name: '/',
        redirect: '/Login',
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
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes,
})

let loginStatus:boolean

router.beforeEach((to, from) => {

    let token = localStorage.getItem('token')
    console.log("beforeEach~")
    if ( to.name != 'Login'){
        if (token === null || token === '') {
            return { name: 'Login'}
        } else {
            checkLogin(token).then( () => {
            changeRouteToLogin(loginStatus)
            })
        }
    } else {
        checkLogin(token).then( () => {
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

async function checkLogin(token: string|null) {
    var data = JSON.stringify({
    "token": token
    });

    var config = {
    method: 'post',
    url: 'portalApi/checkLogin',
    headers: { 
        'Content-Type': 'application/json'
    },
    data : data
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