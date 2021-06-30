package codec

import (
	i "github.com/AgentCoop/peppermint/internal"
)

func NewPacket(nodeId i.NodeId, sId i.SessionId, payload interface{}, encKey []byte) *packet {
	p := &packet{}
	p.payload = payload
	p.nodeId = nodeId
	p.sId = sId
	p.encKey = encKey
	return p
}

func NewPassthroughPacker() *packet {
	p := &packet{}
	return p
}
