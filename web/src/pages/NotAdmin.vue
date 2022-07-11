<template>
  <h1>401 您目前登录的不是管理员账号~</h1>
  <el-button type="primary" @click="backToHome">点击返回主页</el-button>
</template>

<script lang="ts" setup>
import { RouterLink, useRouter } from "vue-router";
const router = useRouter();

function backToHome(){
    router.push("/Home")
}
</script>
