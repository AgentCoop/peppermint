package client

import (
	"github.com/AgentCoop/peppermint/internal/grpc/client"
)

type JoinContext interface {
	JoinHelloRequest() client.ReqChan
	JoinHelloResponse() client.ResChan
	JoinRequest() client.ReqChan
	JoinResponse() client.ResChan
}

