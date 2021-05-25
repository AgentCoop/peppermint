package hub

import (
	"github.com/AgentCoop/peppermint/internal/grpc/server"
)

type JoinContext interface {
	ReqChan() [2]server.PairChan
	ResChan() [2]server.PairChan
}
