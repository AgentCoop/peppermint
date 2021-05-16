package metadata

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"google.golang.org/grpc"
)

type metadata struct {
	context.Context
	server.RequestHeader
	server.ResponseHeader
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		reqHeader := server.NewRequestHeader(ctx)
		resHeader := server.NewResponseHeader()
		newCtx := &metadata{
			Context:        ctx,
			RequestHeader: reqHeader,
			ResponseHeader: resHeader,
		}

		resp, err := handler(newCtx, req)

		newCtx.ResponseHeader.SendHeaders()
		return resp, err
	}
}