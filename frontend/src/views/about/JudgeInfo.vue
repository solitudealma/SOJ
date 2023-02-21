<template>
  <el-row :gutter="20">
    <el-col :span="12">
      <el-card class="container">
        <template #header>
          <span class="panel-title home-title">编译器 & 例题</span>
        </template>
        <div class="content">
          <ul>
            <li v-for="lang in languages" :key="lang.name">
              {{ lang.name }} ( {{ lang.description }} )
              <p style="color: #409eff; font-size: 16px">编译器</p>
              <pre>{{ lang.compileCommand }}</pre>
              <p style="color: #409eff; font-size: 16px">A+B 题目</p>
              <Highlight
                :code="lang.template"
                :language="lang.name"
              ></Highlight>
            </li>
          </ul>
        </div>
      </el-card>
    </el-col>
    <el-col :span="12">
      <el-card class="container">
        <template #header>
          <span class="panel-title home-title">结果说明</span>
        </template>
        <ul class="result">
          <li>
            <el-tag :color="getStatusColor(5)">Pending</el-tag>
            ：您的解答正在排队等待评测中，请等待结果...
          </li>
          <li>
            <el-tag :color="getStatusColor(10)">Submitted Failed</el-tag>
            ：您的此次提交失败，请点击按钮重新提交...
          </li>
          <li>
            <el-tag :color="getStatusColor(6)">Compiling</el-tag>
            ：正在对您的源代码进行编译中，请等待结果...
          </li>
          <li>
            <el-tag :color="getStatusColor(7)">Judging</el-tag>
            ：正在使用测试数据运行您的程序中，请等待结果...
          </li>
          <li>
            <el-tag :color="getStatusColor(-2)">Compile Error</el-tag>
            ：无法编译您的源代码，点击链接查看编译器的输出。
          </li>
          <li>
            <el-tag :color="getStatusColor(-3)">Presentation Error</el-tag>
            ：您提交的代码已经很接近正确答案，请检查代码格式输出是否有多余空格，换行等空白符。
          </li>
          <li>
            <el-tag :color="getStatusColor(8)">Partial Accepted</el-tag>
            ：加油！您提交的代码通过了部分测试点，请考虑其他可能性。
          </li>
          <li>
            <el-tag :color="getStatusColor(0)">Accepted</el-tag>
            ：恭喜! 您的解题方法是正确的。
          </li>
          <li>
            <el-tag :color="getStatusColor(-1)">Wrong Answer</el-tag>
            ：您的程序输出结果与判题程序的答案不符。
          </li>
          <li>
            <el-tag :color="getStatusColor(3)">Runtime Error</el-tag>
            ：您的程序异常终止，可能的原因是：段错误，被零除或用非0的代码退出程序。
          </li>
          <li>
            <el-tag :color="getStatusColor(1)">Time Limit Exceeded</el-tag>
            ：您的程序运行时间已超出题目限制。
          </li>
          <li>
            <el-tag :color="getStatusColor(2)">Memory Limit Exceeded</el-tag>
            ：您的程序实际使用的内存已超出题目限制。
          </li>
          <li>
            <el-tag :color="getStatusColor(4)">System Error</el-tag>
            ：糟糕，判题机系统出了问题。请报告给管理员。
          </li>
          <li>
            <el-tag :color="getStatusColor(-4)">Cancelled</el-tag>
            ：您的此次提交被取消！
          </li>
        </ul>
      </el-card>
      <el-card class="container">
        <template #header>
          <span class="panel-title home-title"> 编译说明 </span>
        </template>
        <ul class="result">
          <li>
            1. __int64不是ANSI标准定义，只能在VC使用，在 GNU C++ 中应写成 long
            long 类型， scanf和printf 请使用%lld作为格式
          </li>
          <li>2. main() 返回值必须定义为 int ，而不是 void</li>
          <li>3. i 在循环外失去定义 "for(int i=0...){...}"</li>
          <li>4. itoa 不是ansi标准函数（标准 C/C++ 中无此函数）</li>
        </ul>
      </el-card>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
let JUDGE_STATUS = {
  "-10": {
    name: "Not Submitted",
    short: "NS",
    color: "gray",
    type: "info",
    rgb: "#909399",
  },
  "-5": {
    name: "Submitted Unknown Result",
    short: "SNR",
    color: "gray",
    type: "info",
    rgb: "#909399",
  },
  "-4": {
    name: "Cancelled",
    short: "CA",
    color: "purple",
    type: "info",
    rgb: "#676fc1",
  },
  "-3": {
    name: "Presentation Error",
    short: "PE",
    color: "orange",
    type: "warning",
    rgb: "#f90",
  },
  "-2": {
    name: "Compile Error",
    short: "CE",
    color: "orange",
    type: "warning",
    rgb: "#f90",
  },
  "-1": {
    name: "Wrong Answer",
    short: "WA",
    color: "red",
    type: "error",
    rgb: "#ed3f14",
  },
  0: {
    name: "Accepted",
    short: "AC",
    color: "green",
    type: "success",
    rgb: "#19be6b",
  },
  1: {
    name: "Time Limit Exceeded",
    short: "TLE",
    color: "red",
    type: "error",
    rgb: "#ed3f14",
  },
  2: {
    name: "Memory Limit Exceeded",
    short: "MLE",
    color: "red",
    type: "error",
    rgb: "#ed3f14",
  },
  3: {
    name: "Runtime Error",
    short: "RE",
    color: "red",
    type: "error",
    rgb: "#ed3f14",
  },
  4: {
    name: "System Error",
    short: "SE",
    color: "gray",
    type: "info",
    rgb: "#909399",
  },
  5: {
    name: "Pending",
    color: "orange",
    type: "warning",
    rgb: "#f90",
  },
  6: {
    name: "Compiling",
    short: "CP",
    color: "green",
    type: "info",
    rgb: "#25bb9b",
  },
  7: {
    name: "Judging",
    color: "blue",
    type: "",
    rgb: "#2d8cf0",
  },
  8: {
    name: "Partial Accepted",
    short: "PAC",
    color: "blue",
    type: "",
    rgb: "#2d8cf0",
  },
  9: {
    name: "Submitting",
    color: "orange",
    type: "warning",
    rgb: "#f90",
  },
  10: {
    name: "Submitted Failed",
    color: "gray",
    short: "SF",
    type: "info",
    rgb: "#909399",
  },
};
let languages = [
  {
    id: 51,
    contentType: "text/x-c++src",
    description: "G++ 7.5.0",
    name: "C++",
    compileCommand:
      "/usr/bin/g++ -DONLINE_JUDGE -w -fmax-errors=3 -std=c++14 {src_path} -lm -o {exe_path}",
    template:
      "#include<iostream>\r\nusing namespace std;\r\nint main()\r\n{\r\n    int a,b;\r\n    cin >> a >> b;\r\n    cout << a + b;\r\n    return 0;\r\n}",
    codeTemplate:
      "//PREPEND BEGIN\r\n#include <iostream>\r\nusing namespace std;\r\n//PREPEND END\r\n\r\n//TEMPLATE BEGIN\r\nint add(int a, int b) {\r\n  // Please fill this blank\r\n  return ___________;\r\n}\r\n//TEMPLATE END\r\n\r\n//APPEND BEGIN\r\nint main() {\r\n  cout << add(1, 2);\r\n  return 0;\r\n}\r\n//APPEND END",
    isSpj: true,
    oj: "ME",
    gmtCreate: "2020-12-12T15:12:44.000+0000",
    gmtModified: "2022-03-10T08:33:19.000+0000",
  },
  {
    id: 52,
    contentType: "text/x-c++src",
    description: "G++ 7.5.0",
    name: "C++ With O2",
    compileCommand:
      "/usr/bin/g++ -DONLINE_JUDGE -O2 -w -fmax-errors=3 -std=c++14 {src_path} -lm -o {exe_path}",
    template:
      "#include<iostream>\r\nusing namespace std;\r\nint main()\r\n{\r\n    int a,b;\r\n    cin >> a >> b;\r\n    cout << a + b;\r\n    return 0;\r\n}",
    codeTemplate:
      "//PREPEND BEGIN\r\n#include <iostream>\r\nusing namespace std;\r\n//PREPEND END\r\n\r\n//TEMPLATE BEGIN\r\nint add(int a, int b) {\r\n  // Please fill this blank\r\n  return ___________;\r\n}\r\n//TEMPLATE END\r\n\r\n//APPEND BEGIN\r\nint main() {\r\n  cout << add(1, 2);\r\n  return 0;\r\n}\r\n//APPEND END",
    isSpj: false,
    oj: "ME",
    gmtCreate: "2020-12-12T15:12:44.000+0000",
    gmtModified: "2022-03-10T08:30:22.000+0000",
  },
  {
    id: 53,
    contentType: "text/x-csrc",
    description: "GCC 7.5.0",
    name: "C",
    compileCommand:
      "/usr/bin/gcc -DONLINE_JUDGE -w -fmax-errors=3 -std=c11 {src_path} -lm -o {exe_path}",
    template:
      '#include <stdio.h>\r\nint main() {\r\n    int a,b;\r\n    scanf("%d %d",&a,&b);\r\n    printf("%d",a+b);\r\n    return 0;\r\n}',
    codeTemplate:
      '//PREPEND BEGIN\r\n#include <stdio.h>\r\n//PREPEND END\r\n\r\n//TEMPLATE BEGIN\r\nint add(int a, int b) {\r\n  // Please fill this blank\r\n  return ___________;\r\n}\r\n//TEMPLATE END\r\n\r\n//APPEND BEGIN\r\nint main() {\r\n  printf("%d", add(1, 2));\r\n  return 0;\r\n}\r\n//APPEND END',
    isSpj: true,
    oj: "ME",
    gmtCreate: "2020-12-12T15:11:44.000+0000",
    gmtModified: "2022-03-10T08:31:36.000+0000",
  },
  {
    id: 54,
    contentType: "text/x-csrc",
    description: "GCC 7.5.0",
    name: "C With O2",
    compileCommand:
      "/usr/bin/gcc -DONLINE_JUDGE -O2 -w -fmax-errors=3 -std=c11 {src_path} -lm -o {exe_path}",
    template:
      '#include <stdio.h>\r\nint main() {\r\n    int a,b;\r\n    scanf("%d %d",&a,&b);\r\n    printf("%d",a+b);\r\n    return 0;\r\n}',
    codeTemplate:
      '//PREPEND BEGIN\r\n#include <stdio.h>\r\n//PREPEND END\r\n\r\n//TEMPLATE BEGIN\r\nint add(int a, int b) {\r\n  // Please fill this blank\r\n  return ___________;\r\n}\r\n//TEMPLATE END\r\n\r\n//APPEND BEGIN\r\nint main() {\r\n  printf("%d", add(1, 2));\r\n  return 0;\r\n}\r\n//APPEND END',
    isSpj: false,
    oj: "ME",
    gmtCreate: "2021-06-14T13:05:57.000+0000",
    gmtModified: "2022-03-10T08:31:43.000+0000",
  },
  {
    id: 55,
    contentType: "text/x-python",
    description: "Python 3.7.5",
    name: "Python3",
    compileCommand: "/usr/bin/python3 -m py_compile {src_path}",
    template: "a, b = map(int, input().split())\r\nprint(a + b)",
    codeTemplate:
      "//PREPEND BEGIN\r\n//PREPEND END\r\n\r\n//TEMPLATE BEGIN\r\ndef add(a, b):\r\n    return a + b\r\n//TEMPLATE END\r\n\r\n\r\nif __name__ == '__main__':  \r\n    //APPEND BEGIN\r\n    a, b = 1, 1\r\n    print(add(a, b))\r\n    //APPEND END",
    isSpj: false,
    oj: "ME",
    gmtCreate: "2020-12-12T15:14:23.000+0000",
    gmtModified: "2022-02-25T12:55:29.000+0000",
  },
  {
    id: 56,
    contentType: "text/x-python",
    description: "Python 2.7.17",
    name: "Python2",
    compileCommand: "/usr/bin/python -m py_compile {src_path}",
    template: "a, b = map(int, raw_input().split())\r\nprint a+b",
    codeTemplate:
      "//PREPEND BEGIN\r\n//PREPEND END\r\n\r\n//TEMPLATE BEGIN\r\ndef add(a, b):\r\n    return a + b\r\n//TEMPLATE END\r\n\r\n\r\nif __name__ == '__main__':  \r\n    //APPEND BEGIN\r\n    a, b = 1, 1\r\n    print add(a, b)\r\n    //APPEND END",
    isSpj: false,
    oj: "ME",
    gmtCreate: "2021-01-26T03:11:44.000+0000",
    gmtModified: "2021-06-14T14:08:05.000+0000",
  },
  {
    id: 57,
    contentType: "text/x-java",
    description: "OpenJDK 1.8",
    name: "Java",
    compileCommand: "/usr/bin/javac {src_path} -d {exe_dir} -encoding UTF8",
    template:
      "import java.util.Scanner;\r\npublic class Main{\r\n    public static void main(String[] args){\r\n        Scanner in=new Scanner(System.in);\r\n        int a=in.nextInt();\r\n        int b=in.nextInt();\r\n        System.out.println((a+b));\r\n    }\r\n}",
    codeTemplate:
      "//PREPEND BEGIN\r\nimport java.util.Scanner;\r\n//PREPEND END\r\n\r\npublic class Main{\r\n    //TEMPLATE BEGIN\r\n    public static Integer add(int a,int b){\r\n        return _______;\r\n    }\r\n    //TEMPLATE END\r\n\r\n    //APPEND BEGIN\r\n    public static void main(String[] args){\r\n        System.out.println(add(a,b));\r\n    }\r\n    //APPEND END\r\n}\r\n",
    isSpj: false,
    oj: "ME",
    gmtCreate: "2020-12-12T15:12:51.000+0000",
    gmtModified: "2021-06-14T14:08:10.000+0000",
  },
  {
    id: 58,
    contentType: "text/x-go",
    description: "Golang 1.16",
    name: "Golang",
    compileCommand: "/usr/bin/go build -o {exe_path} {src_path}",
    template:
      'package main\r\nimport "fmt"\r\n\r\nfunc main(){\r\n    var x int\r\n    var y int\r\n    fmt.Scanln(&x,&y)\r\n    fmt.Printf("%d",x+y)  \r\n}\r\n',
    codeTemplate:
      '\r\npackage main\r\n\r\n//PREPEND BEGIN\r\nimport "fmt"\r\n//PREPEND END\r\n\r\n\r\n//TEMPLATE BEGIN\r\nfunc add(a,b int)int{\r\n    return ______\r\n}\r\n//TEMPLATE END\r\n\r\n//APPEND BEGIN\r\nfunc main(){\r\n    var x int\r\n    var y int\r\n    fmt.Printf("%d",add(x,y))  \r\n}\r\n//APPEND END\r\n',
    isSpj: false,
    oj: "ME",
    gmtCreate: "2021-03-28T15:15:54.000+0000",
    gmtModified: "2021-06-14T14:09:51.000+0000",
  },
  {
    id: 59,
    contentType: "text/x-csharp",
    description: "C# Mono 4.6.2",
    name: "C#",
    compileCommand: "/usr/bin/mcs -optimize+ -out:{exe_path} {src_path}",
    template:
      "using System;\r\nusing System.Linq;\r\n\r\nclass Program {\r\n    public static void Main(string[] args) {\r\n        Console.WriteLine(Console.ReadLine().Split().Select(int.Parse).Sum());\r\n    }\r\n}",
    codeTemplate:
      "//PREPEND BEGIN\r\nusing System;\r\nusing System.Collections.Generic;\r\nusing System.Text;\r\n//PREPEND END\r\n\r\nclass Solution\r\n{\r\n    //TEMPLATE BEGIN\r\n    static int add(int a,int b){\r\n        return _______;\r\n    }\r\n    //TEMPLATE END\r\n\r\n    //APPEND BEGIN\r\n    static void Main(string[] args)\r\n    {\r\n        int a ;\r\n        int b ;\r\n        Console.WriteLine(add(a,b));\r\n    }\r\n    //APPEND END\r\n}",
    isSpj: false,
    oj: "ME",
    gmtCreate: "2021-04-13T08:10:03.000+0000",
    gmtModified: "2022-03-21T11:12:26.000+0000",
  },
  {
    id: 1062,
    contentType: "text/x-php",
    description: "PHP 7.3.33",
    name: "PHP",
    compileCommand: "/usr/bin/php {src_path}",
    template: '<?=array_sum(fscanf(STDIN, "%d %d"));',
    codeTemplate: null,
    isSpj: false,
    oj: "ME",
    gmtCreate: "2022-02-25T12:55:30.000+0000",
    gmtModified: "2022-02-25T12:55:30.000+0000",
  },
  {
    id: 1063,
    contentType: "text/x-python",
    description: "PyPy 2.7.18 (7.3.8)",
    name: "PyPy2",
    compileCommand: "/usr/bin/pypy -m py_compile {src_path}",
    template: "print sum(int(x) for x in raw_input().split(' '))",
    codeTemplate:
      "//PREPEND BEGIN\n//PREPEND END\n\n//TEMPLATE BEGIN\ndef add(a, b):\n    return a + b\n//TEMPLATE END\n\n\nif __name__ == '__main__':  \n    //APPEND BEGIN\n    a, b = 1, 1\n    print add(a, b)\n    //APPEND END",
    isSpj: false,
    oj: "ME",
    gmtCreate: "2022-02-25T12:55:30.000+0000",
    gmtModified: "2022-02-25T12:55:30.000+0000",
  },
  {
    id: 1064,
    contentType: "text/x-python",
    description: "PyPy 3.8.12 (7.3.8)",
    name: "PyPy3",
    compileCommand: "/usr/bin/pypy3 -m py_compile {src_path}",
    template: "print(sum(int(x) for x in input().split(' ')))",
    codeTemplate:
      "//PREPEND BEGIN\n//PREPEND END\n\n//TEMPLATE BEGIN\ndef add(a, b):\n    return a + b\n//TEMPLATE END\n\n\nif __name__ == '__main__':  \n    //APPEND BEGIN\n    a, b = 1, 1\n    print(add(a, b))\n    //APPEND END",
    isSpj: false,
    oj: "ME",
    gmtCreate: "2022-02-25T12:55:30.000+0000",
    gmtModified: "2022-02-25T12:55:30.000+0000",
  },
  {
    id: 1065,
    contentType: "text/javascript",
    description: "Node.js 14.19.0",
    name: "JavaScript Node",
    compileCommand: "/usr/bin/node {src_path}",
    template:
      "var readline = require('readline');\nconst rl = readline.createInterface({\n        input: process.stdin,\n        output: process.stdout\n});\nrl.on('line', function(line){\n   var tokens = line.split(' ');\n    console.log(parseInt(tokens[0]) + parseInt(tokens[1]));\n});",
    codeTemplate: null,
    isSpj: false,
    oj: "ME",
    gmtCreate: "2022-02-25T12:55:30.000+0000",
    gmtModified: "2022-02-25T12:55:30.000+0000",
  },
  {
    id: 1066,
    contentType: "text/javascript",
    description: "JavaScript V8 8.4.109",
    name: "JavaScript V8",
    compileCommand: "/usr/bin/jsv8/d8 {src_path}",
    template:
      "const [a, b] = readline().split(' ').map(n => parseInt(n, 10));\nprint((a + b).toString());",
    codeTemplate: null,
    isSpj: false,
    oj: "ME",
    gmtCreate: "2022-02-25T12:55:30.000+0000",
    gmtModified: "2022-02-25T12:55:30.000+0000",
  },
];

const init = () => {
  languages = languages.filter(function (element) {
    return element.oj == "ME";
  });
};

init();

const getStatusColor = (status) => {
  return JUDGE_STATUS[status].color;
};
</script>

<style lang="less" scoped>
.container {
  margin-bottom: 20px;
}

.container .content {
  font-size: 16px;
  margin: 0 50px 20px 50px;
}

.container .content pre {
  padding: 5px 10px;
  white-space: pre-wrap;
  margin-top: 15px;
  margin-bottom: 15px;
  background: #f8f8f9;
  border: 1px dashed #e9eaec;
}

@media screen and (max-width: 768px) {
  .container .content {
    font-size: 1rem;
    margin: 0 5px;
  }
}

ul {
  list-style: disc;
  padding-inline-start: 0px;
}

li {
  line-height: 2;
}

li .title {
  font-weight: 600;
  font-size: 1rem;
}

.result li {
  list-style-type: none;
  margin-top: 8px;
}
</style>
