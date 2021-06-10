package session

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"time"
)

func (m sessionMap) new(expireInSecs time.Duration, callOrder gRpcCallOder) *sessionDesc {
	now := time.Now().UTC()
	var sId i.SessionId
	sId.Rand()
	desc := &sessionDesc{
		id:        sId,
		ipc:       newIpc(callOrder),
		createdAt: now,
		expireAt:  now.Add(expireInSecs * time.Second),
	}
	desc.Job().SetValue(desc)
	m[sId] = desc
	return desc
}

func NewSession(expireInSecs time.Duration) *sessionDesc {
	return sMap.new(expireInSecs, Sequential)
}

// Session lifetime time in seconds
func newIpc(callOrder gRpcCallOder) *ipc {
	c := &ipc{}
	c.serviceJob = job.NewJob(nil)
	c.serviceJob.WithErrorWrapper(grpc.ErrorWrapper)
	c.serviceJob.WithShutdown(c.shutdown)
	c.callOrder = callOrder
	return c
}

