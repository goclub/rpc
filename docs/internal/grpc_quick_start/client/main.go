package main

import (
	"context"
	"fmt"
	xerr "github.com/goclub/error"
	xhttp "github.com/goclub/http"
	"github.com/goclub/rpc/docs/internal/pbecho"
	"google.golang.org/grpc"
	grpcstatus "google.golang.org/grpc/status"
	"log"
	"net/http"
	"time"
)

func main() {
	// WithInsecure() 不安全的：不使用证书，后续再展开说明证书
	// WithBlock() 堵塞：直到链接成功。当不设置 WithBlock() 时候，如果 server 没有打开。会立即返回 err
	// 可以试试先启动 client 等待几秒后再启动 server,观察输出。删除 grpc.WithBlock() 后再把前面的流程走一遍。
	conn, err := grpc.Dial("localhost:9292", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	log.Print("dial success")
	client := pbecho.NewEchoServiceClient(conn)

	r := xhttp.NewRouter(xhttp.RouterOption{
		// http 拦截器
		OnCatchError: func(c *xhttp.Context, err error) error {
			if status, ok := grpcstatus.FromError(err); ok {
				if status.Code() == 100 {
					return c.WriteBytes([]byte(status.Message()))
				} else {
					c.WriteStatusCode(500)
					log.Print(status.Proto())
					return c.WriteBytes([]byte(fmt.Sprintf("grpc error: code is %d", status.Code())))
				}

			} else {
				// error 不返回错误信息,避免泄露安全信息
				xerr.PrintStack(err)
				c.Writer.WriteHeader(500)
				return nil
			}
		},
	})
	r.HandleFunc(xhttp.Route{xhttp.GET, "/"}, func(c *xhttp.Context) (err error) {
		message := c.Request.URL.Query().Get("m")
		// rpc 的接口响应时间最慢不能应该超过1秒，一般场景下平均速度应该是100ms
		ctx, cancelCtx := context.WithTimeout(c.Request.Context(), time.Second)
		defer cancelCtx()
		reply, err := client.Echo(ctx, &pbecho.MessageRequest{Message: message})
		if err != nil {
			return
		}
		return c.WriteBytes([]byte(reply.Message))
	})
	addr := ":8311"
	origin := "http://127.0.0.1" + addr
	log.Print(origin + "/?m=abc")
	log.Print(origin + "/?m=error")
	log.Print(origin + "/?m=reject")
	server := http.Server{
		Addr:    addr,
		Handler: r,
	}
	log.Print(server.ListenAndServe())

}
