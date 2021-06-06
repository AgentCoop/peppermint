package grpc

import (
	"context"
)

type Stream interface {
	Context() context.Context
	Close()
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	ReceivedCount() int
	SentCount() int
}

type ClientStream interface {
	Stream
	CallDesc() ClientCallDesc
}

type ServerStream interface {
	Stream
	CallDesc() ServerCallDesc
	FullMethod() string
}


