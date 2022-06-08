<template>
    <el-form
    label-position="Right"
    label-width="100px"
    model="user"
    style="max-width: 350px"
  >
    <el-form-item label="账号">
      <el-input 
      v-model="user.stuID" 
      placeholder="请输入学号"
      prefix-icon = "User" />
    </el-form-item>
    <el-form-item label="密码">
              <el-input
            v-model="user.password"
            type="password"
            placeholder="请输入密码"
            show-password
            prefix-icon = "Lock"
        />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="login">登录</el-button>
      <el-button @click="register">注册</el-button>
    </el-form-item>
  </el-form>
        

</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router';
import axios from 'axios'
import {ElMessage} from 'element-plus'
import sha256 from 'crypto-js/sha256'
import Base64 from 'crypto-js/enc-base64'

const labelPosition = ref('right')
const user = reactive({
    stuID: '',
    password: '',
})

const router = useRouter()

const register = () => {
  router.push('/Register')
}


const login = () => {
  var data = JSON.stringify({
    "stuID": user.stuID,
    "password": Base64.stringify(sha256(user.stuID+user.password+user.stuID)),
  });

  var config = {
    method: 'post',
    url: '/portal/login',
    headers: { 
      'Content-Type': 'application/json'
    },
    data : data
  };

  axios(config)
  .then(function (response:any) {
    console.log(JSON.stringify(response.data));
    if (response.data.success == true) {
      ElMessage({
          message: "登录成功!",
          type: 'success',
      })
      router.push('/Home')
    } else {
      ElMessage({
        showClose: true,
          message: "密码错误！",
          type: 'error',
      })
    }
  })
  .catch(function (error:any) {
    console.log(error);
  });

  console.log('submit!')
}

</script>