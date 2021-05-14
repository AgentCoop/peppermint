package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type response struct {
	ctx context.Context
	md metadata.MD
}

type ResponseHeader interface {
	SendHeader()
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

func NewResponseHeader(ctx context.Context) *response {
	r := new(response)
	r.ctx = ctx
	//r.md, _ = metadata.FromOutgoingContext(context.Background())
	r.md = metadata.New(nil)
	return r
}

func (r *response) AddMetaValue(key string, value string) {
	r.md.Append(key, value)
}

func (r *response) AddBinMetaValue(key string, value []byte) {
	r.md.Append(key + "-bin", string(value))
}

func (r *response) ToGrpcResponse() interface{} {
	panic("ToGrpcResponse method must be implemented")
}

func (r *response) SendHeader() {
	grpc.SendHeader(r.ctx, r.md)
}