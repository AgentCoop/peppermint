package calldesc

import (
	"context"
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/pkg/service"
	"google.golang.org/grpc/metadata"
)

type DescriptorType int

const (
	ServerType DescriptorType = iota
	ClientType
)

type common struct {
	context.Context
	typ       DescriptorType
	meta      meta
	secPolicy *secPolicy
	method    service.Method
}

type secPolicy struct {
	encKey  []byte
	e2e_Enc bool
}

type meta struct {
	parent  *common
	header  metadata.MD
	trailer metadata.MD
	sId     internal.SessionId
	nodeId  internal.NodeId
}

type srvDescriptor struct {
	common
	svc     service.Service
	method  service.Method
	reqData grpc.RequestData
	resData grpc.ResponseData
	sess    grpc.Session
}

type clDescriptor struct {
	common
}
