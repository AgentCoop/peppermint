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
	//context.Context
	RequestHeader
	RequestData
	ToGrpcRequest() interface{}
}

type request struct {
	context.Context
	md metadata.MD
	client BaseClient
}

func (r *request) ToGrpcRequest() interface{} {
	//return nil
	panic("implement me")
}

func NewRequest(client BaseClient, ctx context.Context) *request {
	r := new(request)
	r.md = metadata.New(nil)
	r.Context = ctx
	r.client = client
	return r
}

func (r *request) SetSessionId(id g.SessionId) {
	utils.SetSessionId(&r.md, id)
}

func (r *request) SendHeader() {
	grpc.SendHeader(r.Context, r.md)
}
