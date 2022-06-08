<template>
    <el-form
    label-position="Right"
    label-width="100px"
    :model="user"
    style="max-width: 320px"
    :rules="rules"
    ref="ruleFormRef"
  >
    <el-form-item label="账号" prop="stuID">
      <el-input 
      v-model="user.stuID" 
      placeholder="请输入学号"
      prefix-icon = "User" />
    </el-form-item>
    <el-form-item label="密码" prop="password">
              <el-input
            v-model="user.password"
            type="password"
            placeholder="请输入密码"
            show-password
            prefix-icon = "Lock"
        />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="login(ruleFormRef)">登录</el-button>
      <el-button @click="register">注册</el-button>
    </el-form-item>
  </el-form>
        

</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router';
import axios from 'axios'
import { ElMessage, FormInstance, timelineItemProps} from 'element-plus'
import sha256 from 'crypto-js/sha256'
import Base64 from 'crypto-js/enc-base64'

const labelPosition = ref('right')
const user = reactive({
    stuID: '',
    password: '',
})

const ruleFormRef = ref<FormInstance>()

const validateStuID = (rule: any, value: any, callback: any) => {
  if (value ==='') {
    callback(new Error('请输入学号'))
  } else {
    callback()
  }
}

const validatePass = (rule: any, value: any, callback: any) => {
  if (value ==='') {
    callback(new Error('请输入密码'))
  } else {
    callback()
  }
}

const rules = reactive({
  stuID: [{ validator: validateStuID, trigger: 'blur'}],
  password: [{ validator: validatePass, trigger: 'blur'}],
})

const router = useRouter()

const register = () => {
  router.push('/Register')
}


const login = (form: FormInstance | undefined) => {

  if (!form) return
  form.validate((valid) => {
    if (valid) {
      var data = JSON.stringify({
        "stuID": user.stuID,
        "password": Base64.stringify(sha256(user.stuID+user.password+user.stuID)),
      });

      var config = {
        method: 'post',
        url: '/portalApi/login',
        headers: { 
          'Content-Type': 'application/json'
        },
        data : data
      };

      axios(config)
      .then(function (response:any) {
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
    } else {
      console.log('error submit!')
      return false
    }
  })

}

</script>