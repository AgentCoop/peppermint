package codec

import (
	i "github.com/AgentCoop/peppermint/internal"
)

func NewPacket(nodeId i.NodeId, payload interface{}, encKey []byte) *packet {
	p := &packet{}
	p.payload = payload
	p.nodeId = nodeId
	p.encKey = encKey
	return p
}

func NewPassthroughPacker() *packet {
	p := &packet{}
	return p
}
