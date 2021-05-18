package client

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/utils"
	"google.golang.org/grpc/metadata"
)

type ResponseHeader interface {
	GetHeader() *metadata.MD
	GetTrailer() *metadata.MD
}

type ResponseData interface {
	//Populate()
}

type Response interface {
	ResponseHeader
	ResponseData
	Process()
}

type response struct {
	context.Context
	client BaseClient
	header metadata.MD
	trailer metadata.MD
}

func NewResponse(c BaseClient) Response {
	r := new(response)
	r.client = c
	r.header = metadata.New(nil)
	r.trailer = metadata.New(nil)
	return r
}

func (r *response) Process() {
	r.client.SetSessionId(utils.GetSessionId(&r.header))
}

func (r *response) GetHeader() *metadata.MD {
	return &r.header
}

func (r *response) GetTrailer() *metadata.MD {
	return &r.trailer
}
