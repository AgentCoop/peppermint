package grpc

import (
	"context"
	"google.golang.org/grpc/metadata"
)

type Stream interface {
	Context() context.Context
	SendMsg(m interface{}) error
	RecvMsg(m interface{}) error
}

type StreamExtended interface {
	Context() context.Context
	Close()
	Send(interface{}) error
	Recv(interface{}) error
	Header() metadata.MD
	WithNewHeader(*metadata.MD)
	MessagesReceived() int
	MessagesSent() int
	EncKey() []byte
}

type ClientStreamExtended interface {
	StreamExtended
}

type ServerStreamExtended interface {
	StreamExtended
	WithTrailer(*metadata.MD)
	FullMethod() string
}
