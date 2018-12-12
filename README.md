# gend

#### 项目介绍
Go ID生成器
这是Go版本的Vesta，

#### 软件架构
软件架构说明
core:gend核心实现
exporter:gend服务端，目前包含grpc和rest版本
importer:gend客户端，包含嵌入式版本和grpc版本，这里有一个单独的java项目实现了整合嵌入式和grpc的java客户端

#### 安装教程
1. cd $GOPATH/src
2. git clone https://github.com/jsen-joker/gend.git github.com/jsen-joker/gend
3. run server (importer/grpc/main.go importer/rest/main.go)
4. run client test (test/client-test/*.go)

#### 使用说明

1. wait