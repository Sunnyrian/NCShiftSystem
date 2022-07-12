<template>
  <el-row>
    <el-col :span="6">
      <el-menu
        default-active="2"
        class="el-menu-vertical-demo"
        router
      >
        <el-sub-menu index="NA">
          <template #title>
            <el-icon><Avatar /></el-icon>
            <span>网管</span>
          </template>

          <el-menu-item-group>
            <el-menu-item index="/Admin/NAList/1">在职网管</el-menu-item>
            <el-menu-item index="/Admin/NAList/*">所有网管</el-menu-item>
            <el-menu-item index="/Admin/NAList/0">预注册网管</el-menu-item>
          </el-menu-item-group>
        </el-sub-menu>

        <el-menu-item index="shift">
          <el-icon><icon-menu /></el-icon>
          <template #title>班次管理</template>
        </el-menu-item>

        <el-menu-item index="ot">
          <el-icon><List /></el-icon>
          <template #title>空闲系统</template>
        </el-menu-item>
        <el-menu-item index="setting">
          <el-icon><setting /></el-icon>
          <template #title>系统设置</template>
        </el-menu-item>
        <el-menu-item index="dingding">
          <el-icon><Link /></el-icon>
          <template #title>钉钉对接</template>
        </el-menu-item>
      </el-menu>
      <el-button plain @click="logOut">退出登录</el-button>
    </el-col>
    <el-col :span="6">
      <router-view></router-view>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
  Link,
} from "@element-plus/icons-vue";
import cookie from "../api/cookie";
import { useRouter, useRoute } from "vue-router";

const router = useRouter();

function logOut() {
  cookie.remove("token");
  router.push("/Login");
}

const isCollapse = ref(true);

</script>

<style>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>
