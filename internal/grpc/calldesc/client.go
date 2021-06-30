package calldesc

import (
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/pkg/service"
)

func (s *clDescriptor) Method() service.Method {
	return s.common.method
}

func (s *clDescriptor) WithSessionFrom(preceding grpc.ClientDescriptor) {
	s.meta.copySessionId(preceding.(*clDescriptor))
}

func (c *clDescriptor) IsSecure() bool {
	return c.secPolicy.e2e_Enc
}

func (c *clDescriptor) EncKey() []byte {
	return c.secPolicy.encKey
}

func (c *clDescriptor) HandleMeta() {
	c.meta.extractCommonFieldsVals()
}

func (c *clDescriptor) Meta() grpc.Meta {
	return &c.meta
}

func (s *clDescriptor) WithSecPolicy(sec grpc.SecurityPolicy) {
	s.secPolicy = sec.(*secPolicy)
}

func (s *clDescriptor) SecPolicy() grpc.SecurityPolicy {
	return s.secPolicy
}
