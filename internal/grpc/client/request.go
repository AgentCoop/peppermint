package client

import (
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/utils"
	"google.golang.org/grpc/metadata"
)

type RequestHeader interface {
	SetSessionId(g.SessionId)
	MetaData() metadata.MD
}

type RequestData interface {

}

type Request interface {
	RequestHeader
	RequestData
	ToGrpcRequest() interface{}
}

type request struct {
	md metadata.MD
	client BaseClient
}

func (r *request) ToGrpcRequest() interface{} {
	panic("ToGrpcRequest method must be implemented!")
}

func NewRequest(client BaseClient) *request {
	r := new(request)
	r.md = metadata.New(nil)
	r.client = client
	utils.SetSessionId(&r.md, client.SessionId())
	return r
}

func (r *request) SetSessionId(id g.SessionId) {
	utils.SetSessionId(&r.md, id)
}

func (r *request) MetaData() metadata.MD {
	return r.md
}
