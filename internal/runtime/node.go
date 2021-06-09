package runtime

import i "github.com/AgentCoop/peppermint/internal"

type Node interface {
	Id() i.NodeId
	EncKey() []byte
	ServiceEndpointByName(string) ServiceEndpoint
}

type EncKeyStore interface {
	Store
}

type NodeManager interface {
	InquiryHub() Node
	EncKeyStore() EncKeyStore
}
