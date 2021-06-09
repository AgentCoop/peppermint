package client

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
)

func SecureChannelUnaryInterceptor(payloadTyp codec.PayloadType) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		callDesc := ctx.(g.ClientCallDesc)
		if callDesc.SecPolicy().IsSecure() {
			req = codec.NewPacker(req, payloadTyp, callDesc.SecPolicy().EncKey())
		}
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}

func SecureChannelStreamInterceptor(c client.BaseClient) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		clientStream, err := streamer(ctx, desc, cc, method, opts...)
		return clientStream, err
	}
}