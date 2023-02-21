<template>
  <el-card>
    <template #header>
      <div>Frequently Asked Questions</div>
    </template>
    <div class="content">
      <ul class="announcements-container">
        <li>
          Where is the input and the output?
          <p>
            Your program shall read input from <code>stdin</code> ('Standard
            Input') and write output to <code>stdout</code> ('Standard Output')
            For example,you can use <code>scanf</code> in C or
            <code>cin</code>
            in C++ to read from <code>stdin</code>,and use
            <code>printf</code> in C or <code>cout</code> in C++ to write to
            <code>stdout</code>. User programs are not allowed to read or write
            files, or you will get a <code>Runtime Error</code>.
          </p>
        </li>
        <li>
          What's the meaning of the submission execution time?
          <p>
            The OnlineJudge might test your code multiple times with different
            input files. If your code gives the correct answer within the time
            limit for each input file, the execution time displayed is the max
            of the time spent for each test case. Otherwise, the execution time
            will have no sense.
          </p>
        </li>
        <li>
          How can I use C++ Int64?
          <p>
            You should declare as <code>long long</code> and use with
            <code>cin/cout</code> or <code>%lld</code>, using
            <code>__int64</code> will result in <code>Compile Error</code>.
          </p>
        </li>
      </ul>
    </div>
  </el-card>
</template>

<script lang="ts" setup></script>

<style scoped lang="less">
.content {
  font-size: 16px;

  > ul {
    list-style: disc;

    li {
      font-size: 16px;
      margin-top: 20px;

      &:first-child {
        margin-top: 0;
      }

      p {
        font-size: 14px;
        margin-top: 5px;
      }
    }
  }
}
</style>
