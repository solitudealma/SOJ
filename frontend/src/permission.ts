import { RouteLocationNormalized, NavigationGuardNext} from 'vue-router'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { useGlobalStore } from './pinia/modules/global'
import router from '@/router'
import Nprogress from 'nprogress'

router.beforeEach(async(to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
  Nprogress.start()
  const userStore = useUserStore();
  const globalStore = useGlobalStore();
  if (to.matched.some(record => record.meta.requireAuth)) { // 判断该路由是否需要登录权限
    const token = userStore.token;
    if (token) { // 判断当前的token是否存在 ； 登录存入的token
        next()
    } else { // 如果没有token
      if(to.path.split('/')[1]==='admin'){
        next({
          path: '/admin/login'  // 管理端无token认证返回登录页
        })
      }else{
        next({
          path: '/'  // 无token认证的一致返回到主页
        })
        globalStore.changeDialogStatus({mode: 'Login', visible: true})
      }
      userStore.clearUserInfoAndToken();
      ElMessage.error("请先登录!")
    }
  } else { // 不需要登录认证的页面
    next()
  }
})

router.afterEach(() => {
  // 路由加载完成后关闭进度条
  document.getElementsByClassName('main-cont main-right')[0]?.scrollTo(0, 0)
  Nprogress.done()
})

router.onError(() => {
  // 路由发生错误后销毁进度条
  Nprogress.remove()
})
