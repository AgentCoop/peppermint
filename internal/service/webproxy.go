package service

import (
	"github.com/AgentCoop/peppermint/internal/runtime/configurator"
	"github.com/AgentCoop/peppermint/pkg/service/balancer"
	"net"
)

type WebProxyConfigurator interface {
	configurator.Configurator
	balancer.BalancerConfigurator
	Address() net.Addr
	ServerName() string
	X509CertPEM() []byte
	X509KeyPEM() []byte
}
