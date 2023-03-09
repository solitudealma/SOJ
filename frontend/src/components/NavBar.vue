<template>
  <el-menu :default-active="activeRoute" router mode="horizontal" :ellipsis="false">
    <div class="logo">SOJ</div>
    <el-menu-item index="/">Home</el-menu-item>
    <el-menu-item index="/problem">Problem</el-menu-item>
    <el-menu-item index="/solution">Solution</el-menu-item>
    <el-sub-menu index="/about">
      <template #title>About</template>
      <el-menu-item index="/judgeinfo">JudgeInfo</el-menu-item>
    </el-sub-menu>
    <div class="flex-grow" />
    <div class="btn-menu" v-if="!isAuthenticated">
      <el-popover placement="bottom" transition="fade-in-linear" width="200" trigger="hover">
        <span>登录后你可以:</span>
        <br />
        <br />
        <span>
          <el-icon>
            <IEpEditPen />
          </el-icon>
          刷题
        </span>
        <span class="btn-menu-icon-trophy">
          <el-icon>
            <IEpMedal />
          </el-icon>
          比赛
        </span>
        <br /><br />
        <!-- LoginForm -->
        <a class="login-btn" @click="handleBtnClick('Login')">Login</a>
        <!-- LoginForm -->
        <br />
        <span style="position: relative; left: 30px">首次使用?</span>
        <!-- RegisterForm -->
        <a class="register-btn" @click="handleBtnClick('Register')">
          点我注册
        </a>
        <template #reference>
          <el-button type="primary">登录</el-button>
        </template>
      </el-popover>
    </div>
    <template v-else>
      <el-avatar class="header-img" :size="40" :src="userInfo.avatar" :alt="avatarAlt">
      </el-avatar>
      <el-dropdown class="drop-menu" :teleported="false" @command="handleCommand">
        <a style="text-decoration: none;">
          <strong id="id_user_username">
            {{ userInfo.username }}
            <el-icon>
              <IEpArrowDown />
            </el-icon>
          </strong>
        </a>
        <template #dropdown>
          <el-dropdown-menu class="dropdown-menu">
            <el-dropdown-item command="mySpace">我的空间</el-dropdown-item>
            <el-dropdown-item command="info">个人信息</el-dropdown-item>
            <el-dropdown-item command="changePassword">修改密码
            </el-dropdown-item>
            <el-dropdown-item :divided="true" command="myTicket">我的工单
            </el-dropdown-item>
            <el-dropdown-item :divided="true" command="logout">登出
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </template>
  </el-menu>
  <el-dialog :title="dialogTitle" v-model="dialogVisible" width="20%" :close-on-click-modal="false">
    <component :is="dialog" v-if="dialogVisible"></component>
  </el-dialog>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { RouteLocationRaw, useRoute, useRouter } from "vue-router";
import { useGlobalStore } from "@/pinia/modules/global";
import { useUserStore } from "@/pinia/modules/user";
import Login from "@/components/Login.vue";
import Register from "@/components/Register.vue";
import { DialogStatus } from "#/store";

const route = useRoute();
const router = useRouter();
const activeRoute = computed(() => route.path);
const globalStore = useGlobalStore();
const userStore = useUserStore();

let isAuthenticated = computed(() => userStore.isAuthenticated);
let userInfo = computed(() => userStore.userInfo);
let avatarAlt = computed(() => `${userInfo.value.username}的头像`);

let dialog = shallowRef(Login);
let dialogVisible = computed({
  get() {
    return globalStore.dialogStatus.visible;
  },
  set(value) {
    globalStore.changeDialogStatus({
      visible: value,
    } as DialogStatus);
  },
});
let dialogTitle = computed(() =>
  globalStore.dialogStatus.mode === "Login" ? "登录-SOJ" : "注册-SOJ"
);

const handleBtnClick = (mode: string) => {
  if (mode === "Login") {
    dialog.value = Login;
  } else if (mode === "Register") {
    dialog.value = Register;
  }
  globalStore.changeDialogStatus({ mode, visible: true });
};

const handleCommand = (route: RouteLocationRaw) => {
  router.push(route);
};

watch(
  () => globalStore.dialogStatus.mode,
  (value) => {
    let mode = globalStore.dialogStatus.mode;
    if (mode === "Login") {
      dialog.value = Login;
    } else if (mode === "Register") {
      dialog.value = Register;
    }
  }
);
</script>

<style lang="less" scoped>
.logo {
  margin-left: 2%;
  margin-right: 2%;
  font-size: 20px;
  float: left;
  line-height: 60px;
}
.btn-menu {
  float: right;
  margin-right: 20px;
  margin-top: 10px;
}
.header-img {
  margin: 10px 0 -13px 0;
}
.drop-menu {
  margin-top: 24px;
  .dropdown-menu {
    left: 0;
    float: left;
    min-width: 160px;
    padding: 5px 0;
    font-size: 14px;
    text-align: left;
    background-color: #fff;
    background-clip: padding-box;
    border-radius: 4px;
  }
}
.flex-grow {
  flex-grow: 1;
}
.el-popover {
  .btn-menu-icon-trophy {
    margin: 0 0 0 70px;
  }
  .login-btn {
    display: block;
    box-sizing: border-box;
    height: 40px;
    padding: 10px 0;
    line-height: 20px;
    font-size: 14px;
    color: #ffffff;
    letter-spacing: 0;
    text-align: center;
    background: #00a1d6;
    border-radius: 2px;
    cursor: pointer;
  }
  .register-btn {
    color: #00a1d6;
    position: relative;
    left: 40px;
    cursor: pointer; //鼠标手型
  }
}
#id_user_username {
  cursor: pointer;
  color: #909399;
}
#id_user_username:hover {
  color: #303133;
}
:deep(.el-dialog) {
  border-radius: 10px !important;
  text-align: center;
}
:deep(.el-dialog__header .el-dialog__title) {
  font-size: 22px;
  font-weight: 600;
  font-family: Arial, Helvetica, sans-serif;
  line-height: 1em;
  color: #4e4e4e;
}
</style>