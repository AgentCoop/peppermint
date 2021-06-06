package stream

import (
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc"
)

func NewClientStream(cs grpc.ClientStream, desc g.ClientCallDesc) *clientStream {
	s := &clientStream{
		cs:       cs,
		callDesc: desc,
	}
	return s
}

func NewServerStream(ss grpc.ServerStream, desc g.ServerCallDesc) *serverStream {
	fullMethod, _ := grpc.MethodFromServerStream(ss)
	s := &serverStream{
		ss:         ss,
		callDesc:   desc,
		fullMethod: fullMethod,
	}
	return s
}
