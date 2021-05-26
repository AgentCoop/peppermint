package server

import (
	"errors"
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/utils"
	"sync"
	"time"
)

type sessionMap map[i.SessionId]*sessionDesc

var (
	sMap sessionMap
)

func init() {
	sMap = make(sessionMap, 0)
	runtime.GlobalRegistry().SetGrpcSession(sMap)
}

type sessionDesc struct {
	job       job.Job
	createdAt time.Time
	expireAt  time.Time
}

func (m sessionMap) New(j job.Job, expireInSecs time.Duration) i.SessionId {
	now := time.Now().UTC()
	desc := &sessionDesc{
		job:       j,
		createdAt: now,
		expireAt:  now.Add(expireInSecs * time.Second),
	}
	id := utils.RandomGrpcSessionId()
	m[id] = desc
	return id
}

func (m sessionMap) Lookup(id i.SessionId) (runtime.SessionDesc, bool) {
	v, ok := m[id]
	return v, ok
}

func (m sessionMap) Remove(id i.SessionId) {

}

// Session descriptor methods
func (d *sessionDesc) Job() job.Job {
	return d.job
}

func (d *sessionDesc) Expired() bool {
	now := time.Now().UTC()
	return now.After(d.expireAt)
}

//
// The main mechanism of communication between the gRPC and service layer.
//
type gRpcCallOder int

const (
	Sequential gRpcCallOder = iota
	OutOfOrder
)

const CommunicatorMaxChans = 8

var (
	ErrOutOfOrderCall    = errors.New("gRPC call is out of order")
	ErrNotStreamableCall = errors.New("gRPC call is not streamable")
)

type communicator struct {
	callOrder       gRpcCallOder
	accessBitMask   int
	lazyInitBitmask int
	lazyInitMu      [CommunicatorMaxChans]sync.Mutex
	svcChanMu       [CommunicatorMaxChans]sync.Mutex
	svcChan         [CommunicatorMaxChans]chan interface{}
	grpcChan        [CommunicatorMaxChans]chan interface{}
	serviceJob      job.Job
	shutdownOnce    sync.Once
}

func NewCommunicator() *communicator {
	c := &communicator{}
	c.serviceJob = job.NewJob(c)
	c.serviceJob.WithErrorWrapper(utils.GrpcErrorWrapper)
	c.serviceJob.WithShutdown(c.shutdown)
	return c
}

func NewOutOfOrderCommunicator() *communicator {
	c := NewCommunicator()
	c.callOrder = OutOfOrder
	return c
}

func (c *communicator) _shutdown(err interface{}) {
	for i := 0; i < CommunicatorMaxChans; i++ {
		// Terminate job tasks listening on initialized channels
		// and propagate error to the gRPC layer
		if c.lazyInitBitmask&(1<<i) != 0 {
			close(c.svcChan[i])
			c.grpcChan[i] <- err
		}
	}
}

func (c *communicator) shutdown(err interface{}) {
	c.shutdownOnce.Do(func() {
		c._shutdown(err)
	})
}

func (c *communicator) chansLazyInit(chanIdx int) {
	c.lazyInitMu[chanIdx].Lock()
	defer c.lazyInitMu[chanIdx].Unlock()
	mask := (1 << chanIdx)
	if c.lazyInitBitmask&mask == 0 {
		c.grpcChan[chanIdx] = make(chan interface{}, 1)
		c.svcChan[chanIdx] = make(chan interface{}, 1)
		c.lazyInitBitmask |= mask
	}
}

func (c *communicator) grpcTx(chanIdx int, streamable bool, data interface{}) {
	mask := (1 << chanIdx)
	prevAccess := c.accessBitMask
	c.chansLazyInit(chanIdx)
	c.accessBitMask |= mask
	switch {
	case !streamable && prevAccess&mask != 0:
		c.shutdown(ErrNotStreamableCall)
		return
	case c.callOrder == Sequential && prevAccess^(mask-1) != 0 && prevAccess != c.accessBitMask: // out of order call
		c.shutdown(ErrOutOfOrderCall)
		return
	}
	c.svcChanMu[chanIdx].Lock()
	c.svcChan[chanIdx] <- data
}

// grpcTx public methods
func (c *communicator) GrpcTx(chanId int, data interface{}) {
	c.grpcTx(chanId, false, data)
}

func (c *communicator) GrpcTxStreamable(chanId int, data interface{}) {
	c.grpcTx(chanId, true, data)
}

//.

func (c *communicator) grpcRx(chanIdx int) interface{} {
	// Assume that access to the service channel was locked in a preceding grpcTx call,
	// otherwise we are misusing the communication mechanism
	defer c.svcChanMu[chanIdx].Unlock()
	v := <-c.grpcChan[chanIdx]
	return v
}

// grpcRx public method
func (c *communicator) GrpcRx(chanIdx int) interface{} {
	return c.grpcRx(chanIdx)
}

//.

func (c *communicator) serviceTx(chanIdx int, data interface{}) {
	switch data.(type) {
	case error:
		c.shutdown(data.(error))
		return
	default:
		c.grpcChan[chanIdx] <- data
	}
}

// serviceTx public method
func (c *communicator) ServiceTx(chanIdx int, data interface{}) {
	c.serviceTx(chanIdx, data)
}

//.

func (c *communicator) serviceRx(chanIdx int) interface{} {
	c.chansLazyInit(chanIdx)
	v := <-c.svcChan[chanIdx]
	return v
}

// serviceRx public method
func (c *communicator) ServiceRx(chanIdx int) interface{} {
	return c.serviceRx(chanIdx)
}

//.

func (c *communicator) Job() job.Job {
	return c.serviceJob
}
