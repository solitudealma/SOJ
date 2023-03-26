<template>
  <div class="problem-container">
    <el-card>
      <template #header>
        <div>
          <h2>题库</h2>
          <el-row>
            <el-col :xs="{span: 22, push: 1}" :sm="{span: 20, push: 2}" :md="{span: 12, push: 6}" :push="6">
              <el-input v-model="searchInput" placeholder="搜索题号、标题、题目来源、算法、题目描述">
              <template #append>
                <el-button :icon="Search" @click="searchProblem()"/>
              </template>
            </el-input>
            </el-col>
          </el-row>
        </div>
      </template>
      <el-table :data="problemData" stripe>
        <el-table-column label="" width="100">
          <template #default="scope">
            <el-icon v-if="scope.row.isAccepted" color="green">
              <IEpSelect />
            </el-icon>
          </template>
        </el-table-column>
        <el-table-column prop="problemId" label="#" width="150"></el-table-column>
        <el-table-column prop="title" label="标题" width="600">
          <template #default="scope">
            <el-link type="primary" :underline="false" @click="getProblemDetail(scope.row.problemId)">{{
              scope.row.title
            }}</el-link>
          </template>
        </el-table-column>
        <el-table-column prop="passingRate" label="通过率" width="100">
          <template #default="scope">
            {{ scope.row.passingRate.toFixed(2) + "%" }}
          </template>
        </el-table-column>
        <el-table-column prop="difficulty" label="难度">
          <template #default="scope">
            <el-tag v-if="scope.row.difficulty === 0" size="small" hit type="success" effect="dark">简单
            </el-tag>
            <el-tag v-else-if="scope.row.difficulty === 1" size="small" hit type="warning" effect="dark">中等
            </el-tag>
            <el-tag v-else-if="scope.row.difficulty === 2" size="small" hit type="danger" effect="dark">困难
            </el-tag>
          </template>
        </el-table-column>
      </el-table>
      <Pagination :total="total" :page-size="limit" @on-change="pushRouter" v-model:current="query.currentPage"
        :layout="'prev, pager, next'">
      </Pagination>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import { Search } from "@element-plus/icons-vue";
import { getProblemListInfo } from "@/api/problem"
import { GetProblemListInfoRequest, Problem } from "@/api/types/problem";

const route = useRoute();
const router = useRouter();

let searchInput = ref("");
let problemData = ref<Array<Problem>>([]);
let total = ref(100);
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

  pullProblemList({ currentPage: query.value.currentPage })
})

const getProblemDetail = (problemId: number) => {
  router.push({
    name: "ProblemListItem",
    params: {
      problemId
    },
  })
};

const searchProblem = () => {
  console.log(searchInput.value)
}

const pushRouter = () => {
  router.push({
    name: "ProblemList",
    query: query.value,
  });
};

const pullProblemList = async (query: GetProblemListInfoRequest) => {
  let res = await getProblemListInfo(query);
  if (res.code === 200) {
    problemData.value = res.data!.problems;
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
</style>
