<template>
  <el-menu
    :default-active="activeRoute"
    class="el-menu-demo"
    router
    mode="horizontal"
  >
    <div class="logo">SOJ</div> 
    <el-menu-item index="/">Home</el-menu-item>
    <el-menu-item index="/problem">Problem</el-menu-item>
    <el-menu-item index="/solution">Solution</el-menu-item>
    <el-sub-menu index="/about">
        <template #title>About</template>
        <el-menu-item index="/judgeinfo">JudgeInfo</el-menu-item>
    </el-sub-menu>
    <div class="flex-grow" />
  </el-menu>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();
const activeRoute = computed(() => route.path);
</script>
<style lang="less" scoped>
.logo {
  margin-left: 2%;
  margin-right: 2%;
  font-size: 20px;
  float: left;
  line-height: 60px;
}

.flex-grow {
  flex-grow: 1;
}
</style>
