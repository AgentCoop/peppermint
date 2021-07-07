package service

import (
	"github.com/AgentCoop/peppermint/internal/grpc/protobuf"
)

func NewServicePolicy(svcFullName string, svcMethods []string) *svcPolicy {
	policy := new(svcPolicy)
	policy.svcFullName = svcFullName
	policy.methods = make(methodsMap, 0)
	policy.desc = protobuf.NewServiceDescriptor(svcFullName)
	policy.populate(svcMethods)
	return policy
}

func NewBaseService(svcName string) *baseService {
	svc := new(baseService)
	svc.name = svcName
	return svc
}
