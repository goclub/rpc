# 简单熟悉一下 grpc

> 扩展阅读: [gRPC 扩展错误处理](https://blog.cong.moe/post/2021-12-29-grpc-richer-error-handling/)

## 定义接口

编写 [echoproto](../pbecho/echo.proto)

> 参考文档：https://developers.google.com/protocol-buffers/docs/gotutorial

## 安装protoc 并生成pb.go 文件

安装文档： https://www.grpc.io/docs/languages/go/quickstart/

执行命令[gen.sh](../pbecho/gen.sh)

## 实现服务端

[server/main.go](./server/main.go)


## 实现客户端

[client/main.go](./client/main.go)


## 执行


简单测试：

```sh
go run server/main.go
# 新开一个终端
go run client/main.go
```


测试 grpc.WithBlock()

```sh
# 关闭所有中断
go run client/main.go
# 新开一个中断
# 等待2秒
go run server/main.go
# 观察 server/main.go 运行时 client/main.go 终端的输出
```
