package node

import "github.com/AgentCoop/peppermint/internal/runtime/store"

func NewNodeManager() *manager {
	m := new(manager)
	m.encKeyStore = store.NewInMemoryStore()
	return m
}

func NewConfigurator() *cfg {
	cfg := new(cfg)
	return cfg
}

