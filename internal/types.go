package internal

import (
	"encoding/binary"
	"math/rand"
)

type UniqueId uint64
type NodeId UniqueId
type SessionId UniqueId

func (id UniqueId) FromByteArray(arr []byte) UniqueId {
	return UniqueId(binary.BigEndian.Uint64(arr))
}

func (id UniqueId) Rand() UniqueId {
	return UniqueId(rand.Uint64())
}

func (u UniqueId) NodeId() NodeId {
	return NodeId(u)
}

func (u UniqueId) SessionId() SessionId {
	return SessionId(u)
}

type SignalChan chan struct{}

var (
	NotifyEvent = struct{}{}
)

type Application interface {
	//Runtime() run
}