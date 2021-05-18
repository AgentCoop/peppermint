package client

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/utils"
	"google.golang.org/grpc/metadata"
)

type ResponseHeader interface {
}

type ResponseData interface {
	//Populate()
}

type Response interface {
	//ResponseHeader
	ResponseData
	Process()
}

type response struct {
	context.Context
	client BaseClient
}

func NewResponse(c BaseClient, ctx context.Context) Response {
	r := new(response)
	r.client = c
	r.Context = ctx
	return r
}

func (r *response) Process() {
	md, _ := metadata.FromIncomingContext(r.Context)
	r.client.SetSessionId(utils.GetSessionId(&md))
}
