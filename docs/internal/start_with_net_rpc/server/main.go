package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"strings"
	"time"
)

func main() {
	server := rpc.NewServer()
	echoService := EchoService{}
	// 注册服务，第二个参数必须传指针 new(EchoService) 语法也可以快速创建 EchoService 的指针
	err := server.RegisterName("echo", &echoService) ; if err != nil {
		panic(err)
	}
	addr := ":8888"
	log.Print("127.0.0.1" + addr)
	// 监听 tcp
	listener, err := net.Listen("tcp", addr) ; if err != nil {
		panic(err)
	}
	// 死循环获取 Accept，
	for {
		// Accept 是堵塞的，当接收到 tcp 请求时返回conn, err
		log.Print("prepare accept")
		conn, err := listener.Accept()
		log.Print("accepted")
		if err != nil {
			log.Print(err)
			continue
		}
		// 启动 routine 处理 rpc 请求 (为了容易理解这里不处理子 routine panic 的问题 github.com/goclub/sync)
		// 如果不开启 goroutine 运行 server.ServeConn(conn) ,则会堵塞
		go server.ServeConn(conn)
	}
}

type EchoService struct {

}

func (p *EchoService) Message(request string, reply *string) error {
	*reply = "echo:" + request
	// 使用 . 模拟运行时间，. 越多耗时越长
	dotCount := time.Duration(strings.Count(request, "."))
	duration := dotCount * time.Second
	log.Print("执行时间：", duration)
	time.Sleep(duration)
	if strings.Contains(request, "fuck") {
		return errors.New("watch you mouth")
	}
	return nil
}