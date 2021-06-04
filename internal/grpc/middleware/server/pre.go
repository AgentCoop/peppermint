package server

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/metadata"
)

func PreUnaryInterceptor(serviceName string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		desc := runtime.GlobalRegistry().LookupService(serviceName)
		callDesc := server.NewCallDescriptor(ctx, desc.Cfg)
		resp, err := handler(callDesc, req)
		//pair.GetResponse().SendHeader()
		return resp, err
	}
}

