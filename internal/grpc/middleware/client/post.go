package client

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	"google.golang.org/grpc"
)

func PostUnaryInterceptor(c client.BaseClient) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		callDesc := ctx.(g.ClientDescriptor)
		opts = append(opts,
			grpc.Header(callDesc.Meta().Header()),
			grpc.Header(callDesc.Meta().Trailer()),
		)
		err := invoker(ctx, method, req, reply, cc, opts...)
		callDesc.HandleMeta()
		return err
	}
}