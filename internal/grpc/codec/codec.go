package codec

import (
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
	packer := NewPacker(0, data, nil)
	err := packer.Parse()
	if err != nil { return err }
	return packer.Unpack(v)
}

func (codec) Name() string {
	return Name
}
