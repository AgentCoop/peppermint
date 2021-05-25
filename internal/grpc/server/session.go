package server

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/utils"
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
	job job.Job
	createdAt time.Time
	expireAt time.Time
}

func (m sessionMap) New(j job.Job, expireInSecs time.Duration) i.SessionId {
	now := time.Now().UTC()
	desc := &sessionDesc{
		job:             j,
		createdAt:       now,
		expireAt: now.Add(expireInSecs * time.Second),
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