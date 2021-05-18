package metadata

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"google.golang.org/grpc"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		pair := server.NewRequestResponsePair(ctx)
		resp, err := handler(pair, req)
		pair.GetResponse().SendHeader()
		return resp, err
	}
}
