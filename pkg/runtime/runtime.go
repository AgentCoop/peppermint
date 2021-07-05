package runtime

import (
	"github.com/AgentCoop/peppermint/pkg"
	"github.com/AgentCoop/peppermint/pkg/service"
)

type Runtime interface {
	App() pkg.App
	NodeManager() pkg.NodeManager
	CliParser() pkg.CliParser
	RegisterService(string, service.Service)
	Services() []service.Service
	ServicePolicyByName(string) service.ServicePolicy
	ServiceByName(string) service.Service
	EncKeyStore() InMemoryStore
}
