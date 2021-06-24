package codec

import (
	"bytes"
	"encoding/binary"
	"errors"
	i "github.com/AgentCoop/peppermint/internal"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/protobuf/proto"
)

const MagicWordLen = 8

var (
	// head -c 8 /dev/urandom | hexdump -C
	packetMagicWord = [MagicWordLen]byte{0xd6, 0x2f, 0x7b, 0x92, 0x24, 0xfb, 0x37, 0x9c}
	ErrEmptyEncKey  = errors.New("codec: empty encryption key")
)

type packet struct {
	flags   g.PacketFlags
	payload interface{}
	nodeId  i.NodeId
	nonce   []byte
	encKey  []byte
}

func (p *packet) probe() bool {
	data := p.payload.([]byte)
	switch {
	case len(data) <= MagicWordLen:
		return false
	case bytes.Compare(packetMagicWord[:], data[0:MagicWordLen]) == 0:
		return true
	default:
		return false
	}
}

func (p *packet) Parse(data []byte) error {
	p.payload = data
	if !p.probe() || p.flags&g.PassthroughFlag != 0 {
		return nil
	}
	node := data[MagicWordLen : MagicWordLen+8]
	data = data[2*MagicWordLen:]
	p.flags = g.PacketFlags(data[0:1][0])
	data = data[1:]
	nodeReader := bytes.NewReader(node)
	err := binary.Read(nodeReader, binary.BigEndian, &p.nodeId)
	if err != nil {
		return err
	}
	p.payload = data
	return nil
}

func (p *packet) Payload() ([]byte, error) {
	data := p.payload.([]byte)
	switch {
	case p.flags&g.EncryptedFlag != 0:
		return p.decrypt(data)
	default:
		return data, nil
	}
}

func (p *packet) writeFlags(dest *bytes.Buffer, flags g.PacketFlags) {
	dest.Write([]byte{byte(flags)})
}

func (p *packet) Pack() ([]byte, error) {
	if p.flags&g.PassthroughFlag != 0 {
		return p.payload.([]byte), nil
	}
	var out bytes.Buffer
	var flags g.PacketFlags
	// Write packet magic word and node ID
	out.Write(packetMagicWord[:])
	binary.Write(&out, binary.BigEndian, p.nodeId)
	// Marshal message
	data, err := proto.Marshal(p.payload.(proto.Message))
	if len(p.encKey) == 0 {
		p.writeFlags(&out, flags)
		return data, err
	}
	flags |= g.EncryptedFlag
	p.writeFlags(&out, flags)
	encrypted := p.encrypt(data)
	out.Write(encrypted)
	return out.Bytes(), nil
}

func (p *packet) HasFlag(flag g.PacketFlags) bool {
	return p.flags&flag != 0
}

func (p *packet) WithFlags(flags...g.PacketFlags) {
	for _, flag := range flags {
		p.flags |= flag
	}
}

func (p *packet) Unpack(v interface{}) error {
	switch {
	case p.nodeId == 0: // Data are not encrypted
		data := p.payload.([]byte)
		return proto.Unmarshal(data, v.(proto.Message))
	default:
		// Find encryption key
		nodeId := p.nodeId
		rt := runtime.GlobalRegistry().Runtime()
		keyStore := rt.EncKeyStore()
		sk, err := keyStore.Get(nodeId)
		if err != nil {
			return err
		}
		p.encKey = sk.([]byte)
		// Decrypt data
		payload, err := p.decrypt(p.payload.([]byte))
		if err != nil {
			return err
		}
		return proto.Unmarshal(payload, v.(proto.Message))
	}
}
