package codec

import (
	"bytes"
	"encoding/binary"
	"errors"
	i "github.com/AgentCoop/peppermint/internal"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/session"
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
	sId     i.SessionId
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
	// Extract node ID
	data = data[MagicWordLen:]
	nodeIdData := data[0:i.NodeId(0).Size()]
	err := (&p.nodeId).NetRead(nodeIdData)
	if err != nil {
		return err
	}
	// Session ID
	data = data[i.NodeId(0).Size():]
	sIdData := data[0:i.SessionId(0).Size()]
	err = (&p.sId).NetRead(sIdData)
	if err != nil {
		return err
	}
	// Packet flags
	data = data[i.SessionId(0).Size():]
	p.flags = g.PacketFlags(data[0:1][0])
	data = data[1:]
	// Payload
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
	// Write packet magic word, node and session ID
	out.Write(packetMagicWord[:])
	binary.Write(&out, binary.BigEndian, p.nodeId)
	binary.Write(&out, binary.BigEndian, p.sId)
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

func (p *packet) WithFlags(flags ...g.PacketFlags) {
	for _, flag := range flags {
		p.flags |= flag
	}
}

func (p *packet) fetchEncKey() error {
	sess, err := session.FindById(p.sId)
	if err == nil {
		if v, ok := sess.TaskContext().(interface{ EncKey() []byte }); ok {
			p.encKey = v.EncKey()
			return nil
		}
	}
	nodeId := p.nodeId
	rt := runtime.GlobalRegistry().Runtime()
	keyStore := rt.EncKeyStore()
	sk, err := keyStore.Get(nodeId)
	if err != nil {
		return err
	}
	p.encKey = sk.([]byte)
	return nil
}

func (p *packet) Unpack(v interface{}) error {
	switch {
	case p.nodeId == 0: // Data are not encrypted
		data := p.payload.([]byte)
		return proto.Unmarshal(data, v.(proto.Message))
	default:
		p.fetchEncKey()
		// Decrypt data
		payload, err := p.decrypt(p.payload.([]byte))
		if err != nil {
			return err
		}
		return proto.Unmarshal(payload, v.(proto.Message))
	}
}
