<template>
    <div class="footer">
        <div class="copyright">
            <hr />
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

<style scoped>
.footer {
    margin: 20px 78px 0 78px;
    font-size: small;
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
