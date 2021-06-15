package service

import (
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
)

type serviceInfo struct {
	shortName   string
	cfg         deps.ServiceConfigurator
	policy      *svcPolicy
	initializer func() grpc.BaseServer
}
