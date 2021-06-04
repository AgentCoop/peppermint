package session

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime"
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

func (m sessionMap) Lookup(id i.SessionId) (runtime.SessionDesc, bool) {
	v, ok := m[id]
	return v, ok
}

func (m sessionMap) Remove(id i.SessionId) {

}

func (d *sessionDesc) Expired() bool {
	now := time.Now().UTC()
	return now.After(d.expireAt)
}
