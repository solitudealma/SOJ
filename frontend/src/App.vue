<template>
  <el-container>
    <el-header>
      <NavBar></NavBar>
    </el-header>
    <el-main>
      <router-view :key="$route.fullPath" v-slot="{ Component }">
        <transition>
          <component :is="Component" :key="$route.path" />
        </transition>
      </router-view>
    </el-main>
    <el-footer>
      <FooterBar></FooterBar>
    </el-footer>
  </el-container>
</template>

<script setup lang="ts">
import NavBar from "@/components/NavBar.vue";
import FooterBar from "@/components/FooterBar.vue";

</script>

<style lang="less">
.el-main, .el-footer {
  margin-top: 20px;
  padding-right: 15px;
  padding-left: 15px;
  margin-right: auto;
  margin-left: auto;
}

@media (min-width: 768px) {
  .el-main, .el-footer {
    width: 750px;
  }
}

@media (min-width: 992px) {
  .el-main, .el-footer {
    width: 900px;
  }
}

@media (min-width: 1200px) {
  .el-main, .el-footer {
    width: 1000px;
  }
}

@media (min-width: 1300px) {
  .el-main, .el-footer {
    width: 1170px;
  }
}

</style>
