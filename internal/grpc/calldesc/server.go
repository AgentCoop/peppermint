package calldesc

import (
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
)

//
// Common interface
//

func (s *sCallDesc) ServiceConfigurator() deps.ServiceConfigurator {
	return s.svcCfg
}

func (s *sCallDesc) HandleMeta() {
	s.meta.extractCommonFieldsVals()
}

func (s *sCallDesc) Meta() grpc.ServerMeta {
	return &s.meta
}

func (s *sCallDesc) WithSecPolicy(sec grpc.SecurityPolicy) {
	s.secPolicy = sec.(secPolicy)
}

func (s *sCallDesc) SecPolicy() grpc.SecurityPolicy {
	return s.secPolicy
}

//
// Call data interface
//
func (s *sCallDesc) ResponseData() grpc.ResponseData {
	return s.resData
}

func (s *sCallDesc) SetResponseData(data grpc.ResponseData) {
	s.resData = data
}

func (s *sCallDesc) RequestData() grpc.RequestData {
	return s.reqData
}

func (s *sCallDesc) SetRequestData(data grpc.RequestData) {
	s.reqData = data
}
