<template>
  <el-card>
    <template #header>
      <div class="card-header">
        <span>{{ announcementTitle }}</span>
        <el-button v-if="listVisible" type="info" plain size="small" class="btn-refresh">刷新</el-button>
        <el-button v-else @click="goBack" type="info" plain size="small" class="btn-back">
          <el-icon size="20">
            <el-icon>
              <IEpBack />
            </el-icon>
          </el-icon>
        </el-button>
      </div>
    </template>
    <template v-if="listVisible">
      <ul class="announcements-container">
        <li v-for="item in announcementList" :key="item.title">
          <div class="flex-container">
            <div class="title">
              <a class="entry" @click="goAnnouncement(item)">{{ item.title }}</a>
            </div>
            <div class="date">{{ item.date }}</div>
            <div class="creator">{{ item.creator }}</div>
          </div>
        </li>
      </ul>
    </template>
    <template v-else>
      <div v-html="announcement.content" :key="announcement.title" class="content-container"></div>
    </template>
  </el-card>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
let listVisible = ref(true);

interface Announcement {
  title: string;
  content: string;
  link: string;
  date: string;
  creator: string;
}

let announcementList = ref<Array<Announcement>>([
  {
    title: "各位超级管理员请勿随便修改本站配置",
    date: "2021-8-28 19:54:38",
    link: "https://baidu.com",
    creator: "Root",
    content: "<p>asds</p>",
  },
  {
    title: "qwe",
    date: "2021-8-28 19:54:38",
    link: "https://baidu.com",
    creator: "Root",
    content: "<p>asds</p>",
  },
  {
    title: "qw",
    date: "2021-8-28 19:54:38",
    link: "https://baidu.com",
    creator: "Root",
    content: "<p>asds</p>",
  },
]);
let announcement = ref<Announcement>({
    title: "",
    content: "",
    date: "",
    link: "",
    creator: "",
  });
let announcementTitle = computed(() => {
  let title = "Announcement";
  if (listVisible.value) {
    return title;
  } else {
    return announcement.value!.title;
  }
});

const goBack = () => {
  listVisible.value = true;
  announcement.value = <Announcement>{
    title: "",
    content: "",
    date: "",
    creator: "",
  };
};

const goAnnouncement = (val: Announcement) => {
  announcement.value = val;
  listVisible.value = false;
};
</script>

<style scoped lang="less">
a {
    color: #2d8cf0;
    background: 0 0;
    text-decoration: none;
    outline: 0;
    cursor: pointer;
    transition: color .2s ease;
}

.card-header {
  font-size: 21px;
  font-weight: 500;
  line-height: 30px;
  padding: 5px 15px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.announcements-container {
  margin-top: -10px;
  margin-bottom: 10px;
  padding-left: 0 !important;

  li {
    padding-top: 15px;
    list-style: none;
    padding-bottom: 15px;
    margin-left: 20px;
    font-size: 16px;
    border-bottom: 1px solid hsla(0,0%,73%,.5);

    &:last-child {
      border-bottom: none;
    }

    .flex-container {
      display: flex;
      width: 100%;
      max-width: 100%;
      justify-content: space-around;
      align-items: flex-start;
      flex-flow: row nowrap;

      .title {
        -ms-flex: 1 1;
        flex: 1 1;
        text-align: left;
        padding-left: 10px;

        a.entry {
          color: #495060;
        }
        
        a.entry:hover {
          color: #2d8cf0;
          border-bottom: 1px solid #2d8cf0
        }
      }

      .date{
        -ms-flex: none;
        flex: none;
        width: 200px;
        text-align: center;
      }

      .creator {
        -ms-flex: none;
        flex: none;
        width: 200px;
        text-align: center;
      }
    }

    a {
      cursor: pointer;
    }
  }
}

.announcements-container li

.btn-refresh,
.btn-back {
  position: relative;
  top: 1px;
  float: right;
}
</style>
