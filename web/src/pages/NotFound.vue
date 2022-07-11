<template>
  <h1>404 您所输入的页面不存在~</h1>
  <el-button type="primary" @click="backToHome"
    >点击返回主页 {{ counter }}s</el-button
  >
</template>

<script lang="ts">

</script>

<script lang="ts" setup>
import { ref, reactive } from "vue";
import { RouterLink, useRouter } from "vue-router";
const router = useRouter();

let counter = ref(3);

setInterval(() => {
  counter.value--;
  if(counter.value==0){
    backToHome()
  }
}, 1000);


function backToHome() {
  router.push("/Home");
}
</script>
