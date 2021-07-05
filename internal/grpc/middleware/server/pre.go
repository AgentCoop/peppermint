package server

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/calldesc"
	"github.com/AgentCoop/peppermint/internal/grpc/stream"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/metadata"
)

func handleMeta(ctx context.Context, svcName string, methodName string) g.ServerDescriptor {
	rt := runtime.GlobalRegistry().Runtime()
	svc := rt.ServiceByName(svcName)
	svcPolicy := svc.Policy()
	method, _ := svcPolicy.FindMethodByName(methodName)
	//secPolicy := calldesc.NewSecurityPolicyFromMethod(method, nil)
	desc := calldesc.NewServer(ctx, svc, method, nil)
	desc.HandleMeta()
	return desc
}

func PreUnaryInterceptor(svcName string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		callDesc := handleMeta(ctx, svcName, info.FullMethod)
		resp, err := handler(callDesc, req)
		return resp, err
	}
}

func PreStreamInterceptor(svcName string) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		callDesc := handleMeta(ss.Context(), svcName, info.FullMethod)
		extended := stream.NewServerStream(ss, callDesc)
		err := handler(srv, extended)
		return err
	}
}

