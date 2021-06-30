package internal

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"reflect"
)

type UniqueIdValue interface {
	FromByteArray([]byte)
	Rand()
	Size() int
	NetRead([]byte) error
}

type UniqueId int64
type NodeId UniqueId
type SessionId UniqueId

var _ UniqueIdValue = new(NodeId)
var _ UniqueIdValue = new(SessionId)

func (n *NodeId) FromByteArray(arr []byte) {
	i64 := binary.BigEndian.Uint64(arr)>>1
	*n = NodeId(i64)
}

func (n *NodeId) Rand() {
	i64 := rand.Uint64() >> 1
	*n = NodeId(i64)
}

func (n NodeId) Size() int {
	return int(reflect.TypeOf(n).Size())
}

func (n *NodeId) NetRead(data []byte) error {
	reader := bytes.NewReader(data)
	return binary.Read(reader, binary.BigEndian, n)
}

func (s *SessionId) FromByteArray(arr []byte) {
	i64 := binary.BigEndian.Uint64(arr)>>1
	*s = SessionId(i64)
}

func (s *SessionId) Rand() {
	i64 := rand.Uint64() >> 1
	*s = SessionId(i64)
}

func (s SessionId) Size() int {
	return int(reflect.TypeOf(s).Size())
}

func (s *SessionId) NetRead(data []byte) error {
	reader := bytes.NewReader(data)
	return binary.Read(reader, binary.BigEndian, s)
}

type SignalChan chan struct{}

var (
	NotifyEvent = struct{}{}
)

type Application interface {
	//Runtime() run
}