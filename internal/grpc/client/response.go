package client

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
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
}

type response struct {
	context.Context
	sId g.SessionId
	md metadata.MD
}

func NewResponse(ctx context.Context) Response {
	md, _ := metadata.FromIncomingContext(ctx)
	r := new(response)
	r.md = md
	r.Context = ctx
	r.sId = utils.GetSessionId(&r.md)
	return r
}
