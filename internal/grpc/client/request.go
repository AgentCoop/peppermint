package client

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/utils"
	"google.golang.org/grpc/metadata"
)

type RequestHeader interface {

}

type RequestData interface {

}

type Request interface {
	RequestHeader
	RequestData
	ToGrpcRequest() interface{}
}

type requestHeader struct {
	md metadata.MD
}

func NewRequestHeader(ctx context.Context) *requestHeader {
	r := new(requestHeader)
	r.md, _ = metadata.FromOutgoingContext(ctx)
	return r
}

func (r *requestHeader) SetSessionId(id grpc.SessionId) {
	utils.SetSessionId(&r.md, id)
}
