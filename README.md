![kratos](doc/img/kratos3.png)

[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![GoDoc](https://godoc.org/git.huoys.com/middle-end/kratos?status.svg)](https://godoc.org/git.huoys.com/middle-end/kratos)

# Kratos

Kratos是[bilibili](https://www.bilibili.com)开源的一套Go微服务框架，包含大量微服务相关框架及工具。  

> 名字来源于:《战神》游戏以希腊神话为背景，讲述由凡人成为战神的奎托斯（Kratos）成为战神并展开弑神屠杀的冒险历程。

## Goals

我们致力于提供完整的微服务研发体验，整合相关框架及工具后，微服务治理相关部分可对整体业务开发周期无感，从而更加聚焦于业务交付。对每位开发者而言，整套Kratos框架也是不错的学习仓库，可以了解和参考到[bilibili](https://www.bilibili.com)在微服务方面的技术积累和经验。

## v0.8新特性
* 使用Gin替换BM HTTP引擎
* 解决BM底层ctx设计缺陷导致的panic问题
* 新的工具链 支持Gin模板生成

### [**升级到0.8需要做什么**](doc/wiki-cn/v0.8/how-to-update.md)

## v0.8失去了什么
* 失去了框架自带的trace链路追踪 因为这样的CTX在旧BM设计缺陷会导致panic
* 可以使用 [**opengintracing**](https://github.com/gin-contrib/opengintracing) 这样的中间件来替代 或者自己去实现基于CTX的链路追踪
* 失去了全局timeout 同样是因为ctx设置全局deadline导致触发parent done导致panic的原因
* 请求timeout应该在Handler上使用中间件实现 比如 Gin官方社区的中间件 [**timeout**](https://github.com/gin-contrib/timeout)


## v0.7新特性
* 全新的[**服务发现**](doc/wiki-cn/v0.7/service-discovery.md)机制 支持原生K8S
* 全量使用gogo protobuf特性[**兼容新旧版本**](doc/wiki-cn/v0.7/other-compatibility.md) 解决google v2API panic问题
* [**全新的工具链**](doc/wiki-cn/v0.7/new-generator-tools.md) 支持gogo特性 生成模板
* 引入[**RPC依赖注入**](doc/wiki-cn/v0.7/RPC-DI.md) 初始化健康检查

## v0.7教程
* [**如何升级到新版本**](doc/wiki-cn/v0.7/how-to-update.md)
* [**其他服务如何调用kratos wardern RPC**](doc/wiki-cn/v0.7/other-service-use-wardenclient.md)
* [**多个不同协议的RPC共存**](doc/wiki-cn/v0.7/multi-protocol-with-wardenclient.md)
* [**在K8S上进行开发调试**](doc/wiki-cn/v0.7/debug-with-other-service-on-k8s.md)

## Features
* HTTP Blademaster：核心基于[gin](https://github.com/gin-gonic/gin)进行模块化设计，简单易用、核心足够轻量；
* GRPC Warden：基于官方gRPC开发，集成[discovery](https://github.com/bilibili/discovery)服务发现，并融合P2C负载均衡；
* Cache：优雅的接口化设计，非常方便的缓存序列化，推荐结合代理模式[overlord](https://github.com/bilibili/overlord)；
* Database：集成MySQL/HBase/TiDB，添加熔断保护和统计支持，可快速发现数据层压力；
* Config：方便易用的[paladin sdk](doc/wiki-cn/config.md)，可配合远程配置中心，实现配置版本管理和更新；
* Log：类似[zap](https://github.com/uber-go/zap)的field实现高性能日志库，并结合log-agent实现远程日志管理；
* Trace：基于opentracing，集成了全链路trace支持（gRPC/HTTP/MySQL/Redis/Memcached）；
* Kratos Tool：工具链，可快速生成标准项目，或者通过Protobuf生成代码，非常便捷使用gRPC、HTTP、swagger文档；

## Quick start

### Requirments

Go version>=1.13

```shell
go get -u github.com/gogo/protobuf
go get -u github.com/gogo/protobuf/protoc-gen-gogofast
go get -u github.com/gogo/googleapis
```

### Installation
### [v0.7工具链安装 必读!](doc/wiki-cn/v0.7/new-generator-tools.md)
```shell
GO111MODULE=on && go get -u git.huoys.com/middle-end/kratos/tool/kratos
cd $GOPATH/src
kratos new kratos-demo    # 同时生成grpc、http模块
// kratos new demo --game # 生成游戏服，包括tcp、http
// kratos new demo --grpc # 只生成grpc模块
// kratos new demo --http # 只生成http模块
```

通过 `kratos new` 会快速生成基于kratos库的脚手架代码，如生成 [kratos-demo](https://git.huoys.com/middle-end/kratos-demo) 

### Build & Run

```shell
cd kratos-demo/cmd
go build
./cmd -conf ../configs
```

打开浏览器访问：[http://localhost:8000/kratos-demo/start](http://localhost:8000/kratos-demo/start)，你会看到输出了`Golang 大法好 ！！！`

[快速开始](doc/wiki-cn/quickstart.md)  [kratos工具](doc/wiki-cn/kratos-tool.md)

## Documentation

> [简体中文](doc/wiki-cn/summary.md)  
> [FAQ](doc/wiki-cn/FAQ.md)  

## License
Kratos is under the MIT license. See the [LICENSE](./LICENSE) file for details.

-------------

*Please report bugs, concerns, suggestions by issues, or join QQ-group 716486124 to discuss problems around source code.*
