package stream

import (
	"google.golang.org/grpc"
)

func NewClientStream(cs grpc.ClientStream, isSecure bool, encKey []byte) *stream {
	s := &stream{
		Stream:   cs,
		isSecure: isSecure,
		typ:      ClientStream,
		encKey:   encKey,
	}
	return s
}

func NewServerStream(ss grpc.ServerStream, isSecure bool, encKey []byte) *stream {
	fullMethod, _ := grpc.MethodFromServerStream(ss)
	s := &stream{
		Stream:   ss,
		isSecure: isSecure,
		typ:      ServerStream,
		encKey:   encKey,
		fullMethod: fullMethod,
	}
	return s
}

