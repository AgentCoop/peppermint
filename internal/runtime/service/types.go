package service

import (
	"github.com/AgentCoop/peppermint/pkg/grpc"
	"github.com/AgentCoop/peppermint/pkg/service"
)

type serviceInfo struct {
	shortName   string
	cfg         service.ServiceConfigurator
	policy      *svcPolicy
	initializer func() grpc.BaseServer
}
