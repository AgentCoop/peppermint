package grpc

import (
	"context"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/pkg/service"
	"google.golang.org/grpc/metadata"
)

type CallDescriptor interface {
	context.Context
	Method() service.Method
	SecPolicy() SecurityPolicy
	//WithSecPolicy(policy SecurityPolicy)
	HandleMeta()
}

type ClientDescriptor interface {
	CallDescriptor
	WithSessionFrom(ClientDescriptor)
	Meta() Meta
}

type ServerDescriptor interface {
	CallDescriptor
	Data
	Meta() ServerMeta
	Service() service.Service
	WithSession(Session)
	Session() Session
}

type SecurityPolicy interface {
	IsSecure() bool
	EncKey() []byte
}

type Meta interface {
	SetHeader(metadata.MD)
	SendHeader(metadata.MD) error
	Header() *metadata.MD
	Trailer() *metadata.MD
	SessionId() i.SessionId
	NodeId() i.NodeId
}

type RequestData interface {
	Validate() error
}

type ResponseData interface {
	ToGrpc() interface{}
}

type Data interface {
	ResponseData() ResponseData
	SetResponseData(ResponseData)
	RequestData() RequestData
	SetRequestData(data RequestData)
}

type ServerMeta interface {
	Meta
	SetTrailer(metadata.MD)
	SetSessionId(i.SessionId)
}
