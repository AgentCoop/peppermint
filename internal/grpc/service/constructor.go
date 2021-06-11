package service

import (
	"github.com/AgentCoop/peppermint/internal/grpc/protobuf"
)

func NewServicePolicy(svcFullName string, svcMethods []string) *svcPolicy {
	policy := new(svcPolicy)
	policy.desc = protobuf.NewServiceDescriptor(svcFullName)
	policy.mOptsReceiver = make(map[string]*methodOptions, len(svcMethods))
	for i := 0; i < len(svcMethods); i++ {
		policy.mOptsReceiver[svcMethods[i]] = &methodOptions{}
	}
	policy.populate()
	return policy
}

