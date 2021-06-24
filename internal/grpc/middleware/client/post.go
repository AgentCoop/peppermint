package client

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/pkg/service"
	"google.golang.org/grpc"
)

func PostUnaryInterceptor(c g.BaseClient, policy service.ServicePolicy) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		desc := ctx.(g.ClientDescriptor)
		opts = append(opts,
			grpc.Header(desc.Meta().Header()),
			grpc.Header(desc.Meta().Trailer()),
		)
		err := invoker(ctx, method, req, reply, cc, opts...)
		desc.HandleMeta()
		c.SetLastCall(desc)
		return err
	}
}