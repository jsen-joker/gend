<template>
  <div id="quick-start">
    <br />
    <div class="t-tips t-small hey">所有libs目录下的jar包可以直接调用，因此你可以直接调用RootVerticle（joker-core包下的启动类）来定制自己想要的功能。</div>
    <br />
    <br />
    <div class="t-title">joker-boot.jar</div>
    <div class="t-tips">作用</div>
    <div class="t-tips t-small">joker-boot包下包含的是joker的启动器类（main函数），负责处理命令行，并执行相应命令。</div>
    <div class="t-tips">详细介绍</div>
    <div class="t-tips t-smaller">joker下包含了重新的main函数，bin包下的脚本命令通过调用boot包并传入不同的参数（如start，stop）实现joker程序的启动和关闭。</div>
    <div class="t-tips t-smaller">joker的启动（start），joker启动时boot的jar包会调用core包，core包会启动joker相应的主循环，初始化并启动joker。</div>
    <div class="t-tips t-smaller">joker的关闭（stop），joker关闭时会读取joker配置文件，从配置文件中找到manager的端口，通过发送stop命令，因此joker的stop命令需要核心插件manager、和manager-master的支持。</div>
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
