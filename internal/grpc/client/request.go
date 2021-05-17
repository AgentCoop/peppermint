package client

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type RequestHeader interface {
	SetSessionId(g.SessionId)
	SendHeader()
}

type RequestData interface {

}

type Request interface {
	context.Context
	RequestHeader
	RequestData
	ToGrpcRequest() interface{}
}

type request struct {
	context.Context
	md metadata.MD
}

func (r *request) ToGrpcRequest() interface{} {
	panic("implement me")
}

func NewRequest(ctx context.Context) *request {
	r := new(request)
	r.md = metadata.New(nil)
	r.Context = metadata.NewOutgoingContext(ctx, r.md)
	return r
}

func (r *request) SetSessionId(id g.SessionId) {
	utils.SetSessionId(&r.md, id)
}

func (r *request) SendHeader() {
	grpc.SendHeader(r.Context, r.md)
}
