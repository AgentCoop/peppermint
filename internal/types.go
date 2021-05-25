package internal

type NodeId uint64
type SessionId uint64
type SignalChan chan struct{}

var (
	SignalEvent = struct{}{}
)

type Application interface {
	//Runtime() run
}