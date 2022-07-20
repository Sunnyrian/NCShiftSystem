<template>
      <el-table-v2
        :columns="columns"
        :data="data"
        :width="1000"
        :height="900"
        fixed
      />
</template>

<script lang="ts" setup>
import axios from "axios";
import { ref } from "vue";

const columns = [
  {
    dataKey: "key",
    key: "key",
    title: "变量名",
    width: 150,
  },
  {
    dataKey: "value",
    key: "value",
    title: "值",
    width: 150,
  },
];

let data = ref(['']);

function getEnv(){
  var config = {
  method: "get",
  url: "/adminApi/getEnv",
};

axios(config)
  .then(function (response) {
    console.log(JSON.stringify(response.data));
    data.value = response.data;
  })
  .catch(function (error) {
    console.log(error);
  });
}
getEnv();
</script>
