package service

import (
	"github.com/AgentCoop/peppermint/internal/grpc/protobuf"
	"github.com/AgentCoop/peppermint/pkg/grpc"
	"github.com/AgentCoop/peppermint/pkg/service"
)

func NewServicePolicy(svcFullName string, svcMethods []string) *svcPolicy {
	policy := new(svcPolicy)
	policy.svcFullName = svcFullName
	policy.methods = make(methodsMap, 0)
	policy.desc = protobuf.NewServiceDescriptor(svcFullName)
	policy.populate(svcMethods)
	return policy
}

func NewBaseService(srv grpc.BaseServer, ipcSrv grpc.BaseServer, cfg service.ServiceConfigurator, policy service.ServicePolicy) *baseService {
	svc := new(baseService)
	svc.srv = srv
	svc.cfg = cfg
	svc.ipcSrv = ipcSrv
	svc.policy = policy
	return svc
}
