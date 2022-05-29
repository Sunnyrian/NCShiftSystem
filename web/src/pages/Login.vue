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

const labelPosition = ref('right')
const user = reactive({
    stuID: '',
    password: '',
})


const login = () => {

  console.log('submit!')
}

const router = useRouter()

const register = () => {
  router.push('/Register')
}

</script>