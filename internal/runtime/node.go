package runtime

import i "github.com/AgentCoop/peppermint/internal"

type Node interface {
	Id() i.NodeId
	EncKey() []byte
	ServiceEndpointByName(string) ServiceEndpoint
}

type NodeManager interface {
	EncKeyStore() InMemoryStore
}

type ServiceNodeManager interface {
	NodeManager
	InquiryHub() Node
}

