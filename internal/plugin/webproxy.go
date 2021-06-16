package plugin

import (
	"github.com/AgentCoop/peppermint/internal/runtime/configurator"
	"net"
)

type WebProxyConfigurator interface {
	configurator.Configurator
	BalancerConfigurator
	Address() net.Addr
	ServerName() string
	X509CertPEM() []byte
	X509KeyPEM() []byte
}

type BalancerConfigurator interface {

}
