package service

import (
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/protobuf"
)

func NewServicePolicy(svcFullName string, svcMethods []string) *svcPolicy {
	policy := new(svcPolicy)
	policy.svcFullName = svcFullName
	policy.methods = make(methodsMap, len(svcMethods))
	policy.desc = protobuf.NewServiceDescriptor(svcFullName)
	policy.populate(svcMethods)
	return policy
}

func NewBaseService(srv grpc.BaseServer, policy grpc.ServicePolicy) *baseService {
	svc := new(baseService)
	svc.srv = srv
	svc.policy = policy
	return svc
}