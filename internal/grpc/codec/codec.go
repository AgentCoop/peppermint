package codec

import (
	"bytes"
	"github.com/AgentCoop/peppermint/internal/crypto"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/encoding"
)

// Name is the name registered for the proto compressor.
const Name = "proto-enc"

type packetType int
const (
	RawPacket packetType = iota
	SerializedPacket
)

func init() {
	encoding.RegisterCodec(codec{})
}

// codec is a Codec implementation with protobuf. It is the default codec for gRPC.
type codec struct {}

type packet struct {
	typ     packetType
	payload interface{}
	encKey  []byte
}

type Packet interface {
	Payload() interface{}
}

func NewPacket(message interface{}, encKey []byte) *packet {
	return &packet{SerializedPacket, message, encKey}
}

func NewRawPacket(raw []byte, encKey []byte) *packet {
	return &packet{RawPacket,  raw, encKey}
}

func (p *packet) Payload() interface{} {
	return p.payload
}

func (codec) Marshal(v interface{}) ([]byte, error) {
	p, ok := v.(packet)
	if !ok {
		return proto.Marshal(v.(proto.Message))
	}
	var data []byte
	var err error
	switch p.typ {
	case SerializedPacket:
		data, err = proto.Marshal(p.payload.(proto.Message))
		if err != nil { panic(err) }
	case RawPacket:
		data = p.payload.([]byte)
	}
	return encrypt(data, p.encKey), nil
}

// Encrypt payload using symmetric cipher.
// Encrypted payload will be prefixed with a cryptographic nonce preceded by its one-byte length
func encrypt(in []byte, key []byte) []byte {
	if key == nil {
		return in
	}
	cipher := crypto.NewSymCipher(key, nil)
	encrypted := cipher.Encrypt(in)
	nonce := cipher.GetNonce()
	noncel := []byte{byte(len(nonce))}
	var out bytes.Buffer
	out.Write(noncel)
	out.Write(nonce)
	out.Write(encrypted)
	return out.Bytes()
}

func decrypt(in []byte, key []byte) []byte {
	if key == nil {
		return in
	}
	noncel := int(in[0:1][0])
	nonce := in[1:noncel]
	encrypted := in[1+noncel:]
	cipher := crypto.NewSymCipher(key, nonce)
	decrypted := cipher.Decrypt(encrypted)
	return decrypted
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	p, ok := v.(packet)
	if !ok {
		return proto.Unmarshal(data, v.(proto.Message))
	}
	var err error
	switch p.typ {
	case SerializedPacket:
		decrypted := decrypt(data, p.encKey)
		err = proto.Unmarshal(decrypted, p.payload.(proto.Message))
	case RawPacket:
		p.payload = decrypt(data, p.encKey)
	}
	return err
}

func (codec) Name() string {
	return Name
}
