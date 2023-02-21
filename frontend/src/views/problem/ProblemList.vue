<template>
  <div class="problem-container">
    <el-card>
      <template #header>
        <span>题库</span>
        <el-input
          v-model="input"
          placeholder="搜索题号、标题、题目来源、算法、题目描述"
        >
          <template #append>
            <el-button :icon="Search" />
          </template>
        </el-input>
      </template>
      <el-table :data="data" stripe>
        <el-table-column prop="passFlag" label="" width="100">
          <template #default="scope">
            <i
              v-if="scope.row.pass_flag"
              class="fa fa-check"
              aria-hidden="true"
              style="color: #5fba7d"
            ></i>
          </template>
        </el-table-column>
        <el-table-column prop="index" label="#" width="150"> </el-table-column>
        <el-table-column prop="title" label="标题" width="600">
          <template #default="scope">
            <el-link type="primary" :underline="false" @click="pushRouter1">{{
              scope.row.title
            }}</el-link>
          </template>
        </el-table-column>
        <el-table-column prop="rate" label="通过率" width="200">
        </el-table-column>
        <el-table-column prop="difficulty" label="难度">
          <template #default="scope">
            <el-tag
              v-if="scope.row.difficulty === 0"
              size="small"
              hit
              type="success"
              effect="dark"
              >简单
            </el-tag>
            <el-tag
              v-else-if="scope.row.difficulty === 1"
              size="small"
              hit
              type="warning"
              effect="dark"
              >中等
            </el-tag>
            <el-tag
              v-else-if="scope.row.difficulty === 2"
              size="small"
              hit
              type="danger"
              effect="dark"
              >困难
            </el-tag>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { Search } from "@element-plus/icons-vue";
let input = "";
const data = [];

const pushRouter1 = () => {};
</script>

<style scoped lang="less"></style>
