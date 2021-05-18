package metadata

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	"google.golang.org/grpc"
)

func UnaryClientInterceptor(c client.BaseClient) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx := client.NewRequestResponsePair(c, ctx)
		err := invoker(newCtx, method, req, reply, cc, opts...)
		return err
	}
}