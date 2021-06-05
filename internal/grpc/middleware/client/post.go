package client

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	"google.golang.org/grpc"
)

func PostUnaryInterceptor(c client.BaseClient) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		callDesc := ctx.(client.ClientCallDescriptor)
		res := callDesc.(client.Response)
		err := invoker(ctx, method, req, reply, cc,
			grpc.Header(res.Header()),
			grpc.Header(res.Trailer()),
		)
		res.Process()
		return err
	}
}