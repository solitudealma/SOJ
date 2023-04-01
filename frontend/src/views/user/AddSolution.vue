<template>
    <div class="add_a_solution_form">
        <el-card>
            <template #header>
                <el-form :model="formSolution" :rules="ruleSolution" ref="formSolutionRef">
                    <el-row :gutter="80">
                        <el-col :span="8">
                            <el-form-item prop="problemSource">
                                <el-select v-model="formSolution.problemSource">
                                    <el-option v-for="item in sourceOptions" :key="item.label" :label="item.label"
                                        :value="item.value">
                                    </el-option>
                                </el-select>
                            </el-form-item>
                        </el-col>
                        <el-col :span="8">
                            <el-form-item prop="problemId">
                                <el-input v-model="formSolution.problemId" placeholder="编号"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="8">
                            <el-form-item prop="problemDifficulty">
                                <el-select v-model="formSolution.problemDifficulty">
                                    <el-option v-for="item in difficultyOption" :key="item.label" :label="item.label"
                                        :value="item.value">
                                    </el-option>
                                </el-select>
                            </el-form-item>
                        </el-col>
                    </el-row>
                    <el-row :gutter="10">
                        <el-col :span="12">
                            <el-form-item prop="title">
                                <el-input v-model="formSolution.title" placeholder="标题"></el-input>
                            </el-form-item>
                        </el-col>
                        <el-col :span="12">
                            <el-form-item prop="problemLink">
                                <el-input v-model="formSolution.problemLink" placeholder="题目链接"></el-input>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </el-form>
            </template>
            <div class='section-markdown-editor'>
                <el-row>
                    <el-col :span='6'>
                        <el-row>
                            <el-col :span='7'>
                                <a class='editor-tab-content' :class='isEditActive' @click='showEdit'>
                                    <i class="fa fa-pencil-square-o" aria-hidden="true"></i>
                                    编辑
                                </a>
                            </el-col>
                            <el-col :span='7'>
                                <a class='preview-tab-content' :class='isPreviewActive' @click='showPreview'>
                                    <i class='fa fa-eye' aria-hidden='true'></i>
                                    预览
                                </a>
                            </el-col>
                        </el-row>
                    </el-col>
                    <el-col :push='8' :span='18' v-show='!isShowPreview'>
                        <div class='toolbar'>
                            <el-row>
                                <el-col :span='1'>
                                    <div>
                                        <i @click='insertBoldCode' class='fa fa-bold' aria-hidden='true'></i>
                                    </div>
                                </el-col>
                                <el-col :span='1'>
                                    <div>
                                        <i @click='insertItalicCode' class='fa fa-italic' aria-hidden='true'></i>
                                    </div>
                                </el-col>
                                <el-col :span='1'>
                                    <div>
                                        <i @click='insertMinusCode' class='fa fa-minus' aria-hidden='true'></i>
                                    </div>
                                </el-col>
                                <el-col :span='1'>
                                    <el-popover placement='bottom' width='125' transition='fade-in-linear' trigger='click'
                                        content=''>
                                        <template #reference>
                                            <i class='fa fa-header' aria-hidden='true'></i>
                                        </template>
                                        <div>
                                            <div class='header1-btn' :class='isHeader1Active' @click='insertHeader1Code'>
                                                标题 1 (Ctrl+Alt+1)
                                            </div>
                                            <div class='header2-btn' :class='isHeader2Active' @click='insertHeader2Code'>
                                                标题 2 (Ctrl+Alt+2)
                                            </div>
                                            <div class='header3-btn' :class='isHeader3Active' @click='insertHeader3Code'>
                                                标题 3 (Ctrl+Alt+3)
                                            </div>
                                        </div>
                                    </el-popover>
                                </el-col>
                                <el-col :span='1'>
                                    <el-popover placement='bottom' width='125' transition='fade-in-linear' trigger='click'>
                                        <template #reference>
                                            <i class="fa fa-code" aria-hidden="true"></i>
                                        </template>
                                        <div>
                                            <div class='text-btn' :class='isTextActive' @click='insertText'>
                                                文本 (Ctrl+Alt+P)
                                            </div>
                                            <div class='code-btn' :class='isCodeActive' @click='insertCode'>
                                                代码 (Ctrl+Alt+C)
                                            </div>
                                        </div>
                                    </el-popover>
                                </el-col>
                                <el-col :span='1'>
                                    <div>
                                        <i @click='insertQuoteCode' class='fa fa-quote-left' aria-hidden='true'></i>
                                    </div>
                                </el-col>
                                <el-col :span='1'>
                                    <div>
                                        <i @click='insertUlCode' class='fa fa-list-ul' aria-hidden='true'></i>
                                    </div>
                                </el-col>
                                <el-col :span='1'>
                                    <div>
                                        <i @click='insertOlCode' class='fa fa-list-ol' aria-hidden='true'></i>
                                    </div>
                                </el-col>
                                <el-col :span='1'>
                                    <div>
                                        <i @click='insertLinkCode' class='fa fa-link' aria-hidden='true'></i>
                                    </div>
                                </el-col>
                                <el-col :span='1'>
                                    <div>
                                        <i @click='insertImgCode' class='fa fa-picture-o' aria-hidden='true'></i>
                                    </div>
                                </el-col>
                                <el-col :span='1'>
                                    <div>
                                        <el-upload class='upload-demo' action='https://jsonplaceholder.typicode.com/posts/'
                                            :limit='1'>
                                            <i class='fa fa-cloud-upload' aria-hidden='true'></i>
                                        </el-upload>
                                    </div>
                                </el-col>
                                <el-col :span='1'>
                                    <div>
                                        <i @click='selectEmoji' class='fa fa-smile-o' aria-hidden='true'></i>
                                    </div>
                                </el-col>
                                <el-col :span='1'>
                                    <div>
                                        <i @click='toggleMaximize' class='fa fa-arrows-alt' aria-hidden='true'></i>
                                    </div>
                                </el-col>
                                <el-col :span='1'>
                                    <i @click='toggleHelp' class='fa fa-question-circle' aria-hidden='true'></i>
                                    <el-dialog v-model='dialogHelpVisible' :show-close='false' top='5vh' width='60%'
                                        :append-to-body='true' :close-on-press-escape='true'>
                                        <el-card style="margin: -60px -20px -30px -20px">
                                            <template #header>
                                                <i class='fa fa-question-circle' aria-hidden='true'><span>Markdown
                                                        Guide</span></i>
                                            </template>
                                            <div>
                                                <div style="color: red">Warning: 代码块尽可能地指定语言，以加速代码高亮，减少自动检测语言次数</div>
                                                This site is powered by Markdown. For full documentation,
                                                <a href='http://commonmark.org/help/' target='_blank'>click here</a>
                                            </div>
                                            <el-table :data='tableData' stripe border :highlight-current-row='true'
                                                style="width: 100%">
                                                <el-table-column prop='code' label='Code' width='150'>
                                                    <template #default="scope">
                                                        <div v-html='scope.row.code'></div>
                                                    </template>
                                                </el-table-column>
                                                <el-table-column prop='or' label='Or' width='180'>
                                                    <template #default="scope">
                                                        <div v-html='scope.row.or'></div>
                                                    </template>
                                                </el-table-column>
                                                <el-table-column prop='devices' label='Linux/Windows'>
                                                </el-table-column>
                                                <el-table-column prop='device' label='Mac OS' width='180'>
                                                </el-table-column>
                                                <el-table-column prop='showOff' label='... to Get' width='200'>
                                                    <template #default="scope">
                                                        <div v-html='scope.row.showOff'></div>
                                                    </template>
                                                </el-table-column>
                                            </el-table>
                                        </el-card>
                                    </el-dialog>
                                </el-col>
                            </el-row>
                        </div>
                    </el-col>
                </el-row>
            </div>
            <br>
            <div class='show-panel'>
                <transition name='edit'>
                    <MonocaEditor ref="monacoEditor" v-model:modelValue="codeValue" :theme="theme" width="100%"
                        height="75vh" :language="language" :options="codeEditorOptions"
                        @update:modelValue="codeValue = $event" v-show='!isShowPreview'>
                    </MonocaEditor>
                </transition>
                <transition name='preview'>
                    <div class='panel-preview markdown-body' ref='markdownPreview' v-show='isShowPreview'></div>
                </transition>
            </div>
            <el-button class="submit-btn" type="success" round @click="submitSolution(formSolutionRef)">提交</el-button>
            <el-button class="save-btn" type="success" round @click="saveSolution(formSolutionRef)">保存</el-button>
        </el-card>
    </div>
</template>
  
<script lang="ts" setup>
import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage, FormInstance, FormRules } from "element-plus";
import { editor } from 'monaco-editor';
import DOMPurify from 'dompurify';
import markdown from "@/utils/markdown";
import { useUserStore } from '@/pinia/modules/user';
import { getSavedSolutionInfo, createSolutionInfo } from "@/api/solution";

const router = useRouter();
const userStore = useUserStore();

let userInfo = computed(() => userStore.userInfo);
let formSolutionRef = ref<FormInstance>();
let difficultyOption = ref([{
    value: 1,
    label: "简单"
}, {
    value: 2,
    label: "中等"
}, {
    value: 2,
    label: "困难"
}]);
let sourceOptions = ref([{
    value: 1,
    label: "SOJ",
},
{
    value: 1,
    label: 'AcWing'
}, {
    value: 2,
    label: 'LeetCode'
}, {
    value: 3,
    label: 'CodeForces'
}, {
    value: 4,
    label: 'POJ'
}, {
    value: 5,
    label: 'HDU'
}, {
    value: 6,
    label: '洛谷'
}, {
    value: 7,
    label: 'AtCoder'
}, {
    value: 8,
    label: '信息学奥赛一本通'
}, {
    value: 9,
    label: 'PAT'
}, {
    value: 10,
    label: '牛客'
}]);
let formSolution = reactive({
    problemId: "",
    problemSource: "SOJ",
    problemDifficulty: 1,
    title: "",
    content: "",
    authorId: userInfo.value.userId,
    authorName: userInfo.value.username,
    authorAvatar: userInfo.value.avatar,
    problemLink: "",
    type: 0,
});
let ruleSolution = reactive<FormRules>({
    problemId: [
        { required: true, message: "请输入编号", trigger: 'blur' }
    ],
    title: [
        { required: true, message: "请输入标题", trigger: 'blur' }
    ],
    problemLink: [
        { required: true, message: "请输入题目链接", trigger: 'blur' }
    ]
});
let monacoEditor = ref();
let codeValue = ref('');
let theme = ref('vs-dark');
let language = ref("markdown");
let codeEditorOptions = ref<editor.IStandaloneEditorConstructionOptions>({
    tabSize: 4,
});
let tableData = [{
    code: ':emoji_name:',
    or: '—',
    devices: '—',
    device: '—',
    showOff: '🧡'
}, {
    code: '*Italic*',
    or: '_Italic_',
    devices: 'Ctrl+I',
    device: 'Command+I',
    showOff: '<em>Italic</em>'
}, {
    code: '**Bold**',
    or: '__Bold__',
    devices: 'Ctrl+B',
    device: 'Command+B',
    showOff: '<strong>Bold</strong>'
}, {
    code: '++Underscores++',
    or: '—',
    devices: 'Shift+U',
    device: 'Option+U',
    showOff: '<ins>Underscores</ins>'
}, {
    code: '~~Strikethrough~~',
    or: '—',
    devices: 'Shift+S',
    device: 'Option+S',
    showOff: '<del>Strikethrough</del>'
}, {
    code: '# Heading 1',
    or: 'Heading 1<br>=========',
    devices: 'Ctrl+Alt+1',
    device: 'Command+Option+1',
    showOff: '<h1>Heading 1</h1>'
}, {
    code: '## Heading 2',
    or: 'Heading 2<br>-----------',
    devices: 'Ctrl+Alt+2',
    device: 'Command+Option+2',
    showOff: '<h2>Heading 1</h2>'
}, {
    code: '[Link](https://a.com)',
    or: '[Link][1]<br>⁝<br>[1]: https://b.org',
    devices: 'Ctrl+L',
    device: 'Command+L',
    showOff: '<a href="https://commonmark.org/">Link</a>'
}, {
    code: '![Image](http://url/a.png)',
    or: '![Image][1]<br>⁝<br>[1]: http://url/b.jpg',
    devices: 'Ctrl+Shift+I',
    device: 'Command+Option+I',
    showOff: '<img src="https://cdn.acwing.com/static/plugins/images/commonmark.png" width="36" height="36" alt="Markdown">'
}, {
    code: '> Blockquote',
    or: '—',
    devices: 'Ctrl+Q',
    device: 'Command+Q',
    showOff: '<blockquote>Blockquote</blockquote>'
}, {
    code: 'A paragraph.<br><br>A paragraph after 1 blank line.',
    or: '—',
    devices: '—',
    device: '—',
    showOff: '<p>A paragraph.</p><p>A paragraph after 1 blank line.</p>'
}, {
    code: '<p>* List<br> * List<br> * List</p>',
    or: '<p> - List<br> - List<br> - List<br></p>',
    devices: 'Ctrl+U',
    device: 'Command+U',
    showOff: '<ul><li>List</li><li>List</li><li>List</li></ul>'
}, {
    code: '<p> 1. One<br> 2. Two<br> 3. Three</p>',
    or: '<p> 1) One<br> 2) Two<br> 3) Three</p>',
    devices: 'Ctrl+Shift+O',
    device: 'Command+Option+O',
    showOff: '<ol><li>One</li><li>Two</li><li>Three</li></ol>'
}, {
    code: 'Horizontal Rule<br><br>-----------',
    or: 'Horizontal Rule<br><br>***********',
    devices: 'Ctrl+H',
    device: 'Command+H',
    showOff: 'Horizontal Rule<hr>'
}, {
    code: '`Inline code` with backticks',
    or: '—',
    devices: 'Ctrl+Alt+C',
    device: 'Command+Option+C',
    showOff: '<code>Inline code</code>with backticks'
}, {
    code: '```<br> def whatever(foo):<br>&nbsp;&nbsp;&nbsp;&nbsp;return foo<br>```',
    or: '<b>with tab / 4 spaces</b><br>....def whatever(foo):<br>....&nbsp;&nbsp;&nbsp;&nbsp;return foo',
    devices: 'Ctrl+Alt+P',
    device: 'Command+Option+P',
    showOff: `<pre class="hljs"><code class=""><span class="hljs-keyword">def</span> <span class="hljs-title function_">whatever</span>(<span class="hljs-params">foo</span>):
    <span class="hljs-keyword">return</span> foo</code></pre>`
}]
let dialogHelpVisible = ref(false);
let isTextActive = ref('');
let isCodeActive = ref('');
let isHeader1Active = ref('');
let isHeader2Active = ref('');
let isHeader3Active = ref('');
let markdownPreview = ref();
let isShowPreview = ref(false);
let isEditActive = ref('active');
let isPreviewActive = ref('');

const pullSavedSolution = async() => { 
    let res = await getSavedSolutionInfo( { authorId: userInfo.value.userId });
    if (res.code === 200) {
        if (res.data !== null) {
            formSolution = Object.assign(formSolution, res.data!);
            codeValue.value = formSolution.content;
            console.log(formSolution)
        }
    }
}

const saveSolution = (formEl: FormInstance | undefined) => {
    if (!formEl) return;
    formEl.validate(async (valid) => {
        if (valid) {
            formSolution.type = 0;
            formSolution.content = codeValue.value;
            let res = await createSolutionInfo(formSolution);
            if (res.code === 200) {
                ElMessage.success("保存成功～")
            }
        } else {
            ElMessage({
                type: "error",
                message: "编号、标题、题目链接不能为空",
                showClose: true,
            });
        }
    });
};

const submitSolution = (formEl: FormInstance | undefined) => {
    if (!formEl) return;
    formEl.validate(async (valid) => {
        if (valid) {
            formSolution.type = 1;
            formSolution.content = codeValue.value;
            let res = await createSolutionInfo(formSolution);
            if (res.code === 200) {
                ElMessage.success("提交成功～")
                router.push({
                    name: "SolutionListItem",
                    params: {
                        solutionId: res.data!.solutionId
                    },
                });
            }
        } else {
            ElMessage({
                type: "error",
                message: "编号、标题、题目链接不能为空",
                showClose: true,
            });
        }
    });
};

const insertBoldCode = () => {
    monacoEditor.value.insertText(" ****", -2);
}

const insertItalicCode = () => {
    monacoEditor.value.insertText(" __", -1);
}

const insertMinusCode = () => {
    monacoEditor.value.insertText("\n\n----------\n\n", 0, 5);
}

const insertHeader1Code = () => {
    isHeader2Active.value = isHeader3Active.value = '';
    isHeader1Active.value = 'active';
    monacoEditor.value.insertText("\n\n# ", 0, 2);
}

const insertHeader2Code = () => {
    isHeader1Active.value = isHeader3Active.value = '';
    isHeader2Active.value = 'active';
    monacoEditor.value.insertText("\n\n## ", 0, 2);
}

const insertHeader3Code = () => {
    isHeader1Active.value = isHeader2Active.value = '';
    isHeader3Active.value = 'active';
    monacoEditor.value.insertText("\n\n### ", 0, 2);
}

const insertText = () => {
    isCodeActive.value = '';
    isTextActive.value = 'active';
    monacoEditor.value.insertText("\n\n```\n\n```\n", 0, 3);
}

const insertCode = () => {
    isTextActive.value = '';
    isCodeActive.value = 'active';
    monacoEditor.value.insertText(" ``", -1);
}

const insertQuoteCode = () => {
    monacoEditor.value.insertText("\n\n> \n", 0, 2);
}

const insertUlCode = () => {
    monacoEditor.value.insertText("\n\n* ", 0, 2);
}

const insertOlCode = () => {
    monacoEditor.value.insertText("\n\n1. ", 0, 2);
}

const insertLinkCode = () => {
    monacoEditor.value.insertText(" []()", -3);
}

const insertImgCode = () => {
    monacoEditor.value.insertText("![]()", -3);
}

const uploadImg = () => {
    monacoEditor.value.insertText("![]()", -3);
}

const selectEmoji = () => {
    monacoEditor.value.insertText("****", -2);
}

const toggleMaximize = () => {
    monacoEditor.value.insertText("****", -2);
}

const toggleHelp = () => {
    dialogHelpVisible.value = !dialogHelpVisible.value;
}

const showEdit = () => {
    isShowPreview.value = false;
    markdownPreview.value.innerHTML = '';
    isEditActive.value = 'active';
    isPreviewActive.value = '';
}

const showPreview = () => {
    isShowPreview.value = true;
    show();
    isEditActive.value = '';
    isPreviewActive.value = 'active';
}

const show = () => {
    let value = codeValue.value;
    markdownPreview.value.innerHTML = value === '' ? DOMPurify.sanitize('<p>内容为空</p>') : DOMPurify.sanitize(markdown(value));
}

onMounted(() => {
    pullSavedSolution();
});
</script>
  
<style lang="less" scoped>
.add_a_solution_form {
    font-size: 14px;
}

.submit-btn {
    float: right;
    margin: 10px 5px 0 0;
}

.save-btn {
    float: right;
    margin: 10px 40px 0 0;
}

.toolbar {
    cursor: pointer; //鼠标手型
}

.show-panel {
    padding: 5px;
    border: 1px solid lightgray;

    .panel-preview {
        padding: 1rem;
        margin: 0 0 0 0;
        width: auto;
        background-color: white;
    }
}

.editor-tab-content,
.preview-tab-content,
.header1-btn,
.header2-btn,
.header3-btn,
.text-btn,
.code-btn {
    border-bottom-color: transparent;
    border-bottom-style: solid;
    border-radius: 0;
    padding: .85714286em 1.14285714em 1.29999714em 1.14285714em;
    border-bottom-width: 2px;
    transition: color .1s ease;
    cursor: pointer; //鼠标手型
}

.header1-btn,
.header2-btn,
.header3-btn,
.code-btn,
.text-btn {
    font-size: 10px;
    padding: .78571429em 1.14285714em !important;
}

.active {
    background-color: transparent;
    box-shadow: none;
    border-color: #1B1C1D;
    font-weight: 700;
    color: rgba(0, 0, 0, .95);
}

.header1-btn:hover,
.header2-btn:hover,
.header3-btn:hover,
.text-btn:hover,
.code-btn:hover {
    cursor: pointer; //鼠标手型
    background: rgba(0, 0, 0, .05) !important;
    color: rgba(0, 0, 0, .95) !important;
}

.helpHeader {
    font-size: 1.228571rem;
    line-height: 1.2857em;
    font-weight: 700;
    border-top-left-radius: .28571429rem;
    border-top-right-radius: .28571429rem;
    display: block;
    font-family: Lato, 'Helvetica Neue', Arial, Helvetica, sans-serif;
    background: #FFF;
    box-shadow: none;
    color: rgba(0, 0, 0, .85);
}

.edit-enter-active,
.edit-leave-active {
    transition: opacity .5s linear;
}

.edit-enter,
.edit-leave-to {
    opacity: 0;
}

.preview-enter-active,
.preview-leave-active {
    transition: opacity .5s linear;
}

.preview-enter,
.preview-leave-to {
    opacity: 0;
}
</style>
  