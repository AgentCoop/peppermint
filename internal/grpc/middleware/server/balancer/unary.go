package balancer

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md , _ := metadata.FromIncomingContext(ctx)
		_ = md
		pair := server.NewRequestResponsePair(ctx)
		resp, err := handler(pair, req)
		pair.GetResponse().SendHeader()
		return resp, err
	}
}

