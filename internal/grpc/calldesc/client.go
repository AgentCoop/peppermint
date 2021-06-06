package calldesc

import (
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/grpc"
)

func (c cCallDesc) IsSecure() bool {
	return c.secPolicy.e2e_Enc
}

func (c cCallDesc) EncKey() []byte {
	return c.secPolicy.encKey
}

func (c cCallDesc) HandleMeta() {

}

func (c cCallDesc) Meta() grpc.Meta {
	panic("implement me")
}

func (c cCallDesc) SessionId() internal.SessionId {
	panic("implement me")
}

func (c cCallDesc) NodeId() internal.NodeId {
	panic("implement me")
}