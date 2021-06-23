package grpc


type PacketFlags uint8

const (
	EncryptedFlag PacketFlags = 1 << iota
	PassthroughFlag
)

type CodecPacket interface {
	WithFlags(flags...PacketFlags)
	HasFlag(flag PacketFlags) bool
	Payload() ([]byte, error)
	Pack() ([]byte, error)
	Unpack(v interface{}) error
}
