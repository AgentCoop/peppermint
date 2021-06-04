package grpc

import (
	"context"
)

type Stream interface {
	Context() context.Context
	SendMsg(m interface{}) error
	RecvMsg(m interface{}) error
}

type StreamExtended interface {
	Close()
	Send(interface{}) error
	Recv(interface{}) error
}

type StreamInfo interface {
	EncKey() []byte
	FullMethod() string
	MessagesReceived() int
}
