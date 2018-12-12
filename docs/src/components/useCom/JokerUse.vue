<template>
  <div id="quick-start">
    <div class="t-title">Helloworld</div>
    <div class="t-tips">1、joker准备</div>
    <div class="t-tips t-smaller">请确保joker可以正常运行，除了三个joker的核心jar包外，还要在entry下有manager和manager-master两个服务模块，启动后访问9090端口，如果可以正常访问，表示joker准备已经完成。</div>
    <div class="t-tips">2、写一个verticle</div>
    <div class="t-tips t-smaller">首先创建一个maven项目，在开始编程前，需要引入joker项目更目录下的pom.xml中的支持库（大部分是vertx编程所需要的，这些库不会被打包，理论上这些支持库将最终放入lib目录下），然后将program-api安装到本地mvn仓库，并在项目中引入，到此所有准备工作已经完成，接下来就可以开始编写helloworld了。</div>
    <div class="t-tips t-smaller">创建一个helloworld的java类，继承RestVerticle（program-api中的包，简单封装了AbstractVerticle，提供了一些方便rest编程的方法），然后重写start方法，一定要调用super.start，因为在父类中有很多初始化工作。</div>
    <div class="t-tips t-smaller">然后编写router，这是vertx种最简单的router写法，router对象在父类中创建，并配置了跨域、默认异常处理等。config()写入服务的基本信息，一般来说，服务的配置信息可以有一个配置信息服务器提供，joker提供了一个默认的配置服务器，可以方便的动态修改配置信息。关于如何使用配置信息服务器，可以参考joker各个插件，许多插件都是用配置信息服务器来获取其配置。</div>
    <div class="t-tips t-smaller">最后启动http服务器，不要忘记complete future，因为如果不complete，则表示该verticle启动失败，这回导致joker不正确的判断，这点非常重要，这里complete由startServer调用。</div>
    <div class="t-tips t-small">如下是代码</div>
    <div class="t-tips t-smaller">    
      <codemirror
        ref="myCm"
        :value="code"
        :options="cmOptions"
        class="code"
        ></codemirror>
    </div>
    <div class="t-tips t-smaller">其中的Enter注解是一个编译时注解，会在编译的时候讲该类写入到pom中，表示joker应该启动该verticle，推荐一个微服务模块只有一个Entry verticle，其他verticle可以由其启动，应该注意verticle的生命周期管理，要正确启动关闭verticle。</div>
    <div class="t-tips t-smaller">运行第一个微服务（helloworld），启动joker，输入9090可以看到界面后，将helloworld打包成jar包（注意要使用maven Plugins里面的compile命令先编译，lifecycle下的编译命令不会解析编译时注解，这样joker就无法找到启动类，然后使用lifecycle下的打包命令，打包jar包即可在target下出现jar包），拖入joker的entry目录，然后访问http://localhost:9999/就可以看到输出的helloworld了。</div>
  </div>
</template>

<script>

import 'codemirror/mode/javascript/javascript.js'
import 'codemirror/theme/base16-light.css'
import { codemirror } from 'vue-codemirror'

import { helloworld0 } from './javas/helloworld'

export default {
  name:'quick-start',
  data: () => {
    return {
      code: helloworld0,
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
