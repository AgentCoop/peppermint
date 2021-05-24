package server

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/model/node"
)

type session struct {}

type Session interface {
	Id()
	NodeId()
	Job() job.Job
	Node(id i.SessionId) (service.Node, error)
}

func(session) Node(id i.SessionId) (service.Node, error) {
	return nil, nil
}

var sessionMap map[i.SessionId]job.Job

func FindSessionById(id i.SessionId) job.Job {
	j, ok := sessionMap[id]
	if ok {
		return j
	} else {
		return nil
	}
}

func StartNewSession(j job.Job) i.SessionId {
	return 2
}