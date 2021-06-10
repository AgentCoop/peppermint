package codec

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
	//"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/encoding"
	"google.golang.org/protobuf/proto"
)

// Name is the name registered for the proto compressor.
const Name = "proto-enc"

func init() {
	encoding.RegisterCodec(codec{})
}

// codec is a Codec implementation with protobuf. It is the default codec for gRPC.
type codec struct{}

func (codec) Marshal(v interface{}) ([]byte, error) {
	p, ok := v.(*packer)
	if !ok {
		return proto.Marshal(v.(proto.Message))
	}
	return p.Pack()
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	if len(data) == 0 {
		return nil
	}
	if !isPacket(data) {
		return proto.Unmarshal(data, v.(proto.Message))
	}
	unpacker, err := NewUnpacker(data);
	if err != nil { return err }

	rt := runtime.GlobalRegistry().Runtime()
	keyStore := rt.NodeManager().EncKeyStore()
	sk, err := keyStore.Get(unpacker.packet.nodeId);
	if err != nil { return err }

	encKey := sk.([]byte)
	payload, err := unpacker.Unpack( encKey);
	if err != nil { return err }

	return proto.Unmarshal(payload, v.(proto.Message))
}

func (codec) Name() string {
	return Name
}
