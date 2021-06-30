package calldesc

import (
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/pkg/service"
)

func (s *ClientDescriptor) Method() service.Method {
	return s.common.method
}

func (s *ClientDescriptor) WithSessionFrom(preceding grpc.ClientDescriptor) {
	s.meta.copySessionId(preceding.(*ClientDescriptor))
}

func (c *ClientDescriptor) IsSecure() bool {
	return c.secPolicy.e2e_Enc
}

func (c *ClientDescriptor) EncKey() []byte {
	return c.secPolicy.encKey
}

func (c *ClientDescriptor) HandleMeta() {
	c.meta.extractCommonFieldsVals()
}

func (c *ClientDescriptor) Meta() grpc.Meta {
	return &c.meta
}

func (s *ClientDescriptor) WithSecPolicy(sec grpc.SecurityPolicy) {
	s.secPolicy = sec.(*secPolicy)
}

func (s *ClientDescriptor) SecPolicy() grpc.SecurityPolicy {
	return s.secPolicy
}
