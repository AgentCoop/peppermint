package metadata

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryServerInterceptor(serviceName string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md , _ := metadata.FromIncomingContext(ctx)
		_ = md
		desc := runtime.GlobalRegistry().LookupService(serviceName)
		pair := server.NewRequestResponsePair(ctx, desc.Cfg)
		resp, err := handler(pair, req)
		pair.GetResponse().SendHeader()
		return resp, err
	}
}
