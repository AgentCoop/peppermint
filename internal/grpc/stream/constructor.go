package stream

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewClientStream(cs grpc.ClientStream, header *metadata.MD, isSecure bool, encKey []byte) *stream {
	s := &stream{
		Stream:   cs,
		header:   header,
		isSecure: isSecure,
		typ:      ClientStream,
		encKey:   encKey,
	}
	return s
}

func NewServerStream(ss grpc.ServerStream, header, trailer *metadata.MD, isSecure bool, encKey []byte) *stream {
	fullMethod, _ := grpc.MethodFromServerStream(ss)
	s := &stream{
		Stream:   ss,
		header:   header,
		trailer: trailer,
		isSecure: isSecure,
		typ:      ServerStream,
		encKey:   encKey,
		fullMethod: fullMethod,
	}
	return s
}
