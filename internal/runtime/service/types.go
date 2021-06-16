package service

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
)

type serviceInfo struct {
	shortName   string
	cfg         runtime.ServiceConfigurator
	policy      *svcPolicy
	initializer func() runtime.BaseServer
}
