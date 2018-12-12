<template>
  <div id="quick-start">
    <div class="t-title">joker-core-{version}.jar</div>
    <div class="t-tips">作用</div>
    <div class="t-tips t-small">joker-core包是joker的核心，实现服务器主循环，环境监测、应用部署卸载等核心功能。</div>
    <div class="t-tips">详细介绍</div>
    <div class="t-tips t-smaller">joker-core包由joker-boot启动，joker-core会加载所有vertex相关的lib（libs目录下的jar包），因此想要依赖第三方jar包而不想打包可以在pom.xml中添加相关依赖，安装即可在joker启动的时候自动加载。</div>
    <div class="t-tips t-smaller">ClassLoader：joker-core下实现了joker自定义的classloader（EntryClassLoader和InnerClassLoader）用于自动加载卸载entry（jar包），每个不同hash的jar包都会生成一个InnerClassLoader，而类的查找策略是有InnerClassLoader向上查找，因此可以保证所有不同的entry可以相互调用。</div>
    <div class="t-tips t-smaller">主循环：joker-core主循环会基于hash监测joker的环境变化（entry下的文件变化），然后通过RootVerticle实现entry的加载和卸载。</div>
    <div class="t-tips t-smaller">Entry解析：joker-core在监测到entry增加后，会解析entry然后部署，解析entry第一步解析entry包下的pom.xml文件查找vertx-boot-class并加载其属性，然后会尝试加载resource/webroot下的资源到static路径下，这些文件会被静态映射，主要是一些css、html等静态资源。</div>
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
