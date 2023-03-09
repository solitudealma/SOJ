<template>
  <el-form :model="formLogin" ref="formLoginRef">
    <el-form-item prop="username">
      <el-input
        placeholder="请输入账号"
        v-model="formLogin.username"
        autocomplete="off"
      ></el-input>
    </el-form-item>
    <el-form-item prop="password">
      <el-input
        placeholder="请输入密码"
        v-model="formLogin.password"
        show-password
      ></el-input>
    </el-form-item>
  </el-form>
  <div class="footer">
    <el-button type="primary" @click="handleLogin(formLoginRef)"
      >登录
    </el-button>
    <el-link type="primary" @click="switchMode('Register')"
      >没有账号？立即注册！</el-link
    >
    <el-link type="primary" style="float: right"
      >忘记密码</el-link
    >
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import { useGlobalStore } from "@/pinia/modules/global";
import { useUserStore } from "@/pinia/modules/user";

const globalStore = useGlobalStore();
const userStore = useUserStore();
const formLoginRef = ref<FormInstance>();
const formLogin = reactive({
  username: "",
  password: "",
});

const login = async () => {
  return await userStore.LogIn(formLogin);
};
const switchMode = (mode: string) => {
  globalStore.changeDialogStatus({
    mode,
    visible: true,
  });
};
const handleLogin = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async (valid) => {
    if (valid) {
      await login();
    } else {
      ElMessage({
        type: "error",
        message: "请正确填写登录信息",
        showClose: true,
      });
      return false;
    }
  });
};
</script>

<style lang="less" scoped>
.footer {
  overflow: auto;
  margin-top: 20px;
  margin-bottom: -15px;
  text-align: left;
}
:deep(.el-button) {
  margin: 0 0 15px 0;
  width: 100%;
}
:deep(.el-form-item__content) {
  margin-left: 0px !important;
}
</style>
