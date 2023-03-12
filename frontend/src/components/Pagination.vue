<template>
    <div class="page">
        <el-pagination background :layout="layout" :page-size="pageSize" :current-page="current" :total="total" @current-change="onChange" @size-change="onPageSizeChange" />
    </div>
</template>

<script lang="ts" setup>
import { defineProps, defineEmits, toRefs } from "vue";

const props = defineProps({
  total: {
    type: Number,
    default: 50
  },
  pageSize: {
    required: false,
    type: Number,
  },
  current: {
    required: false,
    type: Number,
  },
  layout: {
    require: false,
    type: String,
    default: 'prev, pager, next',
  },
})

const { total, layout, pageSize, current } = toRefs(props)

const emit = defineEmits(['update:current', 'on-change', 'update:pageSize', 'on-page-size-change'])
const onChange = (page: number) => {
    if (page < 1) {
        page = 1;
    }
    
    emit('update:current', page);
    emit('on-change', page);
};
const onPageSizeChange = (pageSize: number) => {
    emit('update:pageSize', pageSize);
    emit('on-page-size-change', pageSize);
}
</script>

<style lang="less" scoped>
.page {
  margin: 20px;
  margin-right: 0px;
  float: right;
}
.el-pagination {
  padding-right: 0px !important;
}
:deep(.el-pagination__sizes) {
  margin: 0px !important;
}
:deep(.el-pagination .el-select .el-input) {
  margin-right: 0px !important;
}
</style>