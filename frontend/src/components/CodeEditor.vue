<template>
  <div class='code-editor'>
    <div class='code-tool-bar'>
      <span style="position: relative;left: 10px; top: 16px; font-size: 18px;">写代码啦</span>
      <el-button @click='dialogTableVisible = true' style='float: right;margin: 10px;'>
        <el-icon>
          <IEpTools />
        </el-icon>
      </el-button>
      <el-button @click='refreshCode' style="float: right; margin: 10px;">
        <el-icon>
          <IEpRefresh />
        </el-icon>
      </el-button>
      <el-select v-model='selectLanguageValue' @change='changSelectLangValue' @focus="saveEditorCode"
        style="float: right;margin: 10px;">
        <el-option v-for='item in languagesOptions' :key='item.label' :label='item.label' :value='item.value'>
        </el-option>
      </el-select>
      <el-dialog v-model='dialogTableVisible' :show-close='false' class='code-editor-config-dialog'>
        <el-card style="margin: -59px -20px 0 -20px;" shadow='never'>
          <template #header>
            <span>代码编辑器设置</span>
            <el-button style="float: right; padding-top: 0px; margin-top: -6px;" type="" :text="true"
              @click='dialogTableVisible = false'>x
            </el-button>
          </template>
          <div class='row'>
            <el-row>
              <el-col :span='16'>
                <div class='code-editor-option-title'>主题</div>
                <div class='code-editor-option-description'>对白色界面感到厌倦了吗？可以尝试其他的背景和代码高亮风格。</div>
              </el-col>
              <el-col :span='8'>
                <el-select v-model='selectThemeValue' @change='changSelectThemeValue' filterable>
                  <el-option v-for='item in themesOptions' :key='item.value' :label='item.label' :value='item.value'>
                  </el-option>
                </el-select>
              </el-col>
            </el-row>
            <hr>
          </div>
          <div class='row'>
            <el-row>
              <el-col :span='16'>
                <div class='code-editor-option-title'>缩进长度</div>
                <div class='code-editor-option-description'>选择代码缩进的长度。默认是4个空格。</div>
              </el-col>
              <el-col :span='8'>
                <el-select v-model='selectTabValue' @change='setTabSize' filterable>
                  <el-option v-for='item in tabOption' :key='item.value' :label='item.label' :value='item.value'>
                  </el-option>
                </el-select>
              </el-col>
            </el-row>
            <!-- <hr> -->
          </div>
        </el-card>
        <template #footer>
          <el-button @click='dialogTableVisible = false'>确 定</el-button>
        </template>
      </el-dialog>
    </div>
    <MonocaEditor v-model:modelValue="codeValue" :theme="theme" width="100%" height="75vh" :language="language"
      :options="codeEditorOptions" @update:modelValue="codeValue = $event">
    </MonocaEditor>
    <el-button class='submitBtn' round type='success'><el-icon>
        <IEpUploadFilled />
      </el-icon>&nbsp; 提交代码
    </el-button>
    <el-button class='debuggerBtn' round @click='debuggerCode'><el-icon>
        <IEpVideoPlay />
      </el-icon>&nbsp; 调试代码
    </el-button>
  </div>
  <el-card v-if="hasSubmit || hasDebug" class="run-code-status-card" body-style="background-color:#fff">
    <template #header>
      <span class="submit-code-status-label">代码运行状态：</span>
      <span style="color: #5fba7d; font-size: 21px">{{ submitCodeStatus }}</span>
      <el-button class="close" type="" :text="true" @click="closeRunCodeStatusCard">x</el-button>
    </template>
    <label for="run-code-stdin" style="font-size: 15px; font-weight: normal;">输入</label>
    <el-input id="run-code-stdin" type="textarea" resize="none" :autosize="{ minRows: 1, maxRows: 200 }"
      v-model="userInput">
    </el-input>
    <br />
    <br />
    <span class="run-code-stdin-span">输出</span>
    <pre id="run-code-stdout" style="background-color: #f8f8f8; border-radius: 5px; margin-top: 5px; min-height: 40px;"
      class="markdown-body">{{ stdOutOrErr }}</pre>
    <div v-if="hasSubmit" id="run-code-expected-stdout-block" style="margin-top: 20px">
      <span style="font-weight: normal; font-size: 15px">标准答案</span>
      <pre id="run-code-expected-stdout" class="markdown-body">{{ expectedOutput }}</pre>
    </div>
    <span class="run-code-time">运行时间：{{ runCodeTime }}ms</span>
  </el-card>
</template>
  
<script lang="ts" setup>
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import { editor } from 'monaco-editor';
import MonocaEditor from '@/components/MonocaEditor.vue'
import { useUserStore } from '@/pinia/modules/user';
import { debugProblem } from '@/api/problem';
import { DebugProblemRequest } from '@/api/types/problem';
import storage from '@/utils/storage';

const route = useRoute();
const useStore = useUserStore();

let isAuthenticated = computed(() => useStore.isAuthenticated);
let userInfo = computed(() => useStore.userInfo);
let problemId = computed(() => route.params.problemId);
let codeValue = ref('');
let theme = ref('vs-dark');
let language = ref("cpp");
let codeEditorOptions = ref<editor.IStandaloneEditorConstructionOptions>({
  tabSize: 4,
});
let dialogTableVisible = ref(false);
let selectTabValue = ref(codeEditorOptions.value.tabSize);
let tabOption = [{
  value: 2,
  label: '2个空格'
}, {
  value: 4,
  label: '4个空格'
}, {
  value: 6,
  label: '6个空格'
}];
let selectThemeValue = ref(theme.value);
let themesOptions = [{
  value: 'vs',
  label: 'vs'
}, {
  value: 'vs-dark',
  label: 'vs-dark'
}, {
  value: 'hc-black',
  label: 'hc-black'
}];
let selectLanguageValue = ref(language.value)
let languagesOptions = [{
  value: 'cpp',
  label: 'C++'
}, {
  value: 'c',
  label: 'C'
}, {
  value: 'java',
  label: 'Java'
}, {
  value: 'go',
  label: 'Golang'
}, {
  value: 'python',
  label: 'Python3'
}, {
  value: 'javascript',
  label: 'Javascript'
}];
let userInput = ref("");
let hasSubmit = ref(false);
let hasDebug = ref(false);
let submitCodeStatus = ref("");
let stdOutOrErr = ref("");
let runCodeTime = ref(0);
let expectedOutput = ref("");

const refreshCode = () => {
  codeValue.value = ''
}

const changSelectLangValue = (value: string) => {
  language.value = selectLanguageValue.value = value;
  loadEditorCode();
  saveEditorLanguage();
}

const closeRunCodeStatusCard = () => {
  hasDebug.value = hasSubmit.value = false;
}

const debuggerCode = async () => {
  hasDebug.value = true;
  let req = <DebugProblemRequest>{
    problemId: problemId.value,
    code: codeValue.value,
    language: language.value,
    userInput: userInput.value,
  }

  let res = await debugProblem(req)
  if (res.code === 200) {
    console.log(res.data!.result)
    let result = res.data!.result
    submitCodeStatus.value = result.status === "Accepted" ? "Finished" : result.status;
    runCodeTime.value = result.time;
    stdOutOrErr.value = result.stderr !== "" ? result.stderr : result.userOutput;
    expectedOutput.value = result.expectedOutput;
  }
}

const changSelectThemeValue = (value: string) => {
  theme.value = selectThemeValue.value = value
}

const setTabSize = (value: number) => {
  codeEditorOptions.value.tabSize = value
}

const loadEditorLanguage = () => {
  let hasSetValue = false;
  if (typeof (localStorage) !== 'undefined') {
    let userId = 'anonymous';
    if (isAuthenticated.value) {
      userId = userInfo.value.userId.toString();
    }

    let key = `${userId}-selectedLang`;
    let usedLanguage = storage.get(key);
    if (typeof usedLanguage === 'string' && usedLanguage.length > 0) {
      language.value = selectLanguageValue.value = usedLanguage;
      hasSetValue = true;
    } else if (userId !== 'anonymous') {
      let key = `anonymous-selectedLang`;
      usedLanguage = storage.get(key);
      if (typeof usedLanguage === 'string' && usedLanguage.length > 0) {
        language.value = selectLanguageValue.value = usedLanguage;
        hasSetValue = true;
      }
    }
  }

  if (!hasSetValue) {
    refreshCode();
  }
};

const saveEditorLanguage = () => {
  if (typeof localStorage !== 'undefined') {
    let userId = 'anonymous';
    if (isAuthenticated.value) {
      userId = userInfo.value.userId.toString()
    }

    let key = `${userId}-selectedLang`;
    let value = language.value;
    storage.set(key, value);
  }
};

const loadEditorCode = () => {
  let hasSetValue = false;
  if (typeof (localStorage) !== 'undefined') {
    codeValue.value = '';
    let userId = 'anonymous';
    if (isAuthenticated.value) {
      userId = userInfo.value.userId.toString();
    }
    let language = selectLanguageValue.value;
    let key = `${problemId.value}-${userId}-${language}`;
    let code = storage.get(key);
    if (typeof code === 'string' && code.length > 0) {
      codeValue.value = code;
      hasSetValue = true;
    } else if (userId !== 'anonymous') {
      let key = `${problemId.value}-anonymous-${language}`;
      code = storage.get(key);
      if (typeof code === 'string' && code.length > 0) {
        codeValue.value = code;
        hasSetValue = true;
      }
    }
  }
  if (!hasSetValue) {
    refreshCode();
  }
};

const saveEditorCode = () => {
  if (typeof localStorage !== 'undefined') {
    let userId = 'anonymous';
    if (isAuthenticated.value) {
      userId = userInfo.value.userId.toString()
    }

    let language = selectLanguageValue.value;
    let key = `${problemId.value}-${userId}-${language}`;
    let value = codeValue.value;
    storage.set(key, value);
  }
};

const beforeunloadHandler = () => {
  // 刷新前(页面被摧毁)保存代码和选择的语言
  saveEditorLanguage();
  saveEditorCode();
}

onMounted(() => {
  loadEditorLanguage();
  loadEditorCode();
  window.addEventListener('beforeunload', e => beforeunloadHandler());
})

onUnmounted(() => {
  window.removeEventListener('beforeunload', e => beforeunloadHandler());
})

</script>
  
<style scoped lang='less'>
.code-tool-bar {
  height: 60px;
  width: 100%;
  background: #f8f9fa;
  border: 1px solid #c2c7d0;
  margin-bottom: 0;
}

.code-editor-option-title {
  font-size: 17px;
  margin-bottom: 10px;
}

.code-editor-option-description {
  font-size: 13px;
  color: grey;
}

hr {
  background: none !important;
  height: 1px !important;
  border: 0 !important;
  border-top: 1px solid #eee !important;
  margin-top: 20px;
  margin-bottom: 20px;
}

.code-editor {
  height: 100%;
  width: 100%;
}

.debuggerBtn {
  float: right;
  margin: 13px 20px 0 0;
}

.submitBtn {
  float: right;
  margin: 13px 0 0 0;
}

pre {
  display: block;
  padding: 9.5px;
  margin: 0 0 10px;
  font-size: 13px;
  line-height: 1.42857143;
  color: #333;
  word-break: break-all;
  word-wrap: break-word;
  background-color: #f5f5f5;
  border: 1px solid #ccc;
  border-radius: 4px;
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

#run-code-stdin {
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

#run-code-expected-stdout {
  background-color: #f8f8f8;
  border-radius: 5px;
  margin-top: 5px;
}
</style>
  