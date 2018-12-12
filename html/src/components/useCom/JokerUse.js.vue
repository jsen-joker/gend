<template>
  <div id="quick-start">
    <div class="t-title">Helloworld by JS</div>
    <div class="t-tips">1、joker准备</div>
    <div class="t-tips t-smaller">同上</div>
    <div class="t-tips">2、写一个verticle</div>
    <div class="t-tips t-smaller">目前js不支持let等es6语法，你可以用babel等将你的npm项目打包时转换为es5，js的微服务很简单</div>
    <div class="t-tips t-small">如下是代码</div>
    <div class="t-tips t-smaller">    
      <codemirror
        ref="myCm"
        :value="code"
        :options="cmOptions"
        class="code"
        ></codemirror>
    </div>
    <div class="t-tips t-smaller">拖入joker的entry目录，访问localhost:9998/就可以看到输出的Hello World!了。</div>
  </div>
</template>

<script>

import 'codemirror/mode/javascript/javascript.js'
import 'codemirror/theme/base16-light.css'
import { codemirror } from 'vue-codemirror'

import { helloworld1 } from './javas/helloworld'

export default {
  name:'quick-start',
  data: () => {
    return {
      code: helloworld1,
      cmOptions: {
        // codemirror options
        matchBrackets: true,
        tabSize: 4,
        mode: 'text/javascript',
        theme: 'base16-light',
        lineNumbers: false,
        line: true,
        // more codemirror options, 更多 codemirror 的高级配置...
      }
    }
  },
  components: {
    codemirror,
  },
}
</script>

<style lang="scss" scoped>
#quick-start {
  .t-title {
    color: #222;
    font-size: 1.8rem;
    font-weight: bold;
  }
  .t-tips {
    margin-left: 20px;
    font-size: 1.2rem;
    margin-top: 30px;
    font-weight: bold;
    a {
      color: #147879;
    }
  }
  .t-small {
    margin-left: 40px;
    font-weight: normal;
    font-size: 1rem;
    color: #46535e;
    margin-top: 16px;
    padding-bottom: 2px;
    border-bottom: 2px dashed #147879;
  }
  .hey {
    font-size: 1.4rem;
  }
  .t-smaller {
    margin-left: 40px;
    font-weight: normal;
    font-size: 1rem;
    color: #46535e;
    margin-top: 16px;
    padding-bottom: 2px;
  }
  .t-part {
    margin-left: 40px;
    color: #545454;
    background-color: #f4f4f4;
    margin-top: 16px;
    padding: 10px 20px;
    .t-commen {
      line-height: 1.4rem;
      font-size: 0.8em;
      color: #aaa;
    }
    .t-code {
      line-height: 2rem;
    }
  }
}
</style>
