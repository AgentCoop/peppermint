package join

import (
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/session"
	"time"
)

func CreateSession() grpc.Session {
	joinCtx := new(joinContext)
	sess := session.NewSession(time.Minute)
	sess.Job().AddTask(joinCtx.JoinHelloTask)
	sess.Job().AddTask(joinCtx.JoinTask)
	sess.Job().Run()
	return sess
}
