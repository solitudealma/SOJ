<template>
    <div class="footer">
        <hr />
        <div class="copyright">
            <span>
                <i class="fa fa-copyright"></i>2021-2021 OnlineJudge版权所有 |
                <a href="https://beian.miit.gov.cn">粤ICP备20063723号</a>
            </span>
        </div>
        <div class="links">
            <a href="/faq">常见问题</a>
        </div>
    </div>
</template>

<script setup lang="ts"></script>

<style scoped lang="less">
.footer {
    height: 30px;
    display: block;
    font-size: 14px;
}

hr {
    background: none !important;
    height: 1px !important;
    border: 0 !important;
    border-top: 1px solid #acaaaa !important;
}

a {
    color: #337ab7;
}

.links {
    float: right;
    margin-top: -17px;
}
</style>
