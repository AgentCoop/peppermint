package node

import (
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/pkg"
)

type Configurator interface {
	Fetch(uint) error // fetch configuration data from DB
	Refresh(uint) error
	MergeCliOptions(parser pkg.CliParser)
}

type NodeConfigurator interface {
	Configurator
	ExternalId() internal.NodeId // :)
	E2E_EncryptionEnabled() bool
	EncKey() []byte
}
