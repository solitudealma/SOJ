import { defineStore } from 'pinia'
import { store } from '@/pinia'
import { useGlobalStore } from '@/pinia/modules/global'
import storage from '@/utils/storage'
import { ElMessage } from 'element-plus'
import { login, getUserInfo } from '@/api/user'
import type { LoginRequest } from '@/api/types/user'
import type { DialogStatus, UserInfo } from '#/store'

interface UserState {
    userInfo: UserInfo
    token: string
}

export const useUserStore = defineStore('user', {
    state: (): UserState => ({
        userInfo: storage.get('userInfo') || {} as UserInfo,
        token: storage.get('token') || '',
    }),
    getters: {
        isAuthenticated(state): boolean {
            return !!state.token
        },
    },
    actions: {
        async LogIn(loginInfo: LoginRequest) {
            const globalStore = useGlobalStore();
            const res = await login(loginInfo)
            if (res.code === 200) {
                this.setUserInfo(res.data!.userInfo)
                this.setToken(res.data!.token)
                globalStore.changeDialogStatus({ visible: false } as DialogStatus)
                ElMessage.success("欢迎回来～")
            }
        },
        async GetUserInfo() {
            const res = await getUserInfo()
            if (res.code === 200) {
                this.setUserInfo(res.data!.userInfo)
            }
        },
        setUserInfo(val: UserInfo) {
            this.userInfo = val;
            storage.set("userInfo", val)
        },
        setToken(val: string) {
            this.token = val;
            storage.set("token", val);
        },
        clearUserInfoAndToken() {
            this.token = ''
            this.userInfo = {} as UserInfo
            storage.clear()
        },
    },
})

// Need to be used outside the setup
export function useUserStoreWithOut() {
    return useUserStore(store);
}