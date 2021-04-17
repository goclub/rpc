package main

import (
	"context"
	"github.com/goclub/rpc/docs/internal/grpc/pbecho"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	log.Print("start dial")
	// WithInsecure() 不安全的：禁用传输安全，此处为了简单的演示所以使用WithInsecure
	// WithBlock() 堵塞：直到链接成功。当不设置 WithBlock() 时候，如果 server 没有打开。会立即返回 err
	// 可以试试先启动 client 等待几秒后再启动 server,观察输出。删除 grpc.WithBlock() 后再把前面的流程走一遍。
	conn, err := grpc.Dial("localhost:9292", grpc.WithInsecure(), grpc.WithBlock()) ; if err != nil {
		panic(err)
	}
	defer conn.Close()
	log.Print("dial success")
	client := pbecho.NewEchoServiceClient(conn)
	// rpc 的接口响应时间最慢不能应该超过1秒，一般场景下平均速度应该是100ms
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Second)
	defer cancelCtx()
	// 调用服务
	reply, err := client.Echo(ctx, &pbecho.MessageRequest{Message: "goclub"}) ; if err != nil {
		log.Print(err.Error())
		return
	}
	log.Print("reply:", reply.Message)

}
