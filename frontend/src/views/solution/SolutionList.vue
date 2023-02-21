<template>
  <div class="solution-container">
    <el-card>
      <template #header>
        <span>题解</span>
        <el-input
          v-model="input"
          type="text"
          placeholder="搜索题号、标题、题目来源、算法、题目描述"
        >
          <template #append>
            <el-button :icon="Search" />
          </template>
        </el-input>
        <el-button type="primary" size="large">
          <el-icon>
            <IEpEdit />
          </el-icon>
          写题解
        </el-button>
      </template>
      <el-table :data="solutionList" stripe>
        <el-table-column prop="problemSource" label="题目来源" width="150">
        </el-table-column>
        <el-table-column prop="sourceId" label="编号" width="80">
        </el-table-column>
        <el-table-column prop="solutionTitle" label="标题" width="520">
          <template #default="scope">
            <el-link
              type="primary"
              :underline="false"
              @click="pushRouter1(scope.row.sourceId)"
              >{{ scope.row.solutionTitle }}</el-link
            >
          </template>
        </el-table-column>
        <el-table-column prop="solutionAuthor" label="作者" width="200">
          <template #default="scope">
            <el-link type="primary" :underline="false" @click="pushRouter2(scope.row.solutionAuthor)">
              <el-avatar
                class="header-img"
                :size="24"
                :src="scope.row.solutionAuthorAvatar"
                alt="作者的头像"
              ></el-avatar>
              {{ scope.row.solutionAuthor }}</el-link
            >
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="时间" width="160">
        </el-table-column>
        <el-table-column prop="read" label="阅读"> </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>
<script lang="ts" setup>
import { Search } from "@element-plus/icons-vue";
const solutionList = [];
let input = "";
const pushRouter1 = (sourceId: number) => {}
const pushRouter2 = (author: string) => {}
</script>

<style lang="less" scoped>
.clearfix span {
  margin-left: 40vw;
}

.edit-btn {
  float: right;
  margin: 12px 2px 0 -105px;
}

.solution-container {
  height: auto;
}

.header-img {
  vertical-align: middle;
}
</style>
