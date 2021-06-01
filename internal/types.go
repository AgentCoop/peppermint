package internal

type UniqueId uint64
type NodeId UniqueId
type SessionId UniqueId

func (u UniqueId) NodeId() uint64 { return uint64(u) }
func (u UniqueId) SessionId() uint64 { return uint64(u) }

type SignalChan chan struct{}

var (
	SignalEvent = struct{}{}
)

type Application interface {
	//Runtime() run
}