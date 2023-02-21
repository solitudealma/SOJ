<template>
  <el-card>
    <template #header>
      <div class="clearfix">
        <span class="title">{{ announcementTitle }}</span>
        <el-button
          v-if="listVisible"
          type="info"
          plain
          size="small"
          class="btn-refresh"
          >刷新</el-button
        >
        <el-button
          v-else
          @click="goBack"
          type="info"
          plain
          size="small"
          class="btn-back"
        >
          <el-icon size="20">
            <el-icon>
              <IEpBack />
            </el-icon>
          </el-icon>
        </el-button>
      </div>
    </template>
    <template v-if="listVisible">
      <ul class="announcements-container" hoverable>
        <li v-for="item in announcementList" :key="item.title">
          <div class="flex-container" style="text-align: left">
            <el-row>
              <el-col :span="14">
                <a @click="goAnnouncement(item)">{{ item.title }}</a>
              </el-col>
              <el-col :span="8">
                <div>{{ item.date }}</div>
              </el-col>
              <el-col :span="2">
                <div>{{ item.createBy }}</div>
              </el-col>
            </el-row>
          </div>
        </li>
      </ul>
    </template>
    <template v-else>
      <div
        v-html="announcement.content"
        :key="announcement.title"
        class="content-container"
      ></div>
    </template>
  </el-card>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
let listVisible = ref(true);
let announcementList = ref([
  {
    title: "各位超级管理员请勿随便修改本站配置",
    date: "2021-8-28 19:54:38",
    link: "https://baidu.com",
    createBy: "Root",
    content: "<p>asds</p>",
  },
  {
    title: "qwe",
    date: "2021-8-28 19:54:38",
    link: "https://baidu.com",
    createBy: "Root",
    content: "<p>asds</p>",
  },
  {
    title: "qw",
    date: "2021-8-28 19:54:38",
    link: "https://baidu.com",
    createBy: "Root",
    content: "<p>asds</p>",
  },
]);
let announcement = ref({
  title: "",
});
let announcementTitle = computed(() => {
  let title = "Announcement";
  if (listVisible.value) {
    return title;
  } else {
    return announcement.value.title;
  }
});

const goBack = () => {
  listVisible.value = true;
  announcement.value = {
    title: "",
  };
};
const goAnnouncement = (val) => {
  announcement.value = val;
  listVisible.value = false;
};
</script>

<style scoped lang="less">
a {
  color: #495060;

  &:hover {
    color: #2d8cf0;
    border-bottom: 1px solid #2d8cf0;
  }
}

.title {
  margin-left: 10px;
  font-size: 17px;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both;
}

.announcements-container {
  margin-bottom: 10px;
  padding-left: 20px !important;

  li {
    list-style: none;
    padding-bottom: 15px;
    font-size: 16px;
    border-bottom: 1px solid rgba(187, 187, 187, 0.5);

    a {
      cursor: pointer;
    }
  }
}

.btn-refresh,
.btn-back {
  position: relative;
  top: 1px;
  float: right;
}
</style>
