package pkg

import (
	i "github.com/AgentCoop/peppermint/internal"
)

type Node interface {
	Id() uint
	ExternalId() i.NodeId
	EncKey() []byte
	EncEnabled() bool
	LoadBalancerWeight() int
	//ServiceEndpointByName(string) runtime.ServiceEndpoint
}

type NodeManager interface {
	//EncKeyStore() runtime.InMemoryStore
	//FindByMethodName(fullName string) []Node
	//	FindEncKeyByNodeId(id i.NodeId) []byte
}

type ServiceNodeManager interface {
	NodeManager
	InquiryHub() Node
}
