package balancer

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		srv := info.Server
		balancerJob, err := srv.(service.WebProxyBalancer).Handle(ctx, req, info.FullMethod)
		switch {
		case err != nil:
			return nil, err
		case balancerJob == nil:
			return handler(ctx, req) // Handle a Web proxy API service call
		default:
			<-balancerJob.Run()
			return nil, nil
		}
	}
}

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if info.IsClientStream {
			return status.Error(codes.Unimplemented, "Client streaming is not supported")
		}
		return nil
	}
}