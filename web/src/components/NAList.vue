<template>
    <el-table-v2
        :columns="columns"
        :data="tableData"
        :width="1000"
        :height="600"
        fixed
    />
    

</template>

<script lang="ts" setup>
    import axios from 'axios'
    import { reactive, ref } from 'vue'

    const columns = [
        {
            dataKey: "name",
            key: "name",
            title: "姓名",
            width: 150,
        },
        {
            dataKey: "stuid",
            key: "stuid",
            title: "学号",
            width: 150,
        },
        {
            dataKey: "nickname",
            key: "nickname",
            title: "昵称",
            width: 150,
        },
        {
            dataKey: "telephone",
            key: "telephone",
            title: "电话",
            width: 150,
        },
        {
            dataKey: "email",
            key: "email",
            title: "邮箱",
            width: 150,
        },
        {
            dataKey: "status",
            key: "status",
            title: "账号状态",
            width: 150,
        },
    ]

    let tableData = ref([])




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
          tableData.value = response.data
          console.log(tableData)
        })
        .catch(function (error) {
        console.log(error);
        });
    }
    getNAList()
</script>
