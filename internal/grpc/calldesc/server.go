package calldesc

import (
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/pkg/service"
)

func (s *srvDescriptor) Service() service.Service {
	return s.svc
}

func (s *srvDescriptor) Method() service.Method {
	return s.method
}

func (s *srvDescriptor) WithSession(sess grpc.Session) {
	s.sess = sess
}

func (s *srvDescriptor) Session() grpc.Session {
	return s.sess
}

//
// Common interface
//

func (s *srvDescriptor) HandleMeta() {
	s.meta.extractCommonFieldsVals()
}

func (s *srvDescriptor) Meta() grpc.ServerMeta {
	return &s.meta
}

func (s *srvDescriptor) WithSecPolicy(sec grpc.SecurityPolicy) {
	s.secPolicy = sec.(*secPolicy)
}

func (s *srvDescriptor) SecPolicy() grpc.SecurityPolicy {
	return s.secPolicy
}

//
// Call data interface
//
func (s *srvDescriptor) ResponseData() grpc.ResponseData {
	return s.resData
}

func (s *srvDescriptor) SetResponseData(data grpc.ResponseData) {
	s.resData = data
}

func (s *srvDescriptor) RequestData() grpc.RequestData {
	return s.reqData
}

func (s *srvDescriptor) SetRequestData(data grpc.RequestData) {
	s.reqData = data
}
