<template>
    <el-row>
        <el-col :span="6">
            <el-card :body-style="{ padding: '15px' }">
                <el-avatar id='my-head-photo' class='img-responsive center-block' :src=userInfo.avatar :size='250'
                    shape='square' alt='我的头像'>
                </el-avatar>
            </el-card>
        </el-col>
        <el-col :span="17" :push="1">
            <el-card>
                <template #header>
                    <h3>个人信息</h3>
                </template>
                <el-form :model="formUserInfo" label-width="120px">
                    <el-row>
                        <el-col :span="12" :push="4">
                            <el-form-item label="用户名:">
                                <el-input v-model="formUserInfo.username" />
                            </el-form-item>
                        </el-col>
                        <el-col :span="24" :push="6">
                            <el-form-item>
                                <el-button type="primary" @click="submitUserProfile">更新信息</el-button>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </el-form>
            </el-card>
        </el-col>
        <el-col :span="17" :push="7">
            <div class="col">
                <el-card>
                    <template #header>
                        <h3>账号安全</h3>
                    </template>
                    <el-row>
                        <el-col :span="8">
                            <img class="img" src="@/assets/password.jpg" :width="25" alt="密码" />&nbsp;密码
                        </el-col>
                        <el-col :span="8">
                            已设置
                        </el-col>
                        <el-col :span="8">
                            <a href="/user/account/changepassword/">修改密码</a>
                        </el-col>
                    </el-row>
                </el-card>
            </div>
        </el-col>
    </el-row>
</template>
  
<script lang="ts" setup>
import { reactive } from 'vue';
import { useUserStore } from '@/pinia/modules/user';
import { updateUserInfo, getUserInfo} from '@/api/user';

const userStore = useUserStore();

let userInfo = computed(() => userStore.userInfo);

const formUserInfo = reactive({
  username: '',
});

const pullUserProfile = async () => {
    let res = await getUserInfo();
    if (res.code === 200) {
        userStore.setUserInfo(res.data!.userInfo);
    }
}

const submitUserProfile = async () => {
    let res = await updateUserInfo(formUserInfo);
    if (res.code === 200) {
        pullUserProfile();
        console.log(res.data);
    }
}

onMounted(() => {
    pullUserProfile();
})
</script>
  
<style lang="less" scoped>

.col {
    margin: 25px 0 0 0;
}

hr {
    margin-top: 20px;
    margin-bottom: 20px;
    box-sizing: content-box;
    display: block;
    unicode-bidi: isolate;
    margin-block-start: 0.5em;
    margin-block-end: 0.5em;
    margin-inline-start: auto;
    margin-inline-end: auto;
    overflow: hidden;
    background: none !important;
    height: 1px !important;
    border: 0 !important;
    border-top: 1px solid #eee !important;
}

</style>
  