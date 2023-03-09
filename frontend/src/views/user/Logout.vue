<template></template>

<script lang="ts" setup>
import { useRouter } from "vue-router";
import { useUserStore } from "@/pinia/modules/user";
import { logout } from "@/api/user";
const router = useRouter();
const userStore = useUserStore();
logout().then(() => {
  userStore.clearUserInfoAndToken();
  router.replace({
    path: "/",
  });
});
</script>