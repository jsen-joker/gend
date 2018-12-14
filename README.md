# gend

#### 项目介绍
    Go ID生成器，借助于go优秀的并发性能，打造一款高效的id生成器
    这是Go版本的Vesta，

#### 软件架构
    软件架构说明
    core: gend核心实现
    docs: gend官网
    exporter: gend服务端，目前包含grpc和rest版本
    html: gend官网vue代码
    importer: gend客户端，包含嵌入式版本和grpc版本，这里有一个单独的java项目实现了整合嵌入式和grpc的java客户端
    test: 测试辅助
    
#### 安装教程
    1. cd $GOPATH/src
    2. git clone https://github.com/jsen-joker/gend.git github.com/jsen-joker/gend
    3. run server (importer/grpc/main.go importer/rest/main.go)
    4. run client test (test/client-test/*.go)

### 使用说明
#### gen java版本使用方式(linux or mac os)
    1、配置go环境（过程略），这里假设配置后的GOPATH为~/go
    2、下载gend源码 cd ~/go/src && git clone https://github.com/jsen-joker/gend.git ~/go/src/github.com/jsen-joker/gend
    3、编译服务器端版本 cd ~/go/src/github.com/jsen-joker/gend/exporter/grpc/boot && go build -o gend-grpc-server main.go
    4、运行gend服务器 ~/go/src/github.com/jsen-joker/gend/exporter/grpc/boot/gend-grpc-server
    5、打包客户端java动态库 cd ~/go/src/github.com/jsen-joker/gend/importer && go build -buildmode=c-shared -o gend-client.so ./main.go
    6、复制so文件到java项目的class目录（resources目录）
    7、下载编译java端gend源码 git clone https://github.com/jsen-joker/gend-java.git && mvn clean package
    8、将编译的jar包也添加到java项目，使用方式：
        private static Facade gendFacade = new GendFacade("", "{\"RpcAddress\":\"localhost:50051\"}");
        直接调用即可，这里屏蔽了内嵌版和服务器版，优先调用服务器，调用失败后调用内嵌版的gend，
        也可以直接调用EmbedFacade和GrpcFacade两个类直接使用内嵌的或只是用grpc的，
 
#### 关于配置文件：
无论是内嵌版还是服务器版，客户端初始化的时候都要传入相关的配置文件，配置文件是一个json的字符串，也可以直接使用空字符串表示使用默认设置

gend的所有可配置项：

    type DefaultConfig struct {
        RpcAddress string
        // 默认字段值
        // 版本 扩展 0
        Version int64
        // 0 最大峰值 1 最小粒度
        IdType int64
        // 生成方式 0 嵌入式生成 1 服务器生成
        Gen int64
        // 机器ID
        Machine int64
        // 字段范围
        VersionBits byte
        TypeBits byte
        GenBits byte
        TimeBits byte
        SerialBits byte
        MachineBits byte
    }



#### 性能简单测试（直接go调用）
内嵌版：

    并发    请求此时    平均耗时us
    
    1000    100         26
    1000    500         18
    1000    1000        25
    1000    5000        17
    1000    10000       22

grpc版：

    并发    请求此时    平均耗时us
    
    100    100         23
    100    500         18
    100    1000        18
    100    5000        19
    100    10000       18
