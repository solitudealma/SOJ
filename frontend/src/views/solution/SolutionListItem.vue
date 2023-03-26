<template>
    <div class='solution-container'>
        <el-card>
            <template #header>
                <h3 style="margin-top: 10px;">
                    {{ solutionData.problemSource + ' ' + solutionData.problemId + '. ' + solutionData.title }}&nbsp;&nbsp;
                    <a class='label label-info' :href='solutionData.problemLink' style="font-size:14px;font-weight: normal;"
                        target='_blank'>原题链接
                    </a>
                    &nbsp;&nbsp;
                    <el-tag v-if="solutionData.problemDifficulty === 0" type='success' size='small'>简单
                    </el-tag>
                    <el-tag v-else-if="solutionData.problemDifficulty === 1" type='warning' size='small'>中等
                    </el-tag>
                    <el-tag v-else-if="solutionData.problemDifficulty === 2" type='danger' size='small'>困难
                    </el-tag>
                </h3>
                <div style="color: #999999; font-size:14px;">
                    作者：
                    <a :href="'/user/space/' + solutionData.authorId" target="_self">
                        <el-avatar class='header-img' :size='24' :src="solutionData.authorAvatar" alt='作者的头像'></el-avatar>
                        &nbsp;
                        <span style="font-size: 18px;">
                            {{ solutionData.authorName }}
                        </span>
                    </a>
                    , &nbsp;{{ time.formatDate(solutionData.updateTime) }} ,
                    &nbsp;阅读&nbsp;{{ solutionData.read }}
                </div>
            </template>
            <el-row>
                <el-col :span='2'>
                    <div class='like-icon'>
                        <i ref='notLike' :class='likeClass' @click='like' v-if='!isLike'
                            class='fa fa-thumbs-o-up fa-2x'></i>
                        <i ref='liked' :class='likeClass' @click='like' v-if='isLike' class='fa fa-thumbs-up fa-2x'></i>
                    </div>
                    <div class='votecnt'>0</div>
                    <div class='trample-icon'>
                        <i ref='notTrample' :class='trampleClass' @click='trample' v-if='!isTrample'
                            class='fa fa-thumbs-o-down fa-2x'></i>
                        <i ref='trampled' :class='trampleClass' @click='trample' v-if='isTrample'
                            class='fa fa-thumbs-down fa-2x'></i>
                    </div>
                    <div class='star-icon'>
                        <i ref='notStar' :class='starClass' @click='star' v-if='!isStar' class='fa fa-star-o fa-2x'></i>
                        <i ref='stared' :class='starClass' @click='star' v-if='isStar' class='fa fa-star fa-2x'></i>
                    </div>
                    <div class='favoritecnt'>0</div>
                </el-col>
                <el-col :span='22'>
                    <div class='markdown-body' ref='solutionContent'>
                    </div>
                </el-col>
            </el-row>
        </el-card>
        <div class='comment-card'>
            <el-card>
                <template #header>
                    <h3 style="font-weight: normal; margin: 10px 0 30px 0;">{{ total }} 评论</h3>
                    <div class='my-reply'>
                        <el-row>
                            <el-col :span='2'>
                                <el-avatar class='header-img' :size='50' :src='userInfo.avatar' @click="getUserSpace(userInfo.userId)">
                                </el-avatar>
                            </el-col>
                            <el-col :span='22'>
                                <div class='reply-area'>
                                    <textarea tabindex='0' id='replyInput' maxlength='10000' cols='40' rows='2'
                                        placeholder="在这里写评论...（支持MarkDown和Latex语法）" class='reply-input' ref='replyInput' />
                                </div>
                            </el-col>
                            <el-col :span='4' :push='17'>
                                <button class='send-reply-btn' @click='sendComment'>发表评论</button>
                            </el-col>
                        </el-row>
                    </div>
                </template>
                <div class="comment-list">
                    <div v-for='(comment, i) in comments' :key='comment.commentId' class='comment-item'>
                        <div class="comment-container">
                            <div class='user-avatar'>
                                <el-avatar class='header-img' :src='comment.fromAvatar' :alt="comment.fromName + '的头像'"
                                    :size='50' @click="getUserSpace(comment.fromUserId)">
                                </el-avatar>
                            </div>
                            <div class='content-warp'>
                                <div class='user-info'>
                                    <a :href="'/user/space/' + comment.fromUserId" class='user-name' target="_self">{{ comment.fromName }}</a>
                                </div>
                                <div class='comment-content' v-html="DOMPurify.sanitize(markdown(comment.content))"></div>
                                <div class='comment-info'>
                                    <span :title='time.formatDate(comment.createTime)' class='comment-time'>
                                        {{ dateStr(comment.createTime) }}
                                    </span>
                                    <a :id="'comment-btn-' + comment.commentId" class='reply-btn'
                                        :ref="getReplyBtnRef(comment.commentId)" style="cursor: pointer"
                                        @click="showReplyInput(i, -1, comment.commentId)">回复
                                    </a>
                                    <transition name='comment'>
                                        <div v-show='comment.inputShow' class='my-reply my-comment-reply'>
                                            <div class='reply-area'>
                                                <textarea tabindex='0' contenteditable='true' spellcheck='false'
                                                    maxlength='10000' cols='40' rows='2'
                                                    placeholder='在这里写评论...（支持MarkDown和Latex语法）' class='reply-input'
                                                    :ref="getReplyInputContentRef(comment.commentId)" />
                                            </div>
                                            <button class='send-reply-btn' @click='sendCommentReply(i, -1, comment)'>
                                                发表评论
                                            </button>
                                        </div>
                                    </transition>
                                </div>
                            </div>
                        </div>
                        <div class='reply-container'>
                            <div class="reply-list">
                                <div v-for='(reply, j) in comment.reply' :key='reply.commentId' class='reply-item'
                                    :ref="getReplyRef(reply.commentId)">
                                    <div class='user-avatar'>
                                        <el-avatar class='header-img' :size='40' :src='reply.fromAvatar'
                                            :alt="reply.fromName + '的头像'" @click="getUserSpace(reply.fromUserId)"></el-avatar>
                                    </div>
                                    <div class='content-warp'>
                                        <div class='user-info'>
                                            <a :href="'/user/space/' + reply.fromUserId" class='user-name'>{{ reply.fromName }}</a>
                                        </div>
                                        <div class='reply-content' v-html='DOMPurify.sanitize(markdown(reply.content))'>
                                        </div>
                                        <div class='reply-info'>
                                            <span :title='time.formatDate(reply.createTime)' class='reply-time'>
                                                {{ dateStr(reply.createTime) }}
                                            </span>
                                            <span v-if='reply.toCommentId !== reply.rootCommentId' class='reply-btn'>回复了
                                                <a :id="'reply-link-' + reply.commentId"
                                                    @click='showReplyByAnimation(reply)'
                                                    :ref="'reply-link-' + reply.commentId">
                                                    {{ reply.toName }} 的评论
                                                </a>
                                            </span>
                                            <a :id="'comment-reply-btn-' + reply.commentId" class='reply-btn'
                                                :ref="getReplyBtnRef(reply.commentId)"
                                                @click='showReplyInput(i, j, reply.commentId)' style="cursor: pointer">回复
                                            </a>
                                            <transition name='comment'>
                                                <div v-show='reply.inputShow' class='my-reply my-comment-reply'>
                                                    <div class='reply-area'>
                                                        <textarea tabindex='0' contenteditable='true' spellcheck='false'
                                                            maxlength='10000' cols='40' rows='2'
                                                            placeholder='在这里写评论...（支持MarkDown和Latex语法）' class='reply-input'
                                                            :ref="getReplyInputContentRef(reply.commentId)" />
                                                    </div>
                                                    <button class='send-reply-btn'
                                                        @click='sendCommentReply(i, j, reply)'>发表评论
                                                    </button>
                                                </div>
                                            </transition>
                                        </div>
                                    </div>
                                    <div class='reply-container'>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </el-card>
        </div>
    </div>
</template>
  
<script lang="ts" setup>
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { ElMessage } from "element-plus";
import DOMPurify from 'dompurify';
import { useUserStore } from "@/pinia/modules/user";
import { useGlobalStore } from "@/pinia/modules/global";
import { getSolutionInfo } from "@/api/solution";
import { GetSolutionInfoRequest, SolutionInfo } from "@/api/types/solution";
import { flatCommentListToTree } from '@/utils/comment';
import { getCommentListInfo, createComment } from "@/api/comment";
import { Comment, GetCommentListInfoRequest, CreateCommentReq } from "@/api/types/comment";
import time from '@/utils/time';
import markdown from "@/utils/markdown";

const router = useRouter();
const route = useRoute();
const userStore = useUserStore();
const globalStore = useGlobalStore();

let userInfo = computed(() => userStore.userInfo);
let solutionData = ref<SolutionInfo>({
    problemId: "",
    title: "",
    problemSource: "",
    content: "",
    problemLink: "",
    problemDifficulty: -1,
    updateTime: 0,
    authorId: 0,
    authorName: "",
    authorAvatar: "",
    read: 0,

});
let replyRefs = ref<{ [key: string]: Ref<HTMLDivElement | null> }>({});
let replyInput = ref<HTMLTextAreaElement | null>(null);
let replyBtnRefs = ref<{ [key: string]: Ref<HTMLAnchorElement | null> }>({});
let replyInputContentRefs = ref<{ [key: string]: Ref<HTMLTextAreaElement | null> }>({});
let isLike = ref(false);
let isTrample = ref(false);
let isStar = ref(false);
let likeClass = ref('');
let trampleClass = ref('');
let starClass = ref('');
let comments = ref<Array<Comment>>([]);
let total = ref(0);
let solutionContent = ref<HTMLDivElement | null>(null);

const getUserSpace = (userId: number) => {
    router.push({
        name: "user-space",
        params: { userId }
    })
}

const getReplyRef = (commentId: number) => {
    const refObj = ref(null);
    // 当 ref 创建时，将其存储到一个对象中，对象的属性名为元素的 id
    replyRefs.value["reply-" + commentId] = refObj
    return refObj
}

const getReplyBtnRef = (commentId: number) => {
    const refObj = ref(null);
    // 当 ref 创建时，将其存储到一个对象中，对象的属性名为元素的 id
    replyBtnRefs.value["reply-btn-" + commentId] = refObj
    return refObj
}

const getReplyInputContentRef = (commentId: number) => {
    const refObj = ref(null);
    // 当 ref 创建时，将其存储到一个对象中，对象的属性名为元素的 id
    replyInputContentRefs.value["reply-input-content-" + commentId] = refObj
    return refObj
}

const getSolutionmDetail = async () => {
    let solutionId = route.params.solutionId;
    let res = await getSolutionInfo({ solutionId } as GetSolutionInfoRequest);
    if (res.code === 200) {
        solutionData.value = res.data!.solutionInfo;
        solutionContent.value!.innerHTML = DOMPurify.sanitize(markdown(solutionData.value.content));
    }
};

const sleep = (time: number) => {
    return new Promise((resolve) => setTimeout(resolve, time));
};

const init = () => {
    getSolutionmDetail();
    getCommentList();
}

const getCommentList = () => {
    let solutionId = parseInt(route.params.solutionId as string, 10);
    getCommentListInfo({ solutionId } as GetCommentListInfoRequest).then(res => {
        comments.value = flatCommentListToTree(res.data!.comments);
        total.value = res.data!.total;
    });
}

const like = () => {
    if (!isLike.value) {
        isLike.value = !isLike.value;
        likeClass.value = 'animate__animated animate__swing';
    } else {
        isLike.value = !isLike.value;
        likeClass.value = 'animate__animated animate__flash';
    }
}

const trample = () => {
    if (!isTrample.value) {
        isTrample.value = !isTrample.value;
        trampleClass.value = 'animate__animated animate__swing';
    } else {
        isTrample.value = !isTrample.value;
        trampleClass.value = 'animate__animated animate__flash';
    }
}

const star = () => {
    if (!isStar.value) {
        isStar.value = !isStar.value;
        starClass.value = 'animate__animated animate__swing';
    } else {
        isStar.value = !isStar.value;
        starClass.value = 'animate__animated animate__flash';
    }
}

const hideReplyBtn = () => {
    //btnShow = false
    //replyInput.style.padding= "10px"
    //replyInput.style.border ="none"
}

const showReplyInput = (i: number, j: number, commentId: number) => {
    if (j === -1) {
        comments.value[i].inputShow = !comments.value[i].inputShow;
    } else {
        comments.value[i].reply[j].inputShow = !comments.value[i].reply[j].inputShow;
    }

    let text = replyBtnRefs.value["reply-btn-" + commentId][0].innerText;
    console.log(text);
    text = text === '回复' ? '收起' : '回复';
    replyBtnRefs.value['reply-btn-' + commentId][0].innerText = text;
}

const sendComment = async() => {
    if(!userStore.isAuthenticated) {
        ElMessage({
            showClose: true,
            type: 'warning',
            message: '请先登录'
        })
        globalStore.changeDialogStatus({
            mode: 'login',
            visible: true,
        })
        return;
    }

    let replyComment = DOMPurify.sanitize(markdown(replyInput.value!.value));
    if (!replyComment) {
        ElMessage({
            showClose: true,
            type: 'warning',
            message: '评论不能为空'
        });
    } else {
        let createTime = new Date().getTime();
        let comment = <Comment>{
            commentId: 0,
            layer: 1,
            createTime: createTime,
            fromName: userInfo.value.username,
            inputShow: false,
            content: replyComment,
            fromAvatar: userInfo.value.avatar,
            fromUserId: userInfo.value.userId,
            toName: '',
            toUserId: 0,
            toCommentId: 0,
            rootCommentId: 0,
            reply: [],
        }
        
        let solutionId = parseInt(route.params.solutionId as string);
        let createCommentReq = <CreateCommentReq>{
            solutionId: solutionId,
            layer: 1,
            rootCommentId: 0,
            fromUserId: userInfo.value.userId,
            fromAvatar: userInfo.value.avatar,
            fromName: userInfo.value.username,
            toCommentId: 0,
            toUserId: 0,
            toName: "",
            content: replyComment,
            createTime: createTime,
        }
        
        let res = await createComment({comment: createCommentReq});
        if (res.code === 200) {
            ElMessage({
                showClose: true,
                type: 'success',
                message: '评论成功'
            })
            comments.value.push(comment);
            replyInput.value!.value = '';
            return;
        } 

        ElMessage({
            showClose: true,
            type: 'error',
            message: '评论失败，请检查是否登录或是网络问题～'
        })
    }
}

const sendCommentReply = async(i: number, j: number, reply: Comment) => {
    if(!userStore.isAuthenticated) {
        ElMessage({
            showClose: true,
            type: 'warning',
            message: '请先登录'
        })
        globalStore.changeDialogStatus({
            mode: 'login',
            visible: true,
        })
        return;
    }

    let replyComment = DOMPurify.sanitize(markdown(replyInputContentRefs.value['reply-input-content-' + reply.commentId][0].value)) ;
    if (!replyComment) {
        ElMessage({
            showClose: true,
            type: 'warning',
            message: '回复不能为空'
        });
    } else {
        let createTime = new Date().getTime();
        let comment = <Comment>{
            commentId: 0,
            layer: 2,
            createTime: createTime,
            fromUserId: userInfo.value.userId,
            fromName: userInfo.value.username,
            inputShow: false,
            toName: reply.fromName,
            fromAvatar: userInfo.value.avatar,
            content: replyComment,
            toUserId: reply.fromUserId,
            toCommentId: reply.commentId,
            //如果是一级评论的话，rootCommentId就是CommentId
            rootCommentId: j === -1 ? reply.commentId : reply.rootCommentId,
            reply: [],
        }
        
        let solutionId = parseInt(route.params.solutionId as string);
        let createCommentReq = <CreateCommentReq>{
            solutionId: solutionId,
            layer: 2,
            rootCommentId: j === -1 ? reply.commentId : reply.rootCommentId,
            fromUserId: userInfo.value.userId,
            fromAvatar: userInfo.value.avatar,
            fromName: userInfo.value.username,
            toUserId: reply.fromUserId,
            toName: reply.fromName,
            toCommentId: reply.commentId,
            content: replyComment,
            createTime: createTime,
        }

        console.log(comments.value[i]);
        let res = await createComment({comment: createCommentReq});
        if (res.code === 200) {
            ElMessage({
                showClose: true,
                type: 'success',
                message: '评论成功'
            })
            comments.value[i].reply.push(comment);
            //隐藏当前的输入框
            reply.inputShow = false;
            replyBtnRefs.value['reply-btn-' + reply.commentId][0].innerText = '回复';
            replyInputContentRefs.value['reply-input-content-' + reply.commentId][0].value = '';
            return;
        } 

        ElMessage({
            showClose: true,
            type: 'error',
            message: '评论失败，请检查是否登录或是网络问题～'
        })
    }
}



const showReplyByAnimation = (reply: Comment) => {
    replyRefs.value['reply-' + reply.toCommentId][0].setAttribute('class', 'author-title animate__animated ' +
        'animate__flash animate__faster animate__infinite');
    sleep(1000).then(() => {
        replyRefs.value['reply-' + reply.toCommentId][0].setAttribute('class', 'author-title');
    });
}

const dateStr = (date: number) => {
    //获取js 时间戳
    let timestamp = new Date().getTime();
    //去掉 js 时间戳后三位，与php 时间戳保持一致
    timestamp = (timestamp - date) / 1000;
    //存储转换值
    let s;
    if (timestamp < 60 * 10) {
        //十分钟内
        return '刚刚';
    } else if (timestamp < 60 * 60 && timestamp >= 60 * 10) {
        //超过十分钟少于1小时
        s = Math.floor(timestamp / 60);
        return s + '分钟前';
    } else if (timestamp < 60 * 60 * 24 && timestamp >= 60 * 60) {
        //超过1小时少于24小时
        s = Math.floor(timestamp / 60 / 60);
        return s + '小时前';
    } else if (timestamp < 60 * 60 * 24 * 30 && timestamp >= 60 * 60 * 24) {
        //超过1天少于30天内
        s = Math.floor(timestamp / 60 / 60 / 24);
        return s + '天前';
    } else {
        //超过30天ddd
        let tmpDate = new Date(date).getTime();
        return time.formatDate(tmpDate);
    }
}

onMounted(() => {
    init();
});
</script>
  
<style lang='less' scoped>
:deep(p) {
    margin: 0 0 10px;
}

.my-reply {
    .reply-area {
        display: inline-block;
        width: 95%;

        @media screen and (max-width: 1200px) {
            width: 80%;
        }

        .reply-input {
            min-height: 20px;
            width: 95%;
            padding: 0.78571429em 1em;
            border: 1px solid rgba(34, 36, 38, .15);
            resize: vertical;
            color: rgba(0, 0, 0, .87);
            box-shadow: 0 0 0 0 transparent inset;
            font-size: 1em;
            line-height: 1.2857;
            border-radius: 4px;
            overflow: auto;
            overflow-wrap: break-word;
        }
    }

    .send-reply-btn {
        display: inline-block;
        padding: 6px 12px;
        margin-left: 83.33333333%;
        margin-bottom: 0;
        font-size: 14px;
        font-weight: 400;
        line-height: 1.42857143;
        text-align: center;
        white-space: nowrap;
        vertical-align: middle;
        -ms-touch-action: manipulation;
        touch-action: manipulation;
        cursor: pointer;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;
        background-image: none;
        border: 1px solid transparent;
        border-radius: 4px;
        background-color: transparent;
        color: #95a5a6;
    }
}

.my-comment-reply {
    overflow: hidden;
    height: 90px;
    width: 580px;

    .reply-input {
        overflow: hidden;
        height: 20px;
        width: 600px;
    }
}

.comment-item:not(:last-child) {
    border-bottom: 1px solid rgba(7, 111, 214, 0.3)
}

.comment-item {
    padding: 10px;

    .user-avatar {
        float: left;
        position: relative;
    }

    .content-warp {
        position: relative;
        margin-left: 70px;

        .user-info {
            display: block;
            word-wrap: break-word;
            line-height: 20px;

            .user-name {
                color: #337ab7;
                text-decoration: none;
                background-color: transparent;
                padding-bottom: 4px;
                font-size: 14px;
                font-weight: bold;
            }
        }

        .comment-content,
        .reply-content {
            line-height: 20px;
            padding: 2px 0;
            margin: 0;
            font-size: 14px;
            text-shadow: none;
            overflow: hidden;
        }

        .comment-info,
        .reply-info {

            .comment-time,
            .reply-time {
                margin-right: 20px;
                line-height: 20px;
                color: #999999;
                overflow: hidden;
                padding: 2px 0;
                font-size: 13px;
            }

            .reply-btn {
                a {
                    color: #337ab7;
                    text-decoration: none;
                    background-color: transparent;
                    cursor: pointer;
                }

                margin-right: 20px;
                color: #999999;
                font-size: 13px;
            }
        }
    }

    .reply-container {
        margin: 10px 0 0 50px;
    }
}

.solution-container {
    height: auto;
    font-size: 14px;
}

.like-icon,
.trample-icon,
.star-icon {
    margin: 20px 0 0 14px;
}

.votecnt,
.favoritecnt {
    text-align: center;
    font-size: 16px;
    color: #000000;
    margin: 20px 0 0 -14px;
}

.comment-card {
    margin: 24px 0 0 0;
}

h2 {
    font-size: 25px;
}

.label-info {
    background-color: #5bc0de;
}

.label {
    display: inline;
    padding: 0.2em 0.6em 0.3em;
    font-size: 75%;
    font-weight: 700;
    line-height: 1;
    color: #fff;
    text-align: center;
    white-space: nowrap;
    vertical-align: baseline;
    border-radius: 0.25em;
    text-decoration: none;
}

a {
    color: #337ab7;
    text-decoration: none;
    background-color: transparent;
}

.header-img {
    vertical-align: sub;
    transition: all 1s;
    cursor: pointer;
}

.header-img:hover {
    transform: rotate(360deg);
}

.comment-enter-active,
.comment-leave-active {
    transition: width ease-in-out .8s, height ease-in-out .8s;
}

.comment-enter,
.comment-leave-to {
    height: 0;
    width: 0;
}
</style>
  