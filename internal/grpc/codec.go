package grpc

import i "github.com/AgentCoop/peppermint/internal"



type CodecPacket interface {
	NodeId() i.NodeId
	PayloadType()
}

type Codec interface {
	GetPacket() CodecPacket
}

type CodecPacker interface {
	Pack() ([]byte, error)
}

type CodecUnpacker interface {
	Unpack(v interface{}, encKey []byte) error
}
