package calldesc

import (
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/pkg/service"
)

func (s *ServerDescriptor) Method() service.Method {
	return s.method
}

//
// Common interface
//

func (s *ServerDescriptor) ServiceConfigurator() service.ServiceConfigurator {
	return s.svcCfg
}

func (s *ServerDescriptor) HandleMeta() {
	s.meta.extractCommonFieldsVals()
}

func (s *ServerDescriptor) Meta() grpc.ServerMeta {
	return &s.meta
}

func (s *ServerDescriptor) WithSecPolicy(sec grpc.SecurityPolicy) {
	s.secPolicy = sec.(*secPolicy)
}

func (s *ServerDescriptor) SecPolicy() grpc.SecurityPolicy {
	return s.secPolicy
}

//
// Call data interface
//
func (s *ServerDescriptor) ResponseData() grpc.ResponseData {
	return s.resData
}

func (s *ServerDescriptor) SetResponseData(data grpc.ResponseData) {
	s.resData = data
}

func (s *ServerDescriptor) RequestData() grpc.RequestData {
	return s.reqData
}

func (s *ServerDescriptor) SetRequestData(data grpc.RequestData) {
	s.reqData = data
}
