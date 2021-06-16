package service

import (
	"github.com/AgentCoop/peppermint/internal/grpc/protobuf"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func NewServicePolicy(svcFullName string, svcMethods []string) *svcPolicy {
	policy := new(svcPolicy)
	policy.svcFullName = svcFullName
	policy.methods = make(methodsMap, len(svcMethods))
	policy.desc = protobuf.NewServiceDescriptor(svcFullName)
	policy.populate(svcMethods)
	return policy
}

func NewBaseService(srv runtime.BaseServer, ipcSrv runtime.BaseServer, cfg runtime.ServiceConfigurator, policy runtime.ServicePolicy) *baseService {
	svc := new(baseService)
	svc.srv = srv
	svc.cfg = cfg
	svc.ipcSrv = ipcSrv
	svc.policy = policy
	return svc
}
