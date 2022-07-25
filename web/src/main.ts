import { createApp } from 'vue'
// import ElementPlus from 'element-plus'
// import 'element-plus/dist/index.css'
// import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import App from './App.vue'
import router from './route'
import axios from 'axios'
// import "~/styles/index.scss";
// import 'uno.css'

const app = createApp(App)
// for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
//     app.component(key, component)
//   }
// app.use(ElementPlus)
app.use(router)
app.mount('#app')
axios.defaults.baseURL = 'http://localhost:3500';