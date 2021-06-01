package config

import (
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

type NodeConfigurator interface {
	runtime.Configurator
	NodeId() internal.NodeId
}
