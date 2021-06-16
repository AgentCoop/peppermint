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
	CallDesc() ClientDescriptor
}

type ServerStream interface {
	Stream
	CallDesc() ServerDescriptor
	FullMethod() string
}


