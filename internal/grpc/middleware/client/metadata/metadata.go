package metadata

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	"google.golang.org/grpc"
)

type requestResponsePair struct {
	context.Context
	client.Request
	client.Response
}

func UnaryClientInterceptor(c client.BaseClient) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx := &requestResponsePair{
			ctx,
			client.NewRequest(ctx),
			nil,
		}
		newCtx.SendHeader()
		err := invoker(newCtx, method, req, reply, cc, opts...)
		c.ParseResponseHeader(newCtx)
		if c.SessionId() != 0 {
			newCtx.Request.SetSessionId(c.SessionId())
		}
		return err
	}
}