<template>
  <div id="quick-start">
    <div class="t-title">joker-annotation-{version}.jar</div>
    <div class="t-tips">joker注解</div>
    <div class="t-tips t-small">joker编译时注解，引入该jar包（scope compile）只要在Entry类上加上Entry注解，即可，在使用mvn编译插件编译的时候，会自动帮您将该类写到maven里面。目前该jar包只支持Entry一个注解。</div>
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
