package internal

import (
	"encoding/binary"
	"math/rand"
)

type UniqueId int64
type NodeId UniqueId
type SessionId UniqueId

func (id UniqueId) FromByteArray(arr []byte) UniqueId {
	u64 := binary.BigEndian.Uint64(arr)
	return UniqueId(u64>>1)
}

func (id UniqueId) Rand() UniqueId {
	u64 := rand.Uint64()
	return UniqueId(u64>>1)
}

func (u UniqueId) NodeId() NodeId {
	return NodeId(u)
}

func (u UniqueId) SessionId() SessionId {
	return SessionId(u)
}

func (NodeId) Size() int { return 8 }

type SignalChan chan struct{}

var (
	NotifyEvent = struct{}{}
)

type Application interface {
	//Runtime() run
}