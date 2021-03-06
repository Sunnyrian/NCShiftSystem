import { createRouter, createWebHashHistory, createWebHistory, RouteRecordRaw} from 'vue-router'

import Login from '../pages/Login.vue'
import Register from '../pages/Register.vue'
import Home from '../pages/Index.vue'
import Admin from '../pages/Admin.vue';
import NAList from '../components/NAList.vue';
import NotFound from '../pages/NotFound.vue'
import NotAdmin from '../pages/NotAdmin.vue'
import Setting from '../components/Setting.vue'
import Shift from '../components/Setting.vue'

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
        path: '/NotAdmin',
        name: 'NotAdmin',
        component: NotAdmin,
    },
    {
        path: '/Admin',
        name: 'Admin',
        component: Admin,
        children: [
            {
                path: 'NAList/:status',
                name: 'NAList',
                component: NAList,
            },
            {
                path: 'Setting',
                name: 'Setting',
                component: Setting,
            },
            {
                path: 'Shift',
                name: 'Shift',
                component: Shift,
            },
        ],
        beforeEnter: (to, from) => {
            if(from.name != 'Admin') {
                checkLogin().then(() => {
                    changeRouteToAdmin(adminStatus)
                })
            }
            
        },
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes,
})


let loginStatus:boolean
let adminStatus:boolean

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
    // ???????????????
    if(!loginStatus) {
        router.push('/Login')
    }
}

function changeRouteToHome (loginStatus: boolean) {
    // ???????????????
    if(loginStatus) {
        router.push('/Home')
    }
}

function changeRouteToAdmin (adminStatus: boolean) {
    if(adminStatus){
        router.push('/Admin')
    } else {
        router.push('/NotAdmin')
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
        adminStatus = response.data.admin
        return adminStatus
    })
    .catch(function (error:any) {
    console.log(error);
    });
}



export default router