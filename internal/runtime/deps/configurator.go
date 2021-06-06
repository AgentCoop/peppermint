package deps

import "net"

type Configurator interface {
	Fetch() error // fetch configuration data from DB
	MergeCliOptions(CliParser)
}

type NodeConfigurator interface {
	Configurator
	E2E_EncryptionEnabled() bool
	EncKey() []byte
}

type ServiceConfigurator interface {
	Configurator
	Address() net.Addr
}
