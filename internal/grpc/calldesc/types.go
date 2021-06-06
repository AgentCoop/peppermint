package calldesc

import (
	"context"
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
	"google.golang.org/grpc/metadata"
)

type common struct {
	meta      meta
	secPolicy secPolicy
}

type secPolicy struct {
	encKey  []byte
	e2e_Enc bool
}

type meta struct {
	header  metadata.MD
	trailer metadata.MD
	sId     internal.SessionId
	nodeId  internal.NodeId
}

type sCallDesc struct {
	context.Context
	common
	reqData grpc.RequestData
	resData grpc.ResponseData
	svcCfg  deps.ServiceConfigurator
}

func (s *sCallDesc) IsSecure() bool {
	panic("implement me")
}

func (s *sCallDesc) EncKey() []byte {
	panic("implement me")
}

func (s *sCallDesc) SessionId() internal.SessionId {
	panic("implement me")
}

func (s *sCallDesc) NodeId() internal.NodeId {
	panic("implement me")
}

func (s *sCallDesc) SetTrailer(md metadata.MD) {
	panic("implement me")
}

func (s *sCallDesc) SetSessionId(id internal.SessionId) {
	panic("implement me")
}

type cCallDesc struct {
	context.Context
	common
	client client.BaseClient
}
