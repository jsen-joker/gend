<template>
  <div id="quick-start">
    <div class="t-title">joker-program-api-{version}.jar</div>
    <div class="t-tips">作用</div>
    <div class="t-tips t-small">基于joker提供一些编程封装，和生命周期的封装。</div>
    <div class="t-tips">详细介绍</div>
    <div class="t-tips t-smaller">RestVerticle：joker下编程建议直接继承RestVerticle来进行开发，这里对服务发现注册等都做了相应的规范，是的其他的核心插件可以比较容易的适配监听一些entry的参数。</div>
    <div class="t-tips t-smaller">JdbcRepositoryWrapper：数据库封装，便于进行数据库连接，该类的使用方式在plugin-login里面有介绍，主要可以用service继承该类，让service具有数据库访问能力。</div>
    <div class="t-tips t-smaller">JokerStaticHandlerImpl：joker静态资源解析，JokerStaticHandlerImpl封装了基于joker的相关路径，可以配合joker解压静态资源后访问到页面。</div>
    <div class="t-tips t-smaller">PathHelp：可以帮助用户快速获取joker根目录。</div>
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
