package calldesc

import (
	"context"
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
	"google.golang.org/grpc/metadata"
)

type CallDescType int

const (
	ServerCallDesc CallDescType = iota
	ClientCallDesc
)

type common struct {
	context.Context
	typ       CallDescType
	meta      meta
	secPolicy secPolicy
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

type sCallDesc struct {
	common
	method  string
	reqData grpc.RequestData
	resData grpc.ResponseData
	svcCfg  deps.ServiceConfigurator
}

type cCallDesc struct {
	common
}
