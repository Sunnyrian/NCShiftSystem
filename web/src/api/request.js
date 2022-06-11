// 请求
import axios from 'axios'

// create an axios instance   创建axios实例
const instance = axios.create({
	// baseURL: 'http://192.168.3.6:8082', // api 的 base_url
	// withCredentials: false,//跨域时使用凭证，默认带上cookies
	timeout: 2000, // request timeout  设置请求超时时间
  })

// 添加请求拦截器
instance.interceptors.request.use(
  config => {
  
    return config;
  },
  error => {
    return Promise.reject(error);
  });

export default instance