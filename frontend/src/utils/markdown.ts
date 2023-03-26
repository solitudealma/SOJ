import * as marked from 'marked';
import katex from 'katex';
import highlight from 'highlight.js';
import 'highlight.js/styles/foundation.css'

const renderer = new marked.Renderer();

function toHtml(text: string) {
  let temp = <HTMLDivElement| null>document.createElement('div');
  temp!.innerHTML = text;
  let output = temp!.innerText || temp!.textContent;
  temp = null;
  return output;
}

function mathsExpression(expr: string| null) {
  if (expr!.match(/^\$\$[\s\S]*\$\$$/)) {
    expr = expr!.substring(2, expr!.length - 2);
    return katex.renderToString(expr, { displayMode: true });
  } else if (expr!.match(/^\$[\s\S]*\$$/)) {
    expr = toHtml(expr!); // temp solution
    expr = expr!.substring(1, expr!.length - 1);
    //Does that mean your text is getting dynamically added to the page? If so, someone must be calling KaTeX to render
    // it, and that call needs to have the strict flag set to false as well. 即控制台警告，比如%为转义或者中文
    // link: https://katex.org/docs/options.html
    return katex.renderToString(expr, {
      displayMode: false,
      strict: false,
      throwOnError: false,
      errorColor: '#cc0000'
    });
  }
}

const unchanged = new marked.Renderer();
renderer.code = function (code, language, escaped) {
  language = language === 'c#' ? 'csharp' : language;
  const isMarkup = ['c++', 'cpp', 'golang', 'java', 'js', 'javascript', 'sql', 'python', 'xml', 'properties', 'csharp',
    'django', 'json', 'php', 'ruby', 'rust', 'yaml'].includes(language!);
  let hled = '';
  if (isMarkup) {
    const math = mathsExpression(code);
    if (math) {
      return math;
    } else {
      hled = highlight.highlight(language!, code).value;
    }
  } else {
    hled = highlight.highlightAuto(code).value;
  }
  return `<pre class='hljs ${language}'><code class='${language}'>${hled}</code></pre>`;
  // return unchanged.code(code, language, escaped);
};
renderer.codespan = function (text) {
  const math = mathsExpression(text);
  if (math) {
    return math;
  }
  return unchanged.codespan(text);
};

const markdown = marked.marked.setOptions({
  renderer: renderer,
  gfm: true,//默认为true。 允许 Git Hub标准的markdown.
  breaks: false,//默认为false。 允许回车换行。该选项要求 gfm 为true。
  pedantic: false,//默认为false。 尽可能地兼容 markdown.pl的晦涩部分。不纠正原始模型任何的不良行为和错误。
  // sanitize: false,//对输出进行过滤（清理） 不支持了，用sanitizer 或者直接渲染的时候过滤
  xhtml: true, // If true, emit self-closing HTML tags for void elements (<br/>, <img/>, etc.) with a "/" as required by XHTML.
  silent: true, //If true, the parser does not throw any exception.
  smartLists: true,
  smartypants: false//使用更为时髦的标点，比如在引用语法中加入破折号。
});

export default markdown;

