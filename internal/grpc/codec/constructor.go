package codec

import (
	"bytes"
	"encoding/binary"
	i "github.com/AgentCoop/peppermint/internal"
)

func NewPacker(nodeId i.NodeId, msg interface{}, kind PacketKind, encKey []byte) *packer {
	p := new(packer)
	p.packet = packet{}
	p.packet.payload = msg
	p.packet.nodeId = nodeId
	p.kind = kind
	p.encKey = encKey
	return p
}

func NewUnpacker(data []byte) (*unpacker, error) {
	u := new(unpacker)
	u.packet = packet{}
	node := data[MagicWordLen:MagicWordLen+8]
	data = data[2*MagicWordLen:]
	typ := PayloadType(data[0:1][0])
	data = data[1:]
	nodeReader := bytes.NewReader(node)
	binary.Read(nodeReader, binary.BigEndian, &u.packet.nodeId)
	u.packet.payload = data
	u.packet.typ = typ
	return u, nil
}
