# 先通过 net/rpc 混个脸熟

多的不说了，一切尽在注释中！

启动 rpc 服务

```sh
go run server/main.go
```

发起 rpc 请求

```sh
go run client_call/main.go
```

多个客户端发起多次 rpc 请求

```sh
go run client_multiple_calls/main.go
```

发起异步 rpc 请求

```sh
go run client_go/main.go
```




