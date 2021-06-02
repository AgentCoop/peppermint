package client

import (
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/utils/grpc"
	"google.golang.org/grpc/metadata"
)

type RequestHeader interface {
	SetSessionId(i.SessionId)
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
	grpc.SetGrpcSessionId(&r.md, client.SessionId())
	return r
}

func (r *request) SetSessionId(id i.SessionId) {
	grpc.SetGrpcSessionId(&r.md, id)
}

func (r *request) MetaData() metadata.MD {
	return r.md
}
