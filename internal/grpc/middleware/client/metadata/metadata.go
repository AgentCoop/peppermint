package metadata

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	"google.golang.org/grpc"
)

func UnaryClientInterceptor(c client.BaseClient) grpc.UnaryClientInterceptor {
	return func(pair context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		s := grpc.ServerTransportStreamFromContext(pair)
		_ = s
		//_req := pair.(client.RequestResponsePair).GetRequest()
		resp := pair.(client.RequestResponsePair).GetResponse()
		pair.(client.RequestResponsePair).SendHeader()
		err := invoker(pair, method, req, reply, cc,
			grpc.Header(resp.GetHeader()),
			grpc.Trailer(resp.GetTrailer()),
		)
		resp.Process()
		return err
	}
}