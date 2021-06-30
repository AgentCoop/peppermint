package session

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"time"
)

type sessionMap map[i.SessionId]*sessionDesc

var (
	sMap sessionMap
)

type sessionDesc struct {
	id        i.SessionId
	job       job.Job
	createdAt time.Time
	expireAt  time.Time
	ipc       *ipc
	taskCtx   interface{}
}

func (s *sessionDesc) Id() i.SessionId {
	return s.id
}

func (s *sessionDesc) Ipc() grpc.GrpcServiceLayersIpc {
	return s.ipc
}

func (s *sessionDesc) Job() job.Job {
	return s.ipc.serviceJob
}

func (s *sessionDesc) WithTaskContext(ctx interface{}) {
	s.taskCtx = ctx
}

func (s *sessionDesc) TaskContext() interface{} {
	return s.taskCtx
}

func (m sessionMap) Remove(id i.SessionId) {

}

func (d *sessionDesc) Expired() bool {
	now := time.Now().UTC()
	return now.After(d.expireAt)
}
