package grpc

import (
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
)

type CodecPacket interface {
	NodeId() i.NodeId
	PayloadType() codec.PayloadType
	Payload() interface{}
	Marshal() ([]byte, error)
	Unmarshal(interface{}) error
}

//type Codec interface {
//	GetPacket() CodecPacket
//}

type CodecPacker interface {
	CodecPacket
	Pack() ([]byte, error)
}

type CodecUnpacker interface {
	CodecPacket
	Unpack(v interface{}, encKey []byte) error
}
