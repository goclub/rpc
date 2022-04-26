package main

import (
	"context"
	"encoding/json"
	xerr "github.com/goclub/error"
	"github.com/goclub/rpc/docs/internal/echo"
	"github.com/goclub/rpc/docs/internal/pbecho"
	"google.golang.org/grpc"
	grpcstatus "google.golang.org/grpc/status"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer(
		// 注册错误拦截器
		grpc.UnaryInterceptor(errInterceptor),
	)
	pbecho.RegisterEchoServiceServer(server, &echo.EchoService{})
	addr := ":9292"
	log.Print("tcp localhost" + addr)
	listener, err := net.Listen("tcp", ":9292")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}

func errInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 继续处理请求
	result, err := handler(ctx, req)
	if err != nil {
		if reject, ok := xerr.AsReject(err); ok {
			message := ""
			errResp, err := json.Marshal(reject.Resp())
			if err != nil {
				message = err.Error()
			} else {
				message = string(errResp)
			}
			return result, grpcstatus.New(100, message).Err()
		}
	}
	return result, err
}
