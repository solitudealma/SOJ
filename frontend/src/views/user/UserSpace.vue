<template>
    <el-row>
        <el-col :span='5'>
            <el-card :body-style="{ padding: '15px' }">
                <el-avatar id='my-head-photo' class='img-responsive center-block' :src=userInfo.avatar :size='200'
                    shape='square' alt='我的头像'>
                </el-avatar>
                <hr>
                <div>
                    <p style="font-size: 18px; font-weight: 700;text-align: center;margin: 0 0 10px;">
                        {{ userInfo.username }}
                    </p>
                </div>
                <hr>
            </el-card>
        </el-col>
        <el-col :span='18' :push='1'>
            <el-tabs type='border-card' v-model='activeName' @tab-click='handleClick'>
                <el-tab-pane name='solution' label='题解'>
                    <div class='table'>
                        <el-table :data='solutionList' max-height='800px' stripe highlight-current-row>
                            <el-table-column prop='problemId' label='#' width='80'>
                            </el-table-column>
                            <el-table-column prop='title' label='标题' width='350'>
                                <template #default="scope">
                                    <el-link type='primary' :underline='false' @click='getSolutionDetail(scope.row.solutionId)'>
                                        {{ scope.row.title }}
                                    </el-link>
                                </template>
                            </el-table-column>
                            <el-table-column prop='problemDifficulty' label='难度' width='70'>
                                <template #default="scope">
                                    <el-tag v-if="scope.row.problemDifficulty === 0" size='small' hit type='success'
                                        effect='dark'>
                                        简单
                                    </el-tag>
                                    <el-tag v-else-if="scope.row.problemDifficulty === 1" size='small' hit type='warning'
                                        effect='dark'>
                                        中等
                                    </el-tag>
                                    <el-tag v-else-if="scope.row.problemDifficulty === 2" size='small' hit type='danger'
                                        effect='dark'>
                                        困难
                                    </el-tag>
                                </template>
                            </el-table-column>
                            <el-table-column prop='updateTime' label='时间' width='170'>
                                <template #default="scope">
                                    {{ time.formatDate(scope.row.updateTime * 1000) }}
                                </template>
                            </el-table-column>
                            <el-table-column prop='read' label='阅读' width='80'>
                            </el-table-column>
                            <el-table-column prop='problemSource' label='题目来源'>
                            </el-table-column>
                        </el-table>
                    </div>
                    <Pagination :total='total' :page-size="limit" @on-change="pushRouter"
                        v-model:current="query.currentPage" :layout="'prev, pager, next'">
                    </Pagination>
                </el-tab-pane>
            </el-tabs>
        </el-col>
    </el-row>
</template>
  
<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import type { TabsPaneContext } from 'element-plus'
import { useRoute, useRouter } from 'vue-router';
import { useUserStore } from '@/pinia/modules/user';
import { getSolutionListInfo } from '@/api/solution';
import { Solution } from '@/api/types/solution';
import time from '@/utils/time';

const userStore = useUserStore();
const route = useRoute();
const router = useRouter();

let userInfo = computed(() => userStore.userInfo);

let activeName = ref('solution');
let total = ref(50);
let limit = ref(50);
let query = ref({
    currentPage: 1
});

let solutionList = ref<Array<Solution>>([]);

const pullSolutionList = async () => {
    let q = route.query;
    query.value.currentPage = parseInt(q.currentPage as string) || 1;
    if (query.value.currentPage < 1) {
        query.value.currentPage = 1;
    }

    let res = await getSolutionListInfo({ currentPage: query.value.currentPage });
    if (res.code === 200) {
        solutionList.value = res.data!.solutions;
        total.value = res.data!.total;
    }
};

const init = () => {
    pullSolutionList();
}

const handleClick = (tab: TabsPaneContext, event: Event) => {
    console.log()
    switch (tab.paneName!.toString()) {
        case 'solution': {
            let q = route.query;
            query.value.currentPage = parseInt(q.currentPage as string) || 1;
            if (query.value.currentPage < 1) {
                query.value.currentPage = 1;
            }
            pullSolutionList();
            break;
        }
        default:
            break;
    }
}

const pushRouter = () => {
    pullSolutionList();
};

const getSolutionDetail = (solutionId: number) => {
  router.push({
    name: "SolutionListItem",
    params: {
      solutionId
    },
  })
};

onMounted(() => {
    init();
});
</script>
  
<style lang='less' scoped>
hr {
    background: none !important;
    height: 1px !important;
    border: 0 !important;
    border-top: 1px solid #eee !important;
    margin-top: 20px;
    margin-bottom: 20px;
    box-sizing: content-box;
}

a {
    color: #337ab7;
    text-decoration: none;
    background-color: transparent;
}
</style>
  