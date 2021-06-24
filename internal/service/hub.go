package service

import (
	"github.com/AgentCoop/peppermint/internal/runtime/configurator"
	"net"
)

type HubConfigurator interface {
	configurator.Configurator
	Address() net.Addr
	Secret() string
}
