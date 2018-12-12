<template>
  <div id="main">
    <header class="page-header" role="banner">
      <h1 class="project-name">Gend</h1>
      <h2 class="project-tagline">高性能ID生成器，提供java支持包</h2>
        <!-- <a href="https://jsen-joker.github.io/Joker" target="view_window" class="btn">Joker</a> -->
      
        <!-- <a href="https://github.com/jsen-joker/gend" class="btn space">Github</a> -->
        <a href="https://github.com/jsen-joker/gend" class="btn" target="view_window">Github</a>
        <a href="https://github.com/jsen-joker/gend/zipball/master" class="btn">下载 .zip</a>
        <a href="https://github.com/jsen-joker/gend/tarball/master" class="btn">下载 .tar.gz</a>
    </header>
    <div class="content">
      <p class="start">快速开始</p>
      <div class="page-home-start" style="width: 650px;"> 
        <code> 
          <div class="page-home-comment">// 下载</div> 
          <div> 
            <div class="pl-smi">cd $GOPATH/src</div> 
            <div class="pl-smi">git clone https://github.com/jsen-joker/gend.git github.com/jsen-joker/gend</div> 
            <div class="pl-smi">cd github.com/jsen-joker/gend</div> 
          </div> 
          <div class="page-home-comment">// 启动服务端</div> 
          <div> 
            <div class="pl-smi">run server (importer/grpc/main.go importer/rest/main.go)</div> 
          </div>
          <div class="page-home-comment">// 启动测试</div> 
          <div> 
            <div class="pl-smi">run client test (test/client-test/*.go)</div> 
          </div> 
        </code> 
      </div>
      <p class="start small">特性介绍</p>
      <div class="group">
        <div class="item">
          <img src="../assets/light.svg" />
          <div>
            <span>可追查</span> 所有生成的ID都可以反序列化得到对应的时间戳，序列号等信息，便于在特殊情况下追查数据。
          </div>
        </div>
        <div class="item">
          <img src="../assets/quick.svg" />
          <div>
            <span>快速</span> 基于Go语言编写，拥有更高的执行效率，在高并发的情况下可以表现出更优异的性能。
          </div>
        </div>
        <div class="item">
          <img src="../assets/rest.svg" />
          <div>
            <span>rest</span> 支持rest格式的发布，使得只要支持http的项目都可以快速接入。
          </div>
        </div>
        <div class="item">
          <img src="../assets/smile.svg" />
          <div>
            <span>通用</span> 提供java的支持jar包，可以方便的嵌入到通用的java项目中，java支持项目暂未上传到git上，有需要联系QQ：jsen1922279340@163.com。
          </div>
        </div>
      </div>
    </div>
    <footer class="page-home-footer"> <a href="https://github.com/jsen-joker/Joker">© Joker</a> </footer>
  </div>
</template>

<script>
export default {
  name: 'index'
}
</script>

<style lang="scss" scoped>
#main {
  .page-home-footer {
    margin: 60px 0 20px;
    text-align: center;
    a {
      color: #26a2ff;
    }
  }
  .content {
    .start {
      text-align: center;
      color: #147879;
      font-size: 2rem;
    }
    .small {
      font-size: 1.6rem;
    }
    .group {
      margin: 40px auto 60px;
      width: 720px;
      .item {
        color: #333;
        margin-bottom: 50px;
        display: flex;
        img {
          width: 64px;
          height: 64px;
        }
        >div {
          margin-left: 10px;
          margin-right: 20px;
          span {
            padding-left: 20px;
            font-weight: bolder;
            line-height: 1.6;
            vertical-align: middle;
            font-size: 1.4rem;
          }
        }
      }
    }
    .page-home-start {
      margin: 40px auto 60px;
      background-color: #f8f8f8;
      border-radius: 8px;
      border: 1px solid #ddd;
      padding: 15px;
      code {
        display: block;
        margin-top: 25px;
        div.page-home-comment {
          color: #969896;
        }
        >div {
          display: block !important;
          line-height: 20px;
        }
        .pl-smi {
          color: #333;
        }
      }
      :first-child {
        margin-top: 0;
      }
    }
  }
  .page-header {
    padding: 5rem 6rem;
    color: #fff;
    text-align: center;
    background-color: #159957;
    background-image: linear-gradient(120deg, #155799, #159957);
    .project-name {
      font-size: 3.25rem;
      margin-top: 0;
      margin-bottom: 0.1rem;
    }
    .project-tagline {
      font-size: 1.25rem;
      margin-bottom: 2rem;
      font-weight: normal;
      opacity: 0.7;
    }
    .btn {
      display: inline-block;
      margin-bottom: 1rem;
      color: rgba(255,255,255,0.7);
      background-color: rgba(255,255,255,0.08);
      border-color: rgba(255,255,255,0.2);
      border-style: solid;
      border-width: 1px;
      border-radius: 0.3rem;
      transition: color 0.2s, background-color 0.2s, border-color 0.2s;
      padding: 0.75rem 1rem;
    }
    .btn:hover {
      color: rgba(255,255,255,0.8);
      text-decoration: none;
      background-color: rgba(255,255,255,0.2);
      border-color: rgba(255,255,255,0.3);
    }
    .space {
      margin-left: 2rem;
    }
  }
}
</style>
