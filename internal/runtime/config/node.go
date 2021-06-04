package config

import (
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
)

type NodeConfigurator interface {
	deps.Configurator
	NodeId() internal.NodeId
	Tags() []string
	EncKey() []byte
	IsSecure() bool
}
