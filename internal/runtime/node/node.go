package node

import "github.com/AgentCoop/peppermint/internal"

type node struct {
	externalId internal.NodeId
	encEnabled bool
	tags       []string
	encKey     []byte
}
