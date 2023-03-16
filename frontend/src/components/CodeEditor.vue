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
      <el-select v-model='selectLanguageValue' @change='changSelectLangValue' @focus="saveEditorCode" :filterable='true'
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
      :options="codeEditorOptions" @update:value="codeValue = $event">
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
</template>
  
<script lang="ts" setup>
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import { editor } from 'monaco-editor';
import MonocaEditor from '@/components/MonocaEditor.vue'
import { useUserStore } from '@/pinia/modules/user';
import storage from '@/utils/storage';

const route = useRoute();
const useStore = useUserStore();

let isAuthenticated = computed(() => useStore.isAuthenticated);
let userInfo = computed(() => useStore.userInfo);

let codeValue = ref('');
let theme = ref('vs-dark');
let language = ref("javascript");
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
}, {
  value: 'xcode',
  label: 'Xcode'
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

const refreshCode = () => {
  codeValue.value = ''
}

const changSelectLangValue = (value: string) => {
  language.value = selectLanguageValue.value = value;
  loadEditorCode();
  saveEditorLanguage();
}

const debuggerCode = () => {
  let value = codeValue.value;
  console.log(value);
}

const changSelectThemeValue = (value: string) => {
  theme.value = selectThemeValue.value = value
}

const setTabSize = (value: number) => {
  console.log(typeof value);
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
    let problemId = route.params.problemId;
    if (isAuthenticated.value) {
      userId = userInfo.value.userId.toString();
    }
    let language = selectLanguageValue.value;
    let key = `${problemId}-${userId}-${language}`;
    let code = storage.get(key);
    if (typeof code === 'string' && code.length > 0) {
      codeValue.value = code;
      hasSetValue = true;
    } else if (userId !== 'anonymous') {
      let key = `${problemId}-anonymous-${language}`;
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
    let problemId = route.params.problemId;
    if (isAuthenticated.value) {
      userId = userInfo.value.userId.toString()
    }

    let language = selectLanguageValue.value;
    let key = `${problemId}-${userId}-${language}`;
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
</style>
  