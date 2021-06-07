package calldesc

import (
	"github.com/AgentCoop/peppermint/internal/grpc"
)

func (s *cCallDesc) WithSessionFrom(preceding grpc.ClientCallDesc) {
	s.meta.copySessionId(preceding.(*cCallDesc))
}

func (c *cCallDesc) IsSecure() bool {
	return c.secPolicy.e2e_Enc
}

func (c *cCallDesc) EncKey() []byte {
	return c.secPolicy.encKey
}

func (c *cCallDesc) HandleMeta() {
	c.meta.extractCommonFieldsVals()
}

func (c *cCallDesc) Meta() grpc.Meta {
	return &c.meta
}

func (s *cCallDesc) WithSecPolicy(sec grpc.SecurityPolicy) {
	s.secPolicy = sec.(secPolicy)
}

func (s *cCallDesc) SecPolicy() grpc.SecurityPolicy {
	return s.secPolicy
}
