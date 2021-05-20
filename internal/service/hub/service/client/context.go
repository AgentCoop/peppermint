package client

import (
	"github.com/AgentCoop/peppermint/internal/grpc/client"
)

type JoinContext interface {
	ReqChan(int) client.ReqChan
	ResChan(int) client.ResChan
}

