package main

import (
	"context"
	"errors"
	"github.com/goclub/rpc/docs/internal/grpc/pbecho"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

func main() {
	server := grpc.NewServer()
	pbecho.RegisterEchoServiceServer(server, &EchoService{})
	addr := ":9292"
	log.Print("tcp localhost" + addr)
	listener, err := net.Listen("tcp", ":9292") ; if err != nil {
		panic(err)
	}
	err = server.Serve(listener) ; if err != nil {
		panic(err)
	}
}



type EchoService struct {

}

func (p *EchoService) Echo(ctx context.Context, in *pbecho.MessageRequest) (reply *pbecho.MessageReply, err error) {
	// reply 不能为 nil ,否则会报错 grpc : error while marshaling: proto: Marshal called with nil
	reply = &pbecho.MessageReply{}
	reply.Message = "echo:" + in.Message
	if strings.Contains(in.Message, "fuck") {
		return reply, errors.New("watch you mouth")
	}
	return reply, nil
}