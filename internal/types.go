package internal

import (
	"encoding/binary"
	"math/rand"
	"unsafe"
)

type UniqueId uint64
type NodeId UniqueId
type SessionId UniqueId

func (id *UniqueId) FromByteArray(arr []byte) {
	uniqId := binary.BigEndian.Uint64(arr)
	*id = *(*UniqueId)(unsafe.Pointer(&uniqId))
}

func (id *UniqueId) FromRand() {
	uniqId := UniqueId(rand.Uint64())
	*id = *(*UniqueId)(unsafe.Pointer(&uniqId))
}

func (u UniqueId) NodeId() uint64 {
	return uint64(u)
}

func (u UniqueId) SessionId() uint64 {
	return uint64(u)
}

type SignalChan chan struct{}

var (
	SignalEvent = struct{}{}
)

type Application interface {
	//Runtime() run
}