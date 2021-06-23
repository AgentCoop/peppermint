package codec

import (
	i "github.com/AgentCoop/peppermint/internal"
)

func NewPacker(nodeId i.NodeId, payload interface{}, encKey []byte) *packer {
	p := &packer{}
	p.packet.payload = payload
	p.packet.nodeId = nodeId
	p.kind = Packable
	p.encKey = encKey
	return p
}

func NewPassthroughPacker(payload []byte) *packer {
	p := &packer{}
	p.packet.payload = payload
	p.kind = Passthrough
	return p
}
