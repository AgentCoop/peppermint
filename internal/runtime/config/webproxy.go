package config

import (
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
	"net"
)

type WebProxyConfigurator interface {
	deps.Configurator
	BalancerConfigurator
	Address() net.Addr
	ServerName() string
	X509CertPEM() []byte
	X509KeyPEM() []byte
}

type BalancerConfigurator interface {

}
