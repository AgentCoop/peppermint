package session

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"time"
)

type sessionMap map[i.SessionId]*sessionDesc

var (
	sMap sessionMap
)

type sessionDesc struct {
	job       job.Job
	createdAt time.Time
	expireAt  time.Time
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
