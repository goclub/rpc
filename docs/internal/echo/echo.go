package echo

import (
	"context"
	xerr "github.com/goclub/error"
	"github.com/goclub/rpc/docs/internal/pbecho"
	"strings"
)

type EchoService struct{}

func (p *EchoService) Echo(ctx context.Context, in *pbecho.MessageRequest) (reply *pbecho.MessageReply, err error) {
	// reply 不能为 nil ,否则会报错 grpc : error while marshaling: proto: Marshal called with nil
	reply = &pbecho.MessageReply{}
	reply.Message = "echo:" + in.Message
	if strings.Contains(in.Message, "error") {
		return nil, xerr.New("some go error")
	}
	if strings.Contains(in.Message, "reject") {
		return nil, xerr.Reject(1, "业务错误消息", false)
	}
	return reply, nil
}
