package config

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
	"net"
)

type HubConfigurator interface {
	runtime.Configurator
	Address() net.Addr
	Secret() string
}
