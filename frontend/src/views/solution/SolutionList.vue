<template>
  <div class="solution-container">
    <el-card>
      <template #header>
        <div>
          <h2>题解</h2>
          <el-row>
            <el-col :xs="{ span: 22, push: 1}" :sm="{ span: 20, push: 2 }" :md="{ span: 12, push: 6 }">
              <el-input v-model="searchInput" placeholder="搜索题号、标题、题目来源、算法、题目描述">
                <template #append>
                  <el-button :icon="Search" @click="searchProblem()" />
                </template>
              </el-input>
            </el-col>
          </el-row>
          <br>
          <el-row>
            <el-col :xs="{ span: 1, push: 22}" :sm="{ span: 1, push: 20 }" :md="{ span: 1, push: 22}">
              <el-button type="primary" size="large">
            <el-icon>
              <IEpEdit />
              </el-icon>
              写题解
            </el-button>
            </el-col>
          </el-row>
        </div>
      </template>
      <el-table :data="solutionData" stripe>
        <el-table-column prop="problemSource" label="题目来源" width="100">
        </el-table-column>
        <el-table-column prop="problemId" label="编号" width="80">
        </el-table-column>
        <el-table-column prop="title" label="标题" width="480">
          <template #default="scope">
            <el-link type="primary" :underline="false" @click="getSolutionDetail(scope.row.solutionId)">
              {{ scope.row.title }}
            </el-link>
          </template>
        </el-table-column>
        <el-table-column prop="authorName" label="作者" width="180">
          <template #default="scope">
            <el-link type="primary" :underline="false" @click="getSolutionAuthorDetail(scope.row.authorId)">
              <el-avatar class="header-img" :size="24" :src="scope.row.authorAvatar" alt="作者的头像"></el-avatar>
              {{ scope.row.authorName }}</el-link>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="时间" width="170">
          <template #default="scope">
            {{ time.formatDate(scope.row.createTime) }}
          </template>
        </el-table-column>
        <el-table-column prop="read" label="阅读"> </el-table-column>
      </el-table>
      <Pagination :total="total" :page-size="limit" @on-change="pushRouter" v-model:current="query.currentPage"
        :layout="'prev, pager, next'">
      </Pagination>
    </el-card>
  </div>
</template>
<script lang="ts" setup>
import { ref, onMounted} from "vue";
import { useRouter, useRoute } from "vue-router";
import { Search } from "@element-plus/icons-vue";
import { getSolutionListInfo } from "@/api/solution";
import { Solution, GetSolutionListInfoRequest} from "@/api/types/solution";
import  time from '@/utils/time'

const route = useRoute();
const router = useRouter();

let solutionData = ref<Array<Solution>>([]);
let searchInput = ref("");
let total = ref(50);
let limit = ref(50);
let query = ref({
  currentPage: 1
});

onMounted(() => {
  let q = route.query;
  query.value.currentPage = parseInt(q.currentPage as string) || 1;
  if (query.value.currentPage < 1) {
    query.value.currentPage = 1;
  }

  pullSolutionList({ currentPage: query.value.currentPage })
})

const searchProblem = () => {
  console.log(searchInput.value)
 }
 

const getSolutionAuthorDetail = (author: string) => { }

const getSolutionDetail = (solutionId: number) => {
  router.push({
    name: "SolutionListItem",
    params: {
      solutionId
    },
  })
};


const pushRouter = () => {
  router.push({
    name: "SolutionList",
    query: query.value,
  });
};

const pullSolutionList = async (query: GetSolutionListInfoRequest) => {
  let res = await getSolutionListInfo(query);
  if (res.code === 200) {
    solutionData.value = res.data!.solutions;
    total.value = res.data!.total;
  }
};
</script>

<style scoped lang="less">
h2 {
  text-align: center;
  font-size: 36px;
  font-family: inherit;
  font-weight: 500;
  line-height: 1.1;
  color: inherit;
}

.solution-container {
  height: auto;
}

.header-img {
  vertical-align: middle;
}
</style>
