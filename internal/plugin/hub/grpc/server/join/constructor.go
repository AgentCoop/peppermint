package join

import (
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/session"
	"time"
)

func CreateSession() grpc.Session {
	sess := session.NewSession(time.Minute)
	return sess
}
