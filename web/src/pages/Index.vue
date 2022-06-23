<template>
    <el-card class="box-card">
    <template #header>
      <div class="card-header"
      >
        <span>{{user.name}}</span>
        <el-button class="button" text>Operation button</el-button>
      </div>
    </template>
    <div>email:{{user.email}}</div>
    <div>学号:{{user.stuID}}</div>
    <div>昵称:{{user.nickname}}</div>
    <div>电话:{{user.telephone}}</div>
    <el-button type="primary" plain @click="logOut">退出登录</el-button>
  </el-card>
</template>

<script lang="ts" setup>
  import { reactive, ref } from 'vue'
  import { RouterLink, useRouter } from 'vue-router'
  import {
  Check,
  Delete,
  Edit,
  Message,
  Search,
  Star,
  } from '@element-plus/icons-vue'
  import axios from 'axios'
  import cookie from '../api/cookie'


  const router = useRouter()

  const user = reactive({
    stuID: '',
    name: '',
    nickname: '',
    telephone: '',
    email: '',
})

  getUserInformation()

  function getUserInformation() {
    var config = {
      method: 'get',
      url: 'http://localhost:3500/userApi/getUserInfo',
    };
    axios(config)
    .then(function (response) {
      user.stuID = response.data.stuID
      user.name  = response.data.name
      user.nickname = response.data.nickname
      user.telephone = response.data.telephone
      user.email = response.data.email
    })
    .catch(function (error) {
      console.log(error);
    });
  }
  
  function logOut() {
    cookie.remove("token")
    router.push('/Login')
  }
  
</script>

<style>
  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .text {
    font-size: 14px;
  }

  .item {
    margin-bottom: 18px;
  }

  .box-card {
    width: 480px;
  }
</style>