package codec

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/encoding"
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
	if !isPacket(data) {
		return proto.Unmarshal(data, v.(proto.Message))
	}
	var (
		payload []byte
		err error
		unpacker *unpacker
	)
	if unpacker, err = NewUnpacker(data); err != nil {
		return err
	}
	if payload, err = unpacker.Unpack( nil); err != nil {
		return err
	}
	return proto.Unmarshal(payload, v.(proto.Message))
}

func (codec) Name() string {
	return Name
}
