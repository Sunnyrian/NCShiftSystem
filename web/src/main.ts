import { createApp } from 'vue'
// import ElementPlus from 'element-plus'
// import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import App from './App.vue'
import router from './route'
import Login from '../pages/Login.vue'
import axios from 'axios'


const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }
// app.use(ElementPlus)
app.use(router)
app.mount('#app')
// axios.defaults.baseURL = 'http://localhost:80';