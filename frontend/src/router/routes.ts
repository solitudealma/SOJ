import { RouteRecordRaw } from 'vue-router'

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Home',
        component: () => import('@/views/general/HomePage.vue'),
    },
    {
        path: '/problem',
        name: 'ProblemList',
        component: () => import('@/views/problem/ProblemList.vue'),
    },
    {
        path: '/problem/:problemId',
        name: 'ProblemListItem',
        component: () => import('@/views/problem/ProblemListItem.vue'),
    },
    {
        path: '/solution',
        name: 'SolutionList',
        component: () => import('@/views/solution/SolutionList.vue'),
    },
    {
        path: '/judgeinfo',
        name: 'JudgeInfo',
        component: () => import('@/views/about/JudgeInfo.vue'),
    },
    {
        path: '/logout',
        name: 'Logout',
        component: () => import('@/views/user/Logout.vue'),
        meta: { requireAuth: true },
    },
    {
        path: '/faq',
        name: 'Faq',
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
