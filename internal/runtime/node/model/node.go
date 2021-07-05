package model

import (
	"github.com/AgentCoop/peppermint/internal"
)

type node struct {
	*Node
}

func (n *node) Id() uint {
	return n.Node.ID
}

func (n *node) ExternalId() internal.NodeId {
	return n.Node.ExternalId
}

func (n *node) LoadBalancerWeight() int {
	return n.Node.LbWeight
}

func (n *node) Tags() []string {
	tags := make([]string, len(n.Node.Tags))
	for i, v := range n.Node.Tags {
		tags[i] = v.Name
	}
	return tags
}

func (n *node) EncKey() []byte {
	return n.Node.EncKey
}

func (n *node) EncEnabled() bool {
	return n.Node.EncEnabled > 0
}
