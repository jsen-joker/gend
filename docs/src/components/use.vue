<template>
  <div id="use">
    <el-container>
      <el-aside class="sider sidebar">
        <div class="search">
          <div class="input-wrap">
            <input type="search" value="" placeholder="搜索">
            <div class="clear-button">
              <svg width="26" height="24">
                <circle cx="12" cy="12" r="11" fill="#ccc"></circle>
                <path stroke="white" stroke-width="2" d="M8.25,8.25,15.75,15.75"></path>
                <path stroke="white" stroke-width="2" d="M8.25,15.75,15.75,8.25"></path>
              </svg>
            </div>
          </div>
          <div class="results-panel"></div>
        </div>
        <h1><a class="app-name-link" data-nosearch="" href="/">Joker</a></h1>
        <div class="sidebar-nav">
          <ul>
            <li>
              <p>入门</p>
              <ul>
                <li><a href="#/use/quick-start">快速开始</a></li>
              </ul>
            </li>
            <li>
              <p>模块介绍</p>
              <ul>
                <li><a href="#/use/joker-boot">boot</a></li>
                <li><a href="#/use/joker-core">core</a></li>
                <li><a href="#/use/joker-program">program</a></li>
                <li><a href="#/use/joker-annotation">annotation</a></li>
              </ul>
            </li>
            <li>
              <p>写一个helloworld</p>
              <ul>
                <li><a href="#/use/joker-use">写一个java微服务</a></li>
                <li><a href="#/use/joker-use-js">写一个js微服务</a></li>
              </ul>
            </li>
            <li>
              <p>写一个helloworld</p>
              <ul>
                <li><a href="#/use/joker-use">写一个java微服务</a></li>
                <li><a href="#/use/joker-use-js">写一个js微服务</a></li>
              </ul>
            </li>
          </ul>
          </div>
      </el-aside>
      <el-main>
        <router-view class="u-c-t"/>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name:'useComponent'
}
</script>

<style lang="scss" scoped>
#use {
  height: 100%;
  .u-c-t {
    padding-top: 40px;
    margin-left: 40px;
    margin-right: 40px;
    padding-bottom: 40px;
  }
  section {
    height: 100%;
  }
  aside {
    height: 100%;
  }
  .sider {
    line-height: 2em;
    width: 260px;
    border-right: 1px solid rgba(0,0,0,0.07);
    .search {
      margin-bottom: 20px;
      padding: 6px;
      border-bottom: 1px solid #eee;
      .input-wrap {
        display: flex;
        align-items: center;
        input {
          outline: none;
          border: none;
          width: 100%;
          padding: 0 7px;
          line-height: 36px;
          font-size: 14px;
        }
        .clear-button {
          width: 36px;
          text-align: right;
          display: none;
        }
      }
    }
    > h1 {
      margin: 0 auto 1rem;
      font-size: 1.5rem;
      font-weight: 300;
      text-align: center;
      a {
        color: inherit;
        text-decoration: none;
      }
    }
    ul, li {
      list-style: none;
    }
    ul {
      margin: 0 0 0 15px;
      padding: 0;
      li {
        margin: 6px 0 6px 0;
        p {
          font-weight: 700;
          margin: 0;
        }
        a {
          color: #505d6b;
          font-size: 14px;
          font-weight: normal;
          overflow: hidden;
          text-decoration: none;
          text-overflow: ellipsis;
          white-space: nowrap;
          border-bottom: none;
          display: block;
        }
        a:hover {
          color: #147879;
        }
      }
    }
  }
  .sidebar {
    border-right: 1px solid rgba(0,0,0,.07);
    overflow-y: auto;
    transition: transform .25s ease-out;
    width: 300px;
    z-index: 3;
  }
  .sidebar::-webkit-scrollbar {
      width: 4px
  }

  .sidebar::-webkit-scrollbar-thumb {
      background: transparent;
      border-radius: 4px
  }

  .sidebar:hover::-webkit-scrollbar-thumb {
      background: hsla(0,0%,53%,.4)
  }

  .sidebar:hover::-webkit-scrollbar-track {
      background: hsla(0,0%,53%,.1)
  }
}
</style>
