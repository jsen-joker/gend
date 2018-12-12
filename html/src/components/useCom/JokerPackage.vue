<template>
  <div id="quick-start">
    <div class="t-title">打包Joker为jar包</div>
    <div class="t-tips">打包核心jar</div>
    <div class="t-tips t-small">joker的核心jar包由boot、core和program-api，这些jar包只要在idea打开，maven正确引入依赖后，使用maven打包命令（mven面板）打包即可。插件打包完成后，boot jar包必须要重命名为joker-boot.jar放在joker目录下，另外两个拖入libs目录下。</div>
    <div class="t-tips">打包插件</div>
    <div class="t-tips t-smaller">joker的插件往往依赖于program-api、core等jar包，所以首先要把上一步打包的三个jar包安装到本地mvn库，mvn install 命令，然后才能编译其他插件，如果有其他的本地依赖也需要安装才能编译。</div>
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
