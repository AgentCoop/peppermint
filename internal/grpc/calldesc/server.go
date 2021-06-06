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
	s.meta.sId = grpc.ExtractSessionId(&s.meta.header)
}

func (s *sCallDesc) Meta() *meta {
	return &s.meta
}


// Security policy
func (s *sCallDesc) WithEncKey(key []byte) {
	s.secPolicy.encKey = key
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
