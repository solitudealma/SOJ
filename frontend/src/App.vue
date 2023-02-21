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

<style scoped>
@media screen and (min-width: 1050px) {
  .el-main {
    margin-top: 20px;
    padding: 0 5%;
  }
}

@media screen and (max-width: 1050px) {
  .el-main {
    margin-top: 20px;
    padding: 0 5px;
  }
}
</style>
