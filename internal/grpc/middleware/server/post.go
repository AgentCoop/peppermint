package server

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"google.golang.org/grpc"
)

func PostUnaryInterceptor(serviceName string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		callDesc := ctx.(server.GrpcCallDescriptor)
		callDesc.GetResponse().SendHeader()
		return resp, err
	}
}


