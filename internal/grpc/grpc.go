package grpc

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
)

const (
	META_FIELD_NODE_ID    = "gs-node-id"
	META_FIELD_SESSION_ID = "gs-session-id"
)

type stream struct {
	runtime.Stream
	encKey     []byte
	recvx      int
	fullMethod string
}

func NewStream(grpcStream runtime.Stream, fullMethod string, encKey []byte) *stream {
	s := &stream{
		Stream:     grpcStream,
		encKey:     encKey,
		recvx:      0,
		fullMethod: fullMethod,
	}
	return s
}

func (s *stream) EncKey() []byte {
	return s.encKey
}

func (s *stream) FullMethod() string {
	return s.fullMethod
}

func (s *stream) MessagesReceived() int {
	return s.recvx
}
