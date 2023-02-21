import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import routes from '@/router/routes'

const router = createRouter({
    history: createWebHistory(),
    routes: routes as RouteRecordRaw[],
})

export default router
