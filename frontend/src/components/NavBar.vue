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
    <div class="btn-menu" v-if="!isAuthenticated">
      <el-popover
        placement="bottom"
        transition="fade-in-linear"
        title=""
        width="200"
        trigger="hover"
        content=""
      >
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
        <a class="login-btn" @click="loginDialogFormVisible = true">Login</a>
        <template>
          <el-dialog
            title="登录"
            @close="loginDialogReset(formLoginRef)"
            v-model="loginDialogFormVisible"
            :append-to-body="true"
            :modal="true"
            width="30%"
            close-on-press-escape
            :destroy-on-close="true"
            center
          >
            <el-form :model="formLogin" :rules="ruleLogin" ref="formLoginRef">
              <el-form-item
                label="账号"
                :label-width="formLabelWidth"
                prop="username"
              >
                <el-input
                  placeholder="请输入账号"
                  v-model="formLogin.username"
                  autocomplete="off"
                ></el-input>
              </el-form-item>
              <el-form-item
                label="密码"
                :label-width="formLabelWidth"
                prop="password"
              >
                <el-input
                  placeholder="请输入密码"
                  v-model="formLogin.password"
                  show-password
                ></el-input>
              </el-form-item>
            </el-form>
            <template #footer>
              <el-button @click="loginDialogFormVisible = false"
                >取 消
              </el-button>
              <el-button type="primary" @click="handleLogin(formLoginRef)"
                >确 定
              </el-button>
            </template>
          </el-dialog>
        </template>
        <!-- LoginForm -->
        <br />
        <span style="position: relative; left: 30px">首次使用?</span>
        <!-- RegisterForm -->
        <a class="register-btn" @click="registerDialogFormVisible = true">
          点我注册
        </a>
        <template>
          <el-dialog
            title="注册"
            v-model="registerDialogFormVisible"
            :append-to-body="true"
            width="30%"
            close-on-press-escape
            :destroy-on-close="true"
            center
          >
            <el-form :model="formRegister">
              <el-form-item label="账号" :label-width="formLabelWidth">
                <el-input
                  placeholder="请输入账号"
                  v-model="formRegister.username"
                  autocomplete="off"
                ></el-input>
              </el-form-item>
              <el-form-item label="密码" :label-width="formLabelWidth">
                <el-input
                  placeholder="请输入密码"
                  v-model="formRegister.password"
                  show-password
                ></el-input>
              </el-form-item>
            </el-form>
            <template #footer>
              <el-button @click="registerDialogFormVisible = false"
                >取 消
              </el-button>
              <el-button
                type="primary"
                @click="registerDialogFormVisible = true"
                >确 定
              </el-button>
            </template>
          </el-dialog>
        </template>
        <!-- RegisterForm -->
        <template #reference>
          <el-button type="primary">登录</el-button>
        </template>
      </el-popover>
    </div>
    <template v-else>
      <el-dropdown size="medium" class="drop-menu" @command="handleCommand">
        <a style="text-decoration: none">
          <el-avatar
            class="header-img"
            :size="40"
            :src="myHeader"
            :alt="avatarAlt"
          >
          </el-avatar>
          <strong id="id_user_username">
            {{ user.username }}
            <i class="fa fa-caret-down"></i>
          </strong>
        </a>
        <el-dropdown-menu slot="dropdown" class="dropdown-menu">
          <el-dropdown-item command="mySpace">我的空间</el-dropdown-item>
          <el-dropdown-item command="info">个人信息</el-dropdown-item>
          <el-dropdown-item command="changePassword">修改密码</el-dropdown-item>
          <el-dropdown-item :divided="true" command="myTicket"
            >我的工单
          </el-dropdown-item>
          <el-dropdown-item :divided="true" command="logout"
            >登出
          </el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </template>
  </el-menu>
</template>

<script setup lang="ts">
import { computed, ref, reactive } from "vue";
import type { FormInstance, FormRules } from "element-plus";
import { useRoute } from "vue-router";

const route = useRoute();
const activeRoute = computed(() => route.path);

let isAuthenticated = ref(false);
let loginDialogFormVisible = ref(false);
let registerDialogFormVisible = ref(false);
let formLabelWidth = "60px";
let myHeader = ref(
  "https://ae01.alicdn.com/kf/Hd60a3f7c06fd47ae85624badd32ce54dv.jpg"
);
let avatarAlt = ref("SolitudeAlma的头像");
let user = ref({
  username: "",
});
const formLoginRef = ref<FormInstance>();

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

const formLogin = reactive({
  username: "",
  password: "",
});

const formRegister = reactive({
  username: "",
  password: "",
});
const ruleLogin = reactive<FormRules>({
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

const handleLogin = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate((valid) => {
    if (valid) {
      console.log("submit!");
    } else {
      console.log("error submit!");
      return false;
    }
  });
};

const loginDialogReset = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};

const handleCommand = (command: string | number | object) => {
  switch (command) {
    case "changPassword":
      break;
    case "mySpace":
      break;
    case "myTicket":
      break;
    case "info":
      break;
    default:
      break;
  }
};
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

.drop-menu {
  float: right;
  margin-right: 20px;

  .el-dropdown-menu .dropdown-menu {
    min-width: 160px;
  }

  .header-img {
    margin: 10px 0 -13px 0;
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

.dropdown-menu {
  position: absolute;
  top: 100%;
  left: 0;
  float: left;
  min-width: 160px;
  padding: 5px 0;
  margin: 2px 0 0;
  font-size: 14px;
  text-align: left;
  background-color: #fff;
  background-clip: padding-box;
  border: 1px solid rgba(0, 0, 0, 0.15);
  border-radius: 4px;
}

#id_user_username {
  cursor: pointer;
  color: #909399;
}

#id_user_username:hover {
  color: #303133;
}

.el-dialog__body {
  padding: 0 25px 0;
}
</style>
