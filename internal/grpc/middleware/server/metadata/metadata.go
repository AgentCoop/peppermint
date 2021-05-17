package metadata

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"google.golang.org/grpc"
)

type requestResponsePair struct {
	context.Context
	server.Request
	server.Response
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx := &requestResponsePair{
			ctx,
			server.NewRequest(ctx),
			server.NewResponse(ctx),
		}
		resp, err := handler(newCtx, req)
		newCtx.Response.SendHeader()
		return resp, err
	}
}