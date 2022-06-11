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
    console.log("loginStatus:",loginStatus)
    // token ===null || token === '' && 
    if (token === null || token === '') {
        if (to.name !== 'Login') {
            return { name: 'Login'}
        } else {
            console.log("haihaihai")
        }
    } else {
        console.log("loginStatus(else):",loginStatus)
        checkLogin(token).then( () => {
            if(!loginStatus){
                console.log("让我check check")
                return { name: 'Login'}
            }
        })
    }
    
})

async function checkLogin(token: string) {
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

    axios(config)
    .then(function (response:any) {
        console.log(JSON.stringify(response.data));
        response.data.success = loginStatus
    })
    .catch(function (error:any) {
    console.log(error);
    });
    
}



export default router