package deps

import (
	"github.com/AgentCoop/peppermint/internal"
	"net"
)

type Configurator interface {
	Fetch() error // fetch configuration data from DB
	MergeCliOptions(CliParser)
}

type NodeConfigurator interface {
	Configurator
	ExternalId() internal.NodeId // :)
	E2E_EncryptionEnabled() bool
	EncKey() []byte
}

type ServiceConfigurator interface {
	Configurator
	Address() net.Addr
}
