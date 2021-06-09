package codec

import (
	"bytes"
	"encoding/binary"
	"errors"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/crypto"
	"github.com/golang/protobuf/proto"
)

const MagicWordLen = 8

var (
	// head -c 8 /dev/urandom | hexdump -C
	packetMagicWord = [MagicWordLen]byte{0xd6, 0x2f, 0x7b, 0x92, 0x24, 0xfb, 0x37, 0x9c}
	ErrEmptyEncKey  = errors.New("codec: empty encryption key")
)

type PacketKind int
type PayloadType int

const (
	SerializedPacket PacketKind = iota
	RawPacket
)

// Type of payload depends on the packet kind and provided encryption key
const (
	SerializedPayload PayloadType = iota + 1
	RawPayload
	RawEncryptedPayload
)

type packet struct {
	typ     PayloadType
	payload interface{}
	nodeId  i.NodeId
	nonce   []byte
}

type packer struct {
	packet packet
	encKey []byte
	kind   PacketKind
}

type unpacker struct {
	packet packet
}

func isPacket(data []byte) bool {
	return bytes.Compare(packetMagicWord[:], data[0:MagicWordLen]) == 0
}

func (p *packet) Payload() interface{} {
	return p.payload
}

func (p *packer) typeToByte(typ PayloadType) []byte {
	return []byte{byte(typ)}
}

func (p *packer) encrypt(data []byte) []byte {
	var out bytes.Buffer
	cipher, _ := crypto.NewSymCipher(p.encKey, nil)
	encrypted := cipher.Encrypt(data)
	nonce := cipher.GetNonce()
	noncel := []byte{byte(len(nonce))}
	out.Write(noncel)
	out.Write(nonce)
	out.Write(encrypted)
	return out.Bytes()
}

func (p *unpacker) decrypt(data []byte, encKey []byte) ([]byte, error) {
	noncel := byte(data[0:1][0])
	nonce := data[1:noncel+1]
	encrypted := data[1+noncel:]
	cipher, err := crypto.NewSymCipher(encKey, nonce)
	if err != nil { return nil, err }
	decrypted := cipher.Decrypt(encrypted)
	return decrypted, nil
}

func (p *packer) nodeIdToByte(dest *bytes.Buffer) {
	binary.Write(dest, binary.BigEndian, p.packet.nodeId)
}

func (p *packer) Pack() ([]byte, error) {
	var out bytes.Buffer
	// Write packet magic word and node ID
	out.Write(packetMagicWord[:])
	p.nodeIdToByte(&out)

	switch p.kind {
	case RawPacket:
		payload := p.packet.payload.([]byte)
		if len(p.encKey) == 0 {
			out.Write(p.typeToByte(RawPayload))
			out.Write(p.packet.payload.([]byte))
			return out.Bytes(), nil
		} else {
			out.Write(p.typeToByte(RawEncryptedPayload))
			encrypted := p.encrypt(payload)
			out.Write(encrypted)
			return out.Bytes(), nil
		}
	case SerializedPacket:
		if len(p.encKey) == 0 {
			return nil, ErrEmptyEncKey
		}
		data, err := proto.Marshal(p.packet.payload.(proto.Message))
		if err != nil {
			return nil, err
		}
		out.Write(p.typeToByte(SerializedPayload))
		encrypted := p.encrypt(data)
		out.Write(encrypted)
		return out.Bytes(), nil
	default:
		return nil, nil
	}
}

func (p *unpacker) Unpack(encKey []byte) ([]byte, error){
	var (
		payload []byte
		err error
	)
	switch p.packet.typ {
	case RawEncryptedPayload:
		fallthrough
	case SerializedPayload:
		payload, err = p.decrypt(p.packet.payload.([]byte), encKey)
		if err != nil { return nil, err }
	case RawPayload:
		payload = p.packet.payload.([]byte)
	}
	return payload, nil
}
