package codec

import (
	"bytes"
	"encoding/binary"
	"errors"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/protobuf/proto"
)

const MagicWordLen = 8

var (
	// head -c 8 /dev/urandom | hexdump -C
	packetMagicWord = [MagicWordLen]byte{0xd6, 0x2f, 0x7b, 0x92, 0x24, 0xfb, 0x37, 0x9c}
	ErrEmptyEncKey  = errors.New("codec: empty encryption key")
)

type PacketKind int
type PacketType int
type PayloadType int

const (
	Packable PacketKind = iota
	Passthrough
)

type packet struct {
	encryptedFlag byte
	payload       interface{}
	nodeId        i.NodeId
	nonce         []byte
}

type packer struct {
	packet packet
	encKey []byte
	kind   PacketKind
}

func (p *packer) probe() bool {
	data := p.packet.payload.([]byte)
	switch {
	case len(data) <= MagicWordLen:
		return false
	case bytes.Compare(packetMagicWord[:], data[0:MagicWordLen]) == 0:
		return true
	default:
		return false
	}
}

func (p *packer) Parse() error {
	if !p.probe() { return nil }
	data := p.packet.payload.([]byte)
	node := data[MagicWordLen : MagicWordLen+8]
	data = data[2*MagicWordLen:]
	encFlag := data[0:1][0]
	data = data[1:]
	nodeReader := bytes.NewReader(node)
	err := binary.Read(nodeReader, binary.BigEndian, &p.packet.nodeId)
	if err != nil {	return err }
	p.packet.payload = data
	p.packet.encryptedFlag = encFlag
	return nil
}

func (p *packet) Payload() interface{} {
	return p.payload
}

func (p *packer) Pack() ([]byte, error) {
	if p.kind == Passthrough {
		return p.packet.payload.([]byte), nil
	}
	var out bytes.Buffer
	// Write packet magic word and node ID
	out.Write(packetMagicWord[:])
	binary.Write(&out, binary.BigEndian, p.packet.nodeId)
	// Marshal message
	data, err := proto.Marshal(p.packet.payload.(proto.Message))
	if len(p.encKey) == 0 {
		return data, err
	}
	out.Write([]byte{p.packet.encryptedFlag})
	encrypted := encrypt(data, p.encKey)
	out.Write(encrypted)
	return out.Bytes(), nil
}

func (p *packer) Unpack(v interface{}) error {
	switch {
	case p.kind == Passthrough: // Do nothing, leave it as it is
		return nil
	case p.packet.nodeId == 0: // Data are not encrypted
		data := p.packet.payload.([]byte)
		return proto.Unmarshal(data, v.(proto.Message))
	default:
		// Find encryption key
		nodeId := p.packet.nodeId
		rt := runtime.GlobalRegistry().Runtime()
		keyStore := rt.NodeManager().EncKeyStore()
		sk, err := keyStore.Get(nodeId)
		if err != nil { return err }
		encKey := sk.([]byte)
		// Decrypt data
		payload, err := decrypt(p.packet.payload.([]byte), encKey)
		if err != nil { return err }
		return proto.Unmarshal(payload, v.(proto.Message))
	}
}
