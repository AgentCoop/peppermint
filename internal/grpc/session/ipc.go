package session

import (
	"errors"
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"sync"
)

type gRpcCallOder int
type StreamType int

var (
	ErrOutOfOrderCall    = errors.New("gRPC call is out of order")
	ErrNotStreamableCall = errors.New("gRPC call is not streamable")
	ErrStreamRequired    = errors.New("server-side streaming RPC requires a stream")
)

const (
	Sequential gRpcCallOder = iota
	OutOfOrder
)

const MaxChans = 8

type ipc struct {
	sId                 i.SessionId
	callOrder           gRpcCallOder
	accessBitMask       int
	lazyInitBitmask     int
	lazyInitMu          [MaxChans]sync.Mutex
	svcChanMu           [MaxChans]sync.Mutex
	svcChan             [MaxChans]chan interface{}
	grpcChan            [MaxChans]chan interface{}
	streams             [MaxChans]g.Stream
	serviceJob          job.Job
	shutdownOnce        sync.Once
}

func (c *ipc) Svc_StreamSend(chanIdx uint, data interface{}) {
	err := c.streams[chanIdx].SendMsg(data)
	if err != nil { c.shutdown(err) }
}

func (c *ipc) Svc_StreamClose(chanIdx uint) {
	c.streams[chanIdx].Close()
	c.grpcChan[chanIdx] <- nil
}

func (c *ipc) Grpc_MakeStreamable(chanIdx uint, stream g.Stream) {
	c.streams[chanIdx] = stream
}

func (c *ipc) Grpc_WaitForStreamClose(chanIdx uint) error {
	defer c.svcChanMu[chanIdx].Unlock()
	err := <-c.grpcChan[chanIdx]
	switch v := err.(type) {
	case error:
		return v
	default:
		return nil
	}
}

func (c *ipc) SessionId() i.SessionId {
	return c.sId
}

func (c *ipc) _shutdown(err interface{}) {
	for i := 0; i < MaxChans; i++ {
		// Terminate job tasks listening on the initialized channels
		// and propagate error to the gRPC layer
		if c.lazyInitBitmask&(1<<i) != 0 {
			close(c.svcChan[i])
			c.grpcChan[i] <- err
		}
	}
}

func (c *ipc) shutdown(err interface{}) {
	c.shutdownOnce.Do(func() {
		c._shutdown(err)
	})
}

func (c *ipc) chansLazyInit(chanIdx int) {
	c.lazyInitMu[chanIdx].Lock()
	defer c.lazyInitMu[chanIdx].Unlock()
	mask := (1 << chanIdx)
	if c.lazyInitBitmask&mask == 0 {
		c.grpcChan[chanIdx] = make(chan interface{}, 1)
		c.svcChan[chanIdx] = make(chan interface{}, 1)
		c.lazyInitBitmask |= mask
	}
}

func (c *ipc) grpc_Send(chanIdx int, streamable bool, data interface{}) {
	c.svcChanMu[chanIdx].Lock()
	c.chansLazyInit(chanIdx)
	mask := (1 << chanIdx)
	prevAccess := c.accessBitMask
	c.accessBitMask |= mask
	switch {
	case !streamable && prevAccess&mask != 0:
		c.shutdown(ErrNotStreamableCall)
		return
	case c.callOrder == Sequential && prevAccess^(mask-1) != 0 && prevAccess != c.accessBitMask: // out of order call
		c.shutdown(ErrOutOfOrderCall)
		return
	}
	c.svcChan[chanIdx] <- data
}

func (c *ipc) Grpc_Send(chanIdx int, data interface{}) {
	c.grpc_Send(chanIdx, false, data)
}

// A wrapper method to allow client to make pseudo-streaming calls
func (c *ipc) Grpc_SendStreamable(chanId int, data interface{}) {
	c.grpc_Send(chanId, true, data)
}

func (c *ipc) Grpc_Recv(chanIdx int) interface{} {
	// Assume that access to the service channel was locked in a preceding Grpc_Send call,
	// otherwise we are misusing the communication mechanism
	defer c.svcChanMu[chanIdx].Unlock()
	v := <-c.grpcChan[chanIdx]
	return v
}

func (c *ipc) Svc_Send(chanIdx int, data interface{}) {
	switch v := data.(type) {
	case error:
		c.shutdown(v)
		return
	default:
		c.grpcChan[chanIdx] <- data
	}
}

func (c *ipc) Svc_Recv(chanIdx int) interface{} {
	c.chansLazyInit(chanIdx)
	v := <-c.svcChan[chanIdx]
	return v
}

func (c *ipc) Job() job.Job {
	return c.serviceJob
}
