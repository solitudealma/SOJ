import { RouteRecordRaw } from 'vue-router'

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'home',
        component: () => import('@/views/general/HomePage.vue'),
    },
    {
        path: '/problem',
        name: 'problemList',
        component: () => import('@/views/problem/ProblemList.vue'),
    },
    {
        path: '/solution',
        name: 'solutionList',
        component: () => import('@/views/solution/SolutionList.vue'),
    },
    {
        path: '/judgeinfo',
        name: 'judgeInfo',
        component: () => import('@/views/about/JudgeInfo.vue'),
    },
    {
        path: '/faq',
        name: 'faq',
        component: () => import('@/views/about/FAQ.vue'),
    },
    {
        path: '/404',
        name: '404',
        component: () => import('@/views/general/404.vue'),
    },
    {
        path: '/:catchAll(.*)',
        redirect: '/404',
    },
];

export default routes;
