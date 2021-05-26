package config

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
	"net"
)

type WebProxyConfigurator interface {
	runtime.Configurator
	Address() net.Addr
	ServerName() string
	X509CertPEM() []byte
	X509KeyPEM() []byte
}