<template>
  <el-table
    :data="tableData"
    style="width: 100%"
    :row-class-name="tableRowClassName"
  >
    
    <el-table-column prop="name" label="姓名" width="180" />
    <el-table-column prop="stuID" label="学号" width="180" />
    <el-table-column prop="nickname" label="昵称" width="180" />
    <el-table-column prop="telephone" label="电话" width="180" />
    <el-table-column prop="email" label="邮箱" width="180"/>
    <el-table-column prop="status" label="账号状态" width="180" />
  </el-table>
</template>

<script lang="ts" setup>
    import axios from 'axios'
    import { reactive, ref } from 'vue'

    interface User {
    name: string
    stuID: string
    nickname: string
    email: string
    telephone: string
    status: boolean
    }

    const tableRowClassName = ({
    row,
    rowIndex,
    }: {
    row: User
    rowIndex: number
    }) => {
    if (rowIndex === 1) {
        return 'warning-row'
    } else if (rowIndex === 3) {
        return 'success-row'
    }
    return ''
    }

    getNAList()

    function getNAList() {
        var config = {
        method: 'get',
        url: '/adminApi/NAList',
        };

        axios(config)
        .then(function (response) {
        console.log(JSON.stringify(response.data));
        const tableData: User[] = response.data
        })
        .catch(function (error) {
        console.log(error);
        });
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
