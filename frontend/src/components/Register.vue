<template>
  <el-form :model="formRegister" :rules="ruleRegister" ref="formRegisterRef">
    <el-form-item props="username">
      <el-input
        placeholder="请输入账号"
        v-model="formRegister.username"
        autocomplete="off"
      ></el-input>
    </el-form-item>
    <el-form-item props="password">
      <el-input
        placeholder="请输入密码"
        v-model="formRegister.password"
        show-password
      ></el-input>
    </el-form-item>
    <el-form-item props="passwordConfirmed">
      <el-input
        placeholder="请再次输入密码"
        v-model="formRegister.passwordConfirmed"
        show-password
      ></el-input>
    </el-form-item>
  </el-form>
  <div class="footer">
    <el-button type="primary" @click="handleRegister(formRegisterRef)">
      注册
    </el-button>
    <el-link type="primary" @click="switchMode('Login')"
      >已有账号？立即登录！</el-link
    >
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from "vue";
import { ElMessage, FormInstance, FormRules } from "element-plus";
import { register } from "@/api/user";
import { useGlobalStore } from "@/pinia/modules/global";

const globalStore = useGlobalStore();
const usernameReg = /^[a-zA-Z0-9_-]{4,16}$/;

const validateUsername = (rule: any, value: any, callback: any) => {
  if (!value) {
    return callback(new Error("用户名不能为空!!"));
  }
  if (!usernameReg.test(value)) {
    return callback(new Error("格式有误"));
  } else {
    return callback();
  }
};
const formRegisterRef = ref<FormInstance>();
const formRegister = reactive({
  username: "",
  password: "",
  passwordConfirmed: "",
});
const ruleRegister = reactive<FormRules>({
  username: [
    { required: true, trigger: "blur" },
    { validator: validateUsername, trigger: "blur" },
  ],
  password: [
    {
      required: true,
      message: "请输入6-20位的密码",
      trigger: "change",
      min: 6,
      max: 20,
    },
  ],
});
const switchMode = (mode: string) => {
  globalStore.changeDialogStatus({
    mode,
    visible: true,
  });
};
const handleRegister = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async (valid) => {
    if (valid) {
      let res = await register(formRegister);
      if (res.code === 200) {
        ElMessage.success("注册成功～")
        switchMode('Login')
      }
    } else {
      ElMessage({
        type: "error",
        message: "请正确填写注册信息",
        showClose: true,
      });
    }
  });
};
</script>

<style lang="less" scoped>
.footer {
  overflow: auto;
  margin-top: 20px;
  margin-bottom: -15px;
  text-align: center;
}
.el-input-group__append {
  color: #fff;
  background: #25bb9b;
}
.footer :deep(.el-button--primary) {
  margin: 0 0 15px 0;
  width: 100%;
}
:deep(.el-form-item__content) {
  margin-left: 0px !important;
}
</style>