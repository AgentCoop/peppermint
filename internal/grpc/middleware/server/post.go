package server

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc"
)

func PostUnaryInterceptor(svcName string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		callDesc := ctx.(g.ServerDescriptor)
		callDesc.Meta().SendHeader(nil)
		return resp, err
	}
}

func PostStreamInterceptor(svcName string) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := handler(srv, ss)
		return err
	}
}

