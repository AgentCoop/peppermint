package metadata

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/client"

	//	"github.com/AgentCoop/peppermint/internal/grpc/client"

	//"github.com/AgentCoop/peppermint/internal/grpc/client"
	"google.golang.org/grpc"
)

type metadata struct {
	context.Context
	client.RequestHeader
}

func UnaryClientInterceptor(client client.BaseClient) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		//reqHeader := client.NewRequestHeader(ctx)
		//newCtx := &metadata{
		//	Context:        ctx,
		//	RequestHeader: reqHeader,
		//}
		//if client.SessionId() != 0 {
		//
		//}
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}