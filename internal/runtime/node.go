package runtime

import i "github.com/AgentCoop/peppermint/internal"

type Node interface {
	Id() i.NodeId
	EncKey() []byte
	LoadBalancerWeight() int
 	ServiceEndpointByName(string) ServiceEndpoint
}

type NodeManager interface {
	EncKeyStore() InMemoryStore
	//FindByMethodName(fullName string) []Node
//	FindEncKeyByNodeId(id i.NodeId) []byte
}

type ServiceNodeManager interface {
	NodeManager
	InquiryHub() Node
}

