<template>
  <div id="quick-start">
    <div class="t-title">快速开始</div>
    <div class="t-tips">1、下载源码</div>
    <div class="t-part">
      <div class="t-code">git clone https://github.com/jsen-joker/Joker.git</div>
    </div>
    <div class="t-tips">2、目录解释</div>
    <div class="t-tips t-small">下载源码，进入根目录后你将得到一个joker的目录，在改目录下是joker服务器的最终目录结构，在该目录下，要运行joker容器的核心jar包包括joker-boot.jar、lib/joker-core.jar，lib/joker-program-api（此处省略版本名字）。</div>
    <div class="t-tips t-small">joker的目录结构类似于tomcat，bin是joker的命令文件，entry下是部署的微服务模块，static下会动态解析符合joker格式微服务的静态资源（html等）</div>
    <div class="t-tips t-small">一个纯净的Joker容器至少包含boot.jar，lib目录下core.jar，promram-api.jar及vertx的支持包，entry目录下默认至少应该有manager和manager-master两个服务模块，这是joker完整运行必须的，不建议移除，但移除后joker可照常运行</div>
    <div class="t-tips">3、快速体验</div>
    <div class="t-tips t-smaller">进入到joker目录下并依次执行以下命令</div>
    <div class="t-part">
      <div class="t-code">cd joker/bin</div>
      <div class="t-commen">/**  启动joker  **/</div>
      <div class="t-code">./startup.sh</div>
      <div class="t-commen">/**  停止joker  **/</div>
      <div class="t-code">./shutdown.sh</div>
    </div>
    <div class="t-tips t-smaller">访问管理台，joker提供了一些默认的微服务，一般命名为joker-core-plugin-xxx.jar，这里joker提供了一些方便的微服务可以来管理和查看后台运行的服务。</div>
    <div class="t-tips t-smaller">启动joker后访问 <a href="http://localhost:9090">http://localhost:9090</a> 即可看到joker的管理台，所有这些插件的源码都在源码包下可查看。</div>
    <div class="t-tips">4、joker核心插件简介</div>
    <div class="t-tips t-smaller">joker核心插件的端口可以在conf目录下设置，以下按joker的默认设置进行介绍。</div>
    <div class="t-tips t-smaller"><a href="http://localhost:9090">http://localhost:9090</a> manager-master 的微服务和管理界面</div>
    <div class="t-tips t-smaller"><a href="http://localhost:9091">http://localhost:9091</a> manager节点的端口，收到manager-master节点管理，一个manager节点代表一个joker容器</div>
    <div class="t-tips t-smaller"><a href="http://localhost:9092">http://localhost:9092</a> joker微服务上传下载的微服务，用户可以把自己的微服务上传到该服务下，并在manager-master节点下分发服务，自动部署运行</div>

  </div>
</template>

<script>
export default {
  name:'quick-start'
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
