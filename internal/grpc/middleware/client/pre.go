package client

import (
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	"google.golang.org/grpc"
	"context"
)

func PreUnaryInterceptor(c client.BaseClient) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc)
		return err
	}
}
