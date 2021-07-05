package node

import (
	"github.com/AgentCoop/peppermint/pkg"
)

type Configurator interface {
	Fetch(uint) error // fetch configuration data from DB
	Refresh(uint) error
	MergeCliOptions(parser pkg.CliParser)
}
