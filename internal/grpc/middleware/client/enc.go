package client

import (
	"context"
	"fmt"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
)

func encryptMessage(desc g.ClientDescriptor, req interface{}) interface{} {
	if ! desc.SecPolicy().IsSecure() {
		return req
	}
	fmt.Printf("client enc key: %x\n", desc.SecPolicy().EncKey())
	packer := codec.NewPacker(desc.Meta().NodeId(), req, codec.SerializedPacket, desc.SecPolicy().EncKey())
	return packer
}

func SecureChannelUnaryInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		desc := ctx.(g.ClientDescriptor)
		req = encryptMessage(desc, req)
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