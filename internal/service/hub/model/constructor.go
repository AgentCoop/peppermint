package model

import (
	"github.com/AgentCoop/peppermint/pkg"
)

type hubDb struct {
	pkg.Db
}

func NewDb(db pkg.Db) hubDb {
	return hubDb{db}
}

func NewConfigurator(hubDb hubDb) *cfg {
	cfg := &cfg{hubDb: hubDb}
	return cfg
}
