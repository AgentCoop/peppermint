package config

import (
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
	"net"
)

type HubConfigurator interface {
	deps.Configurator
	Address() net.Addr
	Secret() string
}
