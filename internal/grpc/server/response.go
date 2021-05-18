package server

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type response struct {
	context.Context
	md metadata.MD
}

type ResponseHeader interface {
	//context.Context
	SendHeader()
	SetSessionId(id g.SessionId)
	//AddMetaValue(key string, value string)
	//AddBinMetaValue(key string, value []byte)
}

type ResponseData interface {
	ToGrpcResponse() interface{}
}

type Response interface {
	ResponseHeader
	ResponseData
}

func NewResponse(ctx context.Context) *response {
	r := new(response)
	r.md = metadata.New(nil)
	r.Context = ctx
	return r
}

func (r *response) AddMetaValue(key string, value string) {
	r.md.Append(key, value)
}

func (r *response) AddBinMetaValue(key string, value []byte) {
	r.md.Append(key + "-bin", string(value))
}

func (r *response) SetSessionId(id g.SessionId) {
	utils.SetSessionId(&r.md, id)
}

func (r *response) ToGrpcResponse() interface{} {
	panic("ToGrpcResponse method must be implemented")
}

func (r *response) SendHeader() {
	grpc.SendHeader(r.Context, r.md)
}
