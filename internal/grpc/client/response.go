package client

import (
	"github.com/AgentCoop/peppermint/internal/utils"
	"google.golang.org/grpc/metadata"
)

type ResponseHeader interface {
	Header() *metadata.MD
	Trailer() *metadata.MD
}

type Response interface {
	ResponseHeader
	Process()
}

type response struct {
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
	r.client.SetSessionId(utils.Grpc_ExtractSessionId(&r.header))
}

func (r *response) Header() *metadata.MD {
	return &r.header
}

func (r *response) Trailer() *metadata.MD {
	return &r.trailer
}
