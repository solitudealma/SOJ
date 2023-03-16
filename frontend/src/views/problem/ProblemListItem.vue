<template>
  <div class="tab-container">
    <el-card class="problem-card" shadow="never">
      <div class="question-name">{{ problemData.problemId }}.{{ problemData.title }}</div>
      <el-tabs type="border-card" v-model="activeTabName" @tab-click="handleClickTab">
        <el-tab-pane name="problemDetail" :lazy="true">
          <template #label>
            <el-icon>
              <IEpDocument />
            </el-icon>题目
          </template>
          <el-row>
            <el-col :span="17">
              <div class="markdown-body">
                <template v-if="problemData.description">
                  <div v-html="problemData.description"></div>
                </template>
                <h4>输入格式</h4>
                <template v-if="problemData.input">
                  <div v-html="problemData.input"></div>
                </template>
                <h4>输出格式</h4>
                <template v-if="problemData.output">
                  <div v-html="problemData.output"></div>
                </template>
                <template v-if="examples">
                  <div v-for="(row, index) in examples" :key="index">
                    <h4>样例输入{{ examples.length === 1 ? "" : index + 1 }}：</h4>
                    <pre class="hljs">{{ row.input }}</pre>
                    <h4>样例输出{{ examples.length === 1 ? "" : index + 1 }}：</h4>
                    <pre class="hljs">{{ row.output }}</pre>
                  </div>
                </template>
              </div>
            </el-col>
            <el-col :span="6" :push="1">
              <el-collapse accordion style="float: right; width: 200px">
                <el-collapse-item title="难度:">
                  <el-tag v-if="problemData.difficulty === 0" type="success" effect="dark">简单
                  </el-tag>
                  <el-tag v-else-if="problemData.difficulty === 1" type="warning" effect="dark">中等
                  </el-tag>
                  <el-tag v-else-if="problemData.difficulty === 2" type="danger" effect="dark">困难
                  </el-tag>
                </el-collapse-item>
                <el-collapse-item title="时/空限制：">
                  <span>{{ problemData.timeLimit / 1000 }}s /
                    {{ problemData.memoryLimit }}MB
                  </span>
                </el-collapse-item>
                <el-collapse-item title="总通过数：">
                  <span>{{ problemData.totalNumberOfPasses }}</span>
                </el-collapse-item>
                <el-collapse-item title="总尝试数：">
                  <div>{{ problemData.totalAttempts }}</div>
                </el-collapse-item>
                <el-collapse-item title="来源：">
                  <div>{{ problemData.source }}</div>
                </el-collapse-item>
                <el-collapse-item title="算法标签">
                  <el-tag v-for="(item, index) in tags" :key="index" :color="item.tagColor" effect="dark">
                    {{ item.tagName }}
                  </el-tag>
                </el-collapse-item>
              </el-collapse>
            </el-col>
          </el-row>
          <CodeEditor v-model="value"></CodeEditor>
          <el-card class="run-code-status-card" body-style="background-color:#fff">
            <div slot="header" class="clearfix">
              <span class="submit-code-status-label">代码运行状态：</span>
              <span style="color: #5fba7d; font-size: 21px">Finished</span>
              <el-button class="close" type="" :text="true">x</el-button>
            </div>
            <div class="card-body">
              <label for="run-code-stdin" style="font-size: 15px">输入</label>
              <el-input id="run-code-stdin" type="textarea" resize="none" :autosize="{ minRows: 1, maxRows: 200 }"
                v-model="textarea">
              </el-input>
              <br />
              <br />
              <span class="run-code-stdin-span">输出</span>
              <pre id="run-code-stdout" style="background-color: #f8f8f8; border-radius: 5px; margin-top: 5px"
                class="markdown-body">
  a.cpp: In function 'int main()':
      a.cpp:6:13: error: expected '}' at end of input
          6 |     return 0;
            |             ^
      a.cpp:2:11: note: to match this '{'
          2 | int main(){
            |           ^</pre>
              <div id="run-code-expected-stdout-block" style="margin-top: 20px">
                <span style="font-weight: normal; font-size: 15px">标准答案</span>
                <pre id="run-code-expected-stdout" style="
                          background-color: #f8f8f8;
                          border-radius: 5px;
                          margin-top: 5px;
                        " class="markdown-body">
  1==1</pre>
              </div>
              <span class="run-code-time">运行时间：1ms</span>
            </div>
          </el-card>
        </el-tab-pane>
        <el-tab-pane name="mySubmission" :lazy="true">
          <template #label>
            <el-icon>
              <IEpList />
            </el-icon> 提交记录
          </template>
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRoute } from "vue-router";
import CodeEditor from "@/components/CodeEditor.vue";
import { getProblemInfo } from "@/api/problem";
import { GetProblemInfoRequest, ProblemInfo } from "@/api/types/problem";
import utils, { Example } from "@/utils/utils";

const route = useRoute();

let textarea = ref("");
let value = ref("");
let activeTabName = "problemDetail";
let problemData = ref<ProblemInfo>({
  problemId: "",
  title: "",
  description: "",
  input: "",
  output: "",
  examples: "",
  difficulty: 0,
  timeLimit: 0,
  memoryLimit: 0,
  totalNumberOfPasses: 0,
  totalAttempts: 0,
  source: "",
  tags: "",
});

interface tagIntem {
  tagName: string;
  tagColor: string;
}

let tags = ref<Array<tagIntem>>([]);
let examples = ref<Array<Example>>([]);

const tagColors = [
  "#00BFFF",
  "#D3D3D3",
  "#FFB6C1",
  "#D8BFD8",
  "#FFA500",
  "#9400D3",
  "#90EE90",
  "#FA8072",
];

const handleClickTab = () => { };

const getProblemDetail = async () => {
  let problemId = route.params.problemId;
  let res = await getProblemInfo({ problemId } as GetProblemInfoRequest);
  if (res.code === 200) {
    examples.value = utils.stringToExamples(res.data!.problemInfo.examples);
    let tmpTags = res.data?.problemInfo.tags.split(",");
    tmpTags!.forEach((item) => {
      tags.value.push({
        tagName: item,
        tagColor: tagColors[(Math.random() * tagColors.length) | 0],
      });
    });
    problemData.value = res.data!.problemInfo;
  }
};

getProblemDetail();
</script>

<style scoped lang="less">
.tab-container {
  height: auto;
  font-size: 14px;
}

.question-name {
  font-size: 28px;
  margin: 20px 0 30px 0;
}

.run-code-status-card {
  margin: 75px 0 0 0;
  background-color: #f5f5f5;
}

.close {
  float: right;
  font-size: 17px;
  font-weight: 700;
  line-height: 1;
  color: #000;
  text-shadow: 0 1px 0 #fff;
  filter: alpha(opacity=20);
  opacity: 0.2;
  -webkit-appearance: none;
  padding: 0;
  cursor: pointer;
  background: 0 0;
  border: 0;
}

.submit-code-status-label {
  font-size: 18px;
}

label {
  display: inline-block;
  max-width: 100%;
  margin-bottom: 5px;
}

span.run-code-stdin-span {
  font-weight: normal;
  font-size: 15px;
}

:deaep(#run-code-stdin) {
  background-color: rgb(248, 248, 248);
  color: #555;
  border-radius: 5px;
  overflow: hidden;
}

#run-code-stdout {
  display: block;
  overflow-x: auto;
  padding: 0.5em;
  background: white;
  color: black;
}
</style>
