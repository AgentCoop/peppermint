package deps

import "net"

type Configurator interface {
	Fetch() error // fetch configuration data from DB
	MergeCliOptions(CliParser)
}

type NodeConfigurator interface {
	Configurator
	EncKey() []byte
}

type ServiceConfigurator interface {
	Configurator
	Address() net.Addr
}
