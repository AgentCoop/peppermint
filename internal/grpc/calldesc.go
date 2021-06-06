package grpc

import (
	"context"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
	"google.golang.org/grpc/metadata"
)

type CallDesc interface {
	context.Context
	SecurityPolicy
	Meta() Meta
	HandleMeta()
	SessionId() i.SessionId
	NodeId() i.NodeId
}

type ClientCallDesc interface {
	CallDesc
	//Client() client.BaseClient
}

type ServerCallDesc interface {
	CallDesc
	Data
	ServerMeta
	ServiceConfigurator() deps.ServiceConfigurator
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
	SetTrailer(metadata.MD)
	SetSessionId(i.SessionId)
}
