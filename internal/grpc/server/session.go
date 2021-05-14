package server

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc"
)

type session struct {

}

type Session interface {
	Id()
	NodeId()
	Job() job.Job
}

var sessionMap map[grpc.SessionId]job.Job

func FindSessionById(id grpc.SessionId) job.Job {
	j, ok := sessionMap[id]
	if ok {
		return j
	} else {
		return nil
	}
}

func StartNewSession(j job.Job) grpc.SessionId {
	return 0
}