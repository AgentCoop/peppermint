package server

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/calldesc"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/stream"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/metadata"
)

func handleMeta(ctx context.Context, svcName string) g.ServerCallDesc {
	// Assume all calls by default are insecure
	// The call can be upgraded to a secure one by an inner middleware
	desc := runtime.GlobalRegistry().LookupService(svcName)
	callDesc := calldesc.NewServerInsecure(ctx, desc.Cfg)
	callDesc.HandleMeta()
	return callDesc
}

func PreUnaryInterceptor(svcName string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		callDesc := handleMeta(ctx, svcName)
		resp, err := handler(callDesc, req)
		return resp, err
	}
}

func PreStreamInterceptor(svcName string) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		callDesc := handleMeta(ss.Context(), svcName)
		extended := stream.NewServerStream(ss, callDesc)
		err := handler(srv, extended)
		return err
	}
}

