package session

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime"
	utils "github.com/AgentCoop/peppermint/internal/utils/grpc"
	"time"
)


func (m sessionMap) New(j job.Job, expireInSecs time.Duration) i.SessionId {
	now := time.Now().UTC()
	desc := &sessionDesc{
		job:       j,
		createdAt: now,
		expireAt:  now.Add(expireInSecs * time.Second),
	}
	sId := i.UniqueId(0).Rand().SessionId()
	m[sId] = desc
	return sId
}

// Session lifetime time in seconds
func NewIpc(sessLifetime time.Duration) *ipc {
	c := &ipc{}
	c.serviceJob = job.NewJob(c)
	c.serviceJob.WithErrorWrapper(utils.GrpcErrorWrapper)
	c.serviceJob.WithShutdown(c.shutdown)
	c.sId = runtime.GlobalRegistry().GrpcSession().New(c.serviceJob, sessLifetime)
	return c
}

func NewOutOfOrderIpc(sessLifetime time.Duration) *ipc {
	c := NewIpc(sessLifetime)
	c.callOrder = OutOfOrder
	return c
}
