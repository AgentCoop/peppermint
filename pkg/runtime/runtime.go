package runtime

import (
	"github.com/AgentCoop/peppermint/pkg"
	"github.com/AgentCoop/peppermint/pkg/node"
	"github.com/AgentCoop/peppermint/pkg/service"
)

type Runtime interface {
	NodeManager() node.NodeManager
	CliParser() pkg.CliParser
	NodeConfigurator() node.NodeConfigurator
	RegisterService(string, service.Service)
	Services() []service.Service
	ServicePolicyByName(string) service.ServicePolicy
	ServiceByName(string) service.Service

	EncKeyStore() InMemoryStore
}
