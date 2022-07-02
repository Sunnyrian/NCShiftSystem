<template>
  <el-table
    :data="tableData"
    style="width: 100%"
    stripe
    show-summary
    highlight-current-row
    table-layout="fixed"
  >
    <el-table-column prop="name" label="姓名" width="180" />
    <el-table-column prop="stuID" label="学号" width="180" />
    <el-table-column prop="nickname" label="昵称" width="180" />
    <el-table-column prop="telephone" label="电话" width="180" />
    <el-table-column prop="email" label="邮箱" width="180"/>
    <el-table-column prop="status" label="账号状态" width="180" />
  </el-table>
  <el-button type="primary" @click="update">更新</el-button>
</template>

<script lang="ts" setup>
    import axios from 'axios'
    import { reactive, ref } from 'vue'

    // interface User {
    // name: string
    // stuID: string
    // nickname: string
    // email: string
    // telephone: string
    // status: boolean
    // }

    // let tableData = [{"nickname":"Nakajima Ayano","name":"周隆","password":"","email":"nakajimaayano@icloud.com","telephone":"283585455","status":1,"stuid":""},{"nickname":"chenzefan","name":"陈泽帆","password":"","email":"123123","telephone":"123123123","status":1,"stuid":""},{"nickname":"huangjiaqiang","name":"黄嘉强","password":"","email":"123123","telephone":"123123126","status":1,"stuid":""},{"nickname":"Koon Kwok Ming","name":"李龙","password":"","email":"kmk@gmail.com","telephone":"75518550702","status":1,"stuid":""},{"nickname":"Tin Sze Kwan","name":"林敬桦","password":"","email":"szekwantin@outlook.com","telephone":"14703524633","status":1,"stuid":""},{"nickname":"Victoria Bryant","name":"刘潜","password":"","email":"vbryant@mail.com","telephone":"13409439986","status":1,"stuid":""},{"nickname":"Otsuka Momoka","name":"刘祥和","password":"","email":"mootsuka@gmail.com","telephone":"1032323765","status":1,"stuid":""},{"nickname":"Chow Wai Man","name":"鲁圣阳","password":"","email":"wmcho@icloud.com","telephone":"218356477","status":1,"stuid":""},{"nickname":"Heung Sau Man","name":"王茗琛","password":"","email":"smheung@icloud.com","telephone":"13199410087","status":1,"stuid":""},{"nickname":"Yue Kwok Kuen","name":"王淇巍","password":"","email":"ykwokk@mail.com","telephone":"75589636324","status":1,"stuid":""},{"nickname":"Douglas Watson","name":"杨乐欣","password":"","email":"wadoug@outlook.com","telephone":"2192923903","status":1,"stuid":""},{"nickname":"Ku Chi Yuen","name":"袁佳鸿","password":"","email":"chiyuk@icloud.com","telephone":"75519244047","status":1,"stuid":""},{"nickname":"Ng Siu Wai","name":"张芮","password":"","email":"siuwaing626@outlook.com","telephone":"2042195958","status":1,"stuid":""},{"nickname":"12345","name":"12345","password":"","email":"12345","telephone":"12345","status":0,"stuid":"12345"},{"nickname":"test3","name":"test3","password":"","email":"wo174831@qq.com","telephone":"18022519054","status":0,"stuid":"123"}]
    let tableData : any[] = reactive([])

    getNAList()

    function getNAList() {
        var config = {
        method: 'get',
        url: '/adminApi/NAList',
        };

        axios(config)
        .then(function (response) {
          console.log(response.data)
          // tableData = [];
          // for ( var i = 0; i < response.data.length; i++) {
          //   tableData.push(response.data[i])
          // }
          tableData = response.data
          console.log(tableData)
        })
        .catch(function (error) {
        console.log(error);
        });
    }

    function update() {
      console.log("before",tableData)
      getNAList()
      console.log("after",tableData)
    }

</script>

<style>
    .el-table .warning-row {
    --el-table-tr-bg-color: var(--el-color-warning-light-9);
    }
    .el-table .success-row {
    --el-table-tr-bg-color: var(--el-color-success-light-9);
    }
</style>
