package server

import (
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/utils"

	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type response struct {
	ctx context.Context
	md metadata.MD
}

type ResponseHeader interface {
	SendHeaders()
	SetSessionId(id g.SessionId)
	AddMetaValue(key string, value string)
	AddBinMetaValue(key string, value []byte)
}

type ResponseData interface {
	ToGrpcResponse() interface{}
}

type Response interface {
	ResponseHeader
	ResponseData
}

func NewResponseHeader() *response {
	r := new(response)
	r.ctx = context.Background()
	r.md = metadata.New(nil)
	return r
}

func (r *response) AddMetaValue(key string, value string) {
	r.md.Append(key, value)
}

func (r *response) AddBinMetaValue(key string, value []byte) {
	r.md.Append(key + "-bin", string(value))
}

func (r *response) SetSessionId(id g.SessionId) {
	r.AddMetaValue(g.META_FIELD_SESSION_ID, utils.IntToHex(id, 16))
}

func (r *response) ToGrpcResponse() interface{} {
	panic("ToGrpcResponse method must be implemented")
}

func (r *response) SendHeaders() {
	grpc.SendHeader(r.ctx, r.md)
}
