<template>
  <el-table-v2
    :columns="columns"
    :data="tableData"
    :width="1000"
    :height="900"
    fixed
  />
  <el-button type="primary" @click="getTableData">Update</el-button>
</template>


<script lang="ts" setup>
import axios from "axios";
import { reactive, ref } from "vue";
import { useRouter, useRoute } from "vue-router";

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
];

let tableData = ref([]);
const route = useRoute()

function getNAList(status: string | string[]) {
  var config = {
    method: "get",
    url: "/adminApi/NAList" + "?status=" + status,
  };

  axios(config)
    .then(function (response) {
      console.log("response:", response.data);
      tableData.value = response.data;
      console.log("tableData:", tableData);
    })
    .catch(function (error) {
      console.log(error);
    });
}
getNAList(route.params.status);

function getTableData() {
    console.log("now the TableData is:",tableData)
}
</script>
