<template>
    <el-card class="box-card">
    <template #header>
      <div class="card-header"
      >
        <span>{{user.name}}</span>
        <el-button class="button" text>Operation button</el-button>
      </div>
    </template>
    <div>{{user.email}}</div>
    <div>{{user.stuID}}</div>
    <div>{{user.nickname}}</div>
    <div>{{user.telephone}}</div>
  </el-card>
</template>

<script lang="ts" setup>
  import { reactive, ref } from 'vue'
  import {
  Check,
  Delete,
  Edit,
  Message,
  Search,
  Star,
  } from '@element-plus/icons-vue'
  import axios from 'axios'
  

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
      console.log(response)
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